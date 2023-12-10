// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeRandomRange struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeRandomRange) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeRandomRange) BaseClass() string {
  return "VisualShaderNodeRandomRange"
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

// Properties