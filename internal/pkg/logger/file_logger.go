package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// FileLoggerConfig holds configuration for file-based logging
type FileLoggerConfig struct {
	LogPath      string // Directory path for log files
	LogFileName  string // Base name for log files (e.g., "app")
	MaxSize      int    // Maximum size in megabytes before rotation
	MaxBackups   int    // Maximum number of old log files to retain
	MaxAge       int    // Maximum number of days to retain old log files
	Compress     bool   // Whether to compress rotated log files
	LocalTime    bool   // Use local time for log file names
	DailyRotate  bool   // Enable daily rotation
}

// FileLogger is a file-based logger with daily rotation support
type FileLogger struct {
	config     FileLoggerConfig
	debug      *log.Logger
	info       *log.Logger
	warn       *log.Logger
	error      *log.Logger
	fatal      *log.Logger
	currentDay int
	mu         sync.Mutex
	writer     *lumberjack.Logger
}

// NewFileLogger creates a new file-based logger with daily rotation
func NewFileLogger(config FileLoggerConfig) (Logger, error) {
	// Ensure log directory exists
	if err := os.MkdirAll(config.LogPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	fl := &FileLogger{
		config:     config,
		currentDay: time.Now().Day(),
	}

	// Initialize the log writer
	if err := fl.initWriter(); err != nil {
		return nil, err
	}

	// Create loggers for each level
	fl.debug = log.New(fl.writer, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	fl.info = log.New(fl.writer, "[INFO] ", log.LstdFlags|log.Lshortfile)
	fl.warn = log.New(fl.writer, "[WARN] ", log.LstdFlags|log.Lshortfile)
	fl.error = log.New(fl.writer, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	fl.fatal = log.New(fl.writer, "[FATAL] ", log.LstdFlags|log.Lshortfile)

	// Start daily rotation check goroutine if enabled
	if config.DailyRotate {
		go fl.startDailyRotation()
	}

	return fl, nil
}

// initWriter initializes the log file writer
func (fl *FileLogger) initWriter() error {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	// Generate log file name with date if daily rotation is enabled
	var logFileName string
	if fl.config.DailyRotate {
		dateStr := time.Now().Format("2006-01-02")
		logFileName = fmt.Sprintf("%s-%s.log", fl.config.LogFileName, dateStr)
	} else {
		logFileName = fmt.Sprintf("%s.log", fl.config.LogFileName)
	}

	logFilePath := filepath.Join(fl.config.LogPath, logFileName)

	// Create lumberjack logger for rotation
	fl.writer = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    fl.config.MaxSize,    // megabytes
		MaxBackups: fl.config.MaxBackups, // number of backups
		MaxAge:     fl.config.MaxAge,     // days
		Compress:   fl.config.Compress,
		LocalTime:  fl.config.LocalTime,
	}

	return nil
}

// startDailyRotation checks daily and rotates log files if needed
// Note: This is a background goroutine that monitors for day changes
func (fl *FileLogger) startDailyRotation() {
	ticker := time.NewTicker(1 * time.Hour) // Check every hour
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		fl.mu.Lock()
		currentDay := fl.currentDay
		fl.mu.Unlock()

		if now.Day() != currentDay {
			// Day changed, rotation will happen on next log call
			// This is handled in the log() method to avoid race conditions
		}
	}
}

// rotateDaily rotates the log file to a new daily file
func (fl *FileLogger) rotateDaily() error {
	fl.mu.Lock()
	defer fl.mu.Unlock()

	// Close current writer
	if fl.writer != nil {
		fl.writer.Close()
	}

	// Reinitialize with new date
	return fl.initWriter()
}


func (fl *FileLogger) Debug(msg string, fields ...Field) {
	fl.log(fl.debug, msg, fields...)
}

func (fl *FileLogger) Info(msg string, fields ...Field) {
	fl.log(fl.info, msg, fields...)
}

func (fl *FileLogger) Warn(msg string, fields ...Field) {
	fl.log(fl.warn, msg, fields...)
}

func (fl *FileLogger) Error(msg string, fields ...Field) {
	fl.log(fl.error, msg, fields...)
}

func (fl *FileLogger) Fatal(msg string, fields ...Field) {
	fl.log(fl.fatal, msg, fields...)
	os.Exit(1)
}

func (fl *FileLogger) log(logger *log.Logger, msg string, fields ...Field) {
	if len(fields) > 0 {
		msg += " | "
		for i, field := range fields {
			if i > 0 {
				msg += ", "
			}
			msg += field.Key + "=" + toString(field.Value)
		}
	}
	
	// Check if we need to rotate (for daily rotation)
	if fl.config.DailyRotate {
		now := time.Now()
		fl.mu.Lock()
		if now.Day() != fl.currentDay {
			fl.currentDay = now.Day()
			// Rotate to new file
			if err := fl.rotateDaily(); err == nil {
				// Recreate loggers with new writer
				fl.debug = log.New(fl.writer, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
				fl.info = log.New(fl.writer, "[INFO] ", log.LstdFlags|log.Lshortfile)
				fl.warn = log.New(fl.writer, "[WARN] ", log.LstdFlags|log.Lshortfile)
				fl.error = log.New(fl.writer, "[ERROR] ", log.LstdFlags|log.Lshortfile)
				fl.fatal = log.New(fl.writer, "[FATAL] ", log.LstdFlags|log.Lshortfile)
			}
		}
		fl.mu.Unlock()
	}
	
	logger.Println(msg)
}

// Close closes the log file
func (fl *FileLogger) Close() error {
	fl.mu.Lock()
	defer fl.mu.Unlock()
	
	if fl.writer != nil {
		return fl.writer.Close()
	}
	return nil
}

