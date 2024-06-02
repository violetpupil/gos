// Fast JSON encoder/decoder compatible with encoding/json for Go
package json

import "github.com/goccy/go-json"

var (
	Marshal = json.Marshal
)

func MarshalString(v any) string {
	r, _ := json.Marshal(v)
	return string(r)
}
