// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EngineProfiler struct {
  obj gdc.ObjectPtr
}

func createEngineProfiler(obj gdc.ObjectPtr) *EngineProfiler {
  return &EngineProfiler{
    obj: obj,
  }
}

func (me *EngineProfiler) BaseClass() string {
  return "EngineProfiler"
}