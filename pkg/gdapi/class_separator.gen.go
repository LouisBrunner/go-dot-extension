// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

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
