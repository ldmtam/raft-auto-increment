package auto_increment

// Config represents the config struct for AutoIncrement service
type Config struct {
	NodeName  string
	RaftAddr  string
	RaftDir   string
	Bootstrap bool

	DataDir  string
	HttpAddr string
	GrpcAddr string
}
