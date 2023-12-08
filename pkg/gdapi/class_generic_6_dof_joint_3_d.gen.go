// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Generic6DOFJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *Generic6DOFJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Generic6DOFJoint3D) BaseClass() string {
  return "Generic6DOFJoint3D"
}

type Generic6DOFJoint3DParam int
const (
  Generic6DOFJoint3DParamLinearLowerLimit Generic6DOFJoint3DParam = 0
  Generic6DOFJoint3DParamLinearUpperLimit Generic6DOFJoint3DParam = 1
  Generic6DOFJoint3DParamLinearLimitSoftness Generic6DOFJoint3DParam = 2
  Generic6DOFJoint3DParamLinearRestitution Generic6DOFJoint3DParam = 3
  Generic6DOFJoint3DParamLinearDamping Generic6DOFJoint3DParam = 4
  Generic6DOFJoint3DParamLinearMotorTargetVelocity Generic6DOFJoint3DParam = 5
  Generic6DOFJoint3DParamLinearMotorForceLimit Generic6DOFJoint3DParam = 6
  Generic6DOFJoint3DParamLinearSpringStiffness Generic6DOFJoint3DParam = 7
  Generic6DOFJoint3DParamLinearSpringDamping Generic6DOFJoint3DParam = 8
  Generic6DOFJoint3DParamLinearSpringEquilibriumPoint Generic6DOFJoint3DParam = 9
  Generic6DOFJoint3DParamAngularLowerLimit Generic6DOFJoint3DParam = 10
  Generic6DOFJoint3DParamAngularUpperLimit Generic6DOFJoint3DParam = 11
  Generic6DOFJoint3DParamAngularLimitSoftness Generic6DOFJoint3DParam = 12
  Generic6DOFJoint3DParamAngularDamping Generic6DOFJoint3DParam = 13
  Generic6DOFJoint3DParamAngularRestitution Generic6DOFJoint3DParam = 14
  Generic6DOFJoint3DParamAngularForceLimit Generic6DOFJoint3DParam = 15
  Generic6DOFJoint3DParamAngularErp Generic6DOFJoint3DParam = 16
  Generic6DOFJoint3DParamAngularMotorTargetVelocity Generic6DOFJoint3DParam = 17
  Generic6DOFJoint3DParamAngularMotorForceLimit Generic6DOFJoint3DParam = 18
  Generic6DOFJoint3DParamAngularSpringStiffness Generic6DOFJoint3DParam = 19
  Generic6DOFJoint3DParamAngularSpringDamping Generic6DOFJoint3DParam = 20
  Generic6DOFJoint3DParamAngularSpringEquilibriumPoint Generic6DOFJoint3DParam = 21
  Generic6DOFJoint3DParamMax Generic6DOFJoint3DParam = 22
)

type Generic6DOFJoint3DFlag int
const (
  Generic6DOFJoint3DFlagEnableLinearLimit Generic6DOFJoint3DFlag = 0
  Generic6DOFJoint3DFlagEnableAngularLimit Generic6DOFJoint3DFlag = 1
  Generic6DOFJoint3DFlagEnableLinearSpring Generic6DOFJoint3DFlag = 3
  Generic6DOFJoint3DFlagEnableAngularSpring Generic6DOFJoint3DFlag = 2
  Generic6DOFJoint3DFlagEnableMotor Generic6DOFJoint3DFlag = 4
  Generic6DOFJoint3DFlagEnableLinearMotor Generic6DOFJoint3DFlag = 5
  Generic6DOFJoint3DFlagMax Generic6DOFJoint3DFlag = 6
)
