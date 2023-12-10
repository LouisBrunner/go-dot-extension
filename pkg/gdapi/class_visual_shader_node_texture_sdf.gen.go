// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTextureSDF struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTextureSDF) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTextureSDF) BaseClass() string {
  return "VisualShaderNodeTextureSDF"
}



// Enums

func (me *VisualShaderNodeTextureSDF) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureSDF) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureSDF) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties