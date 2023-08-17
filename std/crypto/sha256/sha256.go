// 摘要是32个字节
package sha256

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Sum256Base64 sha256 base64摘要
func Sum256Base64(s []byte) string {
	a := sha256.Sum256(s)
	return base64.StdEncoding.EncodeToString(a[:])
}

// Sum256 sha256 十六进制摘要
func Sum256(s string) string {
	a := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", a)
}
