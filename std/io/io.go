package io

import "io"

var (
	// EOF is the error returned by Read when no more input is available.
	// Functions should return EOF only to signal a graceful end of input.
	EOF = io.EOF
	// ErrClosedPipe is the error used for read or write operations on a closed pipe.
	ErrClosedPipe = io.ErrClosedPipe
	// EOF occurs unexpectedly
	ErrUnexpectedEOF = io.ErrUnexpectedEOF
)

var (
	// 向writer里写字符串
	WriteString = io.WriteString
	// 读取 io.Reader
	// 正常读取结束，不返回 io.EOF
	ReadAll = io.ReadAll
	// 将 io.Reader 转为 io.ReadCloser
	// Close() 方法直接返回
	NopCloser = io.NopCloser
	// 写入多个 writer
	MultiWriter = io.MultiWriter
)

type (
	// Write writes len(p) bytes from p to the underlying data stream.
	Writer = io.Writer
)
