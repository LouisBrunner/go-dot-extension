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

type HBoxContainer struct {
  BoxContainer
}

func (me *HBoxContainer) BaseClass() string {
  return "HBoxContainer"
}

func NewHBoxContainer() *HBoxContainer {
  str := StringNameFromStr("HBoxContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HBoxContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *HBoxContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HBoxContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HBoxContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
