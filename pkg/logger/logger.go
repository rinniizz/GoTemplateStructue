package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init initializes the logger
func Init(level string, format string) {
	log = logrus.New()

	// Set log level
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// Set log format
	if format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	log.SetOutput(os.Stdout)
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if log == nil {
		Init("info", "text")
	}
	return log
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Info logs an info message
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Fatal logs a fatal message and exits
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// WithFields creates a logger with fields
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return GetLogger().WithFields(fields)
}
