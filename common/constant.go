package common

import "errors"

const (
	BADGER_STORAGE = "badger"
	BOLT_STORAGE   = "bolt"
)

var (
	ErrStorageNotAvailable = errors.New("storage is not available")
)
