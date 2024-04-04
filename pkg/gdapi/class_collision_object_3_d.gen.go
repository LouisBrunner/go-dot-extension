// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type CollisionObject3D struct {
  Node3D
}

func (me *CollisionObject3D) BaseClass() string {
  return "CollisionObject3D"
}

func NewCollisionObject3D() *CollisionObject3D {
  str := StringNameFromStr("CollisionObject3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CollisionObject3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CollisionObject3DDisableMode int
const (
  CollisionObject3DDisableModeDisableModeRemove CollisionObject3DDisableMode = 0
  CollisionObject3DDisableModeDisableModeMakeStatic CollisionObject3DDisableMode = 1
  CollisionObject3DDisableModeDisableModeKeepActive CollisionObject3DDisableMode = 2
)

func (me *CollisionObject3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionObject3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionObject3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CollisionObject3D) SetCollisionLayer(layer int64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCollisionLayer() int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) SetCollisionMask(mask int64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCollisionMask() int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) SetCollisionLayerValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCollisionLayerValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCollisionMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) SetCollisionPriority(priority float64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCollisionPriority() float64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) SetDisableMode(mode CollisionObject3DDisableMode, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1623620376) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetDisableMode() CollisionObject3DDisableMode {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410164780) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CollisionObject3DDisableMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CollisionObject3D) SetRayPickable(ray_pickable bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_pickable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) IsRayPickable() bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject3D) SetCaptureInputOnDrag(enable bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_capture_input_on_drag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetCaptureInputOnDrag() bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_capture_input_on_drag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionObject3D) GetRid() RID {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) CreateShapeOwner(owner Object, ) int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) RemoveShapeOwner(owner_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) GetShapeOwners() PackedInt32Array {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) ShapeOwnerSetTransform(owner_id int64, transform Transform3D, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) ShapeOwnerGetTransform(owner_id int64, ) Transform3D {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&owner_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject3D) ShapeOwnerGetOwner(owner_id int64, ) Object {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) ShapeOwnerSetDisabled(owner_id int64, disabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) IsShapeOwnerDisabled(owner_id int64, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) ShapeOwnerAddShape(owner_id int64, shape Shape3D, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2566676345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) ShapeOwnerGetShapeCount(owner_id int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) ShapeOwnerGetShape(owner_id int64, shape_id int64, ) Shape3D {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015519174) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&shape_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape3D()
  pinner.Pin(&owner_id)
  pinner.Pin(&shape_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionObject3D) ShapeOwnerGetShapeIndex(owner_id int64, shape_id int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

func  (me *CollisionObject3D) ShapeOwnerRemoveShape(owner_id int64, shape_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , gdc.ConstTypePtr(&shape_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) ShapeOwnerClearShapes(owner_id int64, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionObject3D) ShapeFindOwner(shape_index int64, ) int64 {
  classNameV := StringNameFromStr("CollisionObject3D")
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

type CollisionObject3DInputEventSignalFn func(camera Node, event InputEvent, position Vector3, normal Vector3, shape_idx int, )

func (me *CollisionObject3D) ConnectInputEvent(subs SignalSubscribers, fn CollisionObject3DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectInputEvent(subs SignalSubscribers, fn CollisionObject3DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject3DMouseEnteredSignalFn func()

func (me *CollisionObject3D) ConnectMouseEntered(subs SignalSubscribers, fn CollisionObject3DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectMouseEntered(subs SignalSubscribers, fn CollisionObject3DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject3DMouseExitedSignalFn func()

func (me *CollisionObject3D) ConnectMouseExited(subs SignalSubscribers, fn CollisionObject3DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectMouseExited(subs SignalSubscribers, fn CollisionObject3DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
