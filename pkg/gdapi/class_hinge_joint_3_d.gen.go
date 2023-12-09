// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type HingeJoint3DParam int
const (
  HingeJoint3DParamParamBias HingeJoint3DParam = 0
  HingeJoint3DParamParamLimitUpper HingeJoint3DParam = 1
  HingeJoint3DParamParamLimitLower HingeJoint3DParam = 2
  HingeJoint3DParamParamLimitBias HingeJoint3DParam = 3
  HingeJoint3DParamParamLimitSoftness HingeJoint3DParam = 4
  HingeJoint3DParamParamLimitRelaxation HingeJoint3DParam = 5
  HingeJoint3DParamParamMotorTargetVelocity HingeJoint3DParam = 6
  HingeJoint3DParamParamMotorMaxImpulse HingeJoint3DParam = 7
  HingeJoint3DParamParamMax HingeJoint3DParam = 8
)

type HingeJoint3DFlag int
const (
  HingeJoint3DFlagFlagUseLimit HingeJoint3DFlag = 0
  HingeJoint3DFlagFlagEnableMotor HingeJoint3DFlag = 1
  HingeJoint3DFlagFlagMax HingeJoint3DFlag = 2
)

func (me *HingeJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HingeJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *HingeJoint3D) SetParam(param HingeJoint3DParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *HingeJoint3D) GetParam(param HingeJoint3DParam, )  {
  panic("TODO: implement")
}

func  (me *HingeJoint3D) SetFlag(flag HingeJoint3DFlag, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *HingeJoint3D) GetFlag(flag HingeJoint3DFlag, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
