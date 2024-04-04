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

type InputEventJoypadMotion struct {
  InputEvent
}

func (me *InputEventJoypadMotion) BaseClass() string {
  return "InputEventJoypadMotion"
}

func NewInputEventJoypadMotion() *InputEventJoypadMotion {
  str := StringNameFromStr("InputEventJoypadMotion") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventJoypadMotion{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventJoypadMotion) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventJoypadMotion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventJoypadMotion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventJoypadMotion) SetAxis(axis JoyAxis, )  {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1332685170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventJoypadMotion) GetAxis() JoyAxis {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4019121683) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret JoyAxis

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *InputEventJoypadMotion) SetAxisValue(axis_value float64, )  {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis_value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventJoypadMotion) GetAxisValue() float64 {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis_value")
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
