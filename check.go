package main

import (
	"fmt"
)

const (
	typeField       = "__type"
	omitEmptyField  = "__omitEmpty"
	maxLenField     = "__maxLen"
	mandatoryField  = "__mandatory"
	fromCallerField = "__fromCaller"
	timeFormatField = "__timeFormat"
	requiredField   = "__required"
	apiAliasField   = "__apiAlias"
	defaultField    = "__default"
	fromEnvField    = "__fromEnv"
)

var internalFields = []string{
	typeField,
	omitEmptyField,
	maxLenField,
	mandatoryField,
	fromCallerField,
	timeFormatField,
	requiredField,
	apiAliasField,
	defaultField,
	fromEnvField,
}

const (
	messageManField   = "message"
	timestampManField = "timestamp"
	levelManField     = "loglevel"
	callerManField    = "caller"
)

var mandatoryFields = make(map[string]string)

type contextType int32

const (
	internalCtx contextType = iota
	mandatoryMessageCtx
	mandatoryTimestampCtx
	mandatoryLevelCtx
	mandatoryCallerCtx
	ErrorMandatoryCtx
	externalCtx
	mixedCtx
	emptyCtx
)

func isInternalContext(field string) bool {
	for i := 0; i < len(internalFields); i++ {
		if field == internalFields[i] {
			return true
		}
	}
	return false
}

func getContextType(ctx map[string]interface{}) contextType {
	if len(ctx) == 0 {
		return emptyCtx
	}
	var isInternal, isExternal = false, false
	var mandatoryCtx = emptyCtx
	for key := range ctx {
		if isInternalContext(key) {
			isInternal = true
		} else {
			isExternal = true
		}
		if key == mandatoryField {
			if v, ok := ctx[key].(string); ok {
				switch v {
				case messageManField:
					mandatoryCtx = mandatoryMessageCtx
				case timestampManField:
					mandatoryCtx = mandatoryTimestampCtx
				case levelManField:
					mandatoryCtx = mandatoryLevelCtx
				case callerManField:
					mandatoryCtx = mandatoryCallerCtx
				default:
					return ErrorMandatoryCtx
				}
			} else {
				return ErrorMandatoryCtx
			}
		}
	}
	if isInternal && isExternal {
		return mixedCtx
	} else if isInternal {
		if mandatoryCtx != emptyCtx {
			return mandatoryCtx
		}
		return internalCtx
	} else {
		return externalCtx
	}
}

