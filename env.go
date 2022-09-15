package main

import (
	"fmt"
)

var envMethod = ""

func createStringEnvMethod(ctx map[string]interface{}, prefix, root string) {
	envMethod += fmt.Sprintf("env.%s = make([]byte, 0, %d)\n", prefix, int(ctx[maxLenField].(float64)))
	envMethod += fmt.Sprintf("env.%s = appendString(env.%s, os.Getenv(\"%s\"))\n", prefix, prefix, ctx[fromEnvField].(string))
}

func createIntEnvMethod(ctx map[string]interface{}, prefix, root string) {
	envMethod += fmt.Sprintf("if val, err := strconv.Atoi(os.Getenv(\"%s\")); err == nil {\n", ctx[fromEnvField].(string))
	envMethod += fmt.Sprintf("env.%s = int64(val)\n", prefix)
	envMethod += "}\n"
}
