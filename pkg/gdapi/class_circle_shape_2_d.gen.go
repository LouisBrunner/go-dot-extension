// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CircleShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CircleShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CircleShape2D) BaseClass() string {
  return "CircleShape2D"
}



// Enums

func (me *CircleShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CircleShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CircleShape2D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CircleShape2D) GetRadius()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
