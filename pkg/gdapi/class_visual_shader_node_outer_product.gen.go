// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeOuterProduct struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeOuterProduct) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeOuterProduct) BaseClass() string {
  return "VisualShaderNodeOuterProduct"
}



// Enums

func (me *VisualShaderNodeOuterProduct) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeOuterProduct) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeOuterProduct) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties