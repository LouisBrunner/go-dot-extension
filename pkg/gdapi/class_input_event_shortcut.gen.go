// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventShortcut struct {
  obj gdc.ObjectPtr
}

func (me *InputEventShortcut) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventShortcut) BaseClass() string {
  return "InputEventShortcut"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shortcut.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventShortcut) GetShortcut() Shortcut {
  classNameV := StringNameFromStr("InputEventShortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3766804753) // FIXME: should cache?
  var ret Shortcut
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
