// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeProximityFade struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeProximityFade) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeProximityFade) BaseClass() string {
  return "VisualShaderNodeProximityFade"
}



// Enums

func (me *VisualShaderNodeProximityFade) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeProximityFade) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeProximityFade) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
