// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  CharacterBody3DMotionModeGrounded CharacterBody3DMotionMode = 0
  CharacterBody3DMotionModeFloating CharacterBody3DMotionMode = 1
)

type CharacterBody3DPlatformOnLeave int
const (
  CharacterBody3DPlatformOnLeaveAddVelocity CharacterBody3DPlatformOnLeave = 0
  CharacterBody3DPlatformOnLeaveAddUpwardVelocity CharacterBody3DPlatformOnLeave = 1
  CharacterBody3DPlatformOnLeaveDoNothing CharacterBody3DPlatformOnLeave = 2
)
