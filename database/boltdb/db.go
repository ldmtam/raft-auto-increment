package boltdb

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/ldmtam/raft-auto-increment/database"
	"github.com/spaolacci/murmur3"
)

const (
	BUCKET_COUNT = 32
)

var (
	bucketName = "bucket-%v"
)

type autoIncrementDB struct {
	db *bolt.DB
}

// New returns new instance of boltdb-based auto-increment database
func New(path string) (database.AutoIncrement, error) {
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}

	db, err := bolt.Open(filepath.Join(path, "data.db"), 0777, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		for i := 0; i < BUCKET_COUNT; i++ {
			if _, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf(bucketName, i))); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &autoIncrementDB{db: db}, nil
}

func (d *autoIncrementDB) GetOne(key string) (uint64, error) {
	var value uint64

	if err := d.db.Batch(func(tx *bolt.Tx) error {
		bucketName, err := d.getBucket(key)
		if err != nil {
			return err
		}

		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("%v is nil", bucketName)
		}

		oldValueBytes := bucket.Get([]byte(key))
		var oldValue uint64

		if oldValueBytes == nil {
			oldValue = 0
		} else {
			oldValue = database.ByteToUint64(oldValueBytes)
		}

		value = oldValue + 1

		return bucket.Put([]byte(key), database.Uint64ToByte(value))
	}); err != nil {
		return 0, err
	}

	return value, nil
}

func (d *autoIncrementDB) GetMany(key string, quantity uint64) ([]uint64, error) {
	values := make([]uint64, quantity)

	if err := d.db.Batch(func(tx *bolt.Tx) error {
		bucketName, err := d.getBucket(key)
		if err != nil {
			return err
		}

		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("%v is nil", bucketName)
		}

		oldValueBytes := bucket.Get([]byte(key))
		var oldValue uint64

		if oldValueBytes == nil {
			oldValue = 0
		} else {
			oldValue = database.ByteToUint64(oldValueBytes)
		}

		for i := uint64(0); i < quantity; i++ {
			values[i] = oldValue + i + 1
		}

		return bucket.Put([]byte(key), database.Uint64ToByte(oldValue+quantity))
	}); err != nil {
		return nil, err
	}

	return values, nil
}

func (d *autoIncrementDB) GetLastInserted(key string) (uint64, error) {
	var value uint64

	if err := d.db.View(func(tx *bolt.Tx) error {
		bucketName, err := d.getBucket(key)
		if err != nil {
			return err
		}

		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("%v is nil", bucketName)
		}

		oldValueBytes := bucket.Get([]byte(key))
		if oldValueBytes == nil {
			value = 0
		} else {
			value = database.ByteToUint64(oldValueBytes)
		}

		return nil
	}); err != nil {
		return 0, err
	}

	return value, nil
}

func (d *autoIncrementDB) Backup() ([]byte, error) {
	var backupBuffer bytes.Buffer
	backupWriter := bufio.NewWriter(&backupBuffer)

	if err := d.db.View(func(tx *bolt.Tx) error {
		if _, err := tx.WriteTo(backupWriter); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return backupBuffer.Bytes(), nil
}

func (d *autoIncrementDB) Close() error {
	return d.db.Close()
}

func (d *autoIncrementDB) getBucket(key string) (string, error) {
	h32 := murmur3.New32()
	if _, err := h32.Write([]byte(key)); err != nil {
		return "", err
	}
	return fmt.Sprintf(bucketName, uint(h32.Sum32())%uint(BUCKET_COUNT)), nil
}
