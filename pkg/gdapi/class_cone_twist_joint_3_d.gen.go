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

type ptrsForConeTwistJoint3DList struct {
  fnSetParam gdc.MethodBindPtr
  fnGetParam gdc.MethodBindPtr
}

var ptrsForConeTwistJoint3D ptrsForConeTwistJoint3DList

func initConeTwistJoint3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ConeTwistJoint3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_param")
    defer methodName.Destroy()
    ptrsForConeTwistJoint3D.fnSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1062470226))
  }
  {
    methodName := StringNameFromStr("get_param")
    defer methodName.Destroy()
    ptrsForConeTwistJoint3D.fnGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2928790850))
  }
}

type ConeTwistJoint3D struct {
  Joint3D
}

func (me *ConeTwistJoint3D) BaseClass() string {
  return "ConeTwistJoint3D"
}

func NewConeTwistJoint3D() *ConeTwistJoint3D {
  str := StringNameFromStr("ConeTwistJoint3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConeTwistJoint3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *ConeTwistJoint3D) SetParam(param ConeTwistJoint3DParam, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConeTwistJoint3D.fnSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConeTwistJoint3D) GetParam(param ConeTwistJoint3DParam, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConeTwistJoint3D.fnGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
