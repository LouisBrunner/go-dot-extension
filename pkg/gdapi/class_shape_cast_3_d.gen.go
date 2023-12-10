// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ShapeCast3D struct {
  obj gdc.ObjectPtr
}

func (me *ShapeCast3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShapeCast3D) BaseClass() string {
  return "ShapeCast3D"
}



// Enums

func (me *ShapeCast3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShapeCast3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShapeCast3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ShapeCast3D) ResourceChanged(resource Resource, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resource_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) IsEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetShape(shape Shape3D, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549710052) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetShape() Shape3D {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3214262478) // FIXME: should cache?
  var ret Shape3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetTargetPosition(local_point Vector3, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetTargetPosition() Vector3 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetMargin() float32 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetMaxResults(max_results int, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_results")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetMaxResults() int {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_results")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) IsColliding() bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_colliding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetCollisionCount() int {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) ForceShapecastUpdate()  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_shapecast_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetCollider(index int, ) Object {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetColliderRid(index int, ) RID {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetColliderShape(index int, ) int {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetCollisionPoint(index int, ) Vector3 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetCollisionNormal(index int, ) Vector3 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetClosestCollisionSafeFraction() float32 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_collision_safe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) GetClosestCollisionUnsafeFraction() float32 {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_collision_unsafe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) AddExceptionRid(rid RID, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_exception_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) AddException(node CollisionObject3D, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1976431078) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) RemoveExceptionRid(rid RID, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_exception_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) RemoveException(node CollisionObject3D, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1976431078) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) ClearExceptions()  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetExcludeParentBody(mask bool, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_parent_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetExcludeParentBody() bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_parent_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_areas_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShapeCast3D) SetDebugShapeCustomColor(debug_shape_custom_color Color, )  {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_shape_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(debug_shape_custom_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShapeCast3D) GetDebugShapeCustomColor() Color {
  classNameV := StringNameFromStr("ShapeCast3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_shape_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ShapeCast3D) GetPropEnabled() bool {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropEnabled(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropShape() Shape3D {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropShape(value Shape3D) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropExcludeParent() bool {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropExcludeParent(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropTargetPosition() Vector3 {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropTargetPosition(value Vector3) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropMargin() float32 {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropMargin(value float32) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropMaxResults() int {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropMaxResults(value int) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropCollisionResult() Array {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropCollisionResult(value Array) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropCollideWithAreas() bool {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropCollideWithAreas(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropCollideWithBodies() bool {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropCollideWithBodies(value bool) {
  panic("TODO: implement")
}

func (me *ShapeCast3D) GetPropDebugShapeCustomColor() Color {
  panic("TODO: implement")
}

func (me *ShapeCast3D) SetPropDebugShapeCustomColor(value Color) {
  panic("TODO: implement")
}