/*
Package for the generated logger

General Usage:
  - Insert the code into your projects
  - import by `log "<package-name>/<rel-path-to-logger-in-proj>"`

There are several components:
  - Event
  - Logger

Event:
  - after setting a value by log.SetXxx, it will return an event
  - event has methods to continue setting values using (*Event).SetXxx (defined by apiAlias)
  - event has methods to log using (*Event).Info, (*Event).Infof, (*Event).Warn, ... (log in format/plain string in 6 levels)
  - event has method to create logger using (*Event).Logger, it will create logger which uses the values in the Event as default value

Logger:
  - logger is initiated by (*Event).Logger or log.NewLogger
  - logger has methods to continue setting values using (*Event).SetXxx (defined by apiAlias), which will return an Event
  - event has methods to log using (*Event).Info, (*Event).Infof, (*Event).Warn, ... (log in format/plain string in 6 levels)
*/
package log

import (
	"io"
	"os"
	"fmt"
	"strconv"

	ilog "log"
)

type logLevel int8

// Logging Levels: 6 levels in total.
// Trace, Debug, Info, Warn, Error, Fatal
const (
	NULL  logLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelStr = map[logLevel]([]byte){
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
	defaultWriter io.Writer     = os.Stdout
	defaultWarnWriter io.Writer = os.Stderr
	defaultLevel logLevel       = INFO
	warnLogger *ilog.Logger     = ilog.New(os.Stderr, "", ilog.Ldate | ilog.Ltime)
)

// Package Init Method
func init() {
	for i := 0; i <= 0x7e; i++ {
		noEscapeTable[i] = i >= 0x20 && i != '\\' && i != '"'
	}
	collectEnvVar()
	ilog.SetOutput(defaultWarnWriter)
}

// set the default writer for logging.
// default is os.Stdout
func GSetDefaultWriter(writer io.Writer) {
	defaultWriter = writer
}

// set the default writer for warning of logging (unset required fields, etc.).
// default is os.Stderr
func GSetDefaultWarnWriter(writer io.Writer) {
	defaultWarnWriter = writer
	warnLogger = ilog.New(writer, "", ilog.Ldate | ilog.Ltime)
}

// set the default level for logging.
// default is INFO
func GSetDefaultLevel(level logLevel) {
	defaultLevel = level
}

func collectEnvVar() {
	%s
}

// Api Method

%s

// Output Method

// log plain string at TRACE level
func Trace(msg string) {
	if defaultLevel <= TRACE {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, TRACE)
	}
}

// log plain string at DEBUG level
func Debug(msg string) {
	if defaultLevel <= DEBUG {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

// log plain string at INFO level
func Info(msg string) {
	if defaultLevel <= INFO {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, INFO)
	}
}

// log plain string at WARN level
func Warn(msg string) {
	if defaultLevel <= WARN {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, WARN)
	}
}

// log plain string at ERROR level
func Error(msg string) {
	if defaultLevel <= ERROR {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, ERROR)
	}
}

// log plain string at FATAL level
func Fatal(msg string) {
	if defaultLevel <= FATAL {
		e := getEvent()
		e.setMessage(msg)
		e.setCaller()
		putEvent(e, FATAL)
	}
}

// Formatted Output Method

// log format string with arguments at TRACE level
func Tracef(msg string, args... interface{}) {
	if defaultLevel <= TRACE {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, TRACE)
	}
}

// log format string with arguments at DEBUG level
func Debugf(msg string, args... interface{}) {
	if defaultLevel <= DEBUG {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, DEBUG)
	}
}

// log format string with arguments at INFO level
func Infof(msg string, args... interface{}) {
	if defaultLevel <= INFO {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, INFO)
	}
}

// log format string with arguments at WARN level
func Warnf(msg string, args... interface{}) {
	if defaultLevel <= WARN {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, WARN)
	}
}

// log format string with arguments at ERROR level
func Errorf(msg string, args... interface{}) {
	if defaultLevel <= ERROR {
		e := getEvent()
		e.setMessage(fmt.Sprintf(msg, args...))
		e.setCaller()
		putEvent(e, ERROR)
	}
}

// log format string with arguments at FATAL level
func Fatalf(msg string, args... interface{}) {
	if defaultLevel <= FATAL {
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
