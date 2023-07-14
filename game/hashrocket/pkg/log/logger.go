package log

import (
	"github.com/supernet/game/hashrocket/pkg/viper"
	"go.uber.org/zap"

	clog "github.com/supernet/common/log"
	"strings"
)

var (
	ZapLog *zap.Logger
)

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
	return clog.NewLogger(
		clog.SetPath(viper.Vp.GetString("log.path")),
		clog.SetPrefix(viper.Vp.GetString("log.prefix")),
		clog.SetDevelopment(viper.Vp.GetBool("log.development")),
		clog.SetDebugFileSuffix(viper.Vp.GetString("log.debugFileSuffix")),
		clog.SetWarnFileSuffix(viper.Vp.GetString("log.warnFileSuffix")),
		clog.SetErrorFileSuffix(viper.Vp.GetString("log.errorFileSuffix")),
		clog.SetInfoFileSuffix(viper.Vp.GetString("log.infoFileSuffix")),
		clog.SetMaxAge(viper.Vp.GetInt("log.maxAge")),
		clog.SetMaxBackups(viper.Vp.GetInt("log.maxBackups")),
		clog.SetMaxSize(viper.Vp.GetInt("log.maxSize")),
		clog.SetLevel(logLevel),
	)
}
