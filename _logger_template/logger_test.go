package log

import (
	"testing"
	"bytes"
	"strings"
	"reflect"
)

// Logger Method Tests
%s

// Logger Output Method Tests
func TestLoggerTrace(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Trace("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, TRACE)
		logger.Trace("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerDebug(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Debug("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, TRACE)
		logger.Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerInfo(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Info("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerWarn(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Warn("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerError(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, FATAL)
		logger.Error("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerFatal(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, 10)
		logger.Fatal("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, FATAL)
		logger.Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Create Logger
func TestCreateLogger(t *testing.T) {
	t.Run("DefaultLogger", func(t *testing.T) {
		e := getEvent()
		logger := e.Logger(nil, NULL)
		if !reflect.DeepEqual(logger.Writer, DefaultWriter) {
			t.Errorf("Different between created logger.Writer and DefaultWriter")
		}
		if !reflect.DeepEqual(logger.Level, DefaultLevel) {
			t.Errorf("Different between created logger.Level and DefaultLevel")
		}
	})
	t.Run("CustomLogger", func(t *testing.T) {
		e := getEvent()
		var writer bytes.Buffer
		logger := e.Logger(&writer, FATAL)
		if !reflect.DeepEqual(logger.Writer, &writer) {
			t.Errorf("Different between created logger.Writer and DefaultWriter")
		}
		if !reflect.DeepEqual(logger.Level, FATAL) {
			t.Errorf("Different between created logger.Level and DefaultLevel")
		}
	})
}

// Logger Format Output Method Tests
func TestLoggerTracef(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Tracef("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, TRACE)
		logger.Tracef("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerDebugf(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Debugf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, TRACE)
		logger.Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerInfof(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Infof("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, DEBUG)
		logger.Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerWarnf(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Warnf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, INFO)
		logger.Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerErrorf(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, FATAL)
		logger.Errorf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, WARN)
		logger.Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestLoggerFatalf(t *testing.T) {
	var writer bytes.Buffer
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, 10)
		logger.Fatalf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, FATAL)
		logger.Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		logger := NewLogger(&writer, ERROR)
		logger.Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}