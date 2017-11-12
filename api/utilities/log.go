package utilities

import (
	"go.uber.org/zap"
)

var Logger, _ = zap.NewDevelopment()

var Sugar = Logger.Sugar()
