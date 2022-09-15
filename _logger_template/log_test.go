package log

import (
	"io"
	"testing"
	"bytes"
	"strings"
	"os"
)

// Api Method Tests
%s

// Api Output Method Tests
func TestApiTrace(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Trace("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = TRACE
		Trace("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebug(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Debug("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = TRACE
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfo(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Info("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarn(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Warn("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiError(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = FATAL
		Error("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatal(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = 10
		Fatal("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = FATAL
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Api Format Output Method Tests
func TestApiTracef(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Tracef("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = TRACE
		Tracef("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebugf(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Debugf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = TRACE
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfof(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Infof("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = DEBUG
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarnf(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Warnf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = INFO
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiErrorf(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = FATAL
		Errorf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = WARN
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatalf(t *testing.T) {
	var writer bytes.Buffer
	DefaultWriter = &writer
	defer func(){
		DefaultWriter = io.Discard
		DefaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = 10
		Fatalf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = FATAL
		Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		DefaultLevel = ERROR
		Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Collect From Env Test
func TestEnvVar(t *testing.T) {
	%s
}
