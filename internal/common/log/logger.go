package log

type Logger interface {
	LogInfo(message string, details Details)
	LogError(message string, details Details)
	LogWarn(message string, details Details)
	LogDebug(message string, details Details)
}

type LevelProvider interface {
	Level() Level
}

type Printer interface {
	Printf(message string, details ...interface{})
}

type Details map[string]interface{}
