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

type VisualShaderNodeCubemapParameter struct {
  VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeCubemapParameter) BaseClass() string {
  return "VisualShaderNodeCubemapParameter"
}

func NewVisualShaderNodeCubemapParameter() *VisualShaderNodeCubemapParameter {
  str := StringNameFromStr("VisualShaderNodeCubemapParameter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeCubemapParameter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeCubemapParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCubemapParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCubemapParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
