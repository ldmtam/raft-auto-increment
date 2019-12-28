package database

import (
	"encoding/binary"
)

// Uint64ToByte converts uint64 to byte array
func Uint64ToByte(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)
	return b
}

// ByteToUint64 converts byte array to uint64
func ByteToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}
