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

func  (me *InputEventFromWindow) SetWindowId(id int, )  {
  classNameV := StringNameFromStr("InputEventFromWindow")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_window_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventFromWindow) GetWindowId() int {
  classNameV := StringNameFromStr("InputEventFromWindow")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
