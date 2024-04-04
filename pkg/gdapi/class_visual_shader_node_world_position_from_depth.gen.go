// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeWorldPositionFromDepth struct {
  VisualShaderNode
}

func (me *VisualShaderNodeWorldPositionFromDepth) BaseClass() string {
  return "VisualShaderNodeWorldPositionFromDepth"
}

func NewVisualShaderNodeWorldPositionFromDepth() *VisualShaderNodeWorldPositionFromDepth {
  str := StringNameFromStr("VisualShaderNodeWorldPositionFromDepth") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeWorldPositionFromDepth{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeWorldPositionFromDepth) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeWorldPositionFromDepth) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeWorldPositionFromDepth) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
