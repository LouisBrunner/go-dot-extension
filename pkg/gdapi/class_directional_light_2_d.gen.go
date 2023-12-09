// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight2D struct {
  obj gdc.ObjectPtr
}

func (me *DirectionalLight2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirectionalLight2D) BaseClass() string {
  return "DirectionalLight2D"
}



// Enums

func (me *DirectionalLight2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DirectionalLight2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DirectionalLight2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *DirectionalLight2D) SetMaxDistance(pixels float32, )  {
  panic("TODO: implement")
}

func  (me *DirectionalLight2D) GetMaxDistance()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
