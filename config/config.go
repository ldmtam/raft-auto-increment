package config

// Config represents the config struct for AutoIncrement service
type Config struct {
	RaftID    string
	RaftAddr  string
	JoinAddr  string
	RaftDir   string
	Bootstrap bool

	NodeAddr string

	Storage string
}
