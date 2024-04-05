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

type ptrsForInputEventFromWindowList struct {
  fnSetWindowId gdc.MethodBindPtr
  fnGetWindowId gdc.MethodBindPtr
}

var ptrsForInputEventFromWindow ptrsForInputEventFromWindowList

func initInputEventFromWindowPtrs(iface gdc.Interface) {

  className := StringNameFromStr("InputEventFromWindow")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_window_id")
    defer methodName.Destroy()
    ptrsForInputEventFromWindow.fnSetWindowId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_window_id")
    defer methodName.Destroy()
    ptrsForInputEventFromWindow.fnGetWindowId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventFromWindow.fnSetWindowId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventFromWindow) GetWindowId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventFromWindow.fnGetWindowId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
