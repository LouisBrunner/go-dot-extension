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



// Enums

type PinJoint3DParam int
const (
  PinJoint3DParamParamBias PinJoint3DParam = 0
  PinJoint3DParamParamDamping PinJoint3DParam = 1
  PinJoint3DParamParamImpulseClamp PinJoint3DParam = 2
)

func (me *PinJoint3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PinJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PinJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PinJoint3D) SetParam(param PinJoint3DParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PinJoint3D) GetParam(param PinJoint3DParam, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
