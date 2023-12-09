// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type KinematicCollision2D struct {
  obj gdc.ObjectPtr
}

func (me *KinematicCollision2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *KinematicCollision2D) BaseClass() string {
  return "KinematicCollision2D"
}



// Enums

func (me *KinematicCollision2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *KinematicCollision2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *KinematicCollision2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *KinematicCollision2D) GetPosition()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetNormal()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetTravel()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetRemainder()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetAngle(up_direction Vector2, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetDepth()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetLocalShape()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetCollider()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetColliderId()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetColliderRid()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetColliderShape()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetColliderShapeIndex()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision2D) GetColliderVelocity()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
