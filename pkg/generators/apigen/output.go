package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

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

func mapName(n string) string {
	switch n {
	case "type":
		return "type_"
	case "range":
		return "range_"
	case "default":
		return "default_"
	case "func":
		return "func_"
	case "interface":
		return "interface_"
	case "map":
		return "map_"
	case "var":
		return "var_"
	}
	return n
}

var arrayMatcher = regexp.MustCompile(`\[(\d+)\]`)

func mapType(t string) string {
	switch {
	case strings.HasPrefix(t, "enum::"):
		return strings.ReplaceAll(strings.TrimPrefix(t, "enum::"), ".", "")
	case strings.HasPrefix(t, "bitfield::"):
		return strings.ReplaceAll(strings.TrimPrefix(t, "bitfield::"), ".", "")
	case strings.HasPrefix(t, "typedarray::"):
		return strings.ReplaceAll(strings.TrimPrefix(t, "typedarray::"), ".", "")
	case strings.Contains(t, "::"):
		return strings.ReplaceAll(t, "::", "")
	case strings.Contains(t, "_t") || strings.HasPrefix(t, "const ") || strings.Contains(t, "void") || strings.Contains(t, "*"):
		// isConst := strings.HasPrefix(t, "const ")
		t = strings.TrimPrefix(t, "const ")
		ptrs := strings.Count(t, "*")
		t = strings.Trim(t, "*")
		typ := strings.TrimSpace(t)
		switch typ {
		case "float":
			typ = "float32"
		case "double":
			typ = "float64"
		case "void":
			return "unsafe.Pointer"
		case "uint8_t":
			typ = "uint8"
		case "uint16_t":
			typ = "uint16"
		case "uint32_t":
			typ = "uint32"
		case "int32_t":
			typ = "int32"
		case "uint64_t":
			typ = "uint64"
		case "real_t":
			typ = "RealType"
		}
		return fmt.Sprintf("%s%s", strings.Repeat("*", ptrs), typ)
	case strings.Contains(t, "["):
		matches := arrayMatcher.FindStringSubmatch(t)
		if len(matches) != 2 {
			return t
		}
		return fmt.Sprintf("[%s]%s", matches[1], mapType(strings.TrimSuffix(t, matches[0])))
	case t == "float":
		return "float32"
	case t == "double":
		return "float64"
	}
	return t
}

func mapLiteral(t string) string {
	switch {
	case strings.HasSuffix(t, ".f"):
		return strings.TrimSuffix(t, ".f")
	}
	return t
}

func mapMethod(name string) string {
	if strings.HasPrefix(name, "_") {
		name = "X" + name
	}
	return strcase.ToCamel(name)
}

func mapClass(name string) string {
	if unicode.IsLower(rune(name[0])) {
		return strcase.ToCamel(name)
	}
	return name
}

func renderFile[T any](templateName, resultFile string, data T, outputDir string) error {
	fileName := fmt.Sprintf("%s.go", templateName)
	outputFilename := fmt.Sprintf("%s.gen.go", resultFile)
	templateFilename := fmt.Sprintf("%s.tmpl", fileName)
	tmpl, err := template.New(templateFilename).Funcs(map[string]interface{}{
		"pascalCased": strcase.ToCamel,
		"lowerCased":  strings.ToLower,
		"mapMethod":   mapMethod,
		"mapName":     mapName,
		"mapType":     mapType,
		"mapLiteral":  mapLiteral,
		"mapClass":    mapClass,
		"replace":     strings.ReplaceAll,
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

func renderBClass(class gencommon.ExtensionBuiltinClass, outputDir string) error {
	name := strcase.ToSnake(class.Name)
	err := renderFile("class", fmt.Sprintf("bclass_%s", name), class, outputDir)
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
