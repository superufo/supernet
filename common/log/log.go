package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ModOptions func(options *Options)

// Options 日志文件配置选项
type Options struct {
	Path            string        // 文件保存地方
	Prefix          string        // 日志文件前缀
	ErrorFileSuffix string        // error日志文件后缀
	WarnFileSuffix  string        // warn日志文件后缀
	InfoFileSuffix  string        // info日志文件后缀
	DebugFileSuffix string        // debug日志文件后缀
	Level           zapcore.Level // 日志等级
	MaxSize         int           // 日志文件大小（M）
	MaxBackups      int           // 最多存在多少个切片文件
	MaxAge          int           // 保存的最大天数
	Development     bool          // 是否是开发模式
	zap.Config
}

var (
	logger                         *Logger
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts        *Options `json:"opts"`
	zapConfig   zap.Config
	initialized bool
}

func NewLogger(mod ...ModOptions) *zap.Logger {
	logger = &Logger{}
	logger.Lock()
	defer logger.Unlock()
	if logger.initialized {
		logger.Info("[NewLogger] logger initEd")
		return nil
	}
	logger.Opts = &Options{
		Path:            "",
		Prefix:          "app",
		ErrorFileSuffix: "error.log",
		WarnFileSuffix:  "warn.log",
		InfoFileSuffix:  "info.log",
		DebugFileSuffix: "debug.log",
		Level:           zapcore.DebugLevel,
		MaxSize:         100,
		MaxBackups:      60,
		MaxAge:          30,
	}

	if logger.Opts.Path == "" {
		logger.Opts.Path, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		logger.Opts.Path += sp + "logs" + sp
	}
	if logger.Opts.Development {
		logger.zapConfig = zap.NewDevelopmentConfig()
		logger.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		logger.zapConfig = zap.NewProductionConfig()
		logger.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	}

	if logger.Opts.OutputPaths == nil || len(logger.Opts.OutputPaths) == 0 {
		logger.zapConfig.OutputPaths = []string{"stdout"}
	}

	if logger.Opts.ErrorOutputPaths == nil || len(logger.Opts.ErrorOutputPaths) == 0 {
		logger.zapConfig.OutputPaths = []string{"stderr"}
	}

	for _, fn := range mod {
		fn(logger.Opts)
	}
	logger.zapConfig.Level.SetLevel(logger.Opts.Level)
	logger.init()
	logger.initialized = true
	return logger.Logger
}

func (logger *Logger) init() {
	logger.setSyncs()
	var err error
	logger.Logger, err = logger.zapConfig.Build(logger.cores())
	if err != nil {
		panic(err)
	}
	defer logger.Logger.Sync()
}

func (logger *Logger) setSyncs() {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   logger.Opts.Path + sp + logger.Opts.Prefix + "-" + fN,
			MaxSize:    logger.Opts.MaxSize,
			MaxBackups: logger.Opts.MaxBackups,
			MaxAge:     logger.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}
	errWS = f(logger.Opts.ErrorFileSuffix)
	warnWS = f(logger.Opts.WarnFileSuffix)
	infoWS = f(logger.Opts.InfoFileSuffix)
	debugWS = f(logger.Opts.DebugFileSuffix)
}
func SetMaxSize(MaxSize int) ModOptions {
	return func(option *Options) {
		option.MaxSize = MaxSize
	}
}
func SetMaxBackups(MaxBackups int) ModOptions {
	return func(option *Options) {
		option.MaxBackups = MaxBackups
	}
}
func SetMaxAge(MaxAge int) ModOptions {
	return func(option *Options) {
		option.MaxAge = MaxAge
	}
}
func SetPath(Path string) ModOptions {
	return func(option *Options) {
		option.Path = Path
	}
}

func SetPrefix(Prefix string) ModOptions {
	return func(option *Options) {
		option.Prefix = Prefix
	}
}

func SetLevel(Level zapcore.Level) ModOptions {
	return func(option *Options) {
		option.Level = Level
	}
}
func SetErrorFileSuffix(ErrorFileSuffix string) ModOptions {
	return func(option *Options) {
		option.ErrorFileSuffix = ErrorFileSuffix
	}
}
func SetWarnFileSuffix(WarnFileSuffix string) ModOptions {
	return func(option *Options) {
		option.WarnFileSuffix = WarnFileSuffix
	}
}

func SetInfoFileSuffix(InfoFileSuffix string) ModOptions {
	return func(option *Options) {
		option.InfoFileSuffix = InfoFileSuffix
	}
}
func SetDebugFileSuffix(DebugFileSuffix string) ModOptions {
	return func(option *Options) {
		option.DebugFileSuffix = DebugFileSuffix
	}
}
func SetDevelopment(Development bool) ModOptions {
	return func(option *Options) {
		option.Development = Development
	}
}
func (logger *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(logger.zapConfig.EncoderConfig)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && zapcore.ErrorLevel-logger.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-logger.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-logger.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-logger.zapConfig.Level.Level() > -1
	})
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		zapcore.NewCore(fileEncoder, warnWS, warnPriority),
		zapcore.NewCore(fileEncoder, infoWS, infoPriority),
		zapcore.NewCore(fileEncoder, debugWS, debugPriority),
	}
	if logger.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
