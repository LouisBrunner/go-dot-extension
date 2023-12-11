// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EngineProfiler struct {
  obj gdc.ObjectPtr
}

func (me *EngineProfiler) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EngineProfiler) BaseClass() string {
  return "EngineProfiler"
}



// Enums

func (me *EngineProfiler) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EngineProfiler) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EngineProfiler) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
