package main

import (
	gde "github.com/LouisBrunner/go-dot-extension"
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type TestClass struct {
	gdapi.Node
}

func newTestClass() gde.Class {
	return &TestClass{}
}

func init() {
	entry.DebugMode = true
	gde.Register(newTestClass)
}

func main() {}
