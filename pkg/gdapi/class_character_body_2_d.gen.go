// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  CharacterBody2DMotionModeGrounded CharacterBody2DMotionMode = 0
  CharacterBody2DMotionModeFloating CharacterBody2DMotionMode = 1
)

type CharacterBody2DPlatformOnLeave int
const (
  CharacterBody2DPlatformOnLeaveAddVelocity CharacterBody2DPlatformOnLeave = 0
  CharacterBody2DPlatformOnLeaveAddUpwardVelocity CharacterBody2DPlatformOnLeave = 1
  CharacterBody2DPlatformOnLeaveDoNothing CharacterBody2DPlatformOnLeave = 2
)
