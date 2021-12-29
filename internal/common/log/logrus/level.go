package logrus

import (
	"common/log"

	"github.com/sirupsen/logrus"
)

func logLevelToLogrusLevel(level log.Level) logrus.Level {
	switch level {
	case log.LevelFatal:
		return logrus.FatalLevel

	case log.LevelError:
		return logrus.ErrorLevel

	case log.LevelWarning:
		return logrus.WarnLevel

	case log.LevelInfo:
		return logrus.InfoLevel

	case log.LevelDebug:
		return logrus.DebugLevel

	default:
		return logrus.InfoLevel
	}
}
