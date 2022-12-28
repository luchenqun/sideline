package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EmployerKeyPrefix is the prefix to retrieve all Employer
	EmployerKeyPrefix = "Employer/value/"
)

// EmployerKey returns the store key to retrieve a Employer from the index fields
func EmployerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
