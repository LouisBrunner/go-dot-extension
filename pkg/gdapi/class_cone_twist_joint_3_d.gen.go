// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type ConeTwistJoint3DParam int
const (
  ConeTwistJoint3DParamParamSwingSpan ConeTwistJoint3DParam = 0
  ConeTwistJoint3DParamParamTwistSpan ConeTwistJoint3DParam = 1
  ConeTwistJoint3DParamParamBias ConeTwistJoint3DParam = 2
  ConeTwistJoint3DParamParamSoftness ConeTwistJoint3DParam = 3
  ConeTwistJoint3DParamParamRelaxation ConeTwistJoint3DParam = 4
  ConeTwistJoint3DParamParamMax ConeTwistJoint3DParam = 5
)

func (me *ConeTwistJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConeTwistJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConeTwistJoint3D) SetParam(param ConeTwistJoint3DParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *ConeTwistJoint3D) GetParam(param ConeTwistJoint3DParam, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
