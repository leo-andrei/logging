package log

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	ERROR
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "ERROR"}[l]
}

type LogEntry struct {
	Timestamp  time.Time
	Level      LogLevel
	Message    string
	Attributes map[string]interface{}
}

func NewLogEntry(level LogLevel, message string, attributes map[string]interface{}) LogEntry {
	return LogEntry{
		Timestamp:  time.Now(),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}
}

func (entry LogEntry) String() string {
	return fmt.Sprintf("%s [%s] %s - %v", entry.Timestamp.Format(time.RFC3339), entry.Level, entry.Message, entry.Attributes)
}
