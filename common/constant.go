package common

import "errors"

const (
	BADGER_STORAGE = "badgerdb"
	BOLT_STORAGE   = "boltdb"
)

var (
	ErrStorageNotAvailable = errors.New("storage is not available")
)
