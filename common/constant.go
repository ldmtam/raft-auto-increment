package common

import "errors"

const (
	BADGER_STORAGE = "badger"
	BOLT_STORAGE   = "bolt"
	MEMORY_STORAGE = "memory"
)

var (
	ErrStorageNotAvailable = errors.New("storage is not available")
)
