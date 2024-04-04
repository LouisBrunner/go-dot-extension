// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type CollisionObject2D struct {
  Node2D
}

func (me *CollisionObject2D) BaseClass() string {
  return "CollisionObject2D"
}

func NewCollisionObject2D() *CollisionObject2D {
  str := StringNameFromStr("CollisionObject2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CollisionObject2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CollisionObject2DDisableMode int
const (
  CollisionObject2DDisableModeDisableModeRemove CollisionObject2DDisableMode = 0
  CollisionObject2DDisableModeDisableModeMakeStatic CollisionObject2DDisableMode = 1
  CollisionObject2DDisableModeDisableModeKeepActive CollisionObject2DDisableMode = 2
)

func (me *CollisionObject2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionObject2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionObject2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CollisionObject2D) GetRid() RID {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject2D) SetCollisionLayer(layer int64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetCollisionLayer() int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) SetCollisionMask(mask int64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetCollisionMask() int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) SetCollisionLayerValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetCollisionLayerValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetCollisionMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) SetCollisionPriority(priority float64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetCollisionPriority() float64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) SetDisableMode(mode CollisionObject2DDisableMode, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1919204045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetDisableMode() CollisionObject2DDisableMode {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3172846349) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CollisionObject2DDisableMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CollisionObject2D) SetPickable(enabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) IsPickable() bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) CreateShapeOwner(owner Object, ) int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3429307534) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{owner.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) RemoveShapeOwner(owner_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetShapeOwners() PackedInt32Array {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_owners")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject2D) ShapeOwnerSetTransform(owner_id int64, transform Transform2D, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30160968) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) ShapeOwnerGetTransform(owner_id int64, ) Transform2D {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3836996910) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject2D) ShapeOwnerGetOwner(owner_id int64, ) Object {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject2D) ShapeOwnerSetDisabled(owner_id int64, disabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) IsShapeOwnerDisabled(owner_id int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_owner_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollision(owner_id int64, enable bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_one_way_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(owner_id int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_owner_one_way_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(owner_id int64, margin float64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(owner_id int64, ) float64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_owner_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) ShapeOwnerAddShape(owner_id int64, shape Shape2D, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2077425081) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) ShapeOwnerGetShapeCount(owner_id int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) ShapeOwnerGetShape(owner_id int64, shape_id int64, ) Shape2D {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3106725749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&shape_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape2D()
  pinner.Pin(&owner_id)
  pinner.Pin(&shape_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeIndex(owner_id int64, shape_id int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&shape_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&owner_id)
  pinner.Pin(&shape_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject2D) ShapeOwnerRemoveShape(owner_id int64, shape_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&shape_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) ShapeOwnerClearShapes(owner_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject2D) ShapeFindOwner(shape_index int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_find_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&shape_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CollisionObject2DInputEventSignalFn func(viewport Node, event InputEvent, shape_idx int, )

func (me *CollisionObject2D) ConnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseEnteredSignalFn func()

func (me *CollisionObject2D) ConnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseExitedSignalFn func()

func (me *CollisionObject2D) ConnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeEnteredSignalFn func(shape_idx int, )

func (me *CollisionObject2D) ConnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
  sig := StringNameFromStr("mouse_shape_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
  sig := StringNameFromStr("mouse_shape_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeExitedSignalFn func(shape_idx int, )

func (me *CollisionObject2D) ConnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
  sig := StringNameFromStr("mouse_shape_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
  sig := StringNameFromStr("mouse_shape_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
