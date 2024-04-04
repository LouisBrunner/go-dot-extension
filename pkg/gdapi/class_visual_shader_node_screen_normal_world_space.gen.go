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

type VisualShaderNodeScreenNormalWorldSpace struct {
  VisualShaderNode
}

func (me *VisualShaderNodeScreenNormalWorldSpace) BaseClass() string {
  return "VisualShaderNodeScreenNormalWorldSpace"
}

func NewVisualShaderNodeScreenNormalWorldSpace() *VisualShaderNodeScreenNormalWorldSpace {
  str := StringNameFromStr("VisualShaderNodeScreenNormalWorldSpace") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeScreenNormalWorldSpace{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeScreenNormalWorldSpace) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeScreenNormalWorldSpace) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeScreenNormalWorldSpace) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