func checkMandatoryMessage(ctx map[string]interface{}, prefix, root string) error {
	// __mandatory is checked before
	// check if it's unique
	if val, ok := mandatoryFields[messageManField]; ok {
		return fmt.Errorf("repeated mandatory message field: \"%s\" and \"%s\"", val, prefix)
	} else {
		mandatoryFields[messageManField] = prefix
	}

	// check if unnecessary fields are included
	for key := range ctx {
		switch key {
		case typeField, omitEmptyField, maxLenField, mandatoryField:
			continue
		default:
			return fmt.Errorf("unexpected field %s in mandatory message field: %s", key, prefix)
		}
	}

	// check __omitEmpty = false
	if raw, ok := ctx[omitEmptyField]; ok {
		if val, ok := raw.(bool); ok {
			if val {
				return fmt.Errorf("field %s has wrong value in mandatory message field: %s. Expect: %s", omitEmptyField, prefix, "false")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory message field: %s. Expect: %s", omitEmptyField, prefix, "bool")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory message field: %s", omitEmptyField, prefix)
	}

	// check __maxLen
	if raw, ok := ctx[maxLenField]; ok {
		if val, ok := raw.(float64); ok {
			if val < 10 || (val != float64(int(val))) {
				return fmt.Errorf("field %s has wrong value in mandatory message field: %s. Expect: %s", maxLenField, prefix, "int>=10")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory message field: %s. Expect: %s", maxLenField, prefix, "int")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory message field: %s", maxLenField, prefix)
	}

	// check __type
	if raw, ok := ctx[typeField]; ok {
		if val, ok := raw.(string); ok {
			if val != "string" {
				return fmt.Errorf("field %s has wrong value in mandatory message field: %s. Expect: %s", timeFormatField, prefix, "string")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory message field: %s. Expect: %s", typeField, prefix, "string")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory message field: %s", typeField, prefix)
	}

	return nil
}

func checkMandatoryTimestamp(ctx map[string]interface{}, prefix, root string) error {
	// __mandatory is checked before
	// check if it's unique
	if val, ok := mandatoryFields[timestampManField]; ok {
		return fmt.Errorf("repeated mandatory timestamp field: \"%s\" and \"%s\"", val, prefix)
	} else {
		mandatoryFields[timestampManField] = prefix
	}

	// check if unnecessary fields are included
	for key := range ctx {
		switch key {
		case typeField, fromCallerField, timeFormatField, mandatoryField:
			continue
		default:
			return fmt.Errorf("unexpected field %s in mandatory timestamp field: %s", key, prefix)
		}
	}

	// check __omitEmpty = false
	if raw, ok := ctx[fromCallerField]; ok {
		if val, ok := raw.(bool); ok {
			if !val {
				return fmt.Errorf("field %s has wrong value in mandatory timestamp field: %s. Expect: %s", fromCallerField, prefix, "true")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory timestamp field: %s. Expect: %s", fromCallerField, prefix, "bool")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory timestamp field: %s", fromCallerField, prefix)
	}

	// check __timeFormat
	if raw, ok := ctx[timeFormatField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in mandatory timestamp field: %s", timeFormatField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory timestamp field: %s. Expect: %s", timeFormatField, prefix, "string")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory timestamp field: %s", timeFormatField, prefix)
	}

	// check __type
	if raw, ok := ctx[typeField]; ok {
		if val, ok := raw.(string); ok {
			if val != "time" {
				return fmt.Errorf("field %s has wrong value in mandatory timestamp field: %s. Expect: %s", timeFormatField, prefix, "time")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory timestamp field: %s. Expect: %s", typeField, prefix, "string")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory timestamp field: %s", typeField, prefix)
	}

	return nil
}

func checkMandatoryCaller(ctx map[string]interface{}, prefix, root string) error {
	// __mandatory is checked before
	// check if it's unique
	if val, ok := mandatoryFields[callerManField]; ok {
		return fmt.Errorf("repeated mandatory caller field: \"%s\" and \"%s\"", val, prefix)
	} else {
		mandatoryFields[callerManField] = prefix
	}

	// check if unnecessary fields are included
	for key := range ctx {
		switch key {
		case typeField, maxLenField, mandatoryField:
			continue
		default:
			return fmt.Errorf("unexpected field %s in mandatory caller field: %s", key, prefix)
		}
	}

	// check __maxLen
	if raw, ok := ctx[maxLenField]; ok {
		if val, ok := raw.(float64); ok {
			if val < 10 || (val != float64(int(val))) {
				return fmt.Errorf("field %s has wrong value in mandatory caller field: %s. Expect: %s", maxLenField, prefix, "int>=10")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory caller field: %s. Expect: %s", maxLenField, prefix, "int")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory caller field: %s", maxLenField, prefix)
	}

	// check __type
	if raw, ok := ctx[typeField]; ok {
		if val, ok := raw.(string); ok {
			if val != "caller" {
				return fmt.Errorf("field %s has wrong value in mandatory caller field: %s. Expect: %s", timeFormatField, prefix, "caller")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory caller field: %s. Expect: %s", typeField, prefix, "string")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory caller field: %s", typeField, prefix)
	}

	return nil
}

func checkMandatoryLevel(ctx map[string]interface{}, prefix, root string) error {
	// __mandatory is checked before
	// check if it's unique
	if val, ok := mandatoryFields[levelManField]; ok {
		return fmt.Errorf("repeated mandatory loglevel field: \"%s\" and \"%s\"", val, prefix)
	} else {
		mandatoryFields[levelManField] = prefix
	}

	// check if unnecessary fields are included
	for key := range ctx {
		switch key {
		case typeField, maxLenField, mandatoryField:
			continue
		default:
			return fmt.Errorf("unexpected field %s in mandatory loglevel field: %s", key, prefix)
		}
	}

	// check __maxLen
	if raw, ok := ctx[maxLenField]; ok {
		if val, ok := raw.(float64); ok {
			if val < 10 || (val != float64(int(val))) {
				return fmt.Errorf("field %s has wrong value in mandatory loglevel field: %s. Expect: %s", maxLenField, prefix, "int>=10")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory loglevel field: %s. Expect: %s", maxLenField, prefix, "int")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory loglevel field: %s", maxLenField, prefix)
	}

	// check __type
	if raw, ok := ctx[typeField]; ok {
		if val, ok := raw.(string); ok {
			if val != "string" {
				return fmt.Errorf("field %s has wrong value in mandatory loglevel field: %s. Expect: %s", timeFormatField, prefix, "string")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in mandatory loglevel field: %s. Expect: %s", typeField, prefix, "string")
		}
	} else {
		return fmt.Errorf("field %s not found in mandatory loglevel field: %s", typeField, prefix)
	}

	return nil
}

func checkStringContext(ctx map[string]interface{}, prefix, root string) error {
	// __type field is checked before
	// check __omitEmpty
	if raw, ok := ctx[omitEmptyField]; ok {
		if _, ok := raw.(bool); !ok {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", omitEmptyField, prefix, "bool")
		}
	} else {
		return fmt.Errorf("field %s not found in string-typed context: %s", omitEmptyField, prefix)
	}

	// check __maxLen
	if raw, ok := ctx[maxLenField]; ok {
		if val, ok := raw.(float64); ok {
			if val < 10 || (val != float64(int(val))) {
				return fmt.Errorf("field %s has wrong value in string-typed context: %s. Expect: %s", maxLenField, prefix, "int>=10")
			}
		} else {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", maxLenField, prefix, "int")
		}
	} else {
		return fmt.Errorf("field %s not found in string-typed context: %s", maxLenField, prefix)
	}

	// check __default
	if raw, ok := ctx[defaultField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in string-typed context: %s", defaultField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", defaultField, prefix, "string")
		}
	}

	// check __fromEnv
	if raw, ok := ctx[fromEnvField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in string-typed context: %s", fromEnvField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", fromEnvField, prefix, "string")
		}
	}

	// check __required
	if raw, ok := ctx[requiredField]; ok {
		if _, ok := raw.(bool); !ok {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", requiredField, prefix, "bool")
		}
	}

	// check __apiAlias
	if raw, ok := ctx[apiAliasField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in string-typed context: %s", apiAliasField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in string-typed context: %s. Expect: %s", apiAliasField, prefix, "string")
		}
	}

	return nil
}

func checkIntContext(ctx map[string]interface{}, prefix, root string) error {
	// __type field is checked before
	// check __default
	if raw, ok := ctx[defaultField]; ok {
		if _, ok := raw.(float64); !ok {
			return fmt.Errorf("field %s has wrong type in int-typed context: %s. Expect: %s", defaultField, prefix, "int")
		}
	}

	// check __fromEnv
	if raw, ok := ctx[fromEnvField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in int-typed context: %s", fromEnvField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in int-typed context: %s. Expect: %s", fromEnvField, prefix, "string")
		}
	}

	// check __required
	if raw, ok := ctx[requiredField]; ok {
		if _, ok := raw.(bool); !ok {
			return fmt.Errorf("field %s has wrong type in int-typed context: %s. Expect: %s", requiredField, prefix, "bool")
		}
	}

	// check __apiAlias
	if raw, ok := ctx[apiAliasField]; ok {
		if val, ok := raw.(string); ok {
			if len(val) == 0 {
				return fmt.Errorf("field %s has empty value in int-typed context: %s", apiAliasField, prefix)
			}
		} else {
			return fmt.Errorf("field %s has wrong type in int-typed context: %s. Expect: %s", apiAliasField, prefix, "string")
		}
	}

	return nil
}

