package logging

import (
	"log"
	"time"
)

type logger interface {
	rotate()
	printf(string, ...interface{})
	write(string)
	flush()
}

// DefaultLogger default logger by log package
var DefaultLogger Logger = Logger{
	logger:    dummyLogger{},
	lastCheck: time.Now(),
}

type dummyLogger struct{}

func (l dummyLogger) rotate() {}
func (l dummyLogger) printf(fmt string, a ...interface{}) {
	log.Printf(fmt, a...)
}
func (l dummyLogger) write(str string) {
	log.Print(str)
}
func (l dummyLogger) flush() {}
