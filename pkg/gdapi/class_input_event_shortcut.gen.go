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

type InputEventShortcut struct {
  InputEvent
}

func (me *InputEventShortcut) BaseClass() string {
  return "InputEventShortcut"
}

func NewInputEventShortcut() *InputEventShortcut {
  str := StringNameFromStr("InputEventShortcut") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventShortcut{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventShortcut) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventShortcut) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventShortcut) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventShortcut) SetShortcut(shortcut Shortcut, )  {
  classNameV := StringNameFromStr("InputEventShortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 857163497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventShortcut) GetShortcut() Shortcut {
  classNameV := StringNameFromStr("InputEventShortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3766804753) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShortcut()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
