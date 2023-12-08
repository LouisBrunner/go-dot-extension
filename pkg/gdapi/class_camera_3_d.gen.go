// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera3D struct {
  obj gdc.ObjectPtr
}

func (me *Camera3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera3D) BaseClass() string {
  return "Camera3D"
}

type Camera3DProjectionType int
const (
  Camera3DProjectionPerspective Camera3DProjectionType = 0
  Camera3DProjectionOrthogonal Camera3DProjectionType = 1
  Camera3DProjectionFrustum Camera3DProjectionType = 2
)

type Camera3DKeepAspect int
const (
  Camera3DKeepWidth Camera3DKeepAspect = 0
  Camera3DKeepHeight Camera3DKeepAspect = 1
)

type Camera3DDopplerTracking int
const (
  Camera3DDopplerTrackingDisabled Camera3DDopplerTracking = 0
  Camera3DDopplerTrackingIdleStep Camera3DDopplerTracking = 1
  Camera3DDopplerTrackingPhysicsStep Camera3DDopplerTracking = 2
)
