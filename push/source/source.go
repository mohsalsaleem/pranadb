package source

import (
	"fmt"
	"github.com/squareup/pranadb/conf"
	"github.com/squareup/pranadb/errors"
	"github.com/squareup/pranadb/kafka"
	"github.com/squareup/pranadb/meta"
	"github.com/squareup/pranadb/push/mover"
	"github.com/squareup/pranadb/push/sched"
	"github.com/squareup/pranadb/table"
	"log"
	"time"

	"github.com/squareup/pranadb/cluster"
	"github.com/squareup/pranadb/common"
	"github.com/squareup/pranadb/push/exec"
	"github.com/squareup/pranadb/sharder"
)

// TODO make configurable
const (
	numConsumersPerSource = 1
	pollTimeoutMs         = 1000
	maxPollMessages       = 10000
)

type RowProcessor interface {
}

type SchedulerSelector interface {
	ChooseLocalScheduler(key []byte) (*sched.ShardScheduler, error)
}

type Source struct {
	sourceInfo              *common.SourceInfo
	tableExecutor           *exec.TableExecutor
	sharder                 *sharder.Sharder
	cluster                 cluster.Cluster
	mover                   *mover.Mover
	messageParser           *MessageParser
	schedSelector           SchedulerSelector
	msgProvFact             kafka.MessageProviderFactory
	msgConsumers            []*MessageConsumer
	startupCommittedOffsets map[int32]int64
	queryExec               common.SimpleQueryExec
}

func NewSource(sourceInfo *common.SourceInfo, tableExec *exec.TableExecutor, sharder *sharder.Sharder,
	cluster cluster.Cluster, mover *mover.Mover, schedSelector SchedulerSelector, cfg *conf.Config,
	queryExec common.SimpleQueryExec) (*Source, error) {
	// TODO we should validate the sourceinfo - e.g. check that number of col selectors, column names and column types are the same
	var msgProvFact kafka.MessageProviderFactory
	ti := sourceInfo.TopicInfo
	if ti == nil {
		// TODO not sure if we need this... parser should catch it?
		return nil, errors.NewUserErrorF(errors.MissingTopicInfo, "No topic info configured for source %s", sourceInfo.Name)
	}
	if cfg.KafkaBrokers == nil {
		return nil, errors.NewUserError(errors.MissingKafkaBrokers, "No Kafka brokers configured")
	}
	brokerConf, ok := cfg.KafkaBrokers[ti.BrokerName]
	if !ok {
		return nil, errors.NewUserErrorF(errors.UnknownBrokerName, "Unknown broker. Name: %s", ti.BrokerName)
	}
	props := copyAndAddAll(brokerConf.Properties, ti.Properties)
	messageParser, err := NewMessageParser(sourceInfo)
	if err != nil {
		return nil, err
	}
	groupID := GenerateGroupID(cfg.ClusterID, sourceInfo)
	switch brokerConf.ClientType {
	case conf.BrokerClientFake:
		var err error
		msgProvFact, err = kafka.NewFakeMessageProviderFactory(ti.TopicName, props, groupID)
		if err != nil {
			return nil, err
		}
	case conf.BrokerClientDefault:
		msgProvFact = kafka.NewCfltMessageProviderFactory(ti.TopicName, props, groupID)
	default:
		return nil, errors.NewUserErrorF(errors.UnsupportedBrokerClientType, "Unsupported broker client type %d", brokerConf.ClientType)
	}
	return &Source{
		sourceInfo:    sourceInfo,
		tableExecutor: tableExec,
		sharder:       sharder,
		cluster:       cluster,
		mover:         mover,
		schedSelector: schedSelector,
		messageParser: messageParser,
		msgProvFact:   msgProvFact,
		queryExec:     queryExec,
	}, nil
}

func (s *Source) Start() error {

	if err := s.loadStartupCommittedOffsets(); err != nil {
		return err
	}

	for i := 0; i < numConsumersPerSource; i++ {
		msgProvider, err := s.msgProvFact.NewMessageProvider()
		if err != nil {
			return err
		}
		// We choose a local scheduler based on the name of the schema, name of source and number of ordinal of the message consumer
		// The shard of that scheduler is where ingested rows will be staged, ready to be forwarded to their
		// destination shards. Having a message consumer pinned to a shard ensures all messages for a particular
		// Kafka partition are forwarded in order. If different batches used different shards, we couldn't guarantee that.
		// We can't write ingested rows directly into the target shards as we can't commit offsets atomically that way unless
		// we write each in a batch with a single row - this is because messages for the same Kafka partition would be hashed
		// to different target Prana shards.
		// A particular message consumer will always stage all its messages in the same local shard for the life of the consumer
		// We can vary the number of consumers to scale the forwarding.
		sKey := fmt.Sprintf("%s-%s-%d", s.sourceInfo.SchemaName, s.sourceInfo.Name, i)
		scheduler, err := s.schedSelector.ChooseLocalScheduler([]byte(sKey))
		if err != nil {
			return err
		}
		consumer := NewMessageConsumer(msgProvider, pollTimeoutMs*time.Millisecond, maxPollMessages, s, scheduler)
		s.msgConsumers = append(s.msgConsumers, consumer)
		consumer.Start()
	}

	// TODO we should wait for rebalancing to settle down when we start the consumers - there may be churn as
	// the cluster starts up - we do not want to consume messages in that time as the partition allocation could
	// change and we could end up with messages from the same partition being processed out of order as they are
	// forwarded into different staging shard IDs.

	// We will do this by stalling message processing until a fixed delay after the last rebalancing.

	return nil
}

