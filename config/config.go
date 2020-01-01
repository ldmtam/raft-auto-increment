package config

// Config represents the config struct for AutoIncrement service
type Config struct {
	NodeID    string
	RaftAddr  string
	RaftDir   string
	Bootstrap bool

	DataDir  string
	HTTPAddr string
	GRPCAddr string
}
