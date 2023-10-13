package logger

import (
	"log"
	"os"
)

type LoggerStore struct {
	log *log.Logger
}

func New() *LoggerStore {
	logger := log.New(os.Stdout, "Afthab", log.Lshortfile)
	return &LoggerStore{
		log: logger,
	}
}
