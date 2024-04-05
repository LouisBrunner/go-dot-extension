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

type ptrsForVisualShaderNodeDistanceFadeList struct {
}

var ptrsForVisualShaderNodeDistanceFade ptrsForVisualShaderNodeDistanceFadeList

func initVisualShaderNodeDistanceFadePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeDistanceFade")
  defer className.Destroy()
}

type VisualShaderNodeDistanceFade struct {
  VisualShaderNode
}

func (me *VisualShaderNodeDistanceFade) BaseClass() string {
  return "VisualShaderNodeDistanceFade"
}

func NewVisualShaderNodeDistanceFade() *VisualShaderNodeDistanceFade {
  str := StringNameFromStr("VisualShaderNodeDistanceFade") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeDistanceFade{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeDistanceFade) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDistanceFade) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDistanceFade) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
