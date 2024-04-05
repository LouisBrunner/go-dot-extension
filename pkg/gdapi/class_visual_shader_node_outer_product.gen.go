// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForVisualShaderNodeOuterProductList struct {
}

var ptrsForVisualShaderNodeOuterProduct ptrsForVisualShaderNodeOuterProductList

func initVisualShaderNodeOuterProductPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeOuterProduct")
  defer className.Destroy()
}

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
