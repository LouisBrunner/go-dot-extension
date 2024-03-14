// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerDummy struct {
  TextServerExtension
}

func (me *TextServerDummy) BaseClass() string {
  return "TextServerDummy"
}



// Enums

func (me *TextServerDummy) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerDummy) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerDummy) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
