// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForEngineProfilerList struct {
  fnXToggle gdc.MethodBindPtr
  fnXAddFrame gdc.MethodBindPtr
  fnXTick gdc.MethodBindPtr
}

var ptrsForEngineProfiler ptrsForEngineProfilerList

func initEngineProfilerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EngineProfiler")
  defer className.Destroy()
}

type EngineProfiler struct {
  RefCounted
}

func (me *EngineProfiler) BaseClass() string {
  return "EngineProfiler"
}

func NewEngineProfiler() *EngineProfiler {
  str := StringNameFromStr("EngineProfiler") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EngineProfiler{}
  obj.SetBaseObject(objPtr)
  return obj
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
