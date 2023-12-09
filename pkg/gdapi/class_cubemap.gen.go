// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Cubemap struct {
  obj gdc.ObjectPtr
}

func (me *Cubemap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Cubemap) BaseClass() string {
  return "Cubemap"
}



// Enums

func (me *Cubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Cubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Cubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Cubemap) CreatePlaceholder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
