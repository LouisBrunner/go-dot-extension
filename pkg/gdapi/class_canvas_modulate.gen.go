// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasModulate struct {
  obj gdc.ObjectPtr
}

func (me *CanvasModulate) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasModulate) BaseClass() string {
  return "CanvasModulate"
}



// Enums

func (me *CanvasModulate) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasModulate) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CanvasModulate) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *CanvasModulate) GetColor()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
