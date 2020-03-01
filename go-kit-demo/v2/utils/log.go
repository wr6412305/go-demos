package utils

import (
	"go.uber.org/zap"

	zaplog "local.com/log-zap"
)

var logger *zap.Logger

// NewLoggerServer ...
func NewLoggerServer() {
	logger = zaplog.NewLogger(
		zaplog.SetAppName("go-kit"),
		zaplog.SetDevelopment(true),
		zaplog.SetLevel(zap.DebugLevel),
	)
}

// GetLogger ...
func GetLogger() *zap.Logger {
	return logger
}
