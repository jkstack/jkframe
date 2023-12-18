package logging

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jkstack/jkframe/utils"
)

func init() {
	log.SetOutput(os.Stdout)
	rand.Seed(time.Now().UnixNano())
}

// SetLevel set log level
func SetLevel(level Level) {
	DefaultLogger.setLevel(level)
}

// Debug print debug log
func Debug(fmt string, a ...interface{}) string {
	return DefaultLogger.Debug(fmt, a...)
}

// Info print info log
func Info(fmt string, a ...interface{}) string {
	return DefaultLogger.Info(fmt, a...)
}

// Error print error log
func Error(fmt string, a ...interface{}) string {
	return DefaultLogger.Error(fmt, a...)
}

// Warning print warning log
func Warning(fmt string, a ...interface{}) string {
	return DefaultLogger.Warning(fmt, a...)
}

// Printf print log by format
func Printf(fmt string, a ...interface{}) string {
	return DefaultLogger.Printf(fmt, a...)
}

// Println print log from values
func Println(v ...interface{}) string {
	return DefaultLogger.Printf(fmt.Sprintln(v...))
}

// Flush flush log
func Flush() {
	DefaultLogger.flush()
}

// Files get log files
func Files() []string {
	return DefaultLogger.files()
}

// Logger logger interface
type Logger struct {
	sync.RWMutex
	logger
	lastCheck time.Time
}

func (l *Logger) rateLimit() bool {
	l.RLock()
	diff := time.Since(l.lastCheck).Seconds()
	l.RUnlock()
	if diff <= 1 {
		if rand.Intn(100) > 0 {
			return true
		}
	}
	return false
}

func (l *Logger) resetLastCheck() {
	l.Lock()
	l.lastCheck = time.Now()
	l.Unlock()
}

// Debug print debug log
func (l *Logger) Debug(fmt string, a ...interface{}) string {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	if l.currentLevel() >= LevelDebug {
		return l.logger.printf("[DEBUG]"+fmt, a...)
	}
	return ""
}

// Info print info log
func (l *Logger) Info(fmt string, a ...interface{}) string {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	if l.currentLevel() >= LevelInfo {
		return l.logger.printf("[INFO]"+fmt, a...)
	}
	return ""
}

// Error print error log
func (l *Logger) Error(fmt string, a ...interface{}) string {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	if l.currentLevel() >= LevelError {
		trace := strings.Join(utils.Trace("  + "), separator)
		return l.logger.printf("[ERROR]"+fmt+separator+trace, a...)
	}
	return ""
}

// Warning print warning log
func (l *Logger) Warning(fmt string, a ...interface{}) string {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	if l.currentLevel() >= LevelWarn {
		return l.logger.printf("[WARNING]"+fmt, a...)
	}
	return ""
}

// Printf print log with format
func (l *Logger) Printf(fmt string, a ...interface{}) string {
	defer l.resetLastCheck()
	if !l.rateLimit() {
		l.logger.rotate()
	}
	return l.logger.printf(fmt, a...)
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
