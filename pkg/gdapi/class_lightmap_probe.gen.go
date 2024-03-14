// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LightmapProbe struct {
  Node3D
}

func (me *LightmapProbe) BaseClass() string {
  return "LightmapProbe"
}



// Enums

func (me *LightmapProbe) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LightmapProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
