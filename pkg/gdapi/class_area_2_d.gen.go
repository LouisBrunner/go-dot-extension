// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Area2D struct {
  obj gdc.ObjectPtr
}

func (me *Area2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area2D) BaseClass() string {
  return "Area2D"
}



// Enums

type Area2DSpaceOverride int
const (
  Area2DSpaceOverrideSpaceOverrideDisabled Area2DSpaceOverride = 0
  Area2DSpaceOverrideSpaceOverrideCombine Area2DSpaceOverride = 1
  Area2DSpaceOverrideSpaceOverrideCombineReplace Area2DSpaceOverride = 2
  Area2DSpaceOverrideSpaceOverrideReplace Area2DSpaceOverride = 3
  Area2DSpaceOverrideSpaceOverrideReplaceCombine Area2DSpaceOverride = 4
)

func (me *Area2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Area2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Area2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Area2D) SetGravitySpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2879900038) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetGravitySpaceOverrideMode() Area2DSpaceOverride {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3990256304) // FIXME: should cache?
  var ret Area2DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetGravityIsPoint(enable bool, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_is_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) IsGravityAPoint() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_gravity_a_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetGravityPointUnitDistance(distance_scale float32, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_point_unit_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetGravityPointUnitDistance() float32 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_point_unit_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetGravityPointCenter(center Vector2, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_point_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(center.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetGravityPointCenter() Vector2 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_point_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetGravityDirection(direction Vector2, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetGravityDirection() Vector2 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetGravity(gravity float32, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetGravity() float32 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetLinearDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2879900038) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetLinearDampSpaceOverrideMode() Area2DSpaceOverride {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3990256304) // FIXME: should cache?
  var ret Area2DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetAngularDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2879900038) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetAngularDampSpaceOverrideMode() Area2DSpaceOverride {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3990256304) // FIXME: should cache?
  var ret Area2DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetLinearDamp(linear_damp float32, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetLinearDamp() float32 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetAngularDamp(angular_damp float32, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetAngularDamp() float32 {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetPriority(priority int, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetPriority() int {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetMonitoring(enable bool, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_monitoring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) IsMonitoring() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_monitoring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetMonitorable(enable bool, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) IsMonitorable() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) GetOverlappingBodies() Node2D {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_overlapping_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Node2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) GetOverlappingAreas() Area2D {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_overlapping_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Area2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) HasOverlappingBodies() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_overlapping_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) HasOverlappingAreas() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_overlapping_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) OverlapsBody(body Node, ) bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("overlaps_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) OverlapsArea(area Node, ) bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("overlaps_area")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetAudioBusName(name StringName, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) GetAudioBusName() StringName {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area2D) SetAudioBusOverride(enable bool, )  {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_bus_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area2D) IsOverridingAudioBus() bool {
  classNameV := StringNameFromStr("Area2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_overriding_audio_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Area2D) GetPropMonitoring() bool {
  panic("TODO: implement")
}

func (me *Area2D) SetPropMonitoring(value bool) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropMonitorable() bool {
  panic("TODO: implement")
}

func (me *Area2D) SetPropMonitorable(value bool) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropPriority() int {
  panic("TODO: implement")
}

func (me *Area2D) SetPropPriority(value int) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravitySpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravitySpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravityPoint() bool {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravityPoint(value bool) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravityPointUnitDistance() float32 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravityPointUnitDistance(value float32) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravityPointCenter() Vector2 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravityPointCenter(value Vector2) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravityDirection() Vector2 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravityDirection(value Vector2) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropGravity() float32 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropGravity(value float32) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropLinearDampSpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area2D) SetPropLinearDampSpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropLinearDamp() float32 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropLinearDamp(value float32) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropAngularDampSpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area2D) SetPropAngularDampSpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropAngularDamp() float32 {
  panic("TODO: implement")
}

func (me *Area2D) SetPropAngularDamp(value float32) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropAudioBusOverride() bool {
  panic("TODO: implement")
}

func (me *Area2D) SetPropAudioBusOverride(value bool) {
  panic("TODO: implement")
}

func (me *Area2D) GetPropAudioBusName() StringName {
  panic("TODO: implement")
}

func (me *Area2D) SetPropAudioBusName(value StringName) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API