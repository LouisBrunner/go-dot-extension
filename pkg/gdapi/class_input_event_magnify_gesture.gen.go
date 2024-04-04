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

type InputEventMagnifyGesture struct {
  InputEventGesture
}

func (me *InputEventMagnifyGesture) BaseClass() string {
  return "InputEventMagnifyGesture"
}

func NewInputEventMagnifyGesture() *InputEventMagnifyGesture {
  str := StringNameFromStr("InputEventMagnifyGesture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventMagnifyGesture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventMagnifyGesture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventMagnifyGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMagnifyGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventMagnifyGesture) SetFactor(factor float64, )  {
  classNameV := StringNameFromStr("InputEventMagnifyGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMagnifyGesture) GetFactor() float64 {
  classNameV := StringNameFromStr("InputEventMagnifyGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
