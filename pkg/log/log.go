package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _log *zap.SugaredLogger

func init() {
	var err error
	defaultLogFilePath := "logs"
	// create path if not exists
	err = os.MkdirAll(defaultLogFilePath, 0750)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	logFile, err := os.OpenFile(defaultLogFilePath+"/i-nas-tools.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	if err != nil {
		panic("open logFile failed :" + err.Error())
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	stdoutLogLevelString := os.Getenv("LOG_LEVEL")
	stdoutLogLevel, err := zapcore.ParseLevel(stdoutLogLevelString)
	if err != nil {
		stdoutLogLevel = zapcore.InfoLevel
	}
	filterPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= stdoutLogLevel && lvl < zapcore.ErrorLevel
	})
	teecore := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(logFile), zap.DebugLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), filterPriority),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stderr), zap.ErrorLevel),
	)
	logger := zap.New(
		teecore,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	zap.ReplaceGlobals(logger)
	_log = logger.Sugar()
	zap.RedirectStdLog(logger)
}

func Debug(args ...interface{}) {
	_log.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	_log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	_log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	_log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	_log.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	_log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	_log.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	_log.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	_log.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	_log.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	_log.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	_log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	_log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	_log.Fatalf(template, args...)
}
