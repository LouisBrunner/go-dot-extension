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

type ptrsForInputEventPanGestureList struct {
  fnSetDelta gdc.MethodBindPtr
  fnGetDelta gdc.MethodBindPtr
}

var ptrsForInputEventPanGesture ptrsForInputEventPanGestureList

func initInputEventPanGesturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("InputEventPanGesture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_delta")
    defer methodName.Destroy()
    ptrsForInputEventPanGesture.fnSetDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_delta")
    defer methodName.Destroy()
    ptrsForInputEventPanGesture.fnGetDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
}

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
  cargs := []gdc.ConstTypePtr{delta.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventPanGesture.fnSetDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventPanGesture) GetDelta() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventPanGesture.fnGetDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
