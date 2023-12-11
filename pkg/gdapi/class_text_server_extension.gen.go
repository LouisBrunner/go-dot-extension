// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerExtension struct {
  obj gdc.ObjectPtr
}

func (me *TextServerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServerExtension) BaseClass() string {
  return "TextServerExtension"
}



// Enums

func (me *TextServerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
