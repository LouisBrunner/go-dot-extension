// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HFlowContainer struct {
  FlowContainer
}

func (me *HFlowContainer) BaseClass() string {
  return "HFlowContainer"
}

func NewHFlowContainer() *HFlowContainer {
  str := StringNameFromStr("HFlowContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HFlowContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *HFlowContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HFlowContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HFlowContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
