package main

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

func improveTypename(typename string) string {
	return strings.TrimPrefix(typename, prefixTypes)
}

func improveEnumValueName(name string) string {
	newName := strcase.ToCamel(strings.ToLower(strings.TrimPrefix(name, prefixEnum)))
	// FIXME: wonky
	newName = strings.ReplaceAll(newName, "Aabb", "AABB")
	newName = strings.ReplaceAll(newName, "Rid", "RID")
	return newName
}

func makeGoTypeName(name string) string {
	return strcase.ToCamel(name)
}

func makeGoVarName(name string) string {
	return strcase.ToLowerCamel(name)
}

var reservedNames = []string{
	"type",
}

func makeCGoName(name string) string {
	if slices.Contains(reservedNames, name) {
		return fmt.Sprintf("_%s", name)
	}
	return name
}
