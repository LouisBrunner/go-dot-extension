// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *ConeTwistJoint3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConeTwistJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConeTwistJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConeTwistJoint3D) SetParam(param ConeTwistJoint3DParam, value float32, )  {
  classNameV := StringNameFromStr("ConeTwistJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1062470226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConeTwistJoint3D) GetParam(param ConeTwistJoint3DParam, ) float32 {
  classNameV := StringNameFromStr("ConeTwistJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2928790850) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ConeTwistJoint3D) GetPropSwingSpan() float32 {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) SetPropSwingSpan(value float32) {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) GetPropTwistSpan() float32 {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) SetPropTwistSpan(value float32) {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) GetPropBias() float32 {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) SetPropBias(value float32) {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) GetPropSoftness() float32 {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) SetPropSoftness(value float32) {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) GetPropRelaxation() float32 {
  panic("TODO: implement")
}

func (me *ConeTwistJoint3D) SetPropRelaxation(value float32) {
  panic("TODO: implement")
}