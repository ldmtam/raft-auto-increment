package memdb

import (
	"bytes"
	"encoding/gob"
	"io"
	"sync"

	"github.com/spaolacci/murmur3"

	"github.com/ldmtam/raft-auto-increment/database"
)

const (
	SHARD_COUNT = 32
)

type memDB struct {
	db map[string]uint64
	mu sync.RWMutex
}

type autoIncrementDB struct {
	shards []*memDB
}

// New returns new instance of mem-based auto-increment database
func New(backup []byte) (database.AutoIncrement, error) {
	db := &autoIncrementDB{
		shards: make([]*memDB, SHARD_COUNT),
	}
	if backup == nil {
		for i := 0; i < SHARD_COUNT; i++ {
			db.shards[i] = &memDB{db: make(map[string]uint64)}
		}
	} else {
		bytesArr := bytes.Split(backup, []byte("|"))
		for i := 0; i < SHARD_COUNT; i++ {
			e := make(map[string]uint64)
			if len(bytesArr[i]) == 0 {
				db.shards[i] = &memDB{db: make(map[string]uint64)}
			} else {
				buf := bytes.NewBuffer(bytesArr[i])
				decoder := gob.NewDecoder(buf)

				if err := decoder.Decode(&e); err != nil && err != io.EOF {
					return nil, err
				}

				db.shards[i] = &memDB{db: e}
			}
		}
	}

	return db, nil
}

func (db *autoIncrementDB) GetLastInserted(key string) (uint64, error) {
	shard, err := db.getShard(key)
	if err != nil {
		return 0, err
	}

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	return shard.db[key], nil
}

func (db *autoIncrementDB) Set(key string, value uint64) error {
	shard, err := db.getShard(key)
	if err != nil {
		return err
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.db[key] = value

	return nil
}

func (db *autoIncrementDB) Backup() ([]byte, error) {
	b := new(bytes.Buffer)

	for i := 0; i < SHARD_COUNT; i++ {
		db.shards[i].mu.Lock()
		defer db.shards[i].mu.Unlock()

		buf := new(bytes.Buffer)
		encoder := gob.NewEncoder(buf)
		if err := encoder.Encode(db.shards[i].db); err != nil {
			return nil, err
		}

		if _, err := b.Write(buf.Bytes()); err != nil {
			return nil, err
		}
		if i != SHARD_COUNT-1 {
			if _, err := b.Write([]byte("|")); err != nil {
				return nil, err
			}
		}
	}

	return b.Bytes(), nil
}

func (db *autoIncrementDB) Close() error {
	return nil
}

func (db *autoIncrementDB) getShard(key string) (*memDB, error) {
	h32 := murmur3.New32()
	if _, err := h32.Write([]byte(key)); err != nil {
		return nil, err
	}
	return db.shards[h32.Sum32()%uint32(SHARD_COUNT)], nil
}
