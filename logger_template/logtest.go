package template

var TLogTest string = `package log

import (
	"testing"
	"bytes"
	"strings"
	"os"
	"io"
)

// Api Method Tests
%s

// Api Output Method Tests
func TestApiTrace(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Trace("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(TRACE)
		Trace("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebug(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Debug("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(TRACE)
		Debug("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfo(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Info("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Info("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarn(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
		Warn("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Warn("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiError(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(FATAL)
		Error("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Error("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatal(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(10)
		Fatal("test")
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(FATAL)
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
		Fatal("__test__")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

// Api Format Output Method Tests
func TestApiTracef(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Tracef("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(TRACE)
		Tracef("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiDebugf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Debugf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(TRACE)
		Debugf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiInfof(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Infof("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(DEBUG)
		Infof("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiWarnf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
		Warnf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(INFO)
		Warnf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiErrorf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(FATAL)
		Errorf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(WARN)
		Errorf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})
}

func TestApiFatalf(t *testing.T) {
	var writer bytes.Buffer
	GSetDefaultWriter(&writer)
	GSetDefaultWarnWriter(io.Discard)
	defer func(){
		GSetDefaultWriter(os.Stdout)
		GSetDefaultLevel(INFO)
		GSetDefaultWarnWriter(os.Stderr)
	}()
	t.Run("LogBelowThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(10)
		Fatalf("%%s %%d", "test", 123)
		if writer.Len() > 0 {
			t.Errorf("Expect no output, Got: %%s", writer.String())
		}
	})

	t.Run("LogAtThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(FATAL)
		Fatalf("__%%s__", "test")
		if !strings.Contains(writer.String(), "__test__") {
			t.Errorf("Expect message with %%s, Got: %%s", "__test__", writer.String())
		}
	})

	t.Run("LogAboveThres", func(t *testing.T) {
		writer.Reset()
		GSetDefaultLevel(ERROR)
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

// Warning for Unset Required Fields
func TestUnsetRequiredFields(t * testing.T) {
	var writer bytes.Buffer
	GSetDefaultWarnWriter(&writer)
	GSetDefaultWriter(io.Discard)
	defer func() {
		GSetDefaultWarnWriter(os.Stderr)
		GSetDefaultWriter(os.Stdout)
	}()
	t.Run("EmptyLog", func(t *testing.T) {
		writer.Reset()
		Info("__")
		warn := writer.String()
		%s
	})
}
`
