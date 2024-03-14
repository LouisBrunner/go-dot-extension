// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTextureParameterTriplanar struct {
  VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTextureParameterTriplanar) BaseClass() string {
  return "VisualShaderNodeTextureParameterTriplanar"
}



// Enums

func (me *VisualShaderNodeTextureParameterTriplanar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureParameterTriplanar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureParameterTriplanar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
