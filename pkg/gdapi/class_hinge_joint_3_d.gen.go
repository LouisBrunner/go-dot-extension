// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type HingeJoint3D struct {
  Joint3D
}

func (me *HingeJoint3D) BaseClass() string {
  return "HingeJoint3D"
}

func NewHingeJoint3D() *HingeJoint3D {
  str := StringNameFromStr("HingeJoint3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HingeJoint3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func (me *HingeJoint3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HingeJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HingeJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *HingeJoint3D) SetParam(param HingeJoint3DParam, value float64, )  {
  classNameV := StringNameFromStr("HingeJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082977519) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HingeJoint3D) GetParam(param HingeJoint3DParam, ) float64 {
  classNameV := StringNameFromStr("HingeJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4066002676) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HingeJoint3D) SetFlag(flag HingeJoint3DFlag, enabled bool, )  {
  classNameV := StringNameFromStr("HingeJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1083494620) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HingeJoint3D) GetFlag(flag HingeJoint3DFlag, ) bool {
  classNameV := StringNameFromStr("HingeJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841369610) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&flag)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
