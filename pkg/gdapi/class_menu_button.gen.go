// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MenuButton struct {
  obj gdc.ObjectPtr
}

func (me *MenuButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MenuButton) BaseClass() string {
  return "MenuButton"
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
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229722558) // FIXME: should cache?
  var ret PopupMenu
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuButton) ShowPopup()  {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuButton) SetSwitchOnHover(enable bool, )  {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_switch_on_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuButton) IsSwitchOnHover() bool {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_switch_on_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuButton) SetDisableShortcuts(disabled bool, )  {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_shortcuts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuButton) SetItemCount(count int, )  {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuButton) GetItemCount() int {
  classNameV := StringNameFromStr("MenuButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_count")
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

type MenuButtonAboutToPopupSignalFn func()

func (me *MenuButton) ConnectAboutToPopup(subs SignalSubscribers, fn MenuButtonAboutToPopupSignalFn) {
  sig := StringNameFromStr("about_to_popup")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MenuButton) DisconnectAboutToPopup(subs SignalSubscribers, fn MenuButtonAboutToPopupSignalFn) {
  sig := StringNameFromStr("about_to_popup")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
