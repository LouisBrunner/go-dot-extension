// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldBoundaryShape2D struct {
  obj gdc.ObjectPtr
}

func (me *WorldBoundaryShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldBoundaryShape2D) BaseClass() string {
  return "WorldBoundaryShape2D"
}



// Enums

func (me *WorldBoundaryShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WorldBoundaryShape2D) SetNormal(normal Vector2, )  {
  panic("TODO: implement")
}

func  (me *WorldBoundaryShape2D) GetNormal()  {
  panic("TODO: implement")
}

func  (me *WorldBoundaryShape2D) SetDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *WorldBoundaryShape2D) GetDistance()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
