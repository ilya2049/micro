package log

type Level int32

const (
	LevelError Level = iota
	LevelWarning
	LevelInfo
	LevelDebug
)
