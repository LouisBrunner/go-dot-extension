// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters2D) BaseClass() string {
  return "PhysicsPointQueryParameters2D"
}



// Enums

func (me *PhysicsPointQueryParameters2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsPointQueryParameters2D) SetPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) GetPosition()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) SetCanvasInstanceId(canvas_instance_id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) GetCanvasInstanceId()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) SetCollisionMask(collision_mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) SetExclude(exclude RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) GetExclude()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsPointQueryParameters2D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
