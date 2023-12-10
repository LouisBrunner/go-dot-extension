// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeDeterminant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeDeterminant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeDeterminant) BaseClass() string {
  return "VisualShaderNodeDeterminant"
}



// Enums

func (me *VisualShaderNodeDeterminant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDeterminant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDeterminant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties