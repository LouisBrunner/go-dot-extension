// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeWorldPositionFromDepth struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeWorldPositionFromDepth) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeWorldPositionFromDepth) BaseClass() string {
  return "VisualShaderNodeWorldPositionFromDepth"
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
