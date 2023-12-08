package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func renderFile[T any](templateName, resultFile string, data T, outputDir string) error {
	fileName := fmt.Sprintf("%s.go", templateName)
	outputFilename := fmt.Sprintf("%s.gen.go", resultFile)
	templateFilename := fmt.Sprintf("%s.tmpl", fileName)
	tmpl, err := template.New(templateFilename).Funcs(map[string]interface{}{
		"pascalCased": strcase.ToCamel,
		"lowerCased":  strings.ToLower,
	}).ParseFS(templates, filepath.Join("templates", templateFilename), filepath.Join("templates", "fragments", "*.tmpl"))
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

func renderClass(class gencommon.ExtensionClass, outputDir string) error {
	name := strcase.ToSnake(class.Name)
	err := renderFile("class", fmt.Sprintf("class_%s", name), class, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class %s: %w", class.Name, err)
	}
	return nil
}

func output(api *gencommon.ExtensionAPI, outputDir string) error {
	config := fmt.Sprintf("float_%d", bitsPerWord)
	classSizes, ok := getBuiltinClassSize(api, config)
	if !ok {
		return fmt.Errorf("no class size for %s", config)
	}
	err := renderFile("class_sizes", "classes_sizes", classSizes, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class sizes: %w", err)
	}

	// for _, class := range api.BuiltinClasses {
	// 	err := renderClass(class, outputDir)
	// 	if err != nil {
	// 		return fmt.Errorf("error rendering class %s: %w", class.Name, err)
	// 	}
	// }

	for _, class := range api.Classes {
		err := renderClass(class, outputDir)
		if err != nil {
			return fmt.Errorf("error rendering class %s: %w", class.Name, err)
		}
	}

	// TODO: finish
	return nil
}
