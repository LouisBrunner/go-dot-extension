// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World2D struct {
  obj gdc.ObjectPtr
}

func (me *World2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World2D) BaseClass() string {
  return "World2D"
}



// Enums

func (me *World2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *World2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *World2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *World2D) GetCanvas()  {
  panic("TODO: implement")
}

func  (me *World2D) GetSpace()  {
  panic("TODO: implement")
}

func  (me *World2D) GetNavigationMap()  {
  panic("TODO: implement")
}

func  (me *World2D) GetDirectSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
