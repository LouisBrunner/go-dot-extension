// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SliderJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *SliderJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SliderJoint3D) BaseClass() string {
  return "SliderJoint3D"
}

type SliderJoint3DParam int
const (
  SliderJoint3DParamLinearLimitUpper SliderJoint3DParam = 0
  SliderJoint3DParamLinearLimitLower SliderJoint3DParam = 1
  SliderJoint3DParamLinearLimitSoftness SliderJoint3DParam = 2
  SliderJoint3DParamLinearLimitRestitution SliderJoint3DParam = 3
  SliderJoint3DParamLinearLimitDamping SliderJoint3DParam = 4
  SliderJoint3DParamLinearMotionSoftness SliderJoint3DParam = 5
  SliderJoint3DParamLinearMotionRestitution SliderJoint3DParam = 6
  SliderJoint3DParamLinearMotionDamping SliderJoint3DParam = 7
  SliderJoint3DParamLinearOrthogonalSoftness SliderJoint3DParam = 8
  SliderJoint3DParamLinearOrthogonalRestitution SliderJoint3DParam = 9
  SliderJoint3DParamLinearOrthogonalDamping SliderJoint3DParam = 10
  SliderJoint3DParamAngularLimitUpper SliderJoint3DParam = 11
  SliderJoint3DParamAngularLimitLower SliderJoint3DParam = 12
  SliderJoint3DParamAngularLimitSoftness SliderJoint3DParam = 13
  SliderJoint3DParamAngularLimitRestitution SliderJoint3DParam = 14
  SliderJoint3DParamAngularLimitDamping SliderJoint3DParam = 15
  SliderJoint3DParamAngularMotionSoftness SliderJoint3DParam = 16
  SliderJoint3DParamAngularMotionRestitution SliderJoint3DParam = 17
  SliderJoint3DParamAngularMotionDamping SliderJoint3DParam = 18
  SliderJoint3DParamAngularOrthogonalSoftness SliderJoint3DParam = 19
  SliderJoint3DParamAngularOrthogonalRestitution SliderJoint3DParam = 20
  SliderJoint3DParamAngularOrthogonalDamping SliderJoint3DParam = 21
  SliderJoint3DParamMax SliderJoint3DParam = 22
)
