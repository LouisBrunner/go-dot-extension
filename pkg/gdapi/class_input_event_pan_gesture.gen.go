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

type InputEventPanGesture struct {
  InputEventGesture
}

func (me *InputEventPanGesture) BaseClass() string {
  return "InputEventPanGesture"
}

func NewInputEventPanGesture() *InputEventPanGesture {
  str := StringNameFromStr("InputEventPanGesture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventPanGesture{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{delta.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventPanGesture) GetDelta() Vector2 {
  classNameV := StringNameFromStr("InputEventPanGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
