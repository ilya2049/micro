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

func NoDetails() Details {
	return nil
}

const (
	LevelDebug = log.LevelDebug
)
