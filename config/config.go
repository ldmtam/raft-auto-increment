package config

// Config represents the config struct for AutoIncrement service
type Config struct {
	NodeID    string
	RaftAddr  string
	JoinAddr  string
	RaftDir   string
	Bootstrap bool

	DataDir string
	Addr    string
}
