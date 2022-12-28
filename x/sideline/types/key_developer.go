package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeveloperKeyPrefix is the prefix to retrieve all Developer
	DeveloperKeyPrefix = "Developer/value/"
)

// DeveloperKey returns the store key to retrieve a Developer from the index fields
func DeveloperKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
