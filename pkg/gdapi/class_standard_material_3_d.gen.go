// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StandardMaterial3D struct {
  BaseMaterial3D
}

func (me *StandardMaterial3D) BaseClass() string {
  return "StandardMaterial3D"
}



// Enums

func (me *StandardMaterial3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StandardMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StandardMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
