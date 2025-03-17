package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Logger levels
const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	OFF   = "OFF"
)

// Logger represents the logger structure
type Logger struct {
	Level  string
	Output io.Writer
	Prefix string
}

// NewLogger creates a new logger instance
func NewLogger(level, prefix string) *Logger {
	return &Logger{
		Level:  strings.ToUpper(level),
		Output: os.Stdout,
		Prefix: prefix,
	}
}

// SetOutput sets the logger output
func (l *Logger) SetOutput(w io.Writer) {
	l.Output = w
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level string) {
	l.Level = strings.ToUpper(level)
}

// SetPrefix sets the logger prefix
func (l *Logger) SetPrefix(prefix string) {
	l.Prefix = prefix
}

// shouldLog returns true if the given level should be logged
func (l *Logger) shouldLog(level string) bool {
	levels := map[string]int{
		OFF:   0,
		ERROR: 1,
		WARN:  2,
		INFO:  3,
		DEBUG: 4,
	}

	return levels[level] <= levels[l.Level]
}

// log logs a message with the given level
func (l *Logger) log(level, format string, v ...interface{}) {
	if !l.shouldLog(level) {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	prefix := ""
	if l.Prefix != "" {
		prefix = fmt.Sprintf("[%s] ", l.Prefix)
	}

	message := fmt.Sprintf(format, v...)
	logLine := fmt.Sprintf("%s %s%s: %s\n", timestamp, prefix, level, message)

	_, _ = fmt.Fprint(l.Output, logLine)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

// WithPrefix creates a new logger with the given prefix
func (l *Logger) WithPrefix(prefix string) *Logger {
	return &Logger{
		Level:  l.Level,
		Output: l.Output,
		Prefix: prefix,
	}
}

// Global logger
var defaultLogger = NewLogger(INFO, "")

// Default returns the default logger
func Default() *Logger {
	return defaultLogger
}

// SetDefaultLogger sets the default logger
func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

// SetDefaultLevel sets the default logger level
func SetDefaultLevel(level string) {
	defaultLogger.SetLevel(level)
}

// Debug logs a debug message to the default logger
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

// Info logs an info message to the default logger
func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// Warn logs a warning message to the default logger
func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

// Error logs an error message to the default logger
func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}
