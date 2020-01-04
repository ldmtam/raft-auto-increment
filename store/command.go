package store

import (
	"encoding/json"
)

type commandType int

const (
	getOneCmd commandType = iota
	getManyCmd
	getLastInsertedCmd
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

type getOnePayload struct {
	Key string `json:"key,omitempty"`
}

type getManyPayload struct {
	Key      string `json:"key,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}

type getLastInsertedPayload struct {
	Key string `json:"key,omitempty"`
}
