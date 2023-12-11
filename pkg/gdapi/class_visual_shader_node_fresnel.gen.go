// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeFresnel struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFresnel) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFresnel) BaseClass() string {
  return "VisualShaderNodeFresnel"
}



// Enums

func (me *VisualShaderNodeFresnel) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFresnel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFresnel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
