// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CharacterBody3D struct {
  obj gdc.ObjectPtr
}

func (me *CharacterBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CharacterBody3D) BaseClass() string {
  return "CharacterBody3D"
}



// Enums

type CharacterBody3DMotionMode int
const (
  CharacterBody3DMotionModeMotionModeGrounded CharacterBody3DMotionMode = 0
  CharacterBody3DMotionModeMotionModeFloating CharacterBody3DMotionMode = 1
)

type CharacterBody3DPlatformOnLeave int
const (
  CharacterBody3DPlatformOnLeavePlatformOnLeaveAddVelocity CharacterBody3DPlatformOnLeave = 0
  CharacterBody3DPlatformOnLeavePlatformOnLeaveAddUpwardVelocity CharacterBody3DPlatformOnLeave = 1
  CharacterBody3DPlatformOnLeavePlatformOnLeaveDoNothing CharacterBody3DPlatformOnLeave = 2
)

func (me *CharacterBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CharacterBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharacterBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CharacterBody3D) MoveAndSlide() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_slide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) ApplyFloorSnap()  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_floor_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) SetVelocity(velocity Vector3, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetVelocity() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetSafeMargin(margin float32, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetSafeMargin() float32 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsFloorStopOnSlopeEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetFloorStopOnSlopeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) SetFloorConstantSpeedEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) IsFloorConstantSpeedEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetFloorBlockOnWallEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) IsFloorBlockOnWallEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetSlideOnCeilingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) IsSlideOnCeilingEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetPlatformFloorLayers(exclude_layer int, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetPlatformFloorLayers() int {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetPlatformWallLayers(exclude_layer int, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetPlatformWallLayers() int {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetMaxSlides() int {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetMaxSlides(max_slides int, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_slides), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetFloorMaxAngle() float32 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetFloorMaxAngle(radians float32, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetFloorSnapLength() float32 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetFloorSnapLength(floor_snap_length float32, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&floor_snap_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetWallMinSlideAngle() float32 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetWallMinSlideAngle(radians float32, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetUpDirection() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetUpDirection(up_direction Vector3, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(up_direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) SetMotionMode(mode CharacterBody3DMotionMode, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2690739026) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetMotionMode() CharacterBody3DMotionMode {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3529553604) // FIXME: should cache?
  var ret CharacterBody3DMotionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody3DPlatformOnLeave, )  {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1459986142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_leave_apply_velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharacterBody3D) GetPlatformOnLeave() CharacterBody3DPlatformOnLeave {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 996491171) // FIXME: should cache?
  var ret CharacterBody3DPlatformOnLeave
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnFloor() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnFloorOnly() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnCeiling() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnCeilingOnly() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnWall() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) IsOnWallOnly() bool {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetFloorNormal() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetWallNormal() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetLastMotion() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetPositionDelta() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetRealVelocity() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_real_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetFloorAngle(up_direction Vector3, ) float32 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2906300789) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(up_direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetPlatformVelocity() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetPlatformAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetSlideCollisionCount() int {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetSlideCollision(slide_idx int, ) KinematicCollision3D {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107003663) // FIXME: should cache?
  var ret KinematicCollision3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slide_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharacterBody3D) GetLastSlideCollision() KinematicCollision3D {
  classNameV := StringNameFromStr("CharacterBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 186875014) // FIXME: should cache?
  var ret KinematicCollision3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
