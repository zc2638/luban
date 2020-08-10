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
	InfoE  *logrus.Logger
	WarnE  *logrus.Logger
	ErrorE *logrus.Logger
)

func init() {
	if InfoE == nil {
		InfoE = NewEngine(SetLevel(logrus.InfoLevel))
	}
	if WarnE == nil {
		WarnE = NewEngine(SetLevel(logrus.WarnLevel))
	}
	if ErrorE == nil {
		ErrorE = NewEngine(SetLevel(logrus.ErrorLevel))
		ErrorE.AddHook(hook.NewFrameContextHook(0, nil))
	}
}

type Config struct {
	Output string `json:"output"` // 文件输出路径，不填输出终端
}

func New(config *Config) {
	InfoE = NewEngine(SetLevel(logrus.InfoLevel))
	WarnE = NewEngine(SetLevel(logrus.WarnLevel))
	ErrorE = NewEngine(SetLevel(logrus.ErrorLevel))
	if config.Output != "" {
		InfoE.AddHook(hook.NewHook("info").SetDir(config.Output))
		WarnE.AddHook(hook.NewHook("warn").SetDir(config.Output))
		ErrorE.AddHook(hook.NewHook("error").SetDir(config.Output))
	}
	ErrorE.AddHook(hook.NewFrameContextHook(30, nil))
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
	InfoE.Print(args...)
}

func Printf(format string, args ...interface{}) {
	InfoE.Printf(format, args...)
}

func Println(args ...interface{}) {
	InfoE.Println(args...)
}

func Error(args ...interface{}) {
	ErrorE.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	ErrorE.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	ErrorE.Errorln(args...)
}

func Warn(args ...interface{}) {
	WarnE.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	WarnE.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	WarnE.Warnln(args...)
}
