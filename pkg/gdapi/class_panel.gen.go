// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Panel struct {
  Control
}

func (me *Panel) BaseClass() string {
  return "Panel"
}

func NewPanel() *Panel {
  str := StringNameFromStr("Panel") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Panel{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Panel) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Panel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Panel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
