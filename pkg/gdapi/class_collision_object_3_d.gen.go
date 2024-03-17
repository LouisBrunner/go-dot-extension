// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CollisionObject3D struct {
  Node3D
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

func  (me *CollisionObject3D) SetCollisionLayer(layer int, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCollisionLayer() int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCollisionLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetCollisionPriority(priority float32, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCollisionPriority() float32 {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetDisableMode(mode CollisionObject3DDisableMode, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1623620376) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetDisableMode() CollisionObject3DDisableMode {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410164780) // FIXME: should cache?
  var ret CollisionObject3DDisableMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetRayPickable(ray_pickable bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_pickable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) IsRayPickable() bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) SetCaptureInputOnDrag(enable bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_capture_input_on_drag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetCaptureInputOnDrag() bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_capture_input_on_drag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) GetRid() RID {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) CreateShapeOwner(owner Object, ) int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3429307534) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(owner.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) RemoveShapeOwner(owner_id int, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) GetShapeOwners() PackedInt32Array {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_owners")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerSetTransform(owner_id int, transform Transform3D, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) ShapeOwnerGetTransform(owner_id int, ) Transform3D {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerGetOwner(owner_id int, ) Object {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerSetDisabled(owner_id int, disabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) IsShapeOwnerDisabled(owner_id int, ) bool {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_owner_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerAddShape(owner_id int, shape Shape3D, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2566676345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) ShapeOwnerGetShapeCount(owner_id int, ) int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerGetShape(owner_id int, shape_id int, ) Shape3D {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015519174) // FIXME: should cache?
  var ret Shape3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerGetShapeIndex(owner_id int, shape_id int, ) int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject3D) ShapeOwnerRemoveShape(owner_id int, shape_id int, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) ShapeOwnerClearShapes(owner_id int, )  {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject3D) ShapeFindOwner(shape_index int, ) int {
  classNameV := StringNameFromStr("CollisionObject3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_find_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
