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

type ptrsForVisualShaderNodeConstantList struct {
}

var ptrsForVisualShaderNodeConstant ptrsForVisualShaderNodeConstantList

func initVisualShaderNodeConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeConstant")
  defer className.Destroy()
}

type VisualShaderNodeConstant struct {
  VisualShaderNode
}

func (me *VisualShaderNodeConstant) BaseClass() string {
  return "VisualShaderNodeConstant"
}

func NewVisualShaderNodeConstant() *VisualShaderNodeConstant {
  str := StringNameFromStr("VisualShaderNodeConstant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeConstant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
