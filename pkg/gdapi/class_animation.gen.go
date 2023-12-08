// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Animation struct {
  obj gdc.ObjectPtr
}

func (me *Animation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Animation) BaseClass() string {
  return "Animation"
}

type AnimationTrackType int
const (
  AnimationTypeValue AnimationTrackType = 0
  AnimationTypePosition3D AnimationTrackType = 1
  AnimationTypeRotation3D AnimationTrackType = 2
  AnimationTypeScale3D AnimationTrackType = 3
  AnimationTypeBlendShape AnimationTrackType = 4
  AnimationTypeMethod AnimationTrackType = 5
  AnimationTypeBezier AnimationTrackType = 6
  AnimationTypeAudio AnimationTrackType = 7
  AnimationTypeAnimation AnimationTrackType = 8
)

type AnimationInterpolationType int
const (
  AnimationInterpolationNearest AnimationInterpolationType = 0
  AnimationInterpolationLinear AnimationInterpolationType = 1
  AnimationInterpolationCubic AnimationInterpolationType = 2
  AnimationInterpolationLinearAngle AnimationInterpolationType = 3
  AnimationInterpolationCubicAngle AnimationInterpolationType = 4
)

type AnimationUpdateMode int
const (
  AnimationUpdateContinuous AnimationUpdateMode = 0
  AnimationUpdateDiscrete AnimationUpdateMode = 1
  AnimationUpdateCapture AnimationUpdateMode = 2
)

type AnimationLoopMode int
const (
  AnimationLoopNone AnimationLoopMode = 0
  AnimationLoopLinear AnimationLoopMode = 1
  AnimationLoopPingpong AnimationLoopMode = 2
)

type AnimationLoopedFlag int
const (
  AnimationLoopedFlagNone AnimationLoopedFlag = 0
  AnimationLoopedFlagEnd AnimationLoopedFlag = 1
  AnimationLoopedFlagStart AnimationLoopedFlag = 2
)

type AnimationFindMode int
const (
  AnimationFindModeNearest AnimationFindMode = 0
  AnimationFindModeApprox AnimationFindMode = 1
  AnimationFindModeExact AnimationFindMode = 2
)
