// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type KinematicCollision3D struct {
  obj gdc.ObjectPtr
}

func (me *KinematicCollision3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *KinematicCollision3D) BaseClass() string {
  return "KinematicCollision3D"
}



// Enums

func (me *KinematicCollision3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *KinematicCollision3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *KinematicCollision3D) GetTravel()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetRemainder()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetDepth()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetCollisionCount()  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetPosition(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetNormal(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetAngle(collision_index int, up_direction Vector3, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetLocalShape(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetCollider(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetColliderId(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetColliderRid(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetColliderShape(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetColliderShapeIndex(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *KinematicCollision3D) GetColliderVelocity(collision_index int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
