// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CheckBox struct {
  Button
}

func (me *CheckBox) BaseClass() string {
  return "CheckBox"
}



// Enums

func (me *CheckBox) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CheckBox) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CheckBox) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
