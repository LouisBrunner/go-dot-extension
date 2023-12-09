// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast3D struct {
  obj gdc.ObjectPtr
}

func (me *RayCast3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RayCast3D) BaseClass() string {
  return "RayCast3D"
}



// Enums

func (me *RayCast3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RayCast3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RayCast3D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetTargetPosition(local_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) IsColliding()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) ForceRaycastUpdate()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetCollider()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetColliderRid()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetColliderShape()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetCollisionPoint()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetCollisionNormal()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) AddExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) AddException(node CollisionObject3D, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) RemoveExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) RemoveException(node CollisionObject3D, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) ClearExceptions()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetExcludeParentBody(mask bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetExcludeParentBody()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetHitFromInside(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) IsHitFromInsideEnabled()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetDebugShapeCustomColor(debug_shape_custom_color Color, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetDebugShapeCustomColor()  {
  panic("TODO: implement")
}

func  (me *RayCast3D) SetDebugShapeThickness(debug_shape_thickness int, )  {
  panic("TODO: implement")
}

func  (me *RayCast3D) GetDebugShapeThickness()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
