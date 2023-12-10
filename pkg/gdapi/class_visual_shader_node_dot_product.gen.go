// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeDotProduct struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeDotProduct) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeDotProduct) BaseClass() string {
  return "VisualShaderNodeDotProduct"
}



// Enums

func (me *VisualShaderNodeDotProduct) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDotProduct) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDotProduct) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties