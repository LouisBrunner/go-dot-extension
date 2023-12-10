// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSDFToScreenUV struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSDFToScreenUV) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSDFToScreenUV) BaseClass() string {
  return "VisualShaderNodeSDFToScreenUV"
}



// Enums

func (me *VisualShaderNodeSDFToScreenUV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSDFToScreenUV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSDFToScreenUV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties