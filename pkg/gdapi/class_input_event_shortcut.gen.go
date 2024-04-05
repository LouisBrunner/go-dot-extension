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

type ptrsForInputEventShortcutList struct {
  fnSetShortcut gdc.MethodBindPtr
  fnGetShortcut gdc.MethodBindPtr
}

var ptrsForInputEventShortcut ptrsForInputEventShortcutList

func initInputEventShortcutPtrs(iface gdc.Interface) {

  className := StringNameFromStr("InputEventShortcut")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_shortcut")
    defer methodName.Destroy()
    ptrsForInputEventShortcut.fnSetShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 857163497))
  }
  {
    methodName := StringNameFromStr("get_shortcut")
    defer methodName.Destroy()
    ptrsForInputEventShortcut.fnGetShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3766804753))
  }
}

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
  cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventShortcut.fnSetShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventShortcut) GetShortcut() Shortcut {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShortcut()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventShortcut.fnGetShortcut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
