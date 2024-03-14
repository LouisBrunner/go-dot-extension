// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SpotLight3D struct {
  Light3D
}

func (me *SpotLight3D) BaseClass() string {
  return "SpotLight3D"
}



// Enums

func (me *SpotLight3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpotLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpotLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
