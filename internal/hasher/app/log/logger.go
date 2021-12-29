package log

import (
	"common/log"
)

type Logger interface {
	log.Logger
}

type Details = log.Details

func NoDetails() Details {
	return nil
}

type Level = log.Level

const (
	LevelError   = log.LevelError
	LevelWarning = log.LevelWarning
	LevelInfo    = log.LevelInfo
	LevelDebug   = log.LevelDebug
)
