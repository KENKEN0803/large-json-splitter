package largeJsonSplitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// SplitJson reads a JSON file from the specified inputPath and splits it into multiple JSON files
// based on its structure. The generated JSON files will be placed in subdirectories under the
// original input file's location.
//
// Parameters:
//   - inputPath: The path to the input JSON file.
//   - indentString: The string to be used for indentation in the generated JSON files. An empty string
//     results in no indentation.
//
// Returns:
//   - An error if any issues occur during the splitting process.
func SplitJson(inputPath string, indentString string) error {
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
				processErr := processMap(newFolderPath, key, resMap, value.(map[string]interface{}), indentString)
				if processErr != nil {
					return processErr
				}
			}
		}

		writeJsonErr := writeJSONFile(newFolderPath, nameWithoutExtension, resMap, indentString)
		if writeJsonErr != nil {
			return writeJsonErr
		}
	}
	return nil
}

func processMap(currentPath string, key string, originalMap map[string]interface{}, currentMap map[string]interface{}, indentString string) error {
	newMap := make(map[string]interface{})
	nextPath := fmt.Sprintf("%s/%s", currentPath, key)

	for innerKey, value := range currentMap {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			subMap, ok := value.(map[string]interface{})
			if ok {
				// Recursive call
				err := processMap(nextPath, innerKey, newMap, subMap, indentString)
				if err != nil {
					return err
				}
			}
		} else {
			newMap[innerKey] = value
		}
	}

	if len(newMap) > 0 {
		writeErr := writeJSONFile(nextPath, key, newMap, indentString)
		if writeErr != nil {
			return writeErr
		}
	}

	delete(originalMap, key)

	return nil
}

func writeJSONFile(path string, filename string, data map[string]interface{}, indentString string) error {
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
	if indentString != "" {
		encoder.SetIndent("", indentString)
	}
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
