// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CharacterBody2D struct {
  obj gdc.ObjectPtr
}

func (me *CharacterBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CharacterBody2D) BaseClass() string {
  return "CharacterBody2D"
}



// Enums

type CharacterBody2DMotionMode int
const (
  CharacterBody2DMotionModeMotionModeGrounded CharacterBody2DMotionMode = 0
  CharacterBody2DMotionModeMotionModeFloating CharacterBody2DMotionMode = 1
)

type CharacterBody2DPlatformOnLeave int
const (
  CharacterBody2DPlatformOnLeavePlatformOnLeaveAddVelocity CharacterBody2DPlatformOnLeave = 0
  CharacterBody2DPlatformOnLeavePlatformOnLeaveAddUpwardVelocity CharacterBody2DPlatformOnLeave = 1
  CharacterBody2DPlatformOnLeavePlatformOnLeaveDoNothing CharacterBody2DPlatformOnLeave = 2
)

func (me *CharacterBody2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CharacterBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharacterBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CharacterBody2D) MoveAndSlide() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_slide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) ApplyFloorSnap()  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_floor_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) SetVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetSafeMargin(margin float32, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetSafeMargin() float32 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsFloorStopOnSlopeEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetFloorStopOnSlopeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) SetFloorConstantSpeedEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) IsFloorConstantSpeedEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetFloorBlockOnWallEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) IsFloorBlockOnWallEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetSlideOnCeilingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) IsSlideOnCeilingEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetPlatformFloorLayers(exclude_layer int, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetPlatformFloorLayers() int {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetPlatformWallLayers(exclude_layer int, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetPlatformWallLayers() int {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetMaxSlides() int {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetMaxSlides(max_slides int, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_slides), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetFloorMaxAngle() float32 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetFloorMaxAngle(radians float32, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetFloorSnapLength() float32 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetFloorSnapLength(floor_snap_length float32, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&floor_snap_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetWallMinSlideAngle() float32 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetWallMinSlideAngle(radians float32, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetUpDirection() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetUpDirection(up_direction Vector2, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(up_direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) SetMotionMode(mode CharacterBody2DMotionMode, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1224392233) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetMotionMode() CharacterBody2DMotionMode {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1160151236) // FIXME: should cache?
  var ret CharacterBody2DMotionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody2DPlatformOnLeave, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2423324375) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_leave_apply_velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody2D) GetPlatformOnLeave() CharacterBody2DPlatformOnLeave {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4054324341) // FIXME: should cache?
  var ret CharacterBody2DPlatformOnLeave
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnFloor() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnFloorOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnCeiling() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnCeilingOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnWall() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) IsOnWallOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetFloorNormal() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetWallNormal() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetLastMotion() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetPositionDelta() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetRealVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_real_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetFloorAngle(up_direction Vector2, ) float32 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841063350) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(up_direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetPlatformVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetSlideCollisionCount() int {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetSlideCollision(slide_idx int, ) KinematicCollision2D {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 860659811) // FIXME: should cache?
  var ret KinematicCollision2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slide_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody2D) GetLastSlideCollision() KinematicCollision2D {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2161834755) // FIXME: should cache?
  var ret KinematicCollision2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CharacterBody2D) GetPropMotionMode() int {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropMotionMode(value int) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropUpDirection() Vector2 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropUpDirection(value Vector2) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropVelocity() Vector2 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropVelocity(value Vector2) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropSlideOnCeiling() bool {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropSlideOnCeiling(value bool) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropMaxSlides() int {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropMaxSlides(value int) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropWallMinSlideAngle() float32 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropWallMinSlideAngle(value float32) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropFloorStopOnSlope() bool {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropFloorStopOnSlope(value bool) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropFloorConstantSpeed() bool {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropFloorConstantSpeed(value bool) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropFloorBlockOnWall() bool {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropFloorBlockOnWall(value bool) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropFloorMaxAngle() float32 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropFloorMaxAngle(value float32) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropFloorSnapLength() float32 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropFloorSnapLength(value float32) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropPlatformOnLeave() int {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropPlatformOnLeave(value int) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropPlatformFloorLayers() int {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropPlatformFloorLayers(value int) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropPlatformWallLayers() int {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropPlatformWallLayers(value int) {
  panic("TODO: implement")
}

func (me *CharacterBody2D) GetPropSafeMargin() float32 {
  panic("TODO: implement")
}

func (me *CharacterBody2D) SetPropSafeMargin(value float32) {
  panic("TODO: implement")
}