// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HingeJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *HingeJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HingeJoint3D) BaseClass() string {
  return "HingeJoint3D"
}

type HingeJoint3DParam int
const (
  HingeJoint3DParamBias HingeJoint3DParam = 0
  HingeJoint3DParamLimitUpper HingeJoint3DParam = 1
  HingeJoint3DParamLimitLower HingeJoint3DParam = 2
  HingeJoint3DParamLimitBias HingeJoint3DParam = 3
  HingeJoint3DParamLimitSoftness HingeJoint3DParam = 4
  HingeJoint3DParamLimitRelaxation HingeJoint3DParam = 5
  HingeJoint3DParamMotorTargetVelocity HingeJoint3DParam = 6
  HingeJoint3DParamMotorMaxImpulse HingeJoint3DParam = 7
  HingeJoint3DParamMax HingeJoint3DParam = 8
)

type HingeJoint3DFlag int
const (
  HingeJoint3DFlagUseLimit HingeJoint3DFlag = 0
  HingeJoint3DFlagEnableMotor HingeJoint3DFlag = 1
  HingeJoint3DFlagMax HingeJoint3DFlag = 2
)
