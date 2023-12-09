package main

import (
	"embed"
	"fmt"

	"github.com/LouisBrunner/go-dot-extension/pkg/generators/gencommon"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*
var templates embed.FS

const bitsPerWord = 32 << (^uint(0) >> 63)

func getBuiltinClassSize(list []gencommon.ExtensionClassSize, config string) (*gencommon.ExtensionClassSize, bool) {
	for _, class := range list {
		if class.GetBuildConfiguration() == config {
			return &class, true
		}
	}
	return nil, false
}

func getBuiltinClassMembers(list []gencommon.ExtensionClassMemberOffsetConfig, config string) (*gencommon.ExtensionClassMemberOffsetConfig, bool) {
	for _, class := range list {
		if class.GetBuildConfiguration() == config {
			return &class, true
		}
	}
	return nil, false
}

func renderClass(class gencommon.ExtensionClass, outputDir string) error {
	name := strcase.ToSnake(class.Name)
	err := renderFile("class", fmt.Sprintf("class_%s", name), class, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class %s: %w", class.Name, err)
	}
	return nil
}

func renderBClass(class gencommon.ExtensionBuiltinClass, outputDir string) error {
	name := strcase.ToSnake(class.Name)
	err := renderFile("bclass", fmt.Sprintf("bclass_%s", name), class, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class %s: %w", class.Name, err)
	}
	return nil
}

func output(api *gencommon.ExtensionAPI, outputDir string) error {
	config := fmt.Sprintf("float_%d", bitsPerWord)
	classSizes, ok := getBuiltinClassSize(api.BuiltinClassSizes, config)
	if !ok {
		return fmt.Errorf("no class size for %s", config)
	}

	err := renderFile("class_sizes", "classes_sizes", classSizes, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class sizes: %w", err)
	}

	classMembers, ok := getBuiltinClassMembers(api.BuiltinClassMemberOffsets, config)
	if !ok {
		return fmt.Errorf("no class size for %s", config)
	}

	err = renderFile("class_members", "classes_members", classMembers, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class members: %w", err)
	}

	for _, class := range api.BuiltinClasses {
		err := renderBClass(class, outputDir)
		if err != nil {
			return fmt.Errorf("error rendering class %s: %w", class.Name, err)
		}
	}

	for _, class := range api.Classes {
		err := renderClass(class, outputDir)
		if err != nil {
			return fmt.Errorf("error rendering class %s: %w", class.Name, err)
		}
	}

	err = renderFile("enums", "enums", api.GlobalEnums, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering enums: %w", err)
	}

	err = renderFile("native_structures", "native_structures", api.NativeStructures, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering native structures: %w", err)
	}

	err = renderFile("singletons", "singletons", api.Singletons, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering singletons: %w", err)
	}

	err = renderFile("utilities", "utilities", api.UtilityFunctions, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering utilities: %w", err)
	}

	return nil

}
