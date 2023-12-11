// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StyleBoxEmpty struct {
  obj gdc.ObjectPtr
}

func (me *StyleBoxEmpty) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StyleBoxEmpty) BaseClass() string {
  return "StyleBoxEmpty"
}



// Enums

func (me *StyleBoxEmpty) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StyleBoxEmpty) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StyleBoxEmpty) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
