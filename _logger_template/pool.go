package log

import (
	"sync"
	"time"
	"strconv"
	ilog "log"
)

var eventPool = sync.Pool{
	New: func() interface{} {
		e := &Event{}
		// Create New Event
		%s
		return e
	},
}

type outputBuf struct {
	buf []byte
}

var outputPool = sync.Pool{
	New: func() interface{} {
		return &outputBuf{
			buf: make([]byte, 0, %d),
		}
	},
}

func getEvent() *Event {
	e := eventPool.Get().(*Event)
	e.__levelThres = defaultLevel
	e.__writer = defaultWriter
	// Initialize Event
	%s

	return e
}

func checkRequiredFields(e *Event) {
	%s
}

func putEvent(e *Event, level logLevel) {
	if level > NULL {
		checkRequiredFields(e)
		assembleWrite(e, level)
	}
	eventPool.Put(e)
}

func assembleWrite(e *Event, level logLevel) {
	out := outputPool.Get().(*outputBuf)
	
	// Assemble Log Message
	%s

	e.__writer.Write(out.buf)
	outputPool.Put(out)
}

// Dummy Function that Make Imported Packages Useful
func _unused_pool() string {
	ilog.Print("123")
	return time.Now().Format(time.RFC3339Nano) + strconv.Itoa(1)
}
