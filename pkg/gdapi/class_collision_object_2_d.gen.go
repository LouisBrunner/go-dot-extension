// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionObject2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionObject2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionObject2D) BaseClass() string {
  return "CollisionObject2D"
}



// Enums

type CollisionObject2DDisableMode int
const (
  CollisionObject2DDisableModeDisableModeRemove CollisionObject2DDisableMode = 0
  CollisionObject2DDisableModeDisableModeMakeStatic CollisionObject2DDisableMode = 1
  CollisionObject2DDisableModeDisableModeKeepActive CollisionObject2DDisableMode = 2
)

func (me *CollisionObject2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionObject2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionObject2D) XInputEvent(viewport Viewport, event InputEvent, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) XMouseEnter()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) XMouseExit()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) XMouseShapeEnter(shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) XMouseShapeExit(shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetRid()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetCollisionLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetCollisionLayer()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetCollisionLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetCollisionLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetCollisionPriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetCollisionPriority()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetDisableMode(mode CollisionObject2DDisableMode, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetDisableMode()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) SetPickable(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) IsPickable()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) CreateShapeOwner(owner Object, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) RemoveShapeOwner(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetShapeOwners()  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerSetTransform(owner_id int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerGetTransform(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerGetOwner(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerSetDisabled(owner_id int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) IsShapeOwnerDisabled(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollision(owner_id int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(owner_id int, margin float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerAddShape(owner_id int, shape Shape2D, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeCount(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerGetShape(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeIndex(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerRemoveShape(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeOwnerClearShapes(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject2D) ShapeFindOwner(shape_index int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
