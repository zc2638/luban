package logger

import "github.com/sirupsen/logrus"

/**
 * Created by zc on 2020/4/26.
 */
type Option func(*logrus.Logger)

func SetLevel(level logrus.Level) Option {
	return func(logger *logrus.Logger) {
		logger.SetLevel(level)
	}
}

func SetFormatter(formatter logrus.Formatter) Option {
	return func(logger *logrus.Logger) {
		logger.SetFormatter(formatter)
	}
}

func AddHook(hook logrus.Hook) Option {
	return func(logger *logrus.Logger) {
		logger.AddHook(hook)
	}
}
