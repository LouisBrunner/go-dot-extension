// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CubemapArray struct {
  obj gdc.ObjectPtr
}

func (me *CubemapArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CubemapArray) BaseClass() string {
  return "CubemapArray"
}



// Enums

func (me *CubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CubemapArray) CreatePlaceholder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
