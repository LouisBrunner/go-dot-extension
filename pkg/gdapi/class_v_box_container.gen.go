// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VBoxContainer struct {
  BoxContainer
}

func (me *VBoxContainer) BaseClass() string {
  return "VBoxContainer"
}



// Enums

func (me *VBoxContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VBoxContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VBoxContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
