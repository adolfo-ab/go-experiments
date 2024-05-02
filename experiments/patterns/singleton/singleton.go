package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	logLevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

var logger *MyLogger

// Make Singleton thread-safe
var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating logger instance")
		logger = &MyLogger{logLevel: 0}
	})
	fmt.Println("Returning logger instance")
	return logger
}
