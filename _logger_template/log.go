package log

import (
	"io"
	"os"
	"fmt"
	"strconv"
)

type LogLevel int8

// Logging Levels
const (
	NULL  LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var LevelStr = map[LogLevel]([]byte){
	NULL: []byte("NULL"),
	TRACE: []byte("TRACE"),
	DEBUG: []byte("DEBUG"),
	INFO: []byte("INFO"),
	WARN: []byte("WARN"),
	ERROR: []byte("ERROR"),
	FATAL: []byte("FATAL"),
}

// Default Writer
var (
	DefaultWriter io.Writer = io.Discard
	DefaultLevel  LogLevel  = INFO
)

// Package Init Method
func init() {
	for i := 0; i <= 0x7e; i++ {
		noEscapeTable[i] = i >= 0x20 && i != '\\' && i != '"'
	}
	collectEnvVar()
}

func collectEnvVar() {
	%s
}

// Api Method
%s

// Output Method
func Trace(msg string) {
	if DefaultLevel <= TRACE {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	}
}

func Debug(msg string) {
	if DefaultLevel <= DEBUG {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

func Info(msg string) {
	if DefaultLevel <= INFO {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	}
}

func Warn(msg string) {
	if DefaultLevel <= WARN {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	}
}

func Error(msg string) {
	if DefaultLevel <= ERROR {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	}
}

func Fatal(msg string) {
	if DefaultLevel <= FATAL {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Formatted Output Method
func Tracef(msg string, args... interface{}) {
	if DefaultLevel <= TRACE {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	}
}

func Debugf(msg string, args... interface{}) {
	if DefaultLevel <= DEBUG {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

func Infof(msg string, args... interface{}) {
	if DefaultLevel <= INFO {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	}
}

func Warnf(msg string, args... interface{}) {
	if DefaultLevel <= WARN {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	}
}

func Errorf(msg string, args... interface{}) {
	if DefaultLevel <= ERROR {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	}
}

func Fatalf(msg string, args... interface{}) {
	if DefaultLevel <= FATAL {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Dummy Function that Make Imported Packages Useful
func _unused_log() string {
	return os.Getenv("UNUSED") + strconv.Itoa(123)
}
