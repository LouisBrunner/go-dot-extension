// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConeTwistJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *ConeTwistJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConeTwistJoint3D) BaseClass() string {
  return "ConeTwistJoint3D"
}

type ConeTwistJoint3DParam int
const (
  ConeTwistJoint3DParamSwingSpan ConeTwistJoint3DParam = 0
  ConeTwistJoint3DParamTwistSpan ConeTwistJoint3DParam = 1
  ConeTwistJoint3DParamBias ConeTwistJoint3DParam = 2
  ConeTwistJoint3DParamSoftness ConeTwistJoint3DParam = 3
  ConeTwistJoint3DParamRelaxation ConeTwistJoint3DParam = 4
  ConeTwistJoint3DParamMax ConeTwistJoint3DParam = 5
)
