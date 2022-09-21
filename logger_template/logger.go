package template

var TLogger string = `package log

import (
	"io"
	"fmt"
	"time"
)

type Logger struct {
	// Event with pre-set/default values
	Data *Event
	// Writer for the logger
	Writer io.Writer
	// Logging level for the logger
	Level logLevel
}

// Create logger without any preset fields (except for defaults or from envs, etc.)
//
// Inputs:
//  - writer: writer that fulfills io.Writer interface, pass nil if using default writer
//  - level: logging level for the new logger, pass NULL (constant defined in this package) if using default level
func NewLogger(writer io.Writer, level logLevel) *Logger {
	logger := &Logger{}
	logger.Data = getEvent()
	if writer != nil {
		logger.Writer = writer
	} else {
		logger.Writer = defaultWriter
	}
	if level != NULL {
		logger.Level = level
	} else {
		logger.Level = defaultLevel
	}
	return logger
}

// Create logger with preset fields
//
// Inputs:
//  - writer: writer that fulfills io.Writer interface, pass nil if using default writer
//  - level: logging level for the new logger, pass NULL (constant defined in this package) if using default level
func (e *Event) Logger(writer io.Writer, level logLevel) *Logger {
	logger := &Logger{}
	logger.Data = e
	if writer != nil {
		logger.Writer = writer
	} else {
		logger.Writer = defaultWriter
	}
	if level != NULL {
		logger.Level = level
	} else {
		logger.Level = defaultLevel
	}
	return logger
}

func (logger *Logger) getLoggerEvent() *Event {
	e := eventPool.Get().(*Event)
	e.__levelThres = logger.Level
	e.__writer = logger.Writer

	// Copy Values to New Event
	%s

	return e
}

// Logger Api Methods

%s

// Logger Output Methods

// log plain string at TRACE level
func (logger *Logger) Trace(msg string) {
	if logger.Level <= TRACE {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	}
}

// log plain string at DEBUG level
func (logger *Logger) Debug(msg string) {
	if logger.Level <= DEBUG {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

// log plain string at INFO level
func (logger *Logger) Info(msg string) {
	if logger.Level <= INFO {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	}
}

// log plain string at WARN level
func (logger *Logger) Warn(msg string) {
	if logger.Level <= WARN {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	}
}

// log plain string at ERROR level
func (logger *Logger) Error(msg string) {
	if logger.Level <= ERROR {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	}
}

// log plain string at FATAL level
func (logger *Logger) Fatal(msg string) {
	if logger.Level <= FATAL {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Logger Format Output Methods

// log format string with arguments at TRACE level
func (logger *Logger) Tracef(msg string, args ...interface{}) {
	if logger.Level <= TRACE {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	}
}

// log format string with arguments at DEBUG level
func (logger *Logger) Debugf(msg string, args ...interface{}) {
	if logger.Level <= DEBUG {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

// log format string with arguments at INFO level
func (logger *Logger) Infof(msg string, args ...interface{}) {
	if logger.Level <= INFO {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	}
}

// log format string with arguments at WARN level
func (logger *Logger) Warnf(msg string, args ...interface{}) {
	if logger.Level <= WARN {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	}
}

// log format string with arguments at ERROR level
func (logger *Logger) Errorf(msg string, args ...interface{}) {
	if logger.Level <= ERROR {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	}
}

// log format string with arguments at FATAL level
func (logger *Logger) Fatalf(msg string, args ...interface{}) {
	if logger.Level <= FATAL {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Dummy Function that Make Imported Packages Useful
func _unused_logger() string {
	return time.Now().Format(time.RFC3339Nano)
}
`
