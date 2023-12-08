// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  PinJoint3DParamParamBias PinJoint3DParam = 0
  PinJoint3DParamParamDamping PinJoint3DParam = 1
  PinJoint3DParamParamImpulseClamp PinJoint3DParam = 2
)

func  (me *PinJoint3D) SetParam(param PinJoint3DParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PinJoint3D) GetParam(param PinJoint3DParam, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
