package main

import (
	"fmt"
)

const (
	assembleMethodStart        = "out.buf = append(out.buf[:0], \"{\"...)\n"
	assembleMethodEnd          = "out.buf = append(out.buf[:len(out.buf)-1], \"}\"...)\n"
	assembleMethodSectionStart = "out.buf = append(out.buf, \"\\\"%s\\\":{\"...)\ntmp%d := len(out.buf)\n"
	assembleMethodSectionEnd   = "if len(out.buf) != tmp%d {\n" +
		"out.buf = append(out.buf[:len(out.buf)-1], \"},\"...)\n" +
		"} else {\n" +
		"out.buf = out.buf[:tmp%d-%d]\n" +
		"}\n"
	sectionHeaderLen = 4
)

var (
	assembleMethod                = ""
	globalAssembleFlagCounter int = 0
)

func createStringAssembleMethod(ctx map[string]interface{}, prefix, root string) {
	out := ""
	if omit, ok := ctx[omitEmptyField]; ok && omit.(bool) {
		out += fmt.Sprintf("if len(e.%s) != 0 {\n", prefix)
	}
	out += fmt.Sprintf("out.buf = append(out.buf, \"\\\"%s\\\":\\\"\"...)\n", root)
	out += fmt.Sprintf("out.buf = append(out.buf, e.%s...)\n", prefix)
	out += "out.buf = append(out.buf, \"\\\",\"...)\n"
	if omit, ok := ctx[omitEmptyField]; ok && omit.(bool) {
		out += "}\n"
	}
	assembleMethod += out
}

func createTimeAssembleMethod(ctx map[string]interface{}, prefix, root string) {
	out := fmt.Sprintf("out.buf = append(out.buf, \"\\\"%s\\\":\\\"\"...)\n", root)
	out += fmt.Sprintf("out.buf = time.Now().AppendFormat(out.buf, \"%s\")\n", ctx[timeFormatField].(string))
	out += "out.buf = append(out.buf, \"\\\",\"...)\n"
	assembleMethod += out
}

func createIntAssembleMethod(ctx map[string]interface{}, prefix, root string) {
	out := ""
	out += fmt.Sprintf("out.buf = append(out.buf, \"\\\"%s\\\":\"...)\n", root)
	out += fmt.Sprintf("out.buf = strconv.AppendInt(out.buf, e.%s, 10)\n", prefix)
	out += "out.buf = append(out.buf, \",\"...)\n"
	assembleMethod += out
}

func createLevelAssembleMethod(ctx map[string]interface{}, prefix, root string) {
	out := ""
	out += fmt.Sprintf("out.buf = append(out.buf, \"\\\"%s\\\":\\\"\"...)\n", root)
	out += fmt.Sprintf("out.buf = append(out.buf, LevelStr[level]...)\n")
	out += "out.buf = append(out.buf, \"\\\",\"...)\n"
	assembleMethod += out
}
