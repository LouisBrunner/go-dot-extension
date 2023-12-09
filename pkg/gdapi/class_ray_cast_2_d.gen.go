// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast2D struct {
  obj gdc.ObjectPtr
}

func (me *RayCast2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RayCast2D) BaseClass() string {
  return "RayCast2D"
}



// Enums

func (me *RayCast2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RayCast2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RayCast2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RayCast2D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetTargetPosition(local_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) IsColliding()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) ForceRaycastUpdate()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetCollider()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetColliderRid()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetColliderShape()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetCollisionPoint()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetCollisionNormal()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) AddExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) AddException(node CollisionObject2D, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) RemoveExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) RemoveException(node CollisionObject2D, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) ClearExceptions()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetExcludeParentBody(mask bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) GetExcludeParentBody()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast2D) SetHitFromInside(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast2D) IsHitFromInsideEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
