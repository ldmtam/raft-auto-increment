package database

// AutoIncrement represents the AutoIncrement database interface
type AutoIncrement interface {
	GetSingle(key string) (uint64, error)
	GetMultiple(key string, quantity uint64) ([]uint64, error)

	Close() error
}
