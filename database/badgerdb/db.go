package badgerdb

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ldmtam/raft-auto-increment/common"

	"github.com/dgraph-io/badger"
	"github.com/ldmtam/raft-auto-increment/database"
)

type autoIncrementDB struct {
	db   *badger.DB
	path string
}

// New returns new instance of badgerdb-based auto increment database.
func New(path string, backup io.Reader) (database.AutoIncrement, error) {
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}

	options := badger.DefaultOptions(path)
	options.Logger = nil

	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	if backup != nil {
		if err := db.Load(backup, 100); err != nil {
			return nil, err
		}
	}

	return &autoIncrementDB{
		db:   db,
		path: path,
	}, nil
}

func (d *autoIncrementDB) GetLastInserted(key string) (uint64, error) {
	var value uint64

	if err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil && err != badger.ErrKeyNotFound {
			return err
		}

		if err == badger.ErrKeyNotFound {
			value = 0
		} else {
			valueBytes, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			value = common.ByteToUint64(valueBytes)
		}

		return nil
	}); err != nil {
		return 0, err
	}

	return value, nil
}

func (d *autoIncrementDB) Set(key string, value uint64) error {
	if err := d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), common.Uint64ToByte(value))
	}); err != nil {
		return err
	}

	return nil
}

func (d *autoIncrementDB) Backup() ([]byte, error) {
	backupFilePath := filepath.Join(d.path, "badger.bak")

	f, err := os.OpenFile(backupFilePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
		os.RemoveAll(backupFilePath)
	}()

	if _, err := d.db.Backup(f, 0); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(backupFilePath)
}

func (d *autoIncrementDB) Close() error {
	return d.db.Close()
}
