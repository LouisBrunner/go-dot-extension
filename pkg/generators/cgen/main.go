package main

import (
	"fmt"
	"os"

	// FIXME: would prefer to use `modernc.org/cc/v4` directly but it doesn't work
	"github.com/xlab/c-for-go/parser"
)

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func generate(input, outputDir string) error {
	interfaceFuncs, err := findInterfaceFuncs(input)
	if err != nil {
		return err
	}
	ast, err := parser.ParseWith(&parser.Config{
		SourcesPaths: []string{input},
	})
	if err != nil {
		return err
	}
	return output(ast, interfaceFuncs, outputDir)
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fail("Usage: cgen <input-header> <input-api> <output-dir>")
	}
	input := args[0]
	// FIXME: ignored for now
	// inputJSON := args[1]
	outputDir := args[2]
	err := generate(input, outputDir)
	if err != nil {
		fail(fmt.Sprintf("error: %s", err))
	}
}
