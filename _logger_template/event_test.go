package log

import (
	"io"
	"os"
	"strings"
	"testing"
	"bytes"
)

// Event Method Tests
%s

// Event Output Method Tests
func TestEventTrace(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Trace("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = TRACE
		e.__writer = &writer
		e.Trace("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventDebug(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Debug("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = TRACE
		e.__writer = &writer
		e.Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventInfo(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Info("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventWarn(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Warn("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventError(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = FATAL
		e.__writer = &writer
		e.Error("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventFatal(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = 10
		e.__writer = &writer
		e.Fatal("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = FATAL
		e.__writer = &writer
		e.Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Event Format Output Method Tests
func TestEventTracef(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Tracef("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = TRACE
		e.__writer = &writer
		e.Tracef("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventDebugf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Debugf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = TRACE
		e.__writer = &writer
		e.Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventInfof(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Infof("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = DEBUG
		e.__writer = &writer
		e.Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventWarnf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Warnf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = INFO
		e.__writer = &writer
		e.Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventErrorf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = FATAL
		e.__writer = &writer
		e.Errorf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = WARN
		e.__writer = &writer
		e.Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestEventFatalf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = 10
		e.__writer = &writer
		e.Fatalf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = FATAL
		e.__writer = &writer
		e.Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		e := getEvent()
		e.__levelThres = ERROR
		e.__writer = &writer
		e.Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}
