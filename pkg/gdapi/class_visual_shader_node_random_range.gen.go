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

type VisualShaderNodeRandomRange struct {
  VisualShaderNode
}

func (me *VisualShaderNodeRandomRange) BaseClass() string {
  return "VisualShaderNodeRandomRange"
}

func NewVisualShaderNodeRandomRange() *VisualShaderNodeRandomRange {
  str := StringNameFromStr("VisualShaderNodeRandomRange") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeRandomRange{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeRandomRange) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeRandomRange) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeRandomRange) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
