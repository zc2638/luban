package logger

import (
	"github.com/sirupsen/logrus"
	"luban/pkg/logger/hook"
	"os"
	"time"
)

/**
 * Created by zc on 2020-04-16.
 */
var (
	infoE  *logrus.Logger
	warnE  *logrus.Logger
	errorE *logrus.Logger
)

func init() {
	if infoE == nil {
		infoE = NewEngine(SetLevel(logrus.InfoLevel))
	}
	if warnE == nil {
		warnE = NewEngine(SetLevel(logrus.WarnLevel))
	}
	if errorE == nil {
		errorE = NewEngine(SetLevel(logrus.ErrorLevel))
		errorE.AddHook(hook.NewFrameContextHook(0, nil))
	}
}

type Config struct {
	Output string `json:"output"` // 文件输出路径，不填输出终端
}

func New(config *Config) {
	infoE = NewEngine(SetLevel(logrus.InfoLevel))
	warnE = NewEngine(SetLevel(logrus.WarnLevel))
	errorE = NewEngine(SetLevel(logrus.ErrorLevel))
	if config.Output != "" {
		infoE.AddHook(hook.NewHook("info").SetDir(config.Output))
		warnE.AddHook(hook.NewHook("warn").SetDir(config.Output))
		errorE.AddHook(hook.NewHook("error").SetDir(config.Output))
	}
	errorE.AddHook(hook.NewFrameContextHook(30, nil))
}

func NewEngine(options ...Option) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	logger.Out = os.Stdout
	for _, option := range options {
		option(logger)
	}
	return logger
}

func Print(args ...interface{}) {
	infoE.Print(args...)
}

func Printf(format string, args ...interface{}) {
	infoE.Printf(format, args...)
}

func Println(args ...interface{}) {
	infoE.Println(args...)
}

func Error(args ...interface{}) {
	errorE.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	errorE.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	errorE.Errorln(args...)
}

func Warn(args ...interface{}) {
	warnE.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	warnE.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	warnE.Warnln(args...)
}
