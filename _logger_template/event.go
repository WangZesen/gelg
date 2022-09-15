package log

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
	__levelThres LogLevel
	__writer     io.Writer
	%s
}

// Event Method
%s

// Event Output Method
func (e *Event) Trace(msg string) {
	if e.getLogLevel() <= TRACE {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Debug(msg string) {
	if e.getLogLevel() <= DEBUG {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Info(msg string) {
	if e.getLogLevel() <= INFO {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Warn(msg string) {
	if e.getLogLevel() <= WARN {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Error(msg string) {
	if e.getLogLevel() <= ERROR {
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	} else {
		putEvent(e, NULL)
	}
}

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
func (e *Event) Tracef(msg string, args... interface{}) {
	if e.getLogLevel() <= TRACE {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Debugf(msg string, args... interface{}) {
	if e.getLogLevel() <= DEBUG {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Infof(msg string, args... interface{}) {
	if e.getLogLevel() <= INFO {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Warnf(msg string, args... interface{}) {
	if e.getLogLevel() <= WARN {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	} else {
		putEvent(e, NULL)
	}
}

func (e *Event) Errorf(msg string, args... interface{}) {
	if e.getLogLevel() <= ERROR {
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	} else {
		putEvent(e, NULL)
	}
}

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
func (e *Event) getLogLevel() LogLevel {
	return e.__levelThres
}

// Dummy Function that Make Imported Packages Useful
func _unused_event() string {
	_, file, _, _ := runtime.Caller(0)
	return file + strconv.Itoa(123)
}
