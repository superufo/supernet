package kit

import (
	"go.uber.org/zap"

	"github.com/supernet/common/log"
)

var (
	kitLogger *zap.Logger
)

func SetKitLogger(logger *zap.Logger) {
	if logger == nil {
		logger = log.NewLogger(
			log.SetPath("/log"),
			log.SetPrefix("test"),
			log.SetDebugFileSuffix("debug.log"),
			log.SetWarnFileSuffix("warn.log"),
			log.SetErrorFileSuffix("error.log"),
			log.SetInfoFileSuffix("info.log"),
			log.SetMaxAge(2),
			log.SetMaxBackups(30),
			log.SetMaxSize(10),
			log.SetDevelopment(true),
			log.SetLevel(zap.DebugLevel),
		)
	}

	kitLogger = logger
}

func GetKitLogger() *zap.Logger {
	if kitLogger == nil {
		kitLogger = log.NewLogger(
			log.SetPath("/log"),
			log.SetPrefix("test"),
			log.SetDebugFileSuffix("debug.log"),
			log.SetWarnFileSuffix("warn.log"),
			log.SetErrorFileSuffix("error.log"),
			log.SetInfoFileSuffix("info.log"),
			log.SetMaxAge(2),
			log.SetMaxBackups(30),
			log.SetMaxSize(10),
			log.SetDevelopment(true),
			log.SetLevel(zap.DebugLevel),
		)
	}

	return kitLogger
}
