package store

type getOneResponse struct {
	key   string
	value uint64
	err   error
}

type getManyResponse struct {
	key    string
	values []uint64
	err    error
}

type getLastInsertedResponse struct {
	key   string
	value uint64
	err   error
}
