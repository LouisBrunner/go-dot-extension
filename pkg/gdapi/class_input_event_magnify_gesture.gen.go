// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventMagnifyGesture struct {
  InputEventGesture
}

func (me *InputEventMagnifyGesture) BaseClass() string {
  return "InputEventMagnifyGesture"
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

func  (me *InputEventMagnifyGesture) SetFactor(factor float32, )  {
  classNameV := StringNameFromStr("InputEventMagnifyGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMagnifyGesture) GetFactor() float32 {
  classNameV := StringNameFromStr("InputEventMagnifyGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
