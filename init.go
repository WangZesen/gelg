package main

import (
	"fmt"
)

var (
	initMethod       = ""
	loggerInitMethod = ""
)

func createStringInitMethod(ctx map[string]interface{}, prefix, root string) {
	out := ""
	if val, ok := ctx[defaultField]; ok {
		out += fmt.Sprintf("e.%s = appendString(e.%s[:0], \"%s\")\n", prefix, prefix, val.(string))
	} else if _, ok := ctx[fromEnvField]; ok {
		out += fmt.Sprintf("e.%s = append(e.%s[:0], env.%s...)\n", prefix, prefix, prefix)
		createStringEnvMethod(ctx, prefix, root)
		createStringEnvTest(ctx, prefix, root)
	} else {
		out += fmt.Sprintf("e.%s = e.%s[:0]\n", prefix, prefix)
	}
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = false\n", addChangeFlagPrefix(prefix))
	}
	initMethod += out
}

func createStringLoggerInitMethod(ctx map[string]interface{}, prefix, root string) {
	loggerInitMethod += fmt.Sprintf("e.%s = append(e.%s[:0], logger.Data.%s...)\n", prefix, prefix, prefix)
	if needChangeFlag(ctx) {
		loggerInitMethod += fmt.Sprintf("e.%s = logger.Data.%s\n", addChangeFlagPrefix(prefix), addChangeFlagPrefix(prefix))
	}
}

func createIntInitMethod(ctx map[string]interface{}, prefix, root string) {
	out := ""
	if val, ok := ctx[defaultField]; ok {
		out += fmt.Sprintf("e.%s = %d\n", prefix, int(val.(float64)))
	} else if _, ok := ctx[fromEnvField]; ok {
		out += fmt.Sprintf("e.%s = env.%s\n", prefix, prefix)
		createIntEnvMethod(ctx, prefix, root)
		createIntEnvTest(ctx, prefix, root)
	} else {
		out += fmt.Sprintf("e.%s = 0\n", prefix)
	}
	if needChangeFlag(ctx) {
		out += fmt.Sprintf("e.%s = false\n", addChangeFlagPrefix(prefix))
	}
	initMethod += out
}

func createIntLoggerInitMethod(ctx map[string]interface{}, prefix, root string) {
	loggerInitMethod += fmt.Sprintf("e.%s = logger.Data.%s\n", prefix, prefix)
	if needChangeFlag(ctx) {
		loggerInitMethod += fmt.Sprintf("e.%s = logger.Data.%s\n", addChangeFlagPrefix(prefix), addChangeFlagPrefix(prefix))
	}
}
