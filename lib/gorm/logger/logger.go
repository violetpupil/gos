package logger

type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)
