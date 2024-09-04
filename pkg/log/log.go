package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := config.Build(zap.AddCallerSkip(1))
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	log.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}
