// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeVectorLen struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorLen) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorLen) BaseClass() string {
  return "VisualShaderNodeVectorLen"
}



// Enums

func (me *VisualShaderNodeVectorLen) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorLen) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorLen) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
