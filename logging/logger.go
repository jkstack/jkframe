package logging

import (
	"fmt"
	"log"
	"time"
)

type logger interface {
	setLevel(Level)
	currentLevel() Level
	rotate()
	printf(string, ...interface{}) string
	write(string)
	flush()
	files() []string
}

// DefaultLogger default logger by log package
var DefaultLogger Logger = Logger{
	logger:    &dummyLogger{},
	lastCheck: time.Now(),
}

type dummyLogger struct {
	level Level
}

func (l *dummyLogger) setLevel(level Level) {
	l.level = level
}
func (l dummyLogger) currentLevel() Level {
	return l.level
}
func (l dummyLogger) rotate() {}
func (l dummyLogger) printf(format string, a ...interface{}) string {
	log.Printf(format, a...)
	return fmt.Sprintf(format, a...)
}
func (l dummyLogger) write(str string) {
	log.Print(str)
}
func (l dummyLogger) flush() {}
func (l dummyLogger) files() []string {
	return nil
}
