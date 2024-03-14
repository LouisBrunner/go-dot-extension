// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeScreenNormalWorldSpace struct {
  VisualShaderNode
}

func (me *VisualShaderNodeScreenNormalWorldSpace) BaseClass() string {
  return "VisualShaderNodeScreenNormalWorldSpace"
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
