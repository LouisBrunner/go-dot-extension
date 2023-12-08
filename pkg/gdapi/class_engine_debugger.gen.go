// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EngineDebugger struct {
  obj gdc.ObjectPtr
}

func createEngineDebugger(obj gdc.ObjectPtr) *EngineDebugger {
  return &EngineDebugger{
    obj: obj,
  }
}

func (me *EngineDebugger) BaseClass() string {
  return "EngineDebugger"
}
