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

func  (me *CharacterBody2D) MoveAndSlide() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) ApplyFloorSnap() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetVelocity(velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetSafeMargin(margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetSafeMargin() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsFloorStopOnSlopeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetFloorStopOnSlopeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetFloorConstantSpeedEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsFloorConstantSpeedEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetFloorBlockOnWallEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsFloorBlockOnWallEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetSlideOnCeilingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsSlideOnCeilingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetPlatformFloorLayers(exclude_layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetPlatformFloorLayers() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetPlatformWallLayers(exclude_layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetPlatformWallLayers() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetMaxSlides() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetMaxSlides(max_slides int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetFloorMaxAngle() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetFloorMaxAngle(radians float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetFloorSnapLength() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetFloorSnapLength(floor_snap_length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetWallMinSlideAngle() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetWallMinSlideAngle(radians float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetUpDirection() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetUpDirection(up_direction Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetMotionMode(mode CharacterBody2DMotionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetMotionMode() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody2DPlatformOnLeave, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetPlatformOnLeave() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnFloor() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnFloorOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnCeiling() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnCeilingOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnWall() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) IsOnWallOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetFloorNormal() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetWallNormal() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetLastMotion() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetPositionDelta() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetRealVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetFloorAngle(up_direction Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetPlatformVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetSlideCollisionCount() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetSlideCollision(slide_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody2D) GetLastSlideCollision() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
