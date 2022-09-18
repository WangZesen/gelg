package main

import (
	"fmt"
	"math/rand"
)

var (
	benchmarkTest    = ""
	eventMethodTest  = ""
	loggerMethodTest = ""
	apiMethodTest    = ""
	envVarTest       = ""
	requiredTest     = ""
	wholeTestStart   = "func Test%sSet%s(t *testing.T) {\n"
	wholeTestEnd     = "}\n"
	subTestStart     = "t.Run(\"%s\", func(t *testing.T){\n"
	subTestEnd       = "})\n"
	instantiaEvent   = "e := getEvent()\n"
	instantiaLogger  = "logger := NewLogger(%s, %s)\n"
)

func createStringBenchmarkTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	benchmarkTest += fmt.Sprintf(".Set%s(\"test\")", UpperFirst(apiAlias))
}

func createStringEventMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Event", UpperFirst(apiAlias))
	maxLen := int(ctx[maxLenField].(float64))
	// subtest: set normal string
	randStr := RandStringRunes(maxLen / 2)
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr, prefix)
	out += subTestEnd

	// subtest: set too-long string
	randStr = RandStringRunes(maxLen + 10)
	out += fmt.Sprintf(subTestStart, "SetTooLongString")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr[:maxLen])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr[:maxLen], prefix)
	out += subTestEnd

	// subtest: set UTF8 string
	utf8Str := "a你a"
	out += fmt.Sprintf(subTestStart, "SetUTF8String")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str, prefix)
	out += subTestEnd

	// subtest: set faulty UTF8 string
	utf8Str = RandStringRunes(maxLen-1) + "你"
	out += fmt.Sprintf(subTestStart, "SetFaultyUTF8String")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str[:maxLen-1])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str[:maxLen-1], prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case1)
	escapeStr := RandStringRunes(maxLen-6) + "\\\\\\b\\f\\n\\r\\t"
	expectedEscapeStr := escapeStr[:maxLen-6] + "\\\\\\\\\\\\b\\\\f\\\\n\\\\r\\\\t"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase1")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case2)
	escapeStr = RandStringRunes(maxLen-6) + "\\u0012"
	expectedEscapeStr = escapeStr[:maxLen-6] + "\\\\u0012"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase2")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	out += wholeTestEnd
	eventMethodTest += out + "\n"
}

func createStringLoggerMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Logger", UpperFirst(apiAlias))
	maxLen := int(ctx[maxLenField].(float64))
	// subtest: set normal string
	randStr := RandStringRunes(maxLen / 2)
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr, prefix)
	out += subTestEnd

	// subtest: set too-long string
	randStr = RandStringRunes(maxLen + 10)
	out += fmt.Sprintf(subTestStart, "SetTooLongString")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr[:maxLen])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr[:maxLen], prefix)
	out += subTestEnd

	// subtest: set UTF8 string
	utf8Str := "a你a"
	out += fmt.Sprintf(subTestStart, "SetUTF8String")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str, prefix)
	out += subTestEnd

	// subtest: set faulty UTF8 string
	utf8Str = RandStringRunes(maxLen-1) + "你"
	out += fmt.Sprintf(subTestStart, "SetFaultyUTF8String")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str[:maxLen-1])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str[:maxLen-1], prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case1)
	escapeStr := RandStringRunes(maxLen-6) + "\\\\\\b\\f\\n\\r\\t"
	expectedEscapeStr := escapeStr[:maxLen-6] + "\\\\\\\\\\\\b\\\\f\\\\n\\\\r\\\\t"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase1")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case2)
	escapeStr = RandStringRunes(maxLen-6) + "\\u0012"
	expectedEscapeStr = escapeStr[:maxLen-6] + "\\\\u0012"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase2")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	out += wholeTestEnd
	loggerMethodTest += out + "\n"
}

func createStringApiMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Api", UpperFirst(apiAlias))
	maxLen := int(ctx[maxLenField].(float64))
	// subtest: set normal string
	randStr := RandStringRunes(maxLen / 2)
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr, prefix)
	out += subTestEnd

	// subtest: set too-long string
	randStr = RandStringRunes(maxLen + 10)
	out += fmt.Sprintf(subTestStart, "SetTooLongString")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), randStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, randStr[:maxLen])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr[:maxLen], prefix)
	out += subTestEnd

	// subtest: set UTF8 string
	utf8Str := "a你a"
	out += fmt.Sprintf(subTestStart, "SetUTF8String")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str, prefix)
	out += subTestEnd

	// subtest: set faulty UTF8 string
	utf8Str = RandStringRunes(maxLen-1) + "你"
	out += fmt.Sprintf(subTestStart, "SetFaultyUTF8String")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), utf8Str)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, utf8Str[:maxLen-1])
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", utf8Str[:maxLen-1], prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case1)
	escapeStr := RandStringRunes(maxLen-6) + "\\\\\\b\\f\\n\\r\\t"
	expectedEscapeStr := escapeStr[:maxLen-6] + "\\\\\\\\\\\\b\\\\f\\\\n\\\\r\\\\t"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase1")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	// subtest: set string with escape chars (case2)
	escapeStr = RandStringRunes(maxLen-6) + "\\u0012"
	expectedEscapeStr = escapeStr[:maxLen-6] + "\\\\u0012"
	out += fmt.Sprintf(subTestStart, "SetStringWithEscapeCharCase2")
	out += fmt.Sprintf("e := Set%s(\"%s\")\n", UpperFirst(apiAlias), escapeStr)
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {\n", prefix, expectedEscapeStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", expectedEscapeStr, prefix)
	out += subTestEnd

	out += wholeTestEnd
	apiMethodTest += out + "\n"
}

