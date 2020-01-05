package store

import (
	"encoding/json"
)

type commandType int

const (
	setIDCmd commandType = iota + 1
	setLeaderInfoCmd
)

type command struct {
	Type    commandType     `json:"type,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

func newCommand(t commandType, p interface{}) (*command, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return &command{
		Type:    t,
		Payload: b,
	}, nil
}

type setIDPayload struct {
	Key   string `json:"key,omitempty"`
	Value uint64 `json:"value,omitempty"`
}

type setLeaderInfoPayload struct {
	NodeAddr string `json:"nodeAddr,omitempty"`
	RaftAddr string `json:"raftAddr,omitempty"`
	RaftID   string `json:"raftID,omitempty"`
}
