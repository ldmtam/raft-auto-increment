package common

import "errors"

const (
	BADGER_STORAGE  = "badger"
	BITCASK_STORAGE = "bitcask"
)

var (
	ErrStorageNotAvailable = errors.New("storage is not available")
)
