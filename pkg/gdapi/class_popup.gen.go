// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Popup struct {
  obj gdc.ObjectPtr
}

func (me *Popup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Popup) BaseClass() string {
  return "Popup"
}



// Enums

func (me *Popup) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Popup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Popup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
