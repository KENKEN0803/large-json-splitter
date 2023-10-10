package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/KENKEN0803/large-json-splitter/largeJsonSplitter"
)

func main() {
	inputFilePath := flag.String("input", "", "Input JSON file path")
	indent := flag.String("indent", "", "Indentation string for JSON files")
	flag.StringVar(inputFilePath, "i", "", "Input JSON file path")
	flag.StringVar(indent, "d", "", "Indentation string for JSON files")
	flag.Parse()

	if *inputFilePath == "" {
		log.Fatalf("Input file path not specified. Usage: ./large-json-splitter -i <input-file-path> [-d <indentation>] [-indent <indentation>]")
	}

	startTime := time.Now()
	log.Printf("Starting large-json-splitter for file: %s", *inputFilePath)

	err := largeJsonSplitter.SplitJson(*inputFilePath, *indent)
	if err != nil {
		log.Fatalf("Error splitting JSON: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("Output files successfully created in %s/\nElapsed time: %s", strings.TrimSuffix(*inputFilePath, ".json"), elapsedTime)
}
