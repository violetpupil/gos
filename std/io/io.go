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
)
