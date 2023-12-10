// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VFlowContainer struct {
  obj gdc.ObjectPtr
}

func (me *VFlowContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VFlowContainer) BaseClass() string {
  return "VFlowContainer"
}



// Enums

func (me *VFlowContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VFlowContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VFlowContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties