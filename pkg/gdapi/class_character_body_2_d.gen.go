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

type CharacterBody2D struct {
  PhysicsBody2D
}

func (me *CharacterBody2D) BaseClass() string {
  return "CharacterBody2D"
}

func NewCharacterBody2D() *CharacterBody2D {
  str := StringNameFromStr("CharacterBody2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CharacterBody2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) ApplyFloorSnap()  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_floor_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) SetVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) SetSafeMargin(margin float64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetSafeMargin() float64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_safe_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsFloorStopOnSlopeEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetFloorStopOnSlopeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_stop_on_slope_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) SetFloorConstantSpeedEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) IsFloorConstantSpeedEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_constant_speed_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetFloorBlockOnWallEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) IsFloorBlockOnWallEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_floor_block_on_wall_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetSlideOnCeilingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) IsSlideOnCeilingEnabled() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slide_on_ceiling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetPlatformFloorLayers(exclude_layer int64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetPlatformFloorLayers() int64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_floor_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetPlatformWallLayers(exclude_layer int64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetPlatformWallLayers() int64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_wall_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) GetMaxSlides() int64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetMaxSlides(max_slides int64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_slides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_slides) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetFloorMaxAngle() float64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetFloorMaxAngle(radians float64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_max_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetFloorSnapLength() float64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetFloorSnapLength(floor_snap_length float64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_floor_snap_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&floor_snap_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetWallMinSlideAngle() float64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) SetWallMinSlideAngle(radians float64, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wall_min_slide_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetUpDirection() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) SetUpDirection(up_direction Vector2, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_up_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) SetMotionMode(mode CharacterBody2DMotionMode, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1224392233) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetMotionMode() CharacterBody2DMotionMode {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1160151236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CharacterBody2DMotionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CharacterBody2D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody2DPlatformOnLeave, )  {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2423324375) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_leave_apply_velocity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharacterBody2D) GetPlatformOnLeave() CharacterBody2DPlatformOnLeave {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_on_leave")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4054324341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CharacterBody2DPlatformOnLeave

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CharacterBody2D) IsOnFloor() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsOnFloorOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_floor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsOnCeiling() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsOnCeilingOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_ceiling_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsOnWall() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) IsOnWallOnly() bool {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_wall_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) GetFloorNormal() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetWallNormal() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wall_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetLastMotion() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetPositionDelta() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetRealVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_real_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetFloorAngle(up_direction Vector2, ) float64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_floor_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841063350) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) GetPlatformVelocity() Vector2 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_platform_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetSlideCollisionCount() int64 {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharacterBody2D) GetSlideCollision(slide_idx int64, ) KinematicCollision2D {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 860659811) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slide_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewKinematicCollision2D()
  pinner.Pin(&slide_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharacterBody2D) GetLastSlideCollision() KinematicCollision2D {
  classNameV := StringNameFromStr("CharacterBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_slide_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2161834755) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewKinematicCollision2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
