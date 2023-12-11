// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventPanGesture struct {
  obj gdc.ObjectPtr
}

func (me *InputEventPanGesture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventPanGesture) BaseClass() string {
  return "InputEventPanGesture"
}



// Enums

func (me *InputEventPanGesture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventPanGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventPanGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventPanGesture) SetDelta(delta Vector2, )  {
  classNameV := StringNameFromStr("InputEventPanGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(delta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventPanGesture) GetDelta() Vector2 {
  classNameV := StringNameFromStr("InputEventPanGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
