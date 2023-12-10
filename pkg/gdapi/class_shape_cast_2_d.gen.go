// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) IsEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetShape(shape Shape2D, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 771364740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetShape() Shape2D {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 522005891) // FIXME: should cache?
  var ret Shape2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetTargetPosition(local_point Vector2, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetTargetPosition() Vector2 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetMargin() float32 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetMaxResults(max_results int, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_results")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetMaxResults() int {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_results")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) IsColliding() bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_colliding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetCollisionCount() int {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) ForceShapecastUpdate()  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_shapecast_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetCollider(index int, ) Object {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetColliderRid(index int, ) RID {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetColliderShape(index int, ) int {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetCollisionPoint(index int, ) Vector2 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetCollisionNormal(index int, ) Vector2 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetClosestCollisionSafeFraction() float32 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_collision_safe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) GetClosestCollisionUnsafeFraction() float32 {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_collision_unsafe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) AddExceptionRid(rid RID, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_exception_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) AddException(node CollisionObject2D, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3090941106) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) RemoveExceptionRid(rid RID, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_exception_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) RemoveException(node CollisionObject2D, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3090941106) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) ClearExceptions()  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetCollisionMask() int {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetExcludeParentBody(mask bool, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_parent_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) GetExcludeParentBody() bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_parent_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_areas_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast2D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast2D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ShapeCast2D) GetPropEnabled() bool {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropEnabled(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropShape() Shape2D {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropShape(value Shape2D) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropExcludeParent() bool {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropExcludeParent(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropTargetPosition() Vector2 {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropTargetPosition(value Vector2) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropMargin() float32 {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropMargin(value float32) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropMaxResults() int {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropMaxResults(value int) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropCollisionResult() Array {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropCollisionResult(value Array) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropCollideWithAreas() bool {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropCollideWithAreas(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast2D) GetPropCollideWithBodies() bool {
  panic("TODO: implement")
}

func (me *ShapeCast2D) SetPropCollideWithBodies(value bool) {
  panic("TODO: implement")
}