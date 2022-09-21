package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	template "github.com/WangZesen/gelg/logger_template"
)

var args = struct {
	output string
	config string
}{}

func init() {
	flag.StringVar(&(args.output), "output", "./_gen/log", "output directory (if exist, old folder will be moved to <out>.bak for backup)")
	flag.StringVar(&(args.config), "config", "./json_template/sample.json", "json directory")
}

func loadJson(config string) (map[string]interface{}, error) {
	fileb, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{})
	err = json.Unmarshal(fileb, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func recursiveGenerate(ctx map[string]interface{}, prefix, root string) {
	ctxType := getContextType(ctx)
	switch ctxType {
	case externalCtx:
		eventDefinition += fmt.Sprintf("%s struct {\n", UpperFirst(root))

		assembleFlagCount := globalAssembleFlagCounter
		globalAssembleFlagCounter += 1
		assembleMethod += fmt.Sprintf(assembleMethodSectionStart, root, assembleFlagCount)

		outputBufSize += len(root) + sectionHeaderLen

		for key, raw := range ctx {
			val := raw.(map[string]interface{})
			recursiveGenerate(val, prefix+"."+UpperFirst(key), key)
		}

		eventDefinition += "}\n"

		assembleMethod += fmt.Sprintf(assembleMethodSectionEnd, assembleFlagCount, assembleFlagCount, len(root)+sectionHeaderLen)
	case internalCtx:
		switch ctx[typeField] {
		case "string":
			createStringDefinition(ctx, prefix, root)
			createStringNew(ctx, prefix, root)
			createStringInitMethod(ctx, prefix, root)
			createStringLoggerInitMethod(ctx, prefix, root)
			createStringEventMethod(ctx, prefix, root)
			createStringApiMethod(ctx, prefix, root)
			createStringLoggerApiMethod(ctx, prefix, root)
			createStringAssembleMethod(ctx, prefix, root)
			createStringBenchmarkTest(ctx, prefix, root)
			createStringEventMethodTest(ctx, prefix, root)
			createStringLoggerMethodTest(ctx, prefix, root)
			createStringApiMethodTest(ctx, prefix, root)
			createStringRequiredTest(ctx, prefix, root)
		case "int":
			createIntDefinition(ctx, prefix, root)
			createIntNew(ctx, prefix, root)
			createIntInitMethod(ctx, prefix, root)
			createIntLoggerInitMethod(ctx, prefix, root)
			createIntEventMethod(ctx, prefix, root)
			createIntApiMethod(ctx, prefix, root)
			createIntLoggerApiMethod(ctx, prefix, root)
			createIntAssembleMethod(ctx, prefix, root)
			createIntBenchmarkTest(ctx, prefix, root)
			createIntEventMethodTest(ctx, prefix, root)
			createIntLoggerMethodTest(ctx, prefix, root)
			createIntApiMethodTest(ctx, prefix, root)
		}
	case mandatoryMessageCtx:
		createBuiltInMessageDefinition(ctx, prefix, root)
		createStringNew(ctx, prefix, root)
		createStringInitMethod(ctx, prefix, root)
		createMandatoryMessageEventMethod(ctx, prefix, root)
		createStringAssembleMethod(ctx, prefix, root)
	case mandatoryTimestampCtx:
		createTimeNew(ctx, prefix, root)
		createTimeAssembleMethod(ctx, prefix, root)
	case mandatoryCallerCtx:
		createBuiltInCallerDefinition(ctx, prefix, root)
		createStringNew(ctx, prefix, root)
		createStringInitMethod(ctx, prefix, root)
		createMandatoryCallerEventMethod(ctx, prefix, root)
		createStringAssembleMethod(ctx, prefix, root)
	case mandatoryLevelCtx:
		createLevelNew(ctx, prefix, root)
		createLevelAssembleMethod(ctx, prefix, root)
	}
}

func generate(ctx map[string]interface{}) {
	eventDefinition += eventDefinitionStart
	assembleMethod += assembleMethodStart
	for key, raw := range ctx {
		val := raw.(map[string]interface{})
		recursiveGenerate(val, UpperFirst(key), key)
	}
	eventDefinition += eventDefinitionEnd
	assembleMethod += assembleMethodEnd
}

func loadFromFile(filename string) (string, error) {
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(filebytes), nil
}

func writeToFile(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	cmd := exec.Command("gofmt", "-w", filename)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed at formatting file: %s", filename)
	}
	return nil
}

func main() {
	flag.Parse()
	ctx, err := loadJson(args.config)
	if err != nil {
		log.Fatal(err)
	}

	// Check JSON Message Definition File
	err = checkContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Generate Code
	generate(ctx)

	// Check Folder
	if _, err := os.Stat(args.output); !os.IsNotExist(err) {
		found := false
		for i := 0; i < 1000; i++ {
			alter := strings.TrimRight(args.output, "/\\") + ".bak" + strconv.Itoa(i)
			if _, err := os.Stat(alter); os.IsNotExist(err) {
				os.Rename(args.output, alter)
				log.Printf("Move %s to %s", args.output, alter)
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("Please Remove Backup Folders: %s.bak*", args.output)
		}
	}
	log.Printf("Create Directory: %s", args.output)
	os.Mkdir(args.output, os.ModePerm)

	output := fmt.Sprintf(template.TEvent, eventDefinition, eventMethod, mandatoryMessageEventMethod, mandatoryCallerEventMethod)
	err = writeToFile(path.Join(args.output, "event.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TLog, envMethod, apiMethod)
	err = writeToFile(path.Join(args.output, "log.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TPool, newMethod, outputBufSize, initMethod, requiredFieldCheck, assembleMethod)
	err = writeToFile(path.Join(args.output, "pool.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TLogger, loggerInitMethod, loggerApiMethod)
	err = writeToFile(path.Join(args.output, "logger.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TEnv, eventDefinition)
	err = writeToFile(path.Join(args.output, "env.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = template.TUtils
	err = writeToFile(path.Join(args.output, "utils.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TBenchmarkTest, benchmarkTest[1:], benchmarkTest[1:], benchmarkTest[1:], benchmarkTest[1:], benchmarkTest[1:], benchmarkTest[1:], benchmarkTest[1:])
	err = writeToFile(path.Join(args.output, "benchmark_test.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TEventTest, eventMethodTest)
	err = writeToFile(path.Join(args.output, "event_test.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TLoggerTest, loggerMethodTest)
	err = writeToFile(path.Join(args.output, "logger_test.go"), output)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf(template.TLogTest, apiMethodTest, envVarTest, requiredTest)
	err = writeToFile(path.Join(args.output, "log_test.go"), output)
	if err != nil {
		log.Fatal(err)
	}
}
