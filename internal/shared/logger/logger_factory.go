package logger

import (
	"fmt"
)

// LoggerType represents the type of logger
type LoggerType string

const (
	LoggerTypeStdout LoggerType = "stdout"
	LoggerTypeFile   LoggerType = "file"
	LoggerTypeStack  LoggerType = "stack" // Both stdout and file
)

// LoggerFactory creates logger instances based on configuration
type LoggerFactory struct{}

// NewLoggerFactory creates a new logger factory
func NewLoggerFactory() *LoggerFactory {
	return &LoggerFactory{}
}

// CreateLogger creates a logger based on the type and configuration
func (lf *LoggerFactory) CreateLogger(loggerType LoggerType, config interface{}) (Logger, error) {
	switch loggerType {
	case LoggerTypeStdout:
		return NewSimpleLogger(), nil

	case LoggerTypeFile:
		fileConfig, ok := config.(FileLoggerConfig)
		if !ok {
			return nil, fmt.Errorf("invalid file logger config")
		}
		return NewFileLogger(fileConfig)

	case LoggerTypeStack:
		// Create both stdout and file loggers
		fileConfig, ok := config.(FileLoggerConfig)
		if !ok {
			return nil, fmt.Errorf("invalid file logger config")
		}
		return NewStackLogger(fileConfig)

	default:
		return nil, fmt.Errorf("unsupported logger type: %s", loggerType)
	}
}

// StackLogger writes to both stdout and file
type StackLogger struct {
	stdout Logger
	file   Logger
}

// NewStackLogger creates a logger that writes to both stdout and file
func NewStackLogger(fileConfig FileLoggerConfig) (Logger, error) {
	stdoutLogger := NewSimpleLogger()
	fileLogger, err := NewFileLogger(fileConfig)
	if err != nil {
		return nil, err
	}

	return &StackLogger{
		stdout: stdoutLogger,
		file:   fileLogger,
	}, nil
}

func (sl *StackLogger) Debug(msg string, fields ...Field) {
	sl.stdout.Debug(msg, fields...)
	sl.file.Debug(msg, fields...)
}

func (sl *StackLogger) Info(msg string, fields ...Field) {
	sl.stdout.Info(msg, fields...)
	sl.file.Info(msg, fields...)
}

func (sl *StackLogger) Warn(msg string, fields ...Field) {
	sl.stdout.Warn(msg, fields...)
	sl.file.Warn(msg, fields...)
}

func (sl *StackLogger) Error(msg string, fields ...Field) {
	sl.stdout.Error(msg, fields...)
	sl.file.Error(msg, fields...)
}

func (sl *StackLogger) Fatal(msg string, fields ...Field) {
	sl.stdout.Fatal(msg, fields...)
	sl.file.Fatal(msg, fields...)
}

// Close closes the file logger
func (sl *StackLogger) Close() error {
	if closer, ok := sl.file.(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

