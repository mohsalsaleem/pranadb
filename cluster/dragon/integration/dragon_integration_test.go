package integration

import (
	"flag"
	"fmt"
	"github.com/squareup/pranadb/common"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/squareup/pranadb/errors"

	log "github.com/sirupsen/logrus"
	"github.com/squareup/pranadb/conf"

	dragon "github.com/squareup/pranadb/cluster/dragon"
	"github.com/stretchr/testify/require"

	"github.com/squareup/pranadb/cluster"
)

var dragonCluster []cluster.Cluster

const (
	numShards = 10
)

var dataDir string

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Infof("-short: skipped")
		return
	}
	var err error
	dataDir, err = ioutil.TempDir("", "dragon-test")
	if err != nil {
		panic("failed to create temp dir")
	}
	dragonCluster, err = startDragonCluster(dataDir)
	if err != nil {
		panic(fmt.Sprintf("failed to start dragon cluster %+v", err))
	}
	defer func() {
		stopDragonCluster()
		defer func() {
			// We want to do this even if stopDragonCluster fails hence the extra defer
			err := os.RemoveAll(dataDir)
			if err != nil {
				panic("failed to remove test dir")
			}
		}()
	}()
	m.Run()
}

func getLocalNodeAndLocalShard() (cluster.Cluster, uint64) {
	clust := dragonCluster[0]
	shardID := clust.GetLocalShardIDs()[0]
	return clust, shardID
}

func TestLocalPutGet(t *testing.T) {
	node, localShard := getLocalNodeAndLocalShard()

	key := []byte("somekey")
	value := []byte("somevalue")

	kvPair := cluster.KVPair{
		Key:   key,
		Value: value,
	}

	writeBatch := createWriteBatchWithPuts(localShard, kvPair)

	err := node.WriteBatch(&writeBatch, false)
	require.NoError(t, err)

	res, err := node.LocalGet(key)
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, string(value), string(res))
}

func TestLocalPutDelete(t *testing.T) {
	node, localShard := getLocalNodeAndLocalShard()

	key := []byte("somekey")
	value := []byte("somevalue")

	kvPair := cluster.KVPair{
		Key:   key,
		Value: value,
	}

	writeBatch := createWriteBatchWithPuts(localShard, kvPair)

	err := node.WriteBatch(&writeBatch, false)
	require.NoError(t, err)

	res, err := node.LocalGet(key)
	require.NoError(t, err)
	require.NotNil(t, res)

	deleteBatch := createWriteBatchWithDeletes(localShard, key)

	err = node.WriteBatch(&deleteBatch, false)
	require.NoError(t, err)

	res, err = node.LocalGet(key)
	require.NoError(t, err)
	require.Nil(t, res)
}

func TestLocalScan(t *testing.T) {
	testLocalScan(t, -1, 10)
}

func TestLocalScanSmallLimit(t *testing.T) {
	testLocalScan(t, 3, 3)
}

func TestLocalScanBigLimit(t *testing.T) {
	testLocalScan(t, 1000, 10)
}

func testLocalScan(t *testing.T, limit int, expected int) {
	t.Helper()
	node, localShard := getLocalNodeAndLocalShard()

	var kvPairs []cluster.KVPair
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			k := []byte(fmt.Sprintf("foo-%02d/bar-%02d", i, j))
			v := []byte(fmt.Sprintf("somevalue%02d", j))
			kvPairs = append(kvPairs, cluster.KVPair{Key: k, Value: v})
		}
	}
	rand.Shuffle(len(kvPairs), func(i, j int) {
		kvPairs[i], kvPairs[j] = kvPairs[j], kvPairs[i]
	})

	wb := cluster.NewWriteBatch(localShard)
	for _, kvPair := range kvPairs {
		wb.AddPut(kvPair.Key, kvPair.Value)
	}

	err := node.WriteBatch(wb, false)
	require.NoError(t, err)

	keyStart := []byte("foo-06")
	keyEnd := []byte("foo-07")

	var res []cluster.KVPair
	res, err = node.LocalScan(keyStart, keyEnd, limit)
	require.NoError(t, err)

	require.Equal(t, expected, len(res))
	for i, kvPair := range res {
		expectedK := fmt.Sprintf("foo-06/bar-%02d", i)
		expectedV := fmt.Sprintf("somevalue%02d", i)
		require.Equal(t, expectedK, string(kvPair.Key))
		require.Equal(t, expectedV, string(kvPair.Value))
	}
}

