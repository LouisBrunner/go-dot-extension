// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("PinJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2059913726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint3D) GetParam(param PinJoint3DParam, ) float32 {
  classNameV := StringNameFromStr("PinJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1758438771) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
