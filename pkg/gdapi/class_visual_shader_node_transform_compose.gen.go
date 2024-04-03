// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
