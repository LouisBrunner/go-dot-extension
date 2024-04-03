// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HSplitContainer struct {
  SplitContainer
}

func (me *HSplitContainer) BaseClass() string {
  return "HSplitContainer"
}

func NewHSplitContainer() *HSplitContainer {
  str := StringNameFromStr("HSplitContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HSplitContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *HSplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HSplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HSplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
