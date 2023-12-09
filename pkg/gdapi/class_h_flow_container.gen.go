// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HFlowContainer struct {
  obj gdc.ObjectPtr
}

func (me *HFlowContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HFlowContainer) BaseClass() string {
  return "HFlowContainer"
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

// TODO: properties (class)

// TODO: signals (class)
