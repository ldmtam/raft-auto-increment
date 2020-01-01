package leveldb

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/spaolacci/murmur3"

	"github.com/ldmtam/raft-auto-increment/database"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	SHARD_COUNT = 32
)

type autoIncrementShard struct {
	db *leveldb.DB
	mu sync.RWMutex
}

type autoIncrementDB struct {
	shardedDB []*autoIncrementShard
}

// New returns new instance of auto-increment leveldb
func New(path string) (database.AutoIncrement, error) {
	shardedDB := make([]*autoIncrementShard, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		db, err := leveldb.OpenFile(filepath.Join(path, fmt.Sprintf("shard-%v", i)), nil)
		if err != nil {
			return nil, err
		}
		shardedDB[i] = &autoIncrementShard{db: db}
	}
	return &autoIncrementDB{shardedDB: shardedDB}, nil
}

func (d *autoIncrementDB) GetSingle(key string) (uint64, error) {
	shard, err := d.getShard(key)
	if err != nil {
		return 0, err
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	isExist, err := shard.db.Has([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	var value uint64
	if isExist {
		valueBytes, err := shard.db.Get([]byte(key), nil)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		value = database.ByteToUint64(valueBytes)
	}

	if err := shard.db.Put([]byte(key), database.Uint64ToByte(value+1), nil); err != nil {
		return 0, err
	}

	return value + 1, nil
}

func (d *autoIncrementDB) GetMultiple(key string, quantity uint64) ([]uint64, error) {
	shard, err := d.getShard(key)
	if err != nil {
		return nil, err
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	isExist, err := shard.db.Has([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	var value uint64
	if isExist {
		valueBytes, err := shard.db.Get([]byte(key), nil)
		if err != nil {
			return nil, err
		}
		value = database.ByteToUint64(valueBytes)
	}

	result := []uint64{}
	for i := uint64(1); i <= quantity; i++ {
		result = append(result, value+i)
	}

	if err := shard.db.Put([]byte(key), database.Uint64ToByte(value+quantity), nil); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *autoIncrementDB) GetLast(key string) (uint64, error) {
	shard, err := d.getShard(key)
	if err != nil {
		return 0, err
	}

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	isExist, err := shard.db.Has([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	if !isExist {
		return 0, nil
	}

	valueBytes, err := shard.db.Get([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	return database.ByteToUint64(valueBytes), nil
}

func (d *autoIncrementDB) Close() error {
	for i := 0; i < SHARD_COUNT; i++ {
		if err := d.shardedDB[i].db.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (d *autoIncrementDB) getShard(key string) (*autoIncrementShard, error) {
	h32 := murmur3.New32()
	if _, err := h32.Write([]byte(key)); err != nil {
		return nil, err
	}
	return d.shardedDB[uint(h32.Sum32())%uint(SHARD_COUNT)], nil
}
