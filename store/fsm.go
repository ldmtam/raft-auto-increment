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
	fmt.Printf("Received log with term %v, index %v, payload %v\n", l.Term, l.Index, string(l.Data))

	var cmd command
	if err := json.Unmarshal(l.Data, &cmd); err != nil {
		return &fsmResponse{err: fmt.Errorf("failed to unmarshal command: %s", l.Data)}
	}

	switch cmd.Type {
	case setIDCmd:
		var p setIDPayload
		if err := json.Unmarshal(cmd.Payload, &p); err != nil {
			return &fsmResponse{err: fmt.Errorf("failed to unmarshal getIDPayload: %v", cmd.Payload)}
		}
		return &fsmResponse{err: f.db.Set(p.Key, p.Value)}
	default:
		return &fsmResponse{err: fmt.Errorf("unknown command: %v", cmd)}
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
