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

type InputEventMouseButton struct {
  InputEventMouse
}

func (me *InputEventMouseButton) BaseClass() string {
  return "InputEventMouseButton"
}

func NewInputEventMouseButton() *InputEventMouseButton {
  str := StringNameFromStr("InputEventMouseButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventMouseButton{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventMouseButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventMouseButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMouseButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventMouseButton) SetFactor(factor float64, )  {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseButton) GetFactor() float64 {
  classNameV := StringNameFromStr("InputEventMouseButton")
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

func  (me *InputEventMouseButton) SetButtonIndex(button_index MouseButton, )  {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3624991109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseButton) GetButtonIndex() MouseButton {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1132662608) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MouseButton

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *InputEventMouseButton) SetPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseButton) SetCanceled(canceled bool, )  {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canceled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canceled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseButton) SetDoubleClick(double_click bool, )  {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_double_click")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&double_click) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseButton) IsDoubleClick() bool {
  classNameV := StringNameFromStr("InputEventMouseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_double_click")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
