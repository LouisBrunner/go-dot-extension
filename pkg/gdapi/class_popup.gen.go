// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Popup struct {
  obj gdc.ObjectPtr
}

func (me *Popup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Popup) BaseClass() string {
  return "Popup"
}



// Enums

func (me *Popup) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Popup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Popup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API