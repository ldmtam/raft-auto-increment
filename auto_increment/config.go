package auto_increment

// Config represents the config struct for AutoIncrement service
type Config struct {
	NodeName  string
	DataDir   string
	HttpAddr  string
	GrpcAddr  string
	RaftAddr  string
	Bootstrap bool
}
