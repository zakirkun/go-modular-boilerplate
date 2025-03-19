package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log levels
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

// Logger wraps zap logger
type Logger struct {
	zap    *zap.Logger
	sugar  *zap.SugaredLogger
	prefix string
	level  zapcore.Level
}

// Config holds the logger configuration
type Config struct {
	Level      string `json:"level"`
	Encoding   string `json:"encoding"`
	OutputPath string `json:"output_path"`
	MaxSize    int    `json:"max_size"`    // Maximum size in megabytes before log file rotates
	MaxBackups int    `json:"max_backups"` // Maximum number of old log files to retain
	MaxAge     int    `json:"max_age"`     // Maximum number of days to retain old log files
	Compress   bool   `json:"compress"`    // Whether to compress old log files
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		Level:      InfoLevel,
		Encoding:   "json",
		OutputPath: "logs/app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
}

// stringToZapLevel converts a string level to a zapcore level
func stringToZapLevel(level string) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// NewLogger creates a new logger with the given configuration
func NewLogger(config Config, prefix string) (*Logger, error) {
	// Create directory for logs if it doesn't exist
	logDir := filepath.Dir(config.OutputPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// Set up log rotation
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.OutputPath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	// Set up encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Determine the level
	level := stringToZapLevel(config.Level)

	// Create the core
	var core zapcore.Core
	if config.Encoding == "json" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)),
			level,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)),
			level,
		)
	}

	// Create the logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer zapLogger.Sync()

	// If prefix is provided, add it to the logger
	if prefix != "" {
		zapLogger = zapLogger.Named(prefix)
	}

	// Create the sugared logger
	sugarLogger := zapLogger.Sugar()

	// Return the logger
	return &Logger{
		zap:    zapLogger,
		sugar:  sugarLogger,
		prefix: prefix,
		level:  level,
	}, nil
}

// WithPrefix creates a new logger with the given prefix
func (l *Logger) WithPrefix(prefix string) *Logger {
	newLogger := l.zap.Named(prefix)
	return &Logger{
		zap:    newLogger,
		sugar:  newLogger.Sugar(),
		prefix: prefix,
		level:  l.level,
	}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.sugar.Debugw(msg, fields...)
}

// Info logs an info message
func (l *Logger) Info(msg string, fields ...interface{}) {
	l.sugar.Infow(msg, fields...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.sugar.Warnw(msg, fields...)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...interface{}) {
	l.sugar.Errorw(msg, fields...)
}

// Fatal logs a fatal message
func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.sugar.Fatalw(msg, fields...)
}

// Sync flushes the logger buffers
func (l *Logger) Sync() error {
	return l.zap.Sync()
}

// Global logger
var defaultLogger *Logger

// InitDefaultLogger initializes the default logger
func InitDefaultLogger(config Config) error {
	var err error
	defaultLogger, err = NewLogger(config, "")
	return err
}

// Default returns the default logger
func Default() *Logger {
	if defaultLogger == nil {
		config := DefaultConfig()
		defaultLogger, _ = NewLogger(config, "")
	}
	return defaultLogger
}

// Debug logs a debug message to the default logger
func Debug(msg string, fields ...interface{}) {
	Default().Debug(msg, fields...)
}

// Info logs an info message to the default logger
func Info(msg string, fields ...interface{}) {
	Default().Info(msg, fields...)
}

// Warn logs a warning message to the default logger
func Warn(msg string, fields ...interface{}) {
	Default().Warn(msg, fields...)
}

// Error logs an error message to the default logger
func Error(msg string, fields ...interface{}) {
	Default().Error(msg, fields...)
}

// Fatal logs a fatal message to the default logger
func Fatal(msg string, fields ...interface{}) {
	Default().Fatal(msg, fields...)
}
