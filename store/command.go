package store

import (
	"encoding/json"
)

type commandType int

const (
	getSingleCmd commandType = iota
	getMultipleCmd
	getLastCmd
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

type getSinglePayLoad struct {
	Key string `json:"key,omitempty"`
}

type getMultiplePayload struct {
	Key      string `json:"key,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}

type getLastPayload struct {
	Key string `json:"key,omitempty"`
}
