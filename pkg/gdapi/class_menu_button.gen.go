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

type ptrsForMenuButtonList struct {
  fnGetPopup gdc.MethodBindPtr
  fnShowPopup gdc.MethodBindPtr
  fnSetSwitchOnHover gdc.MethodBindPtr
  fnIsSwitchOnHover gdc.MethodBindPtr
  fnSetDisableShortcuts gdc.MethodBindPtr
  fnSetItemCount gdc.MethodBindPtr
  fnGetItemCount gdc.MethodBindPtr
}

var ptrsForMenuButton ptrsForMenuButtonList

func initMenuButtonPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MenuButton")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_popup")
    defer methodName.Destroy()
    ptrsForMenuButton.fnGetPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229722558))
  }
  {
    methodName := StringNameFromStr("show_popup")
    defer methodName.Destroy()
    ptrsForMenuButton.fnShowPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_switch_on_hover")
    defer methodName.Destroy()
    ptrsForMenuButton.fnSetSwitchOnHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_switch_on_hover")
    defer methodName.Destroy()
    ptrsForMenuButton.fnIsSwitchOnHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_disable_shortcuts")
    defer methodName.Destroy()
    ptrsForMenuButton.fnSetDisableShortcuts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_item_count")
    defer methodName.Destroy()
    ptrsForMenuButton.fnSetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_item_count")
    defer methodName.Destroy()
    ptrsForMenuButton.fnGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type MenuButton struct {
  Button
}

func (me *MenuButton) BaseClass() string {
  return "MenuButton"
}

func NewMenuButton() *MenuButton {
  str := StringNameFromStr("MenuButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MenuButton{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MenuButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MenuButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MenuButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MenuButton) GetPopup() PopupMenu {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnGetPopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MenuButton) ShowPopup()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnShowPopup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuButton) SetSwitchOnHover(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnSetSwitchOnHover), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuButton) IsSwitchOnHover() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnIsSwitchOnHover), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuButton) SetDisableShortcuts(disabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnSetDisableShortcuts), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuButton) SetItemCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnSetItemCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuButton) GetItemCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMenuButton.fnGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MenuButtonAboutToPopupSignalFn func()

func (me *MenuButton) ConnectAboutToPopup(subs SignalSubscribers, fn MenuButtonAboutToPopupSignalFn) {
  sig := StringNameFromStr("about_to_popup")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MenuButton) DisconnectAboutToPopup(subs SignalSubscribers, fn MenuButtonAboutToPopupSignalFn) {
  sig := StringNameFromStr("about_to_popup")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
