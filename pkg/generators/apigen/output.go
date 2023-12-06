package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/LouisBrunner/go-dot-extension/pkg/generators/gencommon"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*
var templates embed.FS

const bitsPerWord = 32 << (^uint(0) >> 63)

func getBuiltinClassSize(api *gencommon.ExtensionAPI, config string) (*gencommon.ExtensionClassSize, bool) {
	for _, class := range api.BuiltinClassSizes {
		if class.BuildConfiguration == config {
			return &class, true
		}
	}
	return nil, false
}

func renderFile[T any](templateName string, data T, outputDir string) error {
	fileName := fmt.Sprintf("%s.go", templateName)
	outputFilename := fmt.Sprintf("%s.gen.go", templateName)
	templateFilename := fmt.Sprintf("%s.tmpl", fileName)
	tmpl, err := template.New(templateFilename).Funcs(map[string]interface{}{
		"pascalCased": strcase.ToCamel,
	}).ParseFS(templates, filepath.Join("templates", templateFilename))
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(outputDir, outputFilename))
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, data)
}

func output(api *gencommon.ExtensionAPI, outputDir string) error {
	config := fmt.Sprintf("float_%d", bitsPerWord)
	classSizes, ok := getBuiltinClassSize(api, config)
	if !ok {
		return fmt.Errorf("no class size for %s", config)
	}
	err := renderFile("class_sizes", classSizes, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class sizes: %w", err)
	}

	// TODO: finish
	return nil
}
