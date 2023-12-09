// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorRect struct {
  obj gdc.ObjectPtr
}

func (me *ColorRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ColorRect) BaseClass() string {
  return "ColorRect"
}



// Enums

func (me *ColorRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ColorRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ColorRect) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorRect) GetColor()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
