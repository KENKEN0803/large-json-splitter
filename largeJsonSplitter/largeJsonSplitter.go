package largeJsonSplitter

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func SplitJson(inputPath *string) {
	file, err := os.Open(*inputPath)
	if err != nil {
		panic(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	var data interface{}

	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		resMap, ok := data.(map[string]interface{})
		if !ok {
			panic("JSON decoding error")
		}

		parts := strings.Split(*inputPath, "/")
		originalFilename := parts[len(parts)-1]
		originalFolderPath := strings.TrimSuffix(*inputPath, originalFilename)
		nameWithoutExtension := strings.TrimSuffix(originalFilename, ".json")
		newFolderPath := fmt.Sprintf("%s/%s/", originalFolderPath, nameWithoutExtension)

		for key, value := range resMap {
			if reflect.TypeOf(value).Kind() == reflect.Map {
				processMap(newFolderPath, key, resMap, value.(map[string]interface{}))
			}
		}

		writeJSONFile(newFolderPath, nameWithoutExtension, resMap)
	}
}

func processMap(currentPath string, key string, originalMap map[string]interface{}, currentMap map[string]interface{}) bool {
	newMap := make(map[string]interface{})
	nextPath := fmt.Sprintf("%s/%s", currentPath, key)

	for innerKey, value := range currentMap {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			subMap, ok := value.(map[string]interface{})
			if ok {
				// Recursively process the subMap
				shouldContinue := processMap(nextPath, innerKey, newMap, subMap)
				if !shouldContinue {
					return false
				}
			}
		} else {
			newMap[innerKey] = value
		}
	}

	if len(newMap) > 0 {
		writeJSONFile(nextPath, key, newMap)
	}

	delete(originalMap, key)

	return true
}

func writeJSONFile(path string, filename string, data map[string]interface{}) {
	newPath := fmt.Sprintf("%s/", path)
	err := os.MkdirAll(newPath, 0700)
	if err != nil {
		panic(err)
		return
	}

	newFilePath := fmt.Sprintf("%s%s.json", newPath, filename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		panic(err)
		return
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			panic(err)
		}
	}(newFile)

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}
