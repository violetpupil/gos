package proto

import "google.golang.org/protobuf/proto"

// 转为指针
var (
	Bool    = proto.Bool
	Int32   = proto.Int32
	Int64   = proto.Int64
	Float32 = proto.Float32
	Float64 = proto.Float64
	Uint32  = proto.Uint32
	Uint64  = proto.Uint64
	String  = proto.String
)
