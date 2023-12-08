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

func  (me *CharacterBody3D) MoveAndSlide() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) ApplyFloorSnap() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetVelocity(velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetSafeMargin(margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetSafeMargin() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsFloorStopOnSlopeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetFloorStopOnSlopeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetFloorConstantSpeedEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsFloorConstantSpeedEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetFloorBlockOnWallEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsFloorBlockOnWallEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetSlideOnCeilingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsSlideOnCeilingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetPlatformFloorLayers(exclude_layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPlatformFloorLayers() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetPlatformWallLayers(exclude_layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPlatformWallLayers() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetMaxSlides() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetMaxSlides(max_slides int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetFloorMaxAngle() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetFloorMaxAngle(radians float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetFloorSnapLength() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetFloorSnapLength(floor_snap_length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetWallMinSlideAngle() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetWallMinSlideAngle(radians float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetUpDirection() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetUpDirection(up_direction Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetMotionMode(mode CharacterBody3DMotionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetMotionMode() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody3DPlatformOnLeave, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPlatformOnLeave() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnFloor() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnFloorOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnCeiling() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnCeilingOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnWall() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) IsOnWallOnly() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetFloorNormal() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetWallNormal() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetLastMotion() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPositionDelta() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetRealVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetFloorAngle(up_direction Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPlatformVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetPlatformAngularVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetSlideCollisionCount() { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetSlideCollision(slide_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CharacterBody3D) GetLastSlideCollision() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
