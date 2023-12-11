// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CheckButton struct {
  obj gdc.ObjectPtr
}

func (me *CheckButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CheckButton) BaseClass() string {
  return "CheckButton"
}



// Enums

func (me *CheckButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CheckButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CheckButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
