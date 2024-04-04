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
