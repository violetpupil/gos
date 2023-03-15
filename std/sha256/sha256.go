// 摘要是32个字节
package sha256

import (
	"crypto/sha256"
	"encoding/base64"
)

// Sum256Base64 sha256 base64摘要
func Sum256Base64(s []byte) string {
	a := sha256.Sum256(s)
	return base64.StdEncoding.EncodeToString(a[:])
}
