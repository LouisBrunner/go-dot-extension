// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Engine struct {
  obj gdc.ObjectPtr
}

func createEngine(obj gdc.ObjectPtr) *Engine {
  return &Engine{
    obj: obj,
  }
}

func (me *Engine) BaseClass() string {
  return "Engine"
}
