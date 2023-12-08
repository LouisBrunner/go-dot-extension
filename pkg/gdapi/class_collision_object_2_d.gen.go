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

type CollisionObject2DDisableMode int
const (
  CollisionObject2DDisableModeDisableModeRemove CollisionObject2DDisableMode = 0
  CollisionObject2DDisableModeDisableModeMakeStatic CollisionObject2DDisableMode = 1
  CollisionObject2DDisableModeDisableModeKeepActive CollisionObject2DDisableMode = 2
)

func  (me *CollisionObject2D) XInputEvent(viewport Viewport, event InputEvent, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) XMouseEnter() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) XMouseExit() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) XMouseShapeEnter(shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) XMouseShapeExit(shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetRid() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetCollisionLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetCollisionLayer() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetCollisionMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetCollisionMask() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetCollisionLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetCollisionLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetCollisionMaskValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetCollisionMaskValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetCollisionPriority(priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetCollisionPriority() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetDisableMode(mode CollisionObject2DDisableMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetDisableMode() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) SetPickable(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) IsPickable() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) CreateShapeOwner(owner Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) RemoveShapeOwner(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetShapeOwners() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerSetTransform(owner_id int, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerGetTransform(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerGetOwner(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerSetDisabled(owner_id int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) IsShapeOwnerDisabled(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollision(owner_id int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(owner_id int, margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerAddShape(owner_id int, shape Shape2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeCount(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerGetShape(owner_id int, shape_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeIndex(owner_id int, shape_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerRemoveShape(owner_id int, shape_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeOwnerClearShapes(owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionObject2D) ShapeFindOwner(shape_index int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
