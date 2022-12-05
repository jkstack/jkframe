package logging

import (
	"fmt"
	"log"
	"time"
)

type logger interface {
	rotate()
	printf(string, ...interface{}) string
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
func (l dummyLogger) printf(format string, a ...interface{}) string {
	log.Printf(format, a...)
	return fmt.Sprintf(format, a...)
}
func (l dummyLogger) write(str string) {
	log.Print(str)
}
func (l dummyLogger) flush() {}
