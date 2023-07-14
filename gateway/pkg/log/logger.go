package log

import (
	"github.com/supernet/gateway/pkg/viper"
	"go.uber.org/zap"
	"strings"
)

var (
	ZapLog *zap.Logger
)

//func init() {
//	Logger = InitLogger()
//}

func InitLogger() *zap.Logger {
	level := viper.Vp.GetString("log.level")
	logLevel := zap.DebugLevel
	if strings.EqualFold("debug", level) {
		logLevel = zap.DebugLevel
	}
	if strings.EqualFold("info", level) {
		logLevel = zap.InfoLevel
	}
	if strings.EqualFold("error", level) {
		logLevel = zap.ErrorLevel
	}
	if strings.EqualFold("warn", level) {
		logLevel = zap.WarnLevel
	}
	return NewLogger(
		SetPath(viper.Vp.GetString("log.path")),
		SetPrefix(viper.Vp.GetString("log.prefix")),
		SetDevelopment(viper.Vp.GetBool("log.development")),
		SetDebugFileSuffix(viper.Vp.GetString("log.debugFileSuffix")),
		SetWarnFileSuffix(viper.Vp.GetString("log.warnFileSuffix")),
		SetErrorFileSuffix(viper.Vp.GetString("log.errorFileSuffix")),
		SetInfoFileSuffix(viper.Vp.GetString("log.infoFileSuffix")),
		SetMaxAge(viper.Vp.GetInt("log.maxAge")),
		SetMaxBackups(viper.Vp.GetInt("log.maxBackups")),
		SetMaxSize(viper.Vp.GetInt("log.maxSize")),
		SetLevel(logLevel),
	)
}
