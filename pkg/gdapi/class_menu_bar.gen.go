// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MenuBar struct {
  obj gdc.ObjectPtr
}

func (me *MenuBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MenuBar) BaseClass() string {
  return "MenuBar"
}



// Enums

func (me *MenuBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MenuBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MenuBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MenuBar) SetSwitchOnHover(enable bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_switch_on_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) IsSwitchOnHover() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_switch_on_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetDisableShortcuts(disabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_shortcuts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) SetPreferGlobalMenu(enabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_prefer_global_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) IsPreferGlobalMenu() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_prefer_global_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) IsNativeMenu() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_native_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) GetMenuCount() int {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) GetLanguage() String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetFlat(enabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) IsFlat() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetStartIndex(enabled int, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) GetStartIndex() int {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetMenuTitle(menu int, title String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) GetMenuTitle(menu int, ) String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetMenuTooltip(menu int, tooltip String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(tooltip.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) GetMenuTooltip(menu int, ) String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetMenuDisabled(menu int, disabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) IsMenuDisabled(menu int, ) bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) SetMenuHidden(menu int, hidden bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), gdc.ConstTypePtr(&hidden), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MenuBar) IsMenuHidden(menu int, ) bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MenuBar) GetMenuPopup(menu int, ) PopupMenu {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2100501353) // FIXME: should cache?
  var ret PopupMenu
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *MenuBar) GetPropFlat() bool {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropFlat(value bool) {
  panic("TODO: implement")
}

func (me *MenuBar) GetPropStartIndex() int {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropStartIndex(value int) {
  panic("TODO: implement")
}

func (me *MenuBar) GetPropSwitchOnHover() bool {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropSwitchOnHover(value bool) {
  panic("TODO: implement")
}

func (me *MenuBar) GetPropPreferGlobalMenu() bool {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropPreferGlobalMenu(value bool) {
  panic("TODO: implement")
}

func (me *MenuBar) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *MenuBar) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *MenuBar) SetPropLanguage(value String) {
  panic("TODO: implement")
}