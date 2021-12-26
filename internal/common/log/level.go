package log

type Level int32

const (
	LevelFatal Level = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)
