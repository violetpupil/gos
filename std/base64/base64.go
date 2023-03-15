package base64

import "encoding/base64"

// Encode 每6个bit位编码一个字节，字节数会增大三分之一
func Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
