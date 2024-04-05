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

type ptrsForVisualShaderNodeTransformComposeList struct {
}

var ptrsForVisualShaderNodeTransformCompose ptrsForVisualShaderNodeTransformComposeList

func initVisualShaderNodeTransformComposePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeTransformCompose")
  defer className.Destroy()
}

type VisualShaderNodeTransformCompose struct {
  VisualShaderNode
}

func (me *VisualShaderNodeTransformCompose) BaseClass() string {
  return "VisualShaderNodeTransformCompose"
}

func NewVisualShaderNodeTransformCompose() *VisualShaderNodeTransformCompose {
  str := StringNameFromStr("VisualShaderNodeTransformCompose") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTransformCompose{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeTransformCompose) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformCompose) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformCompose) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
