package main

import (
	"fmt"
	"strings"
)

const (
	eventDefinitionStart = ""
	eventDefinitionEnd   = "\n"
	changeFlagPrefix     = "__Changed"
)

var eventDefinition = ""

func needChangeFlag(ctx map[string]interface{}) bool {
	_, hasDefault := ctx[defaultField]
	_, hasEnvVar := ctx[fromEnvField]
	raw, isRequired := ctx[requiredField]
	required, ok := raw.(bool)
	return (!hasDefault && !hasEnvVar && isRequired && ok && required)
}

func addChangeFlagPrefix(name string) string {
	ret := ""
	fields := strings.Split(name, ".")
	for i := 0; i < len(fields)-1; i++ {
		ret += fields[i] + "."
	}
	return ret + changeFlagPrefix + fields[len(fields)-1]
}

func createBuiltInMessageDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("// [Built-in Message]: %s\n", ctx[descriptionField].(string))
	eventDefinition += fmt.Sprintf("%s []byte\n", UpperFirst(root))
	if needChangeFlag(ctx) {
		eventDefinition += fmt.Sprintf("%s bool\n", addChangeFlagPrefix(UpperFirst(root)))
	}
}

func createBuiltInCallerDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("// [Built-in Caller]: %s\n", ctx[descriptionField].(string))
	eventDefinition += fmt.Sprintf("%s []byte\n", UpperFirst(root))
	if needChangeFlag(ctx) {
		eventDefinition += fmt.Sprintf("%s bool\n", addChangeFlagPrefix(UpperFirst(root)))
	}
}

func createStringDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("// %s\n", ctx[descriptionField].(string))
	eventDefinition += fmt.Sprintf("%s []byte\n", UpperFirst(root))
	if needChangeFlag(ctx) {
		eventDefinition += fmt.Sprintf("%s bool\n", addChangeFlagPrefix(UpperFirst(root)))
	}
}

func createIntDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("// %s\n", ctx[descriptionField].(string))
	eventDefinition += fmt.Sprintf("%s int64\n", UpperFirst(root))
	if needChangeFlag(ctx) {
		eventDefinition += fmt.Sprintf("%s bool\n", addChangeFlagPrefix(UpperFirst(root)))
	}
}
