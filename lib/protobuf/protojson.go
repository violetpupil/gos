package protobuf

// This package produces a different output than the standard "encoding/json" package,
// which does not operate correctly on protocol buffer messages.
import "google.golang.org/protobuf/encoding/protojson"

var (
	// Format formats the message as a multiline string.
	// This function is only intended for human consumption and ignores errors.
	Format = protojson.Format
	// Marshal json编码
	Marshal = protojson.Marshal
	// Unmarshal json解码
	Unmarshal = protojson.Unmarshal
)
