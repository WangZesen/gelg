package log

import (
	"sync"
	"time"
	"strconv"
	ilog "log"
)

var EventPool = sync.Pool{
	New: func() interface{} {
		e := &Event{}
		// Create New Event
		%s
		return e
	},
}

type OutputBuf struct {
	buf []byte
}

var OutputPool = sync.Pool{
	New: func() interface{} {
		return &OutputBuf{
			buf: make([]byte, 0, %d),
		}
	},
}

func getEvent() *Event {
	e := EventPool.Get().(*Event)
	e.__levelThres = defaultLevel
	e.__writer = defaultWriter
	// Initialize Event
	%s

	return e
}

func checkRequiredFields(e *Event) {
	%s
}

func putEvent(e *Event, level LogLevel) {
	if level > NULL {
		checkRequiredFields(e)
		assembleWrite(e, level)
	}
	EventPool.Put(e)
}

func assembleWrite(e *Event, level LogLevel) {
	out := OutputPool.Get().(*OutputBuf)
	
	// Assemble Log Message
	%s

	e.__writer.Write(out.buf)
	OutputPool.Put(out)
}

// Dummy Function that Make Imported Packages Useful
func _unused_pool() string {
	ilog.Print("123")
	return time.Now().Format(time.RFC3339Nano) + strconv.Itoa(1)
}
