package utils

import (
	"fmt"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogHelper struct {
	logger *zap.SugaredLogger
}

func NewLogHandler(logPath string, logLevel zapcore.Level) *LogHelper {

	logWriterSyncer := getLogWriter(fmt.Sprintf("%s/access.log", logPath))
	errorWriterSyncer := getLogWriter(fmt.Sprintf("%s/error.log", logPath))

	encoder := getEncoder()

	// consoleDebugging := zapcore.Lock(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, logWriterSyncer, logLevel),
		zapcore.NewCore(encoder, errorWriterSyncer, zapcore.ErrorLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	appLogger := logger.Sugar()
	defer appLogger.Sync()

	return &LogHelper{
		logger: appLogger,
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (l *LogHelper) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogHelper) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *LogHelper) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogHelper) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *LogHelper) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogHelper) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *LogHelper) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogHelper) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *LogHelper) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *LogHelper) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *LogHelper) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *LogHelper) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *LogHelper) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LogHelper) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
