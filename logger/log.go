package logger

import log "github.com/sirupsen/logrus"

// Fields ...
type Fields map[string]interface{}

// Trace ...
func Trace(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Trace(data)
}

// TraceWithFields ...
func TraceWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Trace(data)
}

// Debug ...
func Debug(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Debug(data)
}

// DebugWithFields ...
func DebugWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Debug(data)
}

// Info ...
func Info(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Info(data)
}

// InfoWithFields ...
func InfoWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Info(data)
}

// Warn ...
func Warn(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Warn(data)
}

// WarnWithFields ...
func WarnWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Warn(data)
}

// Error ...
func Error(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Error(data)
}

// ErrorWithFields ...
func ErrorWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Debug(data)
}

// Fatal ...
func Fatal(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Fatal(data)
}

// FatalWithFields ...
func FatalWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Fatal(data)
}

// Panic ...
func Panic(context string, data string) {
	logger := log.WithField("Context", context)
	logger.Panic(data)
}

// PanicWithFields ...
func PanicWithFields(context string, data string, fields Fields) {
	logger := log.WithField("Context", context)
	logger.WithFields(log.Fields(fields)).Panic(data)
}
