package main

import "fmt"

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

func getLoggerInstance() *MyLogger {
	if logger == nil {
		fmt.Println("Creating logger instance")
		logger = &MyLogger{logLevel: 0}
	}
	fmt.Println("Returning logger instance")
	return logger
}
