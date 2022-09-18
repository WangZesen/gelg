package main

import (
	"fmt"
)

var (
	eventMethod                 = ""
	apiMethod                   = ""
	loggerApiMethod             = ""
	mandatoryMessageEventMethod = ""
	mandatoryCallerEventMethod  = ""
)

const (
	eventMethodStart     = "func (e *Event) Set%s(data %s) *Event {\n"
	eventMethodEnd       = "return e\n}\n\n"
	apiMethodStart       = "func Set%s(data %s) *Event{\ne := getEvent()\n"
	apiMethodEnd         = "return e\n}\n\n"
	loggerApiMethodStart = "func (logger *Logger) Set%s(data %s) *Event{\ne := logger.getLoggerEvent()\n"
	loggerApiMethodEnd   = "return e\n}\n\n"
)

func createStringEventMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	maxLen := int(ctx[maxLenField].(float64))
	out := fmt.Sprintf(eventMethodStart, UpperFirst(apiAlias), "string")
	out += fmt.Sprintf("if len(data) <= %d {\n", maxLen)
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data)\n", prefix, prefix)
	out += fmt.Sprintf("} else {\n")
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data[:%d])\n", prefix, prefix, maxLen)
	out += fmt.Sprintf("}\n")
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	eventMethod += out + eventMethodEnd
}

func createStringApiMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	maxLen := int(ctx[maxLenField].(float64))
	out := fmt.Sprintf(apiMethodStart, UpperFirst(apiAlias), "string")
	out += fmt.Sprintf("if len(data) <= %d {\n", maxLen)
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data)\n", prefix, prefix)
	out += fmt.Sprintf("} else {\n")
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data[:%d])\n", prefix, prefix, maxLen)
	out += fmt.Sprintf("}\n")
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	apiMethod += out + apiMethodEnd
}

func createStringLoggerApiMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	maxLen := int(ctx[maxLenField].(float64))
	out := fmt.Sprintf(loggerApiMethodStart, UpperFirst(apiAlias), "string")
	out += fmt.Sprintf("if len(data) <= %d {\n", maxLen)
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data)\n", prefix, prefix)
	out += fmt.Sprintf("} else {\n")
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], data[:%d])\n", prefix, prefix, maxLen)
	out += fmt.Sprintf("}\n")
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	loggerApiMethod += out + loggerApiMethodEnd
}

func createIntEventMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(eventMethodStart, UpperFirst(apiAlias), "int64")
	out += fmt.Sprintf("e.%s = data\n", prefix)
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	eventMethod += out + eventMethodEnd
}

func createIntApiMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(apiMethodStart, UpperFirst(apiAlias), "int64")
	out += fmt.Sprintf("e.%s = data\n", prefix)
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	apiMethod += out + apiMethodEnd
}

func createIntLoggerApiMethod(ctx map[string]interface{}, prefix, root string) {
	apiAlias := root
	if raw, ok := ctx[apiMethod]; ok {
		apiAlias = raw.(string)
	}
	out := fmt.Sprintf(loggerApiMethodStart, UpperFirst(apiAlias), "int64")
	out += fmt.Sprintf("e.%s = data\n", prefix)
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = true\n", addChangeFlagPrefix(prefix))
	}
	loggerApiMethod += out + loggerApiMethodEnd
}

func createMandatoryMessageEventMethod(ctx map[string]interface{}, prefix, root string) {
	maxLen := int(ctx[maxLenField].(float64))
	out := ""
	out += fmt.Sprintf("if len(msg) <= %d {\n", maxLen)
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], msg)\n", prefix, prefix)
	out += fmt.Sprintf("} else {\n")
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], msg[:%d])\n", prefix, prefix, maxLen)
	out += fmt.Sprintf("}")
	mandatoryMessageEventMethod += out
}

func createMandatoryCallerEventMethod(ctx map[string]interface{}, prefix, root string) {
	out := fmt.Sprintf("if _, file, line, ok := runtime.Caller(stackSkip + extraStackSkip); ok {\n")
	out += fmt.Sprintf("e.%s = appendString(e.%s[:0], file)\n", prefix, prefix)
	out += fmt.Sprintf("e.%s = append(e.%s, \":\"...)\n", prefix, prefix)
	out += fmt.Sprintf("e.%s = strconv.AppendInt(e.%s, int64(line), 10)\n", prefix, prefix)
	out += fmt.Sprintf("}")
	mandatoryCallerEventMethod += out
}
