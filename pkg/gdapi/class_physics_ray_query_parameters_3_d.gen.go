// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsRayQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsRayQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsRayQueryParameters3D) BaseClass() string {
  return "PhysicsRayQueryParameters3D"
}



// Enums

func (me *PhysicsRayQueryParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsRayQueryParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsRayQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  PhysicsRayQueryParameters3DCreate(from Vector3, to Vector3, collision_mask int, exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetFrom(from Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) GetFrom()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetTo(to Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) GetTo()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetCollisionMask(collision_mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetExclude(exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) GetExclude()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetHitFromInside(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) IsHitFromInsideEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) SetHitBackFaces(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsRayQueryParameters3D) IsHitBackFacesEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
