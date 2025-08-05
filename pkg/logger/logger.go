package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Level int

const (
	LevelError Level = iota
	LevelInfo
	LevelDebug
)

type Logger struct {
	level   Level
	prefix  string
	logger  *log.Logger
	verbose bool
}

var globalLogger *Logger

func Init(verbose bool) {
	globalLogger = &Logger{
		level:   LevelInfo,
		prefix:  "[go-fake]",
		logger:  log.New(os.Stderr, "", 0),
		verbose: verbose,
	}
	
	if verbose {
		globalLogger.level = LevelDebug
	}
}

func (l *Logger) log(level Level, format string, args ...interface{}) {
	if level > l.level {
		return
	}

	timestamp := time.Now().Format("15:04:05")
	var levelStr string
	
	switch level {
	case LevelError:
		levelStr = "ERROR"
	case LevelInfo:
		levelStr = "INFO"
	case LevelDebug:
		levelStr = "DEBUG"
	}

	message := fmt.Sprintf(format, args...)
	l.logger.Printf("%s %s [%s] %s", timestamp, l.prefix, levelStr, message)
}

// Global logging functions
func Error(format string, args ...interface{}) {
	if globalLogger != nil {
		globalLogger.log(LevelError, format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if globalLogger != nil {
		globalLogger.log(LevelInfo, format, args...)
	}
}

func Debug(format string, args ...interface{}) {
	if globalLogger != nil {
		globalLogger.log(LevelDebug, format, args...)
	}
}

func Fatal(format string, args ...interface{}) {
	if globalLogger != nil {
		globalLogger.log(LevelError, format, args...)
	}
	os.Exit(1)
}

// Progress logging for long operations
func Progress(current, total int, operation string) {
	if globalLogger != nil && globalLogger.verbose {
		percentage := float64(current) / float64(total) * 100
		globalLogger.log(LevelDebug, "%s: %d/%d (%.1f%%)", operation, current, total, percentage)
	}
}

// Step logging for major operations
func Step(step, operation string) {
	if globalLogger != nil {
		globalLogger.log(LevelInfo, "Step %s: %s", step, operation)
	}
}

// Timing helper
func Time(operation string, fn func()) {
	if globalLogger != nil && globalLogger.verbose {
		start := time.Now()
		globalLogger.log(LevelDebug, "Starting: %s", operation)
		fn()
		duration := time.Since(start)
		globalLogger.log(LevelDebug, "Completed: %s (took %v)", operation, duration)
	} else {
		fn()
	}
}
