package config

// Config represents the config struct for AutoIncrement service
type Config struct {
	RaftID    string
	RaftAddr  string
	JoinAddr  string
	RaftDir   string
	Bootstrap bool

	DataDir  string
	NodeAddr string
}

const (
	DB_FILE_NAME = "data.db"
)
