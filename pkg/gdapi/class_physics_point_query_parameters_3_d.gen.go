// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters3D) BaseClass() string {
  return "PhysicsPointQueryParameters3D"
}



// Enums

func (me *PhysicsPointQueryParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsPointQueryParameters3D) SetPosition(position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) GetPosition()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) SetCollisionMask(collision_mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) SetExclude(exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) GetExclude()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters3D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
