package utilities

import (
	"go.uber.org/zap"
)

// Lightweight logger.
var Logger, _ = zap.NewDevelopment()

// Ergonomic API for logging from Logger.
var Sugar = Logger.Sugar()
