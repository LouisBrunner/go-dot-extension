// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventFromWindow struct {
  obj gdc.ObjectPtr
}

func (me *InputEventFromWindow) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *InputEventFromWindow) GetPropWindowId() int {
  panic("TODO: implement")
}

func (me *InputEventFromWindow) SetPropWindowId(value int) {
  panic("TODO: implement")
}