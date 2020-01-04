package store

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ldmtam/raft-auto-increment/common"

	"github.com/ldmtam/raft-auto-increment/database/boltdb"

	"github.com/ldmtam/raft-auto-increment/config"

	"github.com/hashicorp/raft"
	"github.com/ldmtam/raft-auto-increment/database"
)

type fsm struct {
	db     database.AutoIncrement
	config *config.Config
}

type fsmSnapshot struct {
	data []byte
}

func newFSM(db database.AutoIncrement) *fsm {
	return &fsm{db: db}
}

func (f *fsm) Apply(l *raft.Log) interface{} {
	var cmd command
	if err := json.Unmarshal(l.Data, &cmd); err != nil {
		return &fsmErrorResponse{err: fmt.Errorf("failed to unmarshal command: %s", l.Data)}
	}

	switch cmd.Type {
	case getOneCmd:
		var p getOnePayload
		if err := json.Unmarshal(cmd.Payload, &p); err != nil {
			return &fsmErrorResponse{err: fmt.Errorf("failed to unmarshal getOnePayload: %v", cmd.Payload)}
		}
		value, err := f.db.GetOne(p.Key)
		return &fsmGetOneResponse{
			key:   p.Key,
			value: value,
			err:   err,
		}
	case getManyCmd:
		var p getManyPayload
		if err := json.Unmarshal(cmd.Payload, &p); err != nil {
			return &fsmErrorResponse{err: fmt.Errorf("failed to unmarshal getManyPayload: %v", cmd.Payload)}
		}
		values, err := f.db.GetMany(p.Key, p.Quantity)
		return &fsmGetManyResponse{
			key:    p.Key,
			values: values,
			err:    err,
		}
	case getLastInsertedCmd:
		var p getLastInsertedPayload
		if err := json.Unmarshal(cmd.Payload, &p); err != nil {
			return &fsmErrorResponse{err: fmt.Errorf("failed to unmarshal getLastInsertedPayload: %v", cmd.Payload)}
		}
		value, err := f.db.GetLastInserted(p.Key)
		return &fsmGetLastInsertedResponse{
			key:   p.Key,
			value: value,
			err:   err,
		}
	default:
		return &fsmErrorResponse{err: fmt.Errorf("unknown command: %v", cmd)}
	}
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	var err error
	snapshot := &fsmSnapshot{}

	snapshot.data, err = f.db.Backup()
	if err != nil {
		return nil, err
	}

	return snapshot, nil
}

func (f *fsm) Restore(rc io.ReadCloser) error {
	if err := f.removeOldData(); err != nil {
		return err
	}

	sizeBytes := make([]byte, 8)
	if _, err := io.ReadFull(rc, sizeBytes); err != nil {
		return err
	}
	size := common.ByteToUint64(sizeBytes)

	database := make([]byte, size)
	if _, err := io.ReadFull(rc, database); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(f.config.DataDir, config.DB_FILE_NAME), database, 0777); err != nil {
		return err
	}

	db, err := boltdb.New(filepath.Join(f.config.DataDir, config.DB_FILE_NAME))
	if err != nil {
		return err
	}

	f.db = db

	return nil
}

func (snapshot *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	if err := func() error {
		size := uint64(len(snapshot.data))

		// write size of database first
		if _, err := sink.Write(common.Uint64ToByte(size)); err != nil {
			return err
		}

		// then, write the actual data.
		if _, err := sink.Write(snapshot.data); err != nil {
			return err
		}
		return nil
	}; err != nil {
		return sink.Cancel()
	}

	return sink.Close()
}

func (snapshot *fsmSnapshot) Release() {}

func (f *fsm) removeOldData() error {
	// Close the boltDB
	if err := f.db.Close(); err != nil {
		return err
	}

	// Remove physical boltDB file on disk
	return os.Remove(filepath.Join(f.config.DataDir, config.DB_FILE_NAME))
}
