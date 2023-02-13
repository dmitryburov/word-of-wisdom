package zap

import (
	"errors"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dmitryburov/word-of-wisdom/config"
)

type client struct {
	cfg    config.LoggerConfig
	logger *zap.SugaredLogger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// NewZapLogger instance logger
func NewZapLogger(cfg config.LoggerConfig) *client {
	return &client{cfg: cfg}
}

func (l *client) getLoggerLevel(lv string) zapcore.Level {
	level, exist := loggerLevelMap[lv]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger init logger
func (l *client) InitLogger(name string) {
	logLevel := l.getLoggerLevel(l.cfg.Level)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// write syncers
	//stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)

	l.logger = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			stderrSyncer,
			zap.NewAtomicLevelAt(logLevel)),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.LevelOf(zap.ErrorLevel)),
		zap.AddCallerSkip(2)).
		Sugar().
		Named(name)

	if err := l.logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		l.logger.Error(err)
	}
}

func (l *client) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *client) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *client) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *client) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *client) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *client) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *client) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *client) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *client) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *client) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *client) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *client) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *client) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *client) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
