// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsRayQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsRayQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsRayQueryParameters2D) BaseClass() string {
  return "PhysicsRayQueryParameters2D"
}



// Enums

func (me *PhysicsRayQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsRayQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  PhysicsRayQueryParameters2DCreate(from Vector2, to Vector2, collision_mask int, exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetFrom(from Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) GetFrom()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetTo(to Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) GetTo()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetCollisionMask(collision_mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetExclude(exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) GetExclude()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) SetHitFromInside(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters2D) IsHitFromInsideEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
