package gorm

import "gorm.io/gorm/logger"

// LogLevel log level
type LogLevel = logger.LogLevel

const (
	Silent = logger.Silent
	Error  = logger.Error
	Warn   = logger.Warn
	Info   = logger.Info
)
