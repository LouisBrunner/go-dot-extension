package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/go-dot-extension/pkg/generators/gencommon"
)

func generate(input, outputDir string) error {
	api, err := gencommon.ReadAPI(input)
	if err != nil {
		return err
	}
	return output(api, outputDir)
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fail("Usage: apigen <input> <output-dir>")
	}
	input := args[0]
	outputDir := args[1]
	err := generate(input, outputDir)
	if err != nil {
		fail(fmt.Sprintf("error: %s", err))
	}
}
