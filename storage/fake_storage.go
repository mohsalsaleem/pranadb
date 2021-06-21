package storage

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/btree"
	"github.com/squareup/pranadb/common"
	"log"
	"sync"
)

type FakeStorage struct {
	btree              *btree.BTree
	mu                 sync.RWMutex
	clusterInfo        *ClusterInfo
	tableSequence      uint64
	remoteWriteHandler RemoteWriteHandler
}

func (f *FakeStorage) SetRemoteWriteHandler(handler RemoteWriteHandler) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.remoteWriteHandler = handler
}

func NewFakeStorage(nodeID int, numShards int) Storage {
	btree := btree.New(3)
	return &FakeStorage{
		btree:         btree,
		clusterInfo:   createClusterInfo(nodeID, numShards),
		tableSequence: uint64(100), // First 100 reserved for system tables
	}
}

func (f *FakeStorage) WriteBatch(batch *WriteBatch, localLeader bool) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	log.Printf("Write batch for shard %d, local leader %t", batch.ShardID, localLeader)
	log.Printf("Writing batch, puts %d, deletes %d", len(batch.puts.TheMap), len(batch.deletes.TheMap))
	for k, v := range batch.puts.TheMap {
		kBytes := common.StringToByteSliceZeroCopy(k)
		log.Printf("Putting key %v value %v", kBytes, v)
		f.putInternal(&kvWrapper{
			key:   kBytes,
			value: v,
		})
	}
	for k, _ := range batch.deletes.TheMap {
		kBytes := common.StringToByteSliceZeroCopy(k)
		log.Printf("Deleting key %v", kBytes)
		err := f.deleteInternal(&kvWrapper{
			key: kBytes,
		})
		if err != nil {
			return err
		}
	}
	if !localLeader && f.remoteWriteHandler != nil {
		go f.remoteWriteHandler.RemoteWriteOccurred(batch.ShardID)
	}
	return nil
}

func (f *FakeStorage) InstallExecutors(shardID uint64, plan *ExecutorPlan) {
	panic("implement me")
}

func (f *FakeStorage) GenerateTableID() (uint64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	res := f.tableSequence
	f.tableSequence++
	return res, nil
}

func (f *FakeStorage) Get(shardID uint64, key []byte, localLeader bool) ([]byte, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.getInternal(&kvWrapper{key: key}), nil
}

func (f *FakeStorage) Scan(shardID uint64, startKeyPrefix []byte, whileKeyPrefix []byte, limit int) ([]KVPair, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	if startKeyPrefix == nil {
		panic("startKeyPrefix cannot be nil")
	}
	if whileKeyPrefix == nil {
		panic("whileKeyPrefix cannot be nil")
	}
	whilePrefixLen := len(whileKeyPrefix)
	var result []KVPair
	count := 0
	resFunc := func(i btree.Item) bool {
		wrapper := i.(*kvWrapper)
		if bytes.Compare(wrapper.key[0:whilePrefixLen], whileKeyPrefix) > 0 {
			return false
		}
		result = append(result, KVPair{
			Key:   wrapper.key,
			Value: wrapper.value,
		})
		count++
		return limit == -1 || count < limit
	}
	f.btree.AscendGreaterOrEqual(&kvWrapper{key: startKeyPrefix}, resFunc)
	return result, nil
}

func (f *FakeStorage) AddShard(shardID uint64, callback ShardCallback) error {
	panic("implement me")
}

func (f *FakeStorage) RemoveShard(shardID uint64) error {
	panic("implement me")
}

func (f *FakeStorage) GetClusterInfo() (*ClusterInfo, error) {
	return f.clusterInfo, nil
}

func (f *FakeStorage) GetNodeInfo(nodeID int) (*NodeInfo, error) {
	nodeInfo, ok := f.clusterInfo.NodeInfos[nodeID]
	if !ok {
		return nil, fmt.Errorf("Invalid node id %d", nodeID)
	}
	return nodeInfo, nil
}

type kvWrapper struct {
	key   []byte
	value []byte
}

func (k kvWrapper) Less(than btree.Item) bool {
	otherKVwrapper := than.(*kvWrapper)

	thisKey := k.key
	otherKey := otherKVwrapper.key

	return bytes.Compare(thisKey, otherKey) < 0
}

func (f *FakeStorage) putInternal(item *kvWrapper) {
	f.btree.ReplaceOrInsert(item)
}

func (f *FakeStorage) deleteInternal(item *kvWrapper) error {
	prevItem := f.btree.Delete(item)
	if prevItem == nil {
		return errors.New("didn't find item to delete")
	}
	return nil
}

func (f *FakeStorage) getInternal(key *kvWrapper) []byte {
	item := f.btree.Get(key)
	if item != nil {
		wrapper := item.(*kvWrapper)
		return wrapper.value
	} else {
		return nil
	}
}

func createClusterInfo(nodeID int, numShards int) *ClusterInfo {
	leaders := make([]uint64, numShards)
	for i := 0; i < numShards; i++ {
		leaders[i] = uint64(i)
	}
	nodeInfo := &NodeInfo{
		Leaders:   leaders,
		Followers: nil,
	}
	nodeInfos := make(map[int]*NodeInfo)
	nodeInfos[nodeID] = nodeInfo
	return &ClusterInfo{NodeInfos: nodeInfos}
}
