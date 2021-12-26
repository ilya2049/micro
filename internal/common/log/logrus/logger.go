package logrus

import (
	"common/log"
	"os"

	formatter "github.com/fabienm/go-logrus-formatters"
	graylog "github.com/gemnasium/logrus-graylog-hook"
	"github.com/sirupsen/logrus"
)

type Config struct {
	GraylogHost string
	ServiceHost string
}

func NewLogger(cfg Config) *Logger {
	logrusLogger := logrus.New()

	gelfFormatter := formatter.NewGelf(cfg.ServiceHost)

	logrusLogger.SetFormatter(gelfFormatter)
	logrusLogger.SetLevel(logrus.DebugLevel)

	graylogHook := graylog.NewGraylogHook(cfg.GraylogHost, map[string]interface{}{})
	logrusLogger.AddHook(graylogHook)

	return &Logger{
		logrusLogger: logrusLogger,
	}
}

type Logger struct {
	logrusLogger *logrus.Logger
}

func (lg *Logger) LogInfo(message string, details log.Details) {
	lg.write(logrus.InfoLevel, message, details)
}

func (lg *Logger) LogError(message string, details log.Details) {
	lg.write(logrus.ErrorLevel, message, details)
}

func (lg *Logger) LogWarn(message string, details log.Details) {
	lg.write(logrus.WarnLevel, message, details)
}

func (lg *Logger) LogDebug(message string, details log.Details) {
	lg.write(logrus.DebugLevel, message, details)
}

func (lg *Logger) LogFatal(message string, details log.Details) {
	lg.write(logrus.FatalLevel, message, details)

	os.Exit(1)
}

func (lg *Logger) Printf(message string, details ...interface{}) {
	args := make([]interface{}, 0, len(details)+1)

	args = append(args, message)
	args = append(args, details...)

	lg.logrusLogger.Log(logrus.InfoLevel, args...)
}

func (lg *Logger) write(level logrus.Level, message string, details log.Details) {
	if details == nil {
		lg.logrusLogger.Log(level, message)

		return
	}

	lg.logrusLogger.WithFields(logrus.Fields(details)).Log(level, message)
}

func (lg *Logger) Level() log.Level {
	switch lg.logrusLogger.Level {
	case logrus.FatalLevel:
		return log.LevelFatal

	case logrus.ErrorLevel:
		return log.LevelError

	case logrus.WarnLevel:
		return log.LevelWarning

	case logrus.InfoLevel:
		return log.LevelInfo

	case logrus.DebugLevel:
		return log.LevelDebug

	default:
		return log.LevelInfo
	}
}
