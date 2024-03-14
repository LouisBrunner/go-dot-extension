// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ORMMaterial3D struct {
  BaseMaterial3D
}

func (me *ORMMaterial3D) BaseClass() string {
  return "ORMMaterial3D"
}



// Enums

func (me *ORMMaterial3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ORMMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ORMMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
