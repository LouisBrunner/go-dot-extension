// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type PinJoint3D struct {
  Joint3D
}

func (me *PinJoint3D) BaseClass() string {
  return "PinJoint3D"
}

func NewPinJoint3D() *PinJoint3D {
  str := StringNameFromStr("PinJoint3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PinJoint3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PinJoint3D) SetParam(param PinJoint3DParam, value float64, )  {
  classNameV := StringNameFromStr("PinJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2059913726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint3D) GetParam(param PinJoint3DParam, ) float64 {
  classNameV := StringNameFromStr("PinJoint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1758438771) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
