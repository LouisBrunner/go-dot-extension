// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StyleBoxEmpty struct {
  StyleBox
}

func (me *StyleBoxEmpty) BaseClass() string {
  return "StyleBoxEmpty"
}

func NewStyleBoxEmpty() *StyleBoxEmpty {
  str := StringNameFromStr("StyleBoxEmpty") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StyleBoxEmpty{}
  obj.SetBaseObject(objPtr)
  return obj
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
