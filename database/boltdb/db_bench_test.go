package boltdb

import "testing"

var (
	getOneResult, getManyFrom, getManyTo uint64
)

func BenchmarkGetOne(b *testing.B) {
	var result uint64
	for i := 0; i < b.N; i++ {
		result, _ = testDB.GetOne("foo")
	}
	getOneResult = result
}

func BenchmarkGetMany(b *testing.B) {
	var from, to uint64
	for i := 0; i < b.N; i++ {
		from, to, _ = testDB.GetMany("bar", 1000)
	}
	getManyFrom, getManyTo = from, to
}
