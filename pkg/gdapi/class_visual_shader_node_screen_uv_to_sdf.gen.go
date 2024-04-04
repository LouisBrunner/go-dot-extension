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

type VisualShaderNodeScreenUVToSDF struct {
  VisualShaderNode
}

func (me *VisualShaderNodeScreenUVToSDF) BaseClass() string {
  return "VisualShaderNodeScreenUVToSDF"
}

func NewVisualShaderNodeScreenUVToSDF() *VisualShaderNodeScreenUVToSDF {
  str := StringNameFromStr("VisualShaderNodeScreenUVToSDF") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeScreenUVToSDF{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeScreenUVToSDF) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeScreenUVToSDF) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeScreenUVToSDF) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
