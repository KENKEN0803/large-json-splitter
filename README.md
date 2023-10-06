# Large JSON Splitter

Large JSON Splitter is a Go library that reads JSON data as a stream and extracts objects, saving them as separate JSON files.

![GitHub](https://img.shields.io/github/license/KENKEN0803/large-json-splitter)

## Input
```json5
// Original sample.json
{
  "stringValue": "value",
  "intValue": 1,
  "boolValue": true,
  "arrayValue": [
    "value1",
    "value2"
  ],
  
  "objectValue": {
    "innerStringValue": "value",
    "innerIntValue": 1,
    "innerBoolValue": true,
    "innerArrayValue": [
      "value1",
      "value2"
    ]
  }
  
}
```

## Output
```
├── sample
│   ├── objectValue
│   │   └── objectValue.json
│   └── sample.json
└── sample.json
```
```json5
// sample.json with objectValue key removed
{
  "arrayValue": [
    "value1",
    "value2"
  ],
  "boolValue": true,
  "intValue": 1,
  "stringValue": "value"
}
```

```json5
// objectValue.json
{
  "innerArrayValue": [
    "value1",
    "value2"
  ],
  "innerBoolValue": true,
  "innerIntValue": 1,
  "innerStringValue": "value"
}
```



## Installation

You can install this library using `go get`:

```bash
go get github.com/KENKEN0803/large-json-splitter
```

## Usage

The large-json-splitter library provides both a command-line executable and a Go function for splitting JSON files.

## Executable Usage
Executables can be found on [release page](https://github.com/KENKEN0803/large-json-splitter/releases).
You can use the executable as follows:

```bash
./large-json-splitter --i ./data/141.json
```

- --input, --i: Specifies the input JSON file to process.

## Go Library Usage
```go
package main

import (
    "github.com/KENKEN0803/large-json-splitter"
    "fmt"
)

func main() {
    inputPath := "89.json"
    
    if err := largeJsonSplitter.SplitJson(inputPath); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

```

## GitHub Repository
Find the source code and additional documentation on the [GitHub Repository](https://github.com/KENKEN0803/large-json-splitter)

## License
This project is licensed under the MIT License. See the [LICENSE](https://choosealicense.com/licenses/mit/) file for details.
