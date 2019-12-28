package database

// AutoIncrement represents the AutoIncrement database interface
type AutoIncrement interface {
	GetSingle(key string) (uint64, error)
	GetMultiple(key string, quantity uint64) ([]uint64, error)
	GetLast(key string) (uint64, error)

	// Set `value` to a particular `key`. Only used by Raft
	Set(key string, value uint64) error

	Close() error
}
