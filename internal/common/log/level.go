package log

type Level string

const (
	LevelFatal   Level = "fatal"
	LevelError   Level = "error"
	LevelWarning Level = "warning"
	LevelInfo    Level = "info"
	LevelDebug   Level = "debug"
)
