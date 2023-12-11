// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeIf struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIf) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIf) BaseClass() string {
  return "VisualShaderNodeIf"
}



// Enums

func (me *VisualShaderNodeIf) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIf) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIf) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
