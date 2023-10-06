package main

import (
	"flag"
	"github.com/KENKEN0803/large-json-splitter/largeJsonSplitter"
)

func main() {
	inputPath := flag.String("input", "", "Input JSON file path")
	flag.StringVar(inputPath, "i", "", "Input JSON file path")
	flag.Parse()

	if *inputPath == "" {
		panic("Input file path not specified")
	}

	err := largeJsonSplitter.SplitJson(*inputPath)
	if err != nil {
		panic(err)
	}
}
