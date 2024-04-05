package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/LouisBrunner/go-dot-extension/pkg/generators/gencommon"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
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

var goos_list = map[string]struct{}{
	"aix":       {},
	"android":   {},
	"darwin":    {},
	"dragonfly": {},
	"freebsd":   {},
	"hurd":      {},
	"illumos":   {},
	"ios":       {},
	"js":        {},
	"linux":     {},
	"nacl":      {},
	"netbsd":    {},
	"openbsd":   {},
	"plan9":     {},
	"solaris":   {},
	"wasip1":    {},
	"windows":   {},
	"zos":       {},
}

func renderClass(class gencommon.ExtensionClass, outputDir string) error {
	name := strcase.ToSnake(class.Name)
	for goos := range goos_list {
		if strings.HasSuffix(name, "_"+goos) {
			name += "_"
			break
		}
	}
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

type lookups struct {
	gencommon.ExtensionClassSize
	Classes []string
}

var forbiddenClasses = []string{
	// these singletons are unavailable and we can't load their methods, so let's just drop them
	"JavaClassWrapper",
	"JavaScriptBridge",
	"DisplayServer",
	"RenderingServer",
	"AudioServer",
	"PhysicsServer2D",
	"PhysicsServer3D",
	"NavigationServer2D",
	"NavigationServer3D",
	"XRServer",
	"CameraServer",
}

var forbiddenSingletons = []string{
	"EditorInterface", // unavailable outside of the editor
	// all of those fail for some reasons, they do seem to be included in a "servers" file in Godot
	"JavaClassWrapper",
	"JavaScriptBridge",
	"DisplayServer",
	"RenderingServer",
	"AudioServer",
	"PhysicsServer2D",
	"PhysicsServer3D",
	"NavigationServer2D",
	"NavigationServer3D",
	"XRServer",
	"CameraServer",
}

func output(api *gencommon.ExtensionAPI, outputDir string) error {
	config := fmt.Sprintf("float_%d", bitsPerWord)
	classSizes, ok := getBuiltinClassSize(api.BuiltinClassSizes, config)
	if !ok {
		return fmt.Errorf("no class size for %s", config)
	}

	// filter out things which don't work through the extension but are somehow included in the JSON
	api.Classes = slices.DeleteFunc(api.Classes, func(e gencommon.ExtensionClass) bool {
		return slices.Contains(forbiddenClasses, e.Name)
	})
	api.Singletons = slices.DeleteFunc(api.Singletons, func(e gencommon.ExtensionNameType) bool {
		return slices.Contains(forbiddenSingletons, e.Name)
	})

	classes := make([]string, 0, len(api.Classes))
	for _, class := range api.Classes {
		classes = append(classes, class.Name)
	}

	err := renderFile("lookups", "lookups", lookups{
		ExtensionClassSize: *classSizes,
		Classes:            classes,
	}, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering class sizes: %w", err)
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

	err = renderFile("class_inits", "class_inits", map[string]interface{}{
		"Classes":  api.Classes,
		"BClasses": api.BuiltinClasses,
	}, outputDir)
	if err != nil {
		return fmt.Errorf("error rendering utilities: %w", err)
	}

	return nil

}