func createStringEnvTest(ctx map[string]interface{}, prefix, root string) {
	env := ctx[fromEnvField].(string)
	randStr := RandStringRunes(int(ctx[maxLenField].(float64)) / 2)
	out := fmt.Sprintf(subTestStart, "NormalStringEnvVar")
	out += fmt.Sprintf("tmp := os.Getenv(\"%s\")\n", env)
	out += fmt.Sprintf("defer os.Setenv(\"%s\", tmp)\n", env)
	out += fmt.Sprintf("os.Setenv(\"%s\", \"%s\")\n", env, randStr)
	out += fmt.Sprintf("collectEnvVar()\n")
	out += fmt.Sprintf("e := getEvent()\n")
	out += fmt.Sprintf("if string(e.%s) != \"%s\" {", prefix, randStr)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%s, Got: %%s\", \"%s\", string(e.%s))\n}\n", randStr, prefix)
	out += subTestEnd
	envVarTest += out
}

func createStringRequiredTest(ctx map[string]interface{}, prefix, root string) {
	if needChangeFlag(ctx) {
		requiredTest += fmt.Sprintf("if !strings.Contains(warn, \"Miss Value for %s\") {\n", prefix)
		requiredTest += fmt.Sprintf("t.Errorf(\"Expect warning about unset required fields, Got Warn Msg: %%s\", warn)\n")
		requiredTest += fmt.Sprintf("}\n")
	}
}

func createIntBenchmarkTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	benchmarkTest += fmt.Sprintf(".Set%s(123)", UpperFirst(apiAlias))
}

func createIntEventMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Event", UpperFirst(apiAlias))
	// subtest: set normal int64
	normalInt := 123456
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += instantiaEvent
	out += fmt.Sprintf("e.Set%s(%d)\n", UpperFirst(apiAlias), normalInt)
	out += fmt.Sprintf("if e.%s != %d {\n", prefix, normalInt)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%d, Got: %%d\", %d, e.%s)\n}\n", normalInt, prefix)
	out += subTestEnd

	out += wholeTestEnd
	eventMethodTest += out + "\n"
}

func createIntLoggerMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Logger", UpperFirst(apiAlias))
	// subtest: set normal int64
	normalInt := 123456
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += fmt.Sprintf(instantiaLogger, "nil", "NULL")
	out += fmt.Sprintf("e := logger.Set%s(%d)\n", UpperFirst(apiAlias), normalInt)
	out += fmt.Sprintf("if e.%s != %d {\n", prefix, normalInt)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%d, Got: %%d\", %d, e.%s)\n}\n", normalInt, prefix)
	out += subTestEnd

	out += wholeTestEnd
	loggerMethodTest += out + "\n"
}

func createIntApiMethodTest(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(wholeTestStart, "Api", UpperFirst(apiAlias))
	// subtest: set normal int64
	normalInt := 123456
	out += fmt.Sprintf(subTestStart, "SetNormalString")
	out += fmt.Sprintf("e := Set%s(%d)\n", UpperFirst(apiAlias), normalInt)
	out += fmt.Sprintf("if e.%s != %d {\n", prefix, normalInt)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%d, Got: %%d\", %d, e.%s)\n}\n", normalInt, prefix)
	out += subTestEnd

	out += wholeTestEnd
	loggerMethodTest += out + "\n"
}

func createIntEnvTest(ctx map[string]interface{}, prefix, root string) {
	env := ctx[fromEnvField].(string)
	randInt := rand.Intn(10000000)
	out := fmt.Sprintf(subTestStart, "NormalIntEnvVar")
	out += fmt.Sprintf("tmp := os.Getenv(\"%s\")\n", env)
	out += fmt.Sprintf("defer os.Setenv(\"%s\", tmp)\n", env)
	out += fmt.Sprintf("os.Setenv(\"%s\", \"%d\")\n", env, randInt)
	out += fmt.Sprintf("collectEnvVar()\n")
	out += fmt.Sprintf("e := getEvent()\n")
	out += fmt.Sprintf("if e.%s != %d {", prefix, randInt)
	out += fmt.Sprintf("t.Errorf(\"Expect: %%d, Got: %%d\", %d, e.%s)\n}\n", randInt, prefix)
	out += subTestEnd
	envVarTest += out
}

func createIntRequiredTest(ctx map[string]interface{}, prefix, root string) {
	if needChangeFlag(ctx) {
		requiredTest += fmt.Sprintf("if !strings.Contains(warn, \"Miss Value for %s\") {\n", prefix)
		requiredTest += fmt.Sprintf("t.Errorf(\"Expect warning about unset required fields, Got Warn Msg: %%s\", warn)\n")
		requiredTest += fmt.Sprintf("}\n")
	}
}
