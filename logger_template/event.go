package template

var TEvent string = `package log

import (
	"io"
	"fmt"
	"runtime"
	"strconv"
)

const stackSkip = 2
var extraStackSkip = 0

// Event Definition
type Event struct {
	__levelThres logLevel
	__writer     io.Writer
	%s
}

// Event Method

%s

// Event Output Method

// log plain string at TRACE level
func (e *Event) Trace(msg string) {
	if e.getLogLevel() <= TRACE {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	} else {
		putEvent(e, NULL)
	}
}

// log plain string at DEBUG level
func (e *Event) Debug(msg string) {
	if e.getLogLevel() <= DEBUG {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	} else {
		putEvent(e, NULL)
	}
}

// log plain string at INFO level
func (e *Event) Info(msg string) {
	if e.getLogLevel() <= INFO {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	} else {
		putEvent(e, NULL)
	}
}

// log plain string at WARN level
func (e *Event) Warn(msg string) {
	if e.getLogLevel() <= WARN {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	} else {
		putEvent(e, NULL)
	}
}

// log plain string at ERROR level
func (e *Event) Error(msg string) {
	if e.getLogLevel() <= ERROR {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	} else {
		putEvent(e, NULL)
	}
}

// log plain string at FATAL level
func (e *Event) Fatal(msg string) {
	if e.getLogLevel() <= FATAL {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, FATAL)
	} else {
		putEvent(e, NULL)
	}
}

// Event Format Output Method

// log format string with arguments at TRACE level
func (e *Event) Tracef(msg string, args... interface{}) {
	if e.getLogLevel() <= TRACE {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	} else {
		putEvent(e, NULL)
	}
}

// log format string with arguments at DEBUG level
func (e *Event) Debugf(msg string, args... interface{}) {
	if e.getLogLevel() <= DEBUG {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	} else {
		putEvent(e, NULL)
	}
}

// log format string with arguments at INFO level
func (e *Event) Infof(msg string, args... interface{}) {
	if e.getLogLevel() <= INFO {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	} else {
		putEvent(e, NULL)
	}
}

// log format string with arguments at WARN level
func (e *Event) Warnf(msg string, args... interface{}) {
	if e.getLogLevel() <= WARN {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	} else {
		putEvent(e, NULL)
	}
}

// log format string with arguments at ERROR level
func (e *Event) Errorf(msg string, args... interface{}) {
	if e.getLogLevel() <= ERROR {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	} else {
		putEvent(e, NULL)
	}
}

// log format string with arguments at FATAL level
func (e *Event) Fatalf(msg string, args... interface{}) {
	if e.getLogLevel() <= FATAL {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, FATAL)
	} else {
		putEvent(e, NULL)
	}
}

// Mandatory Fields
func (e *Event) setMessage(msg string) {
	%s
}

func (e *Event) setCaller() {
	%s
}

// Get LogLevel
func (e *Event) getLogLevel() logLevel {
	return e.__levelThres
}

// Dummy Function that Make Imported Packages Useful
func _unused_event() string {
	_, file, _, _ := runtime.Caller(0)
	return file + strconv.Itoa(123)
}
`
