// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeScreenUVToSDF struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeScreenUVToSDF) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeScreenUVToSDF) BaseClass() string {
  return "VisualShaderNodeScreenUVToSDF"
}



// Enums

func (me *VisualShaderNodeScreenUVToSDF) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeScreenUVToSDF) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeScreenUVToSDF) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties