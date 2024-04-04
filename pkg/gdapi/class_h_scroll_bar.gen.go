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

type HScrollBar struct {
  ScrollBar
}

func (me *HScrollBar) BaseClass() string {
  return "HScrollBar"
}

func NewHScrollBar() *HScrollBar {
  str := StringNameFromStr("HScrollBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HScrollBar{}
  obj.SetBaseObject(objPtr)
  return obj
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
