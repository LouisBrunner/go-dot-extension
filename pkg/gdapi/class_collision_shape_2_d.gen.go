// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionShape2D) BaseClass() string {
  return "CollisionShape2D"
}



// Enums

func (me *CollisionShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionShape2D) SetShape(shape Shape2D, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) GetShape()  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) SetDisabled(disabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) IsDisabled()  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) SetOneWayCollision(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) IsOneWayCollisionEnabled()  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) SetOneWayCollisionMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) GetOneWayCollisionMargin()  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) SetDebugColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape2D) GetDebugColor()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
