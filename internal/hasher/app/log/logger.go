package log

import (
	"common/log"
)

type Logger interface {
	log.Logger
}

type Details = log.Details

const (
	LevelDebug = log.LevelDebug
)