func TestGenerateClusterSequence(t *testing.T) {
	for i := 0; i < 10; i++ {
		id, err := dragonCluster[i%len(dragonCluster)].GenerateClusterSequence("sequence1")
		require.NoError(t, err)
		require.Equal(t, uint64(i), id)
	}
	for i := 0; i < 10; i++ {
		id, err := dragonCluster[i%len(dragonCluster)].GenerateClusterSequence("sequence2")
		require.NoError(t, err)
		require.Equal(t, uint64(i), id)
	}
}

func TestGetNodeID(t *testing.T) {
	for i := 0; i < len(dragonCluster); i++ {
		require.Equal(t, i, dragonCluster[i].GetNodeID())
	}
}

func TestGetAllShardIDs(t *testing.T) {
	for i := 0; i < len(dragonCluster); i++ {
		allShardIds := dragonCluster[i].GetAllShardIDs()
		require.Equal(t, numShards, len(allShardIds))
	}
}

func TestGetReleaseLock(t *testing.T) {
	dragon0 := dragonCluster[0]
	dragon1 := dragonCluster[1]
	ok, err := dragon0.GetLock("/schema1")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema1")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon1.GetLock("/schema2")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon0.GetLock("/schema2")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon0.GetLock("/")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon1.GetLock("/")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon1.ReleaseLock("/schema1")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon0.ReleaseLock("/schema2")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon0.GetLock("/")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema1")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon1.ReleaseLock("/")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema1")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.ReleaseLock("/schema1")
	require.NoError(t, err)
	require.True(t, ok)
}

func TestLocksRestart(t *testing.T) {
	dragon0 := dragonCluster[0]
	dragon1 := dragonCluster[1]
	ok, err := dragon0.GetLock("/schema1")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema2")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema3")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.ReleaseLock("/schema2")
	require.NoError(t, err)
	require.True(t, ok)
	stopDragonCluster()
	dragonCluster, err = startDragonCluster(dataDir)
	require.NoError(t, err)

	dragon0 = dragonCluster[0]
	dragon1 = dragonCluster[1]
	ok, err = dragon0.GetLock("/schema1")
	require.NoError(t, err)
	require.False(t, ok)
	ok, err = dragon1.GetLock("/schema2")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dragon1.GetLock("/schema3")
	require.NoError(t, err)
	require.False(t, ok)
}

func stopDragonCluster() {
	for _, dragon := range dragonCluster {
		err := dragon.Stop()
		if err != nil {
			panic(fmt.Sprintf("failed to stop dragon cluster %v", err))
		}
	}
}

func startDragonCluster(dataDir string) ([]cluster.Cluster, error) {

	nodeAddresses := []string{
		"localhost:63101",
		"localhost:63102",
		"localhost:63103",
	}

	chans := make([]chan error, len(nodeAddresses))
	clusterNodes := make([]cluster.Cluster, len(nodeAddresses))
	for i := 0; i < len(chans); i++ {
		ch := make(chan error)
		chans[i] = ch
		cnf := conf.NewDefaultConfig()
		cnf.NodeID = i
		cnf.ClusterID = 123
		cnf.RaftAddresses = nodeAddresses
		cnf.NumShards = numShards
		cnf.DataDir = dataDir
		cnf.ReplicationFactor = 3
		cnf.TestServer = true
		clus, err := dragon.NewDragon(*cnf, &common.AtomicBool{})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		clusterNodes[i] = clus
		clus.RegisterShardListenerFactory(&cluster.DummyShardListenerFactory{})
		clus.SetRemoteQueryExecutionCallback(&cluster.DummyRemoteQueryExecutionCallback{})

		go startDragonNode(clus, ch)
	}

	for i := 0; i < len(chans); i++ {
		err, ok := <-chans[i]
		if !ok {
			return nil, errors.Error("channel was closed")
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	time.Sleep(5 * time.Second)
	return clusterNodes, nil
}

func startDragonNode(clus cluster.Cluster, ch chan error) {
	err := clus.Start()
	ch <- err
}

func createWriteBatchWithPuts(shardID uint64, puts ...cluster.KVPair) cluster.WriteBatch {
	wb := cluster.NewWriteBatch(shardID)
	for _, kvPair := range puts {
		wb.AddPut(kvPair.Key, kvPair.Value)
	}
	return *wb
}

func createWriteBatchWithDeletes(shardID uint64, deletes ...[]byte) cluster.WriteBatch {
	wb := cluster.NewWriteBatch(shardID)
	for _, del := range deletes {
		wb.AddDelete(del)
	}
	return *wb
}
