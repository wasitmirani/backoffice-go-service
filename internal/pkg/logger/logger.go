package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger interface for logging operations
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value interface{}
}

// SimpleLogger is a simple implementation of Logger using standard log package
type SimpleLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

// NewSimpleLogger creates a new simple logger
func NewSimpleLogger() Logger {
	return &SimpleLogger{
		debug: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
		info:  log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
		warn:  log.New(os.Stdout, "[WARN] ", log.LstdFlags|log.Lshortfile),
		error: log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile),
		fatal: log.New(os.Stderr, "[FATAL] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *SimpleLogger) Debug(msg string, fields ...Field) {
	l.log(l.debug, msg, fields...)
}

func (l *SimpleLogger) Info(msg string, fields ...Field) {
	l.log(l.info, msg, fields...)
}

func (l *SimpleLogger) Warn(msg string, fields ...Field) {
	l.log(l.warn, msg, fields...)
}

func (l *SimpleLogger) Error(msg string, fields ...Field) {
	l.log(l.error, msg, fields...)
}

func (l *SimpleLogger) Fatal(msg string, fields ...Field) {
	l.log(l.fatal, msg, fields...)
	os.Exit(1)
}

func (l *SimpleLogger) log(logger *log.Logger, msg string, fields ...Field) {
	if len(fields) > 0 {
		msg += " | "
		for i, field := range fields {
			if i > 0 {
				msg += ", "
			}
			msg += field.Key + "=" + toString(field.Value)
		}
	}
	logger.Println(msg)
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return fmt.Sprintf("%d", val)
	case int8:
		return fmt.Sprintf("%d", val)
	case int16:
		return fmt.Sprintf("%d", val)
	case int32:
		return fmt.Sprintf("%d", val)
	case int64:
		return fmt.Sprintf("%d", val)
	case uint:
		return fmt.Sprintf("%d", val)
	case uint8:
		return fmt.Sprintf("%d", val)
	case uint16:
		return fmt.Sprintf("%d", val)
	case uint32:
		return fmt.Sprintf("%d", val)
	case uint64:
		return fmt.Sprintf("%d", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	case error:
		return val.Error()
	default:
		return fmt.Sprintf("%v", v)
	}
}

