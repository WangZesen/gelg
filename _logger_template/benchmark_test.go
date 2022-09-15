package log

import (
	"testing"
)

func BenchmarkEmptyLogBelowLogThreshold(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Trace("test_message")
	}
}

func BenchmarkEmptyLogWithFormatBelowLogThreshold(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Tracef("This is %%s message %%d", "test", 1)
	}
}

func BenchmarkAllFieldsLogBelowLogThreshold(t *testing.B) {
	for i := 0; i < t.N; i++ {
		%s.Trace("test_message")
	}
}

func BenchmarkAllFieldsLoggerBelowLogThreshold(t *testing.B) {
	logger := %s.Logger(nil, INFO)
	for i := 0; i < t.N; i++ {
		logger.Trace("test_message")
	}
}

func BenchmarkAllFieldsLoggerWithFormatBelowLogThreshold(t *testing.B) {
	logger := %s.Logger(nil, INFO)
	for i := 0; i < t.N; i++ {
		logger.Tracef("This is %%s message %%d", "test", 1)
	}
}

func BenchmarkEmptyLog(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Info("test_message")
	}
}

func BenchmarkAllFieldsLog(t *testing.B) {
	for i := 0; i < t.N; i++ {
		%s.Info("test_message")
	}
}

func BenchmarkAllFieldsLogWithFormat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		%s.Infof("This is %%s message %%d", "test", 321)
	}
}

func BenchmarkAllFieldsLogger(t *testing.B) {
	logger := %s.Logger(nil, INFO)
	for i := 0; i < t.N; i++ {
		logger.Info("test_message")
	}
}

func BenchmarkAllFieldsLoggerWithFormat(t *testing.B) {
	logger := %s.Logger(nil, INFO)
	for i := 0; i < t.N; i++ {
		logger.Infof("This is %%s message %%d", "test", 321)
	}
}
