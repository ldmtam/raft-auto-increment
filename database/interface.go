package database

// AutoIncrement represents the AutoIncrement database interface
type AutoIncrement interface {
	GetOne(key string) (uint64, error)
	GetMany(key string, quantity uint64) (uint64, uint64, error)
	GetLastInserted(key string) (uint64, error)

	Set(key string, value uint64) error

	Backup() ([]byte, error)

	Close() error
}
