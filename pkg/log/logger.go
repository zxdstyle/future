package log

import (
	"context"
	"go.uber.org/zap"
)

type (
	Fields map[string]interface{}

	Logger interface {
		Debug(args ...interface{})
		Debugf(format string, args ...interface{})

		Info(args ...interface{})
		Infof(format string, args ...interface{})

		Warn(args ...interface{})
		Warnf(format string, args ...interface{})

		Error(args ...interface{})
		Errorf(format string, args ...interface{})

		WithFields(keyValues Fields) Logger
	}
)

// WithContext is a logger that can log msg and log span for trace
func WithContext(ctx context.Context) Logger {
	//return zap logger

	//if span := trace.SpanFromContext(ctx); span != nil {
	//	logger := spanLogger{span: span, logger: zl}
	//
	//	spanCtx := span.SpanContext()
	//	logger.spanFields = []zapcore.Field{
	//		zap.String("trace_id", spanCtx.TraceID().String()),
	//		zap.String("span_id", spanCtx.SpanID().String()),
	//	}
	//
	//	return logger
	//}
	return GetLogger()
}

// Debug logger
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logger
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logger
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logger
func Error(args ...interface{}) {
	log.Error(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// WithFields logger
// output more field, eg:
//
//	contextLogger := log.WithFields(log.Fields{"key1": "value1"})
//	contextLogger.Info("print multi field")
//
// or more sample to use:
//
//	log.WithFields(log.Fields{"key1": "value1"}).Info("this is a test log")
//	log.WithFields(log.Fields{"key1": "value1"}).Infof("this is a test log, user_id: %d", userID)
func WithFields(keyValues Fields) Logger {
	return GetLogger().WithFields(keyValues)
}

// GetLogger return a log
func GetLogger() Logger {
	return logger
}

// GetZapLogger return raw zap logger
func GetZapLogger() *zap.Logger {
	return zl
}
