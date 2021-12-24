package log

import (
	"common/log"
)

type Logger interface {
	log.Logger
	log.LevelProvider
	log.Printer
}

type Details = log.Details

const (
	LevelDebug = log.LevelDebug
)
