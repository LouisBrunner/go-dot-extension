// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeLinearSceneDepth struct {
  VisualShaderNode
}

func (me *VisualShaderNodeLinearSceneDepth) BaseClass() string {
  return "VisualShaderNodeLinearSceneDepth"
}



// Enums

func (me *VisualShaderNodeLinearSceneDepth) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeLinearSceneDepth) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeLinearSceneDepth) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
