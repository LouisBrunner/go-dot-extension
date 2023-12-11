// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetCollisionLayer(layer int, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetCollisionLayer() int {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetCollisionMask() int {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetCollisionLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetCollisionLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetCollisionPriority(priority float32, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetCollisionPriority() float32 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetDisableMode(mode CollisionObject2DDisableMode, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1919204045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetDisableMode() CollisionObject2DDisableMode {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3172846349) // FIXME: should cache?
  var ret CollisionObject2DDisableMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) SetPickable(enabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) IsPickable() bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) CreateShapeOwner(owner Object, ) int {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3429307534) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(owner.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) RemoveShapeOwner(owner_id int, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_shape_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetShapeOwners() PackedInt32Array {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_owners")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerSetTransform(owner_id int, transform Transform2D, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30160968) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) ShapeOwnerGetTransform(owner_id int, ) Transform2D {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3836996910) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerGetOwner(owner_id int, ) Object {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerSetDisabled(owner_id int, disabled bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) IsShapeOwnerDisabled(owner_id int, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_owner_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollision(owner_id int, enable bool, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_one_way_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(owner_id int, ) bool {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_owner_one_way_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(owner_id int, margin float32, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_set_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(owner_id int, ) float32 {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_owner_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerAddShape(owner_id int, shape Shape2D, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2077425081) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeCount(owner_id int, ) int {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerGetShape(owner_id int, shape_id int, ) Shape2D {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3106725749) // FIXME: should cache?
  var ret Shape2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerGetShapeIndex(owner_id int, shape_id int, ) int {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_get_shape_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionObject2D) ShapeOwnerRemoveShape(owner_id int, shape_id int, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) ShapeOwnerClearShapes(owner_id int, )  {
  classNameV := StringNameFromStr("CollisionObject2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_owner_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionObject2D) ShapeFindOwner(shape_index int, ) int {
  classNameV := StringNameFromStr("CollisionObject2D")
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

type CollisionObject2DInputEventSignalFn func(viewport Node, event InputEvent, shape_idx int, )

func (me *CollisionObject2D) ConnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
  sig := StringNameFromStr("input_event")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CollisionObject2DMouseEnteredSignalFn func()

func (me *CollisionObject2D) ConnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
  sig := StringNameFromStr("mouse_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CollisionObject2DMouseExitedSignalFn func()

func (me *CollisionObject2D) ConnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
  sig := StringNameFromStr("mouse_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeEnteredSignalFn func(shape_idx int, )

func (me *CollisionObject2D) ConnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
  sig := StringNameFromStr("mouse_shape_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
  sig := StringNameFromStr("mouse_shape_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeExitedSignalFn func(shape_idx int, )

func (me *CollisionObject2D) ConnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
  sig := StringNameFromStr("mouse_shape_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
  sig := StringNameFromStr("mouse_shape_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
