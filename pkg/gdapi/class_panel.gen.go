// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Panel struct {
  obj gdc.ObjectPtr
}

func (me *Panel) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Panel) BaseClass() string {
  return "Panel"
}



// Enums

func (me *Panel) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Panel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Panel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
