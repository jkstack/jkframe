package logging

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jkstack/jkframe/utils"
)

func init() {
	log.SetOutput(os.Stdout)
	rand.Seed(time.Now().UnixNano())
}

// Debug print debug log
func Debug(fmt string, a ...interface{}) {
	DefaultLogger.Debug(fmt, a...)
}

// Info print info log
func Info(fmt string, a ...interface{}) {
	DefaultLogger.Info(fmt, a...)
}

// Error print error log
func Error(fmt string, a ...interface{}) {
	DefaultLogger.Error(fmt, a...)
}

// Warning print warning log
func Warning(fmt string, a ...interface{}) {
	DefaultLogger.Warning(fmt, a...)
}

// Printf print log by format
func Printf(fmt string, a ...interface{}) {
	DefaultLogger.Printf(fmt, a...)
}

// Println print log from values
func Println(v ...interface{}) {
	DefaultLogger.Printf(fmt.Sprintln(v...))
}

// Flush flush log
func Flush() {
	DefaultLogger.flush()
}

// Logger logger interface
type Logger struct {
	logger
	lastCheck time.Time
}

func (l *Logger) rateLimit() bool {
	if time.Since(l.lastCheck).Seconds() <= 1 {
		if rand.Intn(100) > 0 {
			return true
		}
	}
	return false
}

func (l *Logger) resetLastCheck() {
	l.lastCheck = time.Now()
}

// Debug print debug log
func (l *Logger) Debug(fmt string, a ...interface{}) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	if rand.Intn(1000) < 1 {
		l.logger.printf("[DEBUG]"+fmt, a...)
	}
}

// Info print info log
func (l *Logger) Info(fmt string, a ...interface{}) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	l.logger.printf("[INFO]"+fmt, a...)
}

// Error print error log
func (l *Logger) Error(fmt string, a ...interface{}) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	trace := strings.Join(utils.Trace("  + "), separator)
	l.logger.printf("[ERROR]"+fmt+separator+trace, a...)
}

// Warning print warning log
func (l *Logger) Warning(fmt string, a ...interface{}) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	l.logger.printf("[WARN]"+fmt, a...)
}

// Printf print log with format
func (l *Logger) Printf(fmt string, a ...interface{}) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	l.logger.printf(fmt, a...)
}

// Write write log
func (l *Logger) Write(data []byte) (int, error) {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	str := string(data)
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	l.logger.write(str)
	return len(data), nil
}

// Flush flush log
func (l *Logger) Flush() {
	l.logger.flush()
}
