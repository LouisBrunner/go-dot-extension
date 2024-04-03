// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MarginContainer struct {
  Container
}

func (me *MarginContainer) BaseClass() string {
  return "MarginContainer"
}

func NewMarginContainer() *MarginContainer {
  str := StringNameFromStr("MarginContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MarginContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MarginContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MarginContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MarginContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
