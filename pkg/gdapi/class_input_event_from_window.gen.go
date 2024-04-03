// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventFromWindow struct {
  InputEvent
}

func (me *InputEventFromWindow) BaseClass() string {
  return "InputEventFromWindow"
}

func NewInputEventFromWindow() *InputEventFromWindow {
  str := StringNameFromStr("InputEventFromWindow") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventFromWindow{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventFromWindow) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventFromWindow) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventFromWindow) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventFromWindow) SetWindowId(id int64, )  {
  classNameV := StringNameFromStr("InputEventFromWindow")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_window_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventFromWindow) GetWindowId() int64 {
  classNameV := StringNameFromStr("InputEventFromWindow")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
