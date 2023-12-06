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
	return strcase.ToCamel(strings.ToLower(strings.TrimPrefix(name, prefixEnum)))
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
