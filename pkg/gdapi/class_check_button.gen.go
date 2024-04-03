// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CheckButton struct {
  Button
}

func (me *CheckButton) BaseClass() string {
  return "CheckButton"
}

func NewCheckButton() *CheckButton {
  str := StringNameFromStr("CheckButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CheckButton{}
  obj.SetBaseObject(objPtr)
  return obj
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
