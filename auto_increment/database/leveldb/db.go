package leveldb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ldmtam/raft-auto-increment/auto_increment/database"
	"github.com/syndtr/goleveldb/leveldb"
)

type autoIncrementDB struct {
	db *leveldb.DB
	mu sync.Mutex
}

// New returns new instance of auto-increment leveldb
func New(path string) (database.AutoIncrement, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &autoIncrementDB{
		db: db,
	}, nil
}

func (d *autoIncrementDB) GetSingle(key string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	isExist, err := d.db.Has([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	var value uint64
	if isExist {
		valueBytes, err := d.db.Get([]byte(key), nil)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		value = database.ByteToUint64(valueBytes)
	}

	if err := d.db.Put([]byte(key), database.Uint64ToByte(value+1), nil); err != nil {
		return 0, err
	}

	return value + 1, nil
}

func (d *autoIncrementDB) GetMultiple(key string, quantity uint64) ([]uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	isExist, err := d.db.Has([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	var value uint64
	if isExist {
		valueBytes, err := d.db.Get([]byte(key), nil)
		if err != nil {
			return nil, err
		}
		value = database.ByteToUint64(valueBytes)
	}

	result := []uint64{}
	for i := uint64(1); i <= quantity; i++ {
		result = append(result, value+i)
	}

	if err := d.db.Put([]byte(key), database.Uint64ToByte(value+quantity), nil); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *autoIncrementDB) GetLast(key string) (uint64, error) {
	isExist, err := d.db.Has([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	if !isExist {
		return 0, nil
	}

	valueBytes, err := d.db.Get([]byte(key), nil)
	if err != nil {
		return 0, err
	}

	return database.ByteToUint64(valueBytes), nil
}

func (d *autoIncrementDB) Set(key string, value uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	isExist, err := d.db.Has([]byte(key), nil)
	if err != nil {
		return err
	}

	if !isExist {
		return errors.New("key not found")
	}

	return d.db.Put([]byte(key), database.Uint64ToByte(value), nil)
}

func (d *autoIncrementDB) Close() error {
	return d.db.Close()
}
