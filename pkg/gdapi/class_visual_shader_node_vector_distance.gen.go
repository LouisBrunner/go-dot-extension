// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeVectorDistance struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorDistance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorDistance) BaseClass() string {
  return "VisualShaderNodeVectorDistance"
}



// Enums

func (me *VisualShaderNodeVectorDistance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorDistance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorDistance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties