// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *CharacterBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharacterBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CharacterBody2D) MoveAndSlide()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) ApplyFloorSnap()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetVelocity(velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetSafeMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetSafeMargin()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsFloorStopOnSlopeEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetFloorStopOnSlopeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetFloorConstantSpeedEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsFloorConstantSpeedEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetFloorBlockOnWallEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsFloorBlockOnWallEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetSlideOnCeilingEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsSlideOnCeilingEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetPlatformFloorLayers(exclude_layer int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetPlatformFloorLayers()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetPlatformWallLayers(exclude_layer int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetPlatformWallLayers()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetMaxSlides()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetMaxSlides(max_slides int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetFloorMaxAngle()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetFloorMaxAngle(radians float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetFloorSnapLength()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetFloorSnapLength(floor_snap_length float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetWallMinSlideAngle()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetWallMinSlideAngle(radians float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetUpDirection()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetUpDirection(up_direction Vector2, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetMotionMode(mode CharacterBody2DMotionMode, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetMotionMode()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody2DPlatformOnLeave, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetPlatformOnLeave()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnFloor()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnFloorOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnCeiling()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnCeilingOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnWall()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) IsOnWallOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetFloorNormal()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetWallNormal()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetLastMotion()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetPositionDelta()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetRealVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetFloorAngle(up_direction Vector2, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetPlatformVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetSlideCollisionCount()  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetSlideCollision(slide_idx int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody2D) GetLastSlideCollision()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
