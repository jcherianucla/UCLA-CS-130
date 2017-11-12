package utilities

import (
	"go.uber.org/zap"
)

var Logger = zap.NewDevelopment().Sugar()
