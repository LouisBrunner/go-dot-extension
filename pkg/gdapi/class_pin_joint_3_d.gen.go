// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PinJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *PinJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PinJoint3D) BaseClass() string {
  return "PinJoint3D"
}

type PinJoint3DParam int
const (
  PinJoint3DParamBias PinJoint3DParam = 0
  PinJoint3DParamDamping PinJoint3DParam = 1
  PinJoint3DParamImpulseClamp PinJoint3DParam = 2
)
