# Efficient Logger Generator in GoLang (ELG)

HELG is a logger generator which takes a JSON file as input, and generate highly efficient logger code in GoLang allowing insertion in any project with worrying about dependencies.

The JSON file defines the format and the structure of the logging message, also defines customizable APIs for using.

## Features

- Lightweight
    - generate have what is needed only
    - no external dependencies in the generated logger
- Lower memory and cpu footprint
    - Allocation only occurs with caller information and formatted message (Infof, Warnf, ...)
    - Great performance
- Convenient APIs
    - Direct logging like `log.SetXxx().Info()`
    - Constrcut logger with several fields filled
- Set constraints on fields (length, repeatedly set)
- Avoid duplicated fields
- Extendable to any writer fulfilling `io.Writer` interface
- Support caller (file & line number)
- Unit tests
- Banchmark tests
- *Able to accept and merge unstructure data*
- *Error Tracing*

## Thanks

Motivated by the design of [Zerolog](https://github.com/rs/zerolog).

## Great Performance & Quality

```
goos: linux
goarch: amd64
pkg: test/log
cpu: Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz
BenchmarkEmptyLogBelowLogThreshold-2                            461349763                2.181 ns/op           0 B/op          0 allocs/op
BenchmarkEmptyLogWithFormatBelowLogThreshold-2                  334602939                3.596 ns/op           0 B/op          0 allocs/op
BenchmarkAllFieldsLogBelowLogThreshold-2                        21417271                55.31 ns/op            0 B/op          0 allocs/op
BenchmarkAllFieldsLoggerBelowLogThreshold-2                     525708780                2.309 ns/op           0 B/op          0 allocs/op
BenchmarkAllFieldsLoggerWithFormatBelowLogThreshold-2           334947140                3.592 ns/op           0 B/op          0 allocs/op
BenchmarkEmptyLog-2                                               647101              1746 ns/op             216 B/op          2 allocs/op
BenchmarkAllFieldsLog-2                                           585732              1974 ns/op             216 B/op          2 allocs/op
BenchmarkAllFieldsLogWithFormat-2                                 524169              2221 ns/op             240 B/op          3 allocs/op
BenchmarkAllFieldsLogger-2                                        581778              1984 ns/op             216 B/op          2 allocs/op
BenchmarkAllFieldsLoggerWithFormat-2                              512361              2189 ns/op             240 B/op          3 allocs/op
PASS
ok      test/log        12.982s
```

The generated code contains nested structures, string fields, int fields, **log level**, **timestamp**, **caller**.
It comes with nearly zero allocation (`2 allocs/op` is due to caller info by `runtime.Caller(skip)`) and low cpu consumption.

This library provides highly customizable APIs. The users can use any nested structures and insert built-in fields in any position. Meanwhile, it keeps high speed and low memory footprint.

```
ok      test/log        0.004s  coverage: 96.2% of statements
```

Also, the tool generates unit tests along with the code, and high percentage of coverage is achieved.

*The result is based on the input JSON file at [template/sample1.json](./template/sample1.json).*

## How to Use

First design a JSON file describing the log message structure and content, then use the generator to generate code and its associated benchmark tests and unit tests. Then it's ready to go! Feel free to insert the generated code in your projects.

## Supported Types

### string
```
required fields:
{
    "__type": "string", // __type has to be string
    "__omitEmpty": <bool>, // if set to true, omit this field if it's empty
    "__maxLen": <int>, // length in maximum,
    "__description": <stirng>, // short description for the field
}
```

```
optional fields (item comes first has higher priority):
{
    "__default": <string>, // default value
    "__fromEnv": <string>, // grab environment variable as default value
    "__required": <bool>, // if set to true, it must be assigned value before output unless it has a default value
    "__apiAlias": <string>, // unique string for api "Set<string>", if not set, it will be assigned with "Set<Prefix><FieldName>"
}
```

### time.Time
```
required fields:
{
    "__type": "time", // __type has to be time
    "__fromCaller": <bool>, // if set to true, use time at logging as default value. This will set the field to be required and not changable (it means ignoring all optional fields)
    "__timeFormat": <string>, // string to specify time format, like time.RFC3339Nano
    "__omitEmpty": <bool>, // if set to true, omit this field if it's empty
    "__description": <stirng>, // short description for the field
}
```

```
optional fields (item comes first has higher priority):
{
    "__default": <string>, // expect string can be called like time.Parse(__timeFormat, <string>)
    "__fromEnv": <string>, // expect string from env can be called like time.Parse(__timeFormat, <string>)
    "__required": <bool>, // if set to true, it must be assigned value before output unless it has a default value
    "__apiAlias": <string>, // unique string for api "Set<string>", if not set, it will be assigned with "Set<Prefix><FieldName>"
}
```

### int64
```
required fields:
{
    "__type": "int", // __type has to be time
    "__description": <stirng>, // short description for the field
}
```

```
optional fields (item comes first has higher priority):
{
    "__default": <int>, // integer within range of int64
    "__fromEnv": <string>, // expect string from env can be converted to int64
    "__required": <bool>, // if set to true, it must be assigned value before output unless it has a default value
    "__apiAlias": <string>, // unique string for api "Set<string>", if not set, it will be assigned with "Set<Prefix><FieldName>"
}
```

```
Interface of SetXxx:
    Input: (data int64)
```

### float
```
To be implemented
```

### array
```
required fields:
{
    "__type": "array",
    "__elemType": <one of "string", "int64", "float">,
    "__maxLen": <int>,
    "__elemMaxLen": <int>,
    "__omitEmpty": <bool>, // if set to true, omit this field if the array is empty
    "__description": <stirng>, // short description for the field
}
```

```
optional fields:
{
    "__default": <string>, // string that can be used to initialize array with type <__type>: "[\"a\", \"b\"]"
    "__required": <bool>, // if set to true, it must be assigned value before output unless it has a default value,
    "__apiAlias": <string>, // unique string for api "Set<string>", if not set, it will be assigned with "Set<Prefix><FieldName>",
}
```

### any
```
To be implemented
```

## Log Message Definition
Log message is defined using json format with some necessary fields.

### Mandatory Fields

The names of mandatory fields

```
{
    "message": { // Message can only be set with Info/Infof/Warn/... at the end
        "__type": "string",
        "__omitEmpty": false,
        "__maxLen": <int>,
        "__mandatory": "message",
        "__description": <stirng>, // short description for the field
    },
    "timestamp": { // Generate timestamp when calling for logging
        "__type": "time",
        "__fromCaller": true,
        "__timeFormat": <string>,
        "__mandatory": "timestamp",
        "__description": <stirng>, // short description for the field
    },
    "level": { // Severity can only be set with Info/Infof/Warn/... at the end
        "__type": "string",
        "__maxLen": 10,
        "__mandatory": "loglevel",
        "__description": <stirng>, // short description for the field
    },
    "caller": { // Generate caller of logger like "<file>:<line number>"
        "__type": "caller",
        "__maxLen": <int>,
        "__mandatory": "caller",
        "__description": <stirng>, // short description for the field
    }
}
```

The field "__mandatory" is to link the field with one of the mandatory fields (message, timestamp, severity, file). The mandatory fields should be unique with the field "__mandatory".

The design allows users to embed the mandatory fields in any structures and levels.

The mandatory fields can be omitted.

### Body

The other parts of the message definition is json following the supported types.

The definition can be nested, but a field can only have either internal sub-fields (pre-set fields start with "__") or external sub-fields (fields with any name except for internal ones).

## Todo

- [x] Escape detection of string
- [x] Add support for log int
- [x] Add support for caller
- [x] Add support for log level
- [x] Add support for level filter
- [x] Create benchmark tests for generated code
- [x] Parse values from environment variables
- [x] Create unit tests for generated code (90% coverage at least?)
- [x] More precise control on output buffer size
- [x] Add support for Tracef, Debugf, Infof, ... (is it possible to avoid alloc?)
- [ ] Add support for log array
- [ ] Function wrapper API to trace enter/exit of the function calls
- [ ] Remove imported packages in generated code if not needed
- [ ] Add support for log float
