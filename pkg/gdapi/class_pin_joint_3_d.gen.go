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

type ptrsForPinJoint3DList struct {
  fnSetParam gdc.MethodBindPtr
  fnGetParam gdc.MethodBindPtr
}

var ptrsForPinJoint3D ptrsForPinJoint3DList

func initPinJoint3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PinJoint3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_param")
    defer methodName.Destroy()
    ptrsForPinJoint3D.fnSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2059913726))
  }
  {
    methodName := StringNameFromStr("get_param")
    defer methodName.Destroy()
    ptrsForPinJoint3D.fnGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1758438771))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint3D.fnSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PinJoint3D) GetParam(param PinJoint3DParam, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPinJoint3D.fnGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
