package log

type Logger interface {
	LogInfo(message string, details Details)
	LogError(message string, details Details)
	LogWarn(message string, details Details)
	LogDebug(message string, details Details)

	Printf(message string, details ...interface{})
}

type Details map[string]interface{}

func NoDetails() Details {
	return nil
}
