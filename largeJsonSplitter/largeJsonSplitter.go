package largeJsonSplitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func SplitJson(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	decoder := json.NewDecoder(file)
	var data interface{}

	for decoder.More() {
		if decodeErr := decoder.Decode(&data); decodeErr != nil {
			return decodeErr
		}

		resMap, ok := data.(map[string]interface{})
		if !ok {
			return errors.New("JSON is not an object")
		}

		parts := strings.Split(inputPath, "/")
		originalFilename := parts[len(parts)-1]
		originalFolderPath := strings.TrimSuffix(inputPath, originalFilename)
		// If just filename is specified, use current directory
		if originalFolderPath == "" {
			originalFolderPath = "./"
		}
		nameWithoutExtension := strings.TrimSuffix(originalFilename, ".json")
		newFolderPath := fmt.Sprintf("%s%s", originalFolderPath, nameWithoutExtension)

		for key, value := range resMap {
			if reflect.TypeOf(value).Kind() == reflect.Map {
				processErr := processMap(newFolderPath, key, resMap, value.(map[string]interface{}))
				if processErr != nil {
					return processErr
				}
			}
		}

		writeJsonErr := writeJSONFile(newFolderPath, nameWithoutExtension, resMap)
		if writeJsonErr != nil {
			return writeJsonErr
		}
	}
	return nil
}

func processMap(currentPath string, key string, originalMap map[string]interface{}, currentMap map[string]interface{}) error {
	newMap := make(map[string]interface{})
	nextPath := fmt.Sprintf("%s/%s", currentPath, key)

	for innerKey, value := range currentMap {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			subMap, ok := value.(map[string]interface{})
			if ok {
				// Recursive call
				err := processMap(nextPath, innerKey, newMap, subMap)
				if err != nil {
					return err
				}
			}
		} else {
			newMap[innerKey] = value
		}
	}

	if len(newMap) > 0 {
		writeErr := writeJSONFile(nextPath, key, newMap)
		if writeErr != nil {
			return writeErr
		}
	}

	delete(originalMap, key)

	return nil
}

func writeJSONFile(path string, filename string, data map[string]interface{}) error {
	newPath := fmt.Sprintf("%s/", path)
	err := os.MkdirAll(newPath, 0700)
	if err != nil {
		return err
	}

	newFilePath := fmt.Sprintf("%s%s.json", newPath, filename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := newFile.Close(); closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
