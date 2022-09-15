package main

import "fmt"

var newMethod = ""
var outputBufSize = 0

func createStringNew(ctx map[string]interface{}, prefix, root string) {
	out := fmt.Sprintf("e.%s = make([]byte, 0, %d)\n", prefix, 2*int(ctx[maxLenField].(float64)))
	outputBufSize += 2 * int(ctx[maxLenField].(float64))
	newMethod += out
}

func createIntNew(ctx map[string]interface{}, prefix, root string) {
	outputBufSize += 2 * 20
}

func createTimeNew(ctx map[string]interface{}, prefix, root string) {
	outputBufSize += 2 * len(ctx[timeFormatField].(string))
}

func createLevelNew(ctx map[string]interface{}, prefix, root string) {
	outputBufSize += int(ctx[maxLenField].(float64)) + len(root) + 4
}
