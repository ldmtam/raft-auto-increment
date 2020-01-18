package store

type fsmResponse struct {
	err  error
	resp interface{}
}

type getOneIDResponse struct {
	key   string
	value uint64
}

type getManyIDsResponse struct {
	key  string
	from uint64
	to   uint64
}
