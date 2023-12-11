// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventJoypadMotion struct {
  obj gdc.ObjectPtr
}

func (me *InputEventJoypadMotion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventJoypadMotion) BaseClass() string {
  return "InputEventJoypadMotion"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventJoypadMotion) GetAxis() JoyAxis {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4019121683) // FIXME: should cache?
  var ret JoyAxis
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventJoypadMotion) SetAxisValue(axis_value float32, )  {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis_value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventJoypadMotion) GetAxisValue() float32 {
  classNameV := StringNameFromStr("InputEventJoypadMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis_value")
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
