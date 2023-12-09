// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VSplitContainer struct {
  obj gdc.ObjectPtr
}

func (me *VSplitContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VSplitContainer) BaseClass() string {
  return "VSplitContainer"
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

// TODO: properties (class)

// TODO: signals (class)
