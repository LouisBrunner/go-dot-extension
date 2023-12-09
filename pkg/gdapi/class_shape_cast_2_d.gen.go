// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShapeCast2D struct {
  obj gdc.ObjectPtr
}

func (me *ShapeCast2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShapeCast2D) BaseClass() string {
  return "ShapeCast2D"
}



// Enums

func (me *ShapeCast2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShapeCast2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShapeCast2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ShapeCast2D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetShape(shape Shape2D, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetShape()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetTargetPosition(local_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetMargin()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetMaxResults(max_results int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetMaxResults()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) IsColliding()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollisionCount()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) ForceShapecastUpdate()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollider(index int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetColliderRid(index int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetColliderShape(index int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollisionPoint(index int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollisionNormal(index int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetClosestCollisionSafeFraction()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetClosestCollisionUnsafeFraction()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) AddExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) AddException(node CollisionObject2D, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) RemoveExceptionRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) RemoveException(node CollisionObject2D, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) ClearExceptions()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetExcludeParentBody(mask bool, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) GetExcludeParentBody()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetCollideWithAreas(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) IsCollideWithAreasEnabled()  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) SetCollideWithBodies(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ShapeCast2D) IsCollideWithBodiesEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
