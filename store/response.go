package store

type getSingleResponse struct {
	key   string
	value uint64
	err   error
}

type getMultipleResponse struct {
	key    string
	values []uint64
	err    error
}

type getLastResponse struct {
	key   string
	value uint64
	err   error
}
