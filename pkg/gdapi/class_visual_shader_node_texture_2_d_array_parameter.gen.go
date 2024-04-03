// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTexture2DArrayParameter struct {
  VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTexture2DArrayParameter) BaseClass() string {
  return "VisualShaderNodeTexture2DArrayParameter"
}

func NewVisualShaderNodeTexture2DArrayParameter() *VisualShaderNodeTexture2DArrayParameter {
  str := StringNameFromStr("VisualShaderNodeTexture2DArrayParameter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTexture2DArrayParameter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeTexture2DArrayParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture2DArrayParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DArrayParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
