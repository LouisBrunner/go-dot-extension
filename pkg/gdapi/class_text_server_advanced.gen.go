// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerAdvanced struct {
  obj gdc.ObjectPtr
}

func (me *TextServerAdvanced) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServerAdvanced) BaseClass() string {
  return "TextServerAdvanced"
}



// Enums

func (me *TextServerAdvanced) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerAdvanced) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerAdvanced) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
