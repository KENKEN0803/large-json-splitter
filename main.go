package main

import (
	"flag"
	"log"
	"time"

	"github.com/KENKEN0803/large-json-splitter/largeJsonSplitter"
)

func main() {
	inputFilePath := flag.String("input", "", "Input JSON file path")
	flag.StringVar(inputFilePath, "i", "", "Input JSON file path")
	flag.Parse()

	if *inputFilePath == "" {
		log.Fatalf("Input file path not specified. Usage: ./large-json-splitter -i <input-file-path>")
	}

	startTime := time.Now()
	log.Printf("Starting large-json-splitter for file: %s", *inputFilePath)

	err := largeJsonSplitter.SplitJson(*inputFilePath)
	if err != nil {
		log.Fatalf("Error splitting JSON: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("Output files successfully created in %s. Elapsed time: %s", *inputFilePath, elapsedTime)
}
