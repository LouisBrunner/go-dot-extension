// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VSeparator struct {
  Separator
}

func (me *VSeparator) BaseClass() string {
  return "VSeparator"
}

func NewVSeparator() *VSeparator {
  str := StringNameFromStr("VSeparator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VSeparator{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VSeparator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VSeparator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VSeparator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