func (s *Source) loadStartupCommittedOffsets() error {
	rows, err := s.queryExec.ExecuteQuery("sys",
		fmt.Sprintf("select partition_id, offset from %s where schema_name='%s' and source_name='%s'",
			meta.SourceOffsetsTableName, s.sourceInfo.SchemaName, s.sourceInfo.Name))
	if err != nil {
		return err
	}
	for i := 0; i < rows.RowCount(); i++ {
		row := rows.GetRow(i)
		partID := row.GetInt64(0)
		offset := row.GetInt64(1)
		currOff, ok := s.startupCommittedOffsets[int32(partID)]
		if !ok || offset > currOff {
			// It's possible that we might have duplicate rows as the mapping of message consumer to ingesting shard ID
			// changes, so in that case we take the largest offset
			s.startupCommittedOffsets[int32(partID)] = offset
		}
	}
	return nil
}

func (s *Source) Stop() error {
	for _, consumer := range s.msgConsumers {
		consumer.Stop()
	}
	return nil
}

func (s *Source) AddConsumingExecutor(executor exec.PushExecutor) {
	s.tableExecutor.AddConsumingNode(executor)
}

func (s *Source) RemoveConsumingExecutor(executor exec.PushExecutor) {
	s.tableExecutor.RemoveConsumingNode(executor)
}

func (s *Source) handleMessages(messages []*kafka.Message, offsetsToCommit map[int32]int64, scheduler *sched.ShardScheduler) error {
	errChan := scheduler.ScheduleAction(func() error {
		return s.ingestMessages(messages, offsetsToCommit, scheduler.ShardID())
	})
	err, ok := <-errChan
	if !ok {
		panic("channel closed")
	}
	return err
}

func (s *Source) ingestMessages(messages []*kafka.Message, offsetsToCommit map[int32]int64, shardID uint64) error {

	rows, err := s.messageParser.ParseMessages(messages)
	if err != nil {
		return err
	}

	log.Printf("Ingesting rows on shard %d and node %d", shardID, s.cluster.GetNodeID())

	// TODO where Source has no key - need to create one

	// Partition the rows and send them to the appropriate shards
	info := s.sourceInfo.TableInfo
	pkCols := info.PrimaryKeyCols
	colTypes := info.ColumnTypes
	tableID := info.ID
	batch := cluster.NewWriteBatch(shardID, false)

	for i := 0; i < rows.RowCount(); i++ {
		row := rows.GetRow(i)
		key := make([]byte, 0, 8)
		key, err := common.EncodeKeyCols(&row, pkCols, colTypes, key)
		if err != nil {
			return err
		}
		destShardID, err := s.sharder.CalculateShard(sharder.ShardTypeHash, key)
		if err != nil {
			return err
		}
		// TODO we can consider an optimisation where execute on any local shards directly
		err = s.mover.QueueForRemoteSend(destShardID, &row, shardID, tableID, colTypes, batch)
		if err != nil {
			return err
		}
	}

	s.commitOffsetsToPrana(offsetsToCommit, batch)

	if err := s.cluster.WriteBatch(batch); err != nil {
		return err
	}
	return s.mover.TransferData(shardID, true)
}

// We commit the Kafka offsets in the same batch as we stage the rows for forwarding.
// We need to commit the offsets here, as if the node fails after committing in Prana but before committing
// in Kafka, then on recovery the same messages can be delivered again. In order to filter out the duplicates
// we store the last received offsets here and we will reject any in the consumer that we've seen before
func (s *Source) commitOffsetsToPrana(offsets map[int32]int64, batch *cluster.WriteBatch) {
	for partID, offset := range offsets {
		key := table.EncodeTableKeyPrefix(common.OffsetsTableID, batch.ShardID, 40)
		key = common.KeyEncodeString(key, s.sourceInfo.SchemaName)
		key = common.KeyEncodeString(key, s.sourceInfo.Name)
		key = common.KeyEncodeInt64(key, int64(partID))
		val := make([]byte, 8)
		val = common.AppendUint64ToBufferLE(val, uint64(offset))
		batch.AddPut(key, val)
	}
}

func (s *Source) TableExecutor() *exec.TableExecutor {
	return s.tableExecutor
}

func copyAndAddAll(p1 map[string]string, p2 map[string]string) map[string]string {
	m := make(map[string]string, len(p1)+len(p2))
	for k, v := range p2 {
		m[k] = v
	}
	// p1 properties override p2 so we add them last
	for k, v := range p1 {
		m[k] = v
	}
	return m
}

func (s *Source) startupLastOffset(partitionID int32) int64 {
	off, ok := s.startupCommittedOffsets[partitionID]
	if !ok {
		off = 0
	}
	return off
}

func GenerateGroupID(clusterID int, sourceInfo *common.SourceInfo) string {
	return fmt.Sprintf("prana-source-%d-%s-%s", clusterID, sourceInfo.SchemaName, sourceInfo.Name)
}