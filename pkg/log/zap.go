package log

import (
	"fmt"
	"future-admin/pkg/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

var (
	log    Logger
	logger Logger
	zl     *zap.Logger
)

// loadConf load logger config
func loadConf() *Config {
	return &Config{
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "",
		Level:             viper.GetString("log.level"),
		Name:              viper.GetString("name"),
	}
}

func init() {
	cfg := loadConf()

	// new zap logger
	zl, err := newZapLogger(cfg)
	if err != nil {
		_ = fmt.Errorf("init newZapLogger err: %v", err)
	}
	_ = zl

	// log 用于支持模块级的方法调用，所以要比其他 Logger 多跳一层
	log, err = newLoggerWithCallerSkip(cfg, 1)
	if err != nil {
		_ = fmt.Errorf("init newLogger err: %v", err)
	}

	logger, err = newLogger(cfg)
	if err != nil {
		_ = fmt.Errorf("init logger err: %v", err)
	}
}

const (
	// WriterConsole console输出
	WriterConsole = "console"
)

const defaultSkip = 1 // zapLogger 包装了一层 zap.Logger，默认要跳过一层

var (
	hostname string
	logDir   string
)

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Prevent data race from occurring during zap.AddStacktrace
var zapStacktraceMutex sync.Mutex

func getLoggerLevel(cfg *Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// zapLogger logger struct
type zapLogger struct {
	sugarLogger *zap.SugaredLogger
}

// newZapLogger new zap logger
func newZapLogger(cfg *Config, opts ...Option) (*zap.Logger, error) {
	for _, opt := range opts {
		opt(cfg)
	}
	return buildLogger(cfg, defaultSkip), nil
}

// newLoggerWithCallerSkip new logger with caller skip
func newLoggerWithCallerSkip(cfg *Config, skip int, opts ...Option) (Logger, error) {
	for _, opt := range opts {
		opt(cfg)
	}
	return &zapLogger{sugarLogger: buildLogger(cfg, defaultSkip+skip).Sugar()}, nil
}

// newLogger new logger
func newLogger(cfg *Config, opts ...Option) (Logger, error) {
	for _, opt := range opts {
		opt(cfg)
	}
	return newLoggerWithCallerSkip(cfg, 0)
}

func buildLogger(cfg *Config, skip int) *zap.Logger {
	logDir = cfg.LoggerDir
	if strings.HasSuffix(logDir, "/") {
		logDir = strings.TrimRight(logDir, "/")
	}

	var encoderCfg zapcore.EncoderConfig
	if cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == WriterConsole {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	var cores []zapcore.Core
	var options []zap.Option
	// init option
	hostname, _ = os.Hostname()
	option := zap.Fields(
		zap.String("ip", utils.GetLocalIP()),
		//zap.String("app_id", cfg.Name),
		//zap.String("instance_id", hostname),
	)
	options = append(options, option)

	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	if !cfg.DisableCaller {
		caller := zap.AddCaller()
		options = append(options, caller)
	}

	// 跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(skip)
	options = append(options, addCallerSkip)

	// 构造日志
	return zap.New(combinedCore, options...)
}

// Debug logger
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Info logger
func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Warn logger
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Error logger
func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugarLogger.With(f...)
	return &zapLogger{newLogger}
}
