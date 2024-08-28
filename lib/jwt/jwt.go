// https://golang-jwt.github.io/jwt/
// 内部使用 base64.RawURLEncoding.EncodeToString，没有padding
package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

var key = []byte("3")

type MyClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"u"`
	TeamId string `json:"t"`
}

func GenToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		UserId: "1",
		TeamId: "2",
	})
	return token.SignedString(key)
}

func ParseToken(s string) (MyClaims, error) {
	var claims MyClaims
	_, err := jwt.ParseWithClaims(s, &claims, func(*jwt.Token) (any, error) { return key, nil })
	return claims, err
}
