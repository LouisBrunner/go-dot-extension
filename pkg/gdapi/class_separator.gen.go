// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Separator struct {
  Control
}

func (me *Separator) BaseClass() string {
  return "Separator"
}

func NewSeparator() *Separator {
  str := StringNameFromStr("Separator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Separator{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Separator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Separator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Separator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
