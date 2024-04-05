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

type ptrsForVSplitContainerList struct {
}

var ptrsForVSplitContainer ptrsForVSplitContainerList

func initVSplitContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VSplitContainer")
  defer className.Destroy()
}

type VSplitContainer struct {
  SplitContainer
}

func (me *VSplitContainer) BaseClass() string {
  return "VSplitContainer"
}

func NewVSplitContainer() *VSplitContainer {
  str := StringNameFromStr("VSplitContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VSplitContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VSplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VSplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VSplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
