package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

var outputFuncs = map[string]interface{}{
	"pascalCased":    strcase.ToCamel,
	"lowerCased":     strings.ToLower,
	"mapMethod":      mapMethod,
	"mapName":        mapName,
	"mapType":        mapType,
	"mapWideType":    mapWideType,
	"mapLiteral":     mapLiteral,
	"mapClass":       mapClass,
	"mapOperator":    mapOperator,
	"replace":        strings.ReplaceAll,
	"isExported":     isExported,
	"presents":       mapIsPresent,
	"startsWith":     strings.HasPrefix,
	"trimPrefix":     strings.TrimPrefix,
	"mapPackedType":  mapPackedType,
	"castPackedType": castPackedType,
	"ternary": func(cond bool, a, b string) string {
		if cond {
			return a
		}
		return b
	},
}

func renderFile[T any](templateName, resultFile string, data T, outputDir string) error {
	fileName := fmt.Sprintf("%s.go", templateName)
	outputFilename := fmt.Sprintf("%s.gen.go", resultFile)
	templateFilename := fmt.Sprintf("%s.tmpl", fileName)
	tmpl, err := template.New(templateFilename).Funcs(outputFuncs).ParseFS(templates, filepath.Join("templates", templateFilename), filepath.Join("templates", "fragments", "*.tmpl"))
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
	case "string":
		return "string_"
	case "len":
		return "len_"
	case "error":
		return "error_"
	}
	return n
}

var arrayMatcher = regexp.MustCompile(`\[(\d+)\]`)

func mapType(t string) string {
	switch {
	case strings.HasPrefix(t, "enum::"):
		t = strings.TrimPrefix(t, "enum::")
		return strings.ReplaceAll(t, ".", "")
	case strings.HasPrefix(t, "bitfield::"):
		t = strings.TrimPrefix(t, "bitfield::")
		return strings.ReplaceAll(t, ".", "")
	case strings.HasPrefix(t, "typedarray::"):
		t = strings.TrimPrefix(t, "typedarray::")
		first, rest, found := strings.Cut(t, ":")
		if found {
			return fmt.Sprintf("[]%s", rest)
		}
		return fmt.Sprintf("[]%s", first)
	case strings.Contains(t, "::"):
		return strings.ReplaceAll(t, "::", "")
	case strings.Contains(t, ","):
		// FIXME: technically this is a union
		return "any"
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

func mapWideType(t string) string {
	typ := mapType(t)
	switch typ {
	case "int":
		return "int64"
	case "float32":
		return "float64"
	}
	return typ
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
	if !isExported(name) {
		return strcase.ToCamel(name)
	}
	return name
}

func mapOperator(name string) string {
	switch name {
	case "==":
		return "Equal"
	case "!=":
		return "NotEqual"
	case "<":
		return "Less"
	case "<=":
		return "LessEqual"
	case ">":
		return "Greater"
	case ">=":
		return "GreaterEqual"
	case "+":
		return "Add"
	case "-":
		return "Subtract"
	case "unary+":
		return "Positive"
	case "unary-":
		return "Negate"
	case "*":
		return "Multiply"
	case "/":
		return "Divide"
	case "**":
		return "Power"
	case "%":
		return "Module"
	case "&":
		return "BitAnd"
	case "|":
		return "BitOr"
	case "^":
		return "BitXor"
	case "~":
		return "BitNegate"
	case "<<":
		return "ShiftLeft"
	case ">>":
		return "ShiftRight"
	case "and":
		return "And"
	case "or":
		return "Or"
	case "xor":
		return "Xor"
	case "not":
		return "Not"
	case "in":
		return "In"
	case "max":
		return "Max"
	default:
		return name
	}
}

func mapIsPresent(obj interface{}, name string) bool {
	jsonb, err := json.Marshal(obj)
	if err != nil {
		return false
	}
	objMap := make(map[string]interface{})
	err = json.Unmarshal(jsonb, &objMap)
	if err != nil {
		return false
	}
	_, ok := objMap[name]
	return ok
}

func isExported(name string) bool {
	return unicode.IsUpper(rune(name[0]))
}

func mapPackedType(t string) string {
	core := strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(t, "Packed"), "Array"))
	switch core {
	case "color":
		return "Color"
	case "vector2":
		return "Vector2"
	case "vector3":
		return "Vector3"
	}
	return core
}

func castPackedType(t, name string) string {
	core := strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(t, "Packed"), "Array"))
	switch core {
	case "int32":
		return fmt.Sprintf("%s(*%s)", core, name)
	case "string":
		return fmt.Sprintf("StringFromPtr(gdc.ConstTypePtr(%s)).String()", name)
	case "color":
		return fmt.Sprintf("*ColorFromPtr(gdc.ConstTypePtr(%s))", name)
	case "vector2":
		return fmt.Sprintf("*Vector2FromPtr(gdc.ConstTypePtr(%s))", name)
	case "vector3":
		return fmt.Sprintf("*Vector3FromPtr(gdc.ConstTypePtr(%s))", name)
	}
	return fmt.Sprintf("*%s", name)
}
