package main

import "fmt"

var newMethod = ""
var outputBufSize = 0

func createStringNew(ctx map[string]interface{}, prefix, root string) {
	out := fmt.Sprintf("e.%s = make([]byte, 0, %d)\n", prefix, 2*int(ctx[maxLenField].(float64)))
	outputBufSize += 2*int(ctx[maxLenField].(float64)) + len(root) + 4
	newMethod += out
}

func createIntNew(ctx map[string]interface{}, prefix, root string) {
	// 2^64: 20 digits + 1 for sign
	outputBufSize += 21 + len(root) + 4
}

func createTimeNew(ctx map[string]interface{}, prefix, root string) {
	outputBufSize += 2*len(ctx[timeFormatField].(string)) + len(root) + 4
}

func createLevelNew(ctx map[string]interface{}, prefix, root string) {
	outputBufSize += int(ctx[maxLenField].(float64)) + len(root) + 4
}
