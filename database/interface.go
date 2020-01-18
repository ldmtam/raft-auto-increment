package database

// AutoIncrement represents the AutoIncrement database interface
type AutoIncrement interface {
	GetLastInserted(key string) (uint64, error)

	Set(key string, value uint64) error

	Backup() ([]byte, error)

	Close() error
}
