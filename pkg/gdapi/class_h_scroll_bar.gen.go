// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HScrollBar struct {
  obj gdc.ObjectPtr
}

func (me *HScrollBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HScrollBar) BaseClass() string {
  return "HScrollBar"
}



// Enums

func (me *HScrollBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HScrollBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HScrollBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
