// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *CharacterBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharacterBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CharacterBody3D) MoveAndSlide()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) ApplyFloorSnap()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetSafeMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetSafeMargin()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsFloorStopOnSlopeEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetFloorStopOnSlopeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetFloorConstantSpeedEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsFloorConstantSpeedEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetFloorBlockOnWallEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsFloorBlockOnWallEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetSlideOnCeilingEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsSlideOnCeilingEnabled()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetPlatformFloorLayers(exclude_layer int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPlatformFloorLayers()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetPlatformWallLayers(exclude_layer int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPlatformWallLayers()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetMaxSlides()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetMaxSlides(max_slides int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetFloorMaxAngle()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetFloorMaxAngle(radians float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetFloorSnapLength()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetFloorSnapLength(floor_snap_length float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetWallMinSlideAngle()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetWallMinSlideAngle(radians float32, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetUpDirection()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetUpDirection(up_direction Vector3, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetMotionMode(mode CharacterBody3DMotionMode, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetMotionMode()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody3DPlatformOnLeave, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPlatformOnLeave()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnFloor()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnFloorOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnCeiling()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnCeilingOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnWall()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) IsOnWallOnly()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetFloorNormal()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetWallNormal()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetLastMotion()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPositionDelta()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetRealVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetFloorAngle(up_direction Vector3, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPlatformVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetPlatformAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetSlideCollisionCount()  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetSlideCollision(slide_idx int, )  {
  panic("TODO: implement")
}

func  (me *CharacterBody3D) GetLastSlideCollision()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
