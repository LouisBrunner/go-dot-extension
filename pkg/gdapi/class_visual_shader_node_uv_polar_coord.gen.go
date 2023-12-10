// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeUVPolarCoord struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUVPolarCoord) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUVPolarCoord) BaseClass() string {
  return "VisualShaderNodeUVPolarCoord"
}



// Enums

func (me *VisualShaderNodeUVPolarCoord) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUVPolarCoord) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUVPolarCoord) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties