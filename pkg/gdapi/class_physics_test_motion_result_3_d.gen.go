// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionResult3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionResult3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionResult3D) BaseClass() string {
  return "PhysicsTestMotionResult3D"
}



// Enums

func (me *PhysicsTestMotionResult3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionResult3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionResult3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsTestMotionResult3D) GetTravel()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetRemainder()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionSafeFraction()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionUnsafeFraction()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionCount()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionPoint(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionNormal(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetColliderVelocity(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetColliderId(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetColliderRid(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollider(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetColliderShape(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionLocalShape(collision_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionResult3D) GetCollisionDepth(collision_index int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
