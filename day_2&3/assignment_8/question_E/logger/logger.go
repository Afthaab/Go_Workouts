package logger

import (
	"log"
	"os"
)

type LoggerStore struct {
	log *log.Logger
}

func New() LoggerStore {
	l := LoggerStore{
		log: log.New(os.Stdin, "afthab", log.Lshortfile),
	}
	return l
}
