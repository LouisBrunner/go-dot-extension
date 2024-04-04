// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VScrollBar struct {
  ScrollBar
}

func (me *VScrollBar) BaseClass() string {
  return "VScrollBar"
}

func NewVScrollBar() *VScrollBar {
  str := StringNameFromStr("VScrollBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VScrollBar{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
