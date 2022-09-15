package log

import (
	"strings"
	"testing"
	"bytes"
)

// Event Method Tests
%s

// Event Output Method Tests
func TestEventTrace(t *testing.T) {
	var writer bytes.Buffer
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
