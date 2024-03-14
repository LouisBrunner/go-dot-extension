// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SliderJoint3D struct {
  Joint3D
}

func (me *SliderJoint3D) BaseClass() string {
  return "SliderJoint3D"
}



// Enums

type SliderJoint3DParam int
const (
  SliderJoint3DParamParamLinearLimitUpper SliderJoint3DParam = 0
  SliderJoint3DParamParamLinearLimitLower SliderJoint3DParam = 1
  SliderJoint3DParamParamLinearLimitSoftness SliderJoint3DParam = 2
  SliderJoint3DParamParamLinearLimitRestitution SliderJoint3DParam = 3
  SliderJoint3DParamParamLinearLimitDamping SliderJoint3DParam = 4
  SliderJoint3DParamParamLinearMotionSoftness SliderJoint3DParam = 5
  SliderJoint3DParamParamLinearMotionRestitution SliderJoint3DParam = 6
  SliderJoint3DParamParamLinearMotionDamping SliderJoint3DParam = 7
  SliderJoint3DParamParamLinearOrthogonalSoftness SliderJoint3DParam = 8
  SliderJoint3DParamParamLinearOrthogonalRestitution SliderJoint3DParam = 9
  SliderJoint3DParamParamLinearOrthogonalDamping SliderJoint3DParam = 10
  SliderJoint3DParamParamAngularLimitUpper SliderJoint3DParam = 11
  SliderJoint3DParamParamAngularLimitLower SliderJoint3DParam = 12
  SliderJoint3DParamParamAngularLimitSoftness SliderJoint3DParam = 13
  SliderJoint3DParamParamAngularLimitRestitution SliderJoint3DParam = 14
  SliderJoint3DParamParamAngularLimitDamping SliderJoint3DParam = 15
  SliderJoint3DParamParamAngularMotionSoftness SliderJoint3DParam = 16
  SliderJoint3DParamParamAngularMotionRestitution SliderJoint3DParam = 17
  SliderJoint3DParamParamAngularMotionDamping SliderJoint3DParam = 18
  SliderJoint3DParamParamAngularOrthogonalSoftness SliderJoint3DParam = 19
  SliderJoint3DParamParamAngularOrthogonalRestitution SliderJoint3DParam = 20
  SliderJoint3DParamParamAngularOrthogonalDamping SliderJoint3DParam = 21
  SliderJoint3DParamParamMax SliderJoint3DParam = 22
)

func (me *SliderJoint3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SliderJoint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SliderJoint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SliderJoint3D) SetParam(param SliderJoint3DParam, value float32, )  {
  classNameV := StringNameFromStr("SliderJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 918243683) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SliderJoint3D) GetParam(param SliderJoint3DParam, ) float32 {
  classNameV := StringNameFromStr("SliderJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 959925627) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
