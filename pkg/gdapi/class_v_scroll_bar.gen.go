// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VScrollBar struct {
  obj gdc.ObjectPtr
}

func (me *VScrollBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VScrollBar) BaseClass() string {
  return "VScrollBar"
}



// Enums

func (me *VScrollBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VScrollBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VScrollBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties