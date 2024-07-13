package log

type Driver interface {
	Log(entry LogEntry) error
}

type Logger struct {
	driver Driver
}

func NewLogger(driver Driver) *Logger {
	return &Logger{driver: driver}
}

func (l *Logger) Log(level LogLevel, message string, attributes map[string]interface{}) {
	entry := NewLogEntry(level, message, attributes)
	l.driver.Log(entry)
}
