package uuid

import "github.com/google/uuid"

type (
	// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC 4122.
	UUID = uuid.UUID
)

var (
	// NewRandom returns a Random (Version 4) UUID.
	NewRandom = uuid.NewRandom
	// NewString 获取随机 uuid 字符串，出错则 panic
	NewString = uuid.NewString
)
