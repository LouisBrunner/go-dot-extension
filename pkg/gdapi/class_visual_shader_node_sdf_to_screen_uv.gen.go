// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSDFToScreenUV struct {
  VisualShaderNode
}

func (me *VisualShaderNodeSDFToScreenUV) BaseClass() string {
  return "VisualShaderNodeSDFToScreenUV"
}

func NewVisualShaderNodeSDFToScreenUV() *VisualShaderNodeSDFToScreenUV {
  str := StringNameFromStr("VisualShaderNodeSDFToScreenUV") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeSDFToScreenUV{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeSDFToScreenUV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSDFToScreenUV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSDFToScreenUV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
