package main

import "fmt"

const (
	eventDefinitionStart = ""
	eventDefinitionEnd = "\n"
)

var eventDefinition = ""

func createStringDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("%s []byte\n", UpperFirst(root))
}

func createIntDefinition(ctx map[string]interface{}, prefix, root string) {
	eventDefinition += fmt.Sprintf("%s int64\n", UpperFirst(root))
}
