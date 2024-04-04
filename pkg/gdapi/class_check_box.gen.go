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

type CheckBox struct {
  Button
}

func (me *CheckBox) BaseClass() string {
  return "CheckBox"
}

func NewCheckBox() *CheckBox {
  str := StringNameFromStr("CheckBox") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CheckBox{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CheckBox) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CheckBox) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CheckBox) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
