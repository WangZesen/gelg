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
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Trace("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = TRACE
		Trace("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebug(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Debug("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = TRACE
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfo(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Info("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarn(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
		Warn("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiError(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = FATAL
		Error("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatal(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = 10
		Fatal("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = FATAL
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Api Format Output Method Tests
func TestApiTracef(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Tracef("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = TRACE
		Tracef("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebugf(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Debugf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = TRACE
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfof(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Infof("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = DEBUG
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarnf(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
		Warnf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = INFO
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiErrorf(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = FATAL
		Errorf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = WARN
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatalf(t *testing.T) {
	var writer bytes.Buffer
	defaultWriter = &writer
	defer func(){
		defaultWriter = io.Discard
		defaultLevel = INFO
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = 10
		Fatalf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = FATAL
		Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		defaultLevel = ERROR
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
