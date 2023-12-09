// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionObject3D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionObject3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionObject3D) BaseClass() string {
  return "CollisionObject3D"
}



// Enums

type CollisionObject3DDisableMode int
const (
  CollisionObject3DDisableModeDisableModeRemove CollisionObject3DDisableMode = 0
  CollisionObject3DDisableModeDisableModeMakeStatic CollisionObject3DDisableMode = 1
  CollisionObject3DDisableModeDisableModeKeepActive CollisionObject3DDisableMode = 2
)

func (me *CollisionObject3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionObject3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionObject3D) XInputEvent(camera Camera3D, event InputEvent, position Vector3, normal Vector3, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) XMouseEnter()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) XMouseExit()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCollisionLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCollisionLayer()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCollisionLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCollisionPriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCollisionPriority()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetDisableMode(mode CollisionObject3DDisableMode, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetDisableMode()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetRayPickable(ray_pickable bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) IsRayPickable()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) SetCaptureInputOnDrag(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetCaptureInputOnDrag()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetRid()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) CreateShapeOwner(owner Object, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) RemoveShapeOwner(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) GetShapeOwners()  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerSetTransform(owner_id int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerGetTransform(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerGetOwner(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerSetDisabled(owner_id int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) IsShapeOwnerDisabled(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerAddShape(owner_id int, shape Shape3D, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerGetShapeCount(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerGetShape(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerGetShapeIndex(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerRemoveShape(owner_id int, shape_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeOwnerClearShapes(owner_id int, )  {
  panic("TODO: implement")
}

func  (me *CollisionObject3D) ShapeFindOwner(shape_index int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
