package log

import (
	"io"
	"fmt"
	"time"
)

type Logger struct {
	Data *Event
	Writer io.Writer
	Level LogLevel
}

func NewLogger(writer io.Writer, level LogLevel) *Logger {
	logger := &Logger{}
	logger.Data = getEvent()
	if writer != nil {
		logger.Writer = writer
	} else {
		logger.Writer = DefaultWriter
	}
	if level != NULL {
		logger.Level = level
	} else {
		logger.Level = DefaultLevel
	}
	return logger
}

func (e *Event) Logger(writer io.Writer, level LogLevel) *Logger {
	logger := &Logger{}
	logger.Data = e
	if writer != nil {
		logger.Writer = writer
	} else {
		logger.Writer = DefaultWriter
	}
	if level != NULL {
		logger.Level = level
	} else {
		logger.Level = DefaultLevel
	}
	return logger
}

func (logger *Logger) getLoggerEvent() *Event {
	e := EventPool.Get().(*Event)
	e.__levelThres = logger.Level
	e.__writer = logger.Writer

	// Copy Values to New Event
	%s

	return e
}

// Logger Api Methods
%s

// Logger Output Methods
func (logger *Logger) Trace(msg string) {
	if logger.Level <= TRACE {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	}
}

func (logger *Logger) Debug(msg string) {
	if logger.Level <= DEBUG {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

func (logger *Logger) Info(msg string) {
	if logger.Level <= INFO {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	}
}

func (logger *Logger) Warn(msg string) {
	if logger.Level <= WARN {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	}
}

func (logger *Logger) Error(msg string) {
	if logger.Level <= ERROR {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	}
}

func (logger *Logger) Fatal(msg string) {
	if logger.Level <= FATAL {
		e := logger.getLoggerEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Logger Format Output Methods
func (logger *Logger) Tracef(msg string, args ...interface{}) {
	if logger.Level <= TRACE {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	}
}

func (logger *Logger) Debugf(msg string, args ...interface{}) {
	if logger.Level <= DEBUG {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

func (logger *Logger) Infof(msg string, args ...interface{}) {
	if logger.Level <= INFO {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	}
}

func (logger *Logger) Warnf(msg string, args ...interface{}) {
	if logger.Level <= WARN {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	}
}

func (logger *Logger) Errorf(msg string, args ...interface{}) {
	if logger.Level <= ERROR {
		e := logger.getLoggerEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	}
}

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
