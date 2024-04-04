// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeOuterProduct struct {
  VisualShaderNode
}

func (me *VisualShaderNodeOuterProduct) BaseClass() string {
  return "VisualShaderNodeOuterProduct"
}

func NewVisualShaderNodeOuterProduct() *VisualShaderNodeOuterProduct {
  str := StringNameFromStr("VisualShaderNodeOuterProduct") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeOuterProduct{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
