package store

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ldmtam/raft-auto-increment/database/memdb"

	"github.com/ldmtam/raft-auto-increment/common"

	"github.com/hashicorp/raft"
	"github.com/ldmtam/raft-auto-increment/database"
)

type fsm struct {
	db    database.AutoIncrement
	store *Store
}

type fsmSnapshot struct {
	data []byte
}

func newFSM(db database.AutoIncrement, store *Store) *fsm {
	return &fsm{
		db:    db,
		store: store,
	}
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
		lastInserted, err := f.db.GetLastInserted(p.Key)
		if err != nil {
			return &fsmResponse{err: err, resp: nil}
		}
		if err := f.db.Set(p.Key, lastInserted+p.Quantity); err != nil {
			return &fsmResponse{err: err, resp: nil}
		}
		var resp *fsmResponse
		if p.Quantity == 1 {
			resp = &fsmResponse{
				resp: &getOneIDResponse{
					key:   p.Key,
					value: lastInserted + p.Quantity,
				},
				err: nil,
			}
		} else {
			resp = &fsmResponse{
				resp: &getManyIDsResponse{
					key:  p.Key,
					from: lastInserted + 1,
					to:   lastInserted + p.Quantity,
				},
				err: nil,
			}
		}
		return resp
	case setLeaderInfoCmd:
		var p setLeaderInfoPayload
		if err := json.Unmarshal(cmd.Payload, &p); err != nil {
			return &fsmResponse{err: err}
		}
		f.store.leader = &leaderInfo{
			NodeAddr: p.NodeAddr,
			RaftAddr: p.RaftAddr,
			RaftID:   p.RaftID,
		}
		// close connection to old leader, if exists.
		if f.store.leaderConn != nil {
			f.store.leaderConn.Close()
			f.store.leaderConn = nil
		}
		return &fsmResponse{err: nil}
	default:
		return &fsmResponse{err: fmt.Errorf("unknown command: %v", cmd)}
	}
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	data, err := f.db.Backup()
	if err != nil {
		return nil, err
	}

	return &fsmSnapshot{data: data}, nil
}

func (f *fsm) Restore(rc io.ReadCloser) error {
	sizeBytes := make([]byte, 8)
	if _, err := io.ReadFull(rc, sizeBytes); err != nil {
		return err
	}
	size := common.ByteToUint64(sizeBytes)

	database := make([]byte, size)
	if _, err := io.ReadFull(rc, database); err != nil {
		return err
	}

	var err error

	f.db, err = memdb.New(database)
	if err != nil {
		return err
	}

	return nil
}

func (snapshot *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	var err error

	defer func() {
		if err == nil {
			sink.Close()
		} else {
			sink.Cancel()
		}
	}()

	size := uint64(len(snapshot.data))

	// write size of database first
	_, err = sink.Write(common.Uint64ToByte(size))
	if err != nil {
		return err
	}

	// then, write the actual data.
	_, err = sink.Write(snapshot.data)
	if err != nil {
		return err
	}

	return nil
}

func (snapshot *fsmSnapshot) Release() {}
