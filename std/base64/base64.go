package base64

import "encoding/base64"

func Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Decode(s string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
