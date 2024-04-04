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

type VisualShaderNodeOutput struct {
  VisualShaderNode
}

func (me *VisualShaderNodeOutput) BaseClass() string {
  return "VisualShaderNodeOutput"
}

func NewVisualShaderNodeOutput() *VisualShaderNodeOutput {
  str := StringNameFromStr("VisualShaderNodeOutput") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeOutput{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeOutput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeOutput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeOutput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
