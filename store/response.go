package store

type fsmGetOneResponse struct {
	key   string
	value uint64
	err   error
}

type fsmGetManyResponse struct {
	key    string
	values []uint64
	err    error
}

type fsmGetLastInsertedResponse struct {
	key   string
	value uint64
	err   error
}

type fsmErrorResponse struct {
	err error
}