func recursiveCheckContext(_ctx interface{}, prefix, root string) error {
	ctx, ok := _ctx.(map[string]interface{})
	if !ok {
		return fmt.Errorf("%s does not fit in map[string]interface{}", prefix)
	}
	ctxType := getContextType(ctx)
	switch ctxType {
	case externalCtx:
		for key, val := range ctx {
			recursiveCheckContext(val, prefix+"."+UpperFirst(key), key)
		}
	case internalCtx:
		if raw, ok := ctx[typeField]; ok {
			dtype, ok := raw.(string)
			if !ok {
				return fmt.Errorf("wrong type for __type field in %s. Expect %s", prefix, "string")
			}
			switch dtype {
			case "string":
				err := checkStringContext(ctx, prefix, root)
				if err != nil {
					return err
				}
			case "int":
				err := checkIntContext(ctx, prefix, root)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupport context type %s", dtype)
			}
		} else {
			return fmt.Errorf("found context %s without __type field", prefix)
		}
	case mandatoryMessageCtx:
		err := checkMandatoryMessage(ctx, prefix, root)
		if err != nil {
			return err
		}
	case mandatoryTimestampCtx:
		err := checkMandatoryTimestamp(ctx, prefix, root)
		if err != nil {
			return err
		}
	case mandatoryCallerCtx:
		err := checkMandatoryCaller(ctx, prefix, root)
		if err != nil {
			return err
		}
	case mandatoryLevelCtx:
		err := checkMandatoryLevel(ctx, prefix, root)
		if err != nil {
			return err
		}
	case emptyCtx:
		return fmt.Errorf("found empty context %s", prefix)
	case mixedCtx:
		return fmt.Errorf("found context %s having both internal and external fields", prefix)
	default:
		return fmt.Errorf("found unsupport context %s with type %s", prefix, ctxType)
	}
	return nil
}

func checkContext(ctx map[string]interface{}) error {
	ctxType := getContextType(ctx)
	switch ctxType {
	case externalCtx:
		for key, val := range ctx {
			err := recursiveCheckContext(val, UpperFirst(key), key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
