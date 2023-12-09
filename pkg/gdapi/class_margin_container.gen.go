// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MarginContainer struct {
  obj gdc.ObjectPtr
}

func (me *MarginContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MarginContainer) BaseClass() string {
  return "MarginContainer"
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

// TODO: properties (class)

// TODO: signals (class)
