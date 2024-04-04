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

type MenuBar struct {
  Control
}

func (me *MenuBar) BaseClass() string {
  return "MenuBar"
}

func NewMenuBar() *MenuBar {
  str := StringNameFromStr("MenuBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MenuBar{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) IsSwitchOnHover() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_switch_on_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) SetDisableShortcuts(disabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_shortcuts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) SetPreferGlobalMenu(enabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_prefer_global_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) IsPreferGlobalMenu() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_prefer_global_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) IsNativeMenu() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_native_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) GetMenuCount() int64 {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MenuBar) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) GetLanguage() String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MenuBar) SetFlat(enabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) IsFlat() bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) SetStartIndex(enabled int64, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) GetStartIndex() int64 {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) SetMenuTitle(menu int64, title String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , title.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) GetMenuTitle(menu int64, ) String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&menu)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MenuBar) SetMenuTooltip(menu int64, tooltip String, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , tooltip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) GetMenuTooltip(menu int64, ) String {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&menu)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MenuBar) SetMenuDisabled(menu int64, disabled bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) IsMenuDisabled(menu int64, ) bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&menu)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) SetMenuHidden(menu int64, hidden bool, )  {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_menu_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MenuBar) IsMenuHidden(menu int64, ) bool {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&menu)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MenuBar) GetMenuPopup(menu int64, ) PopupMenu {
  classNameV := StringNameFromStr("MenuBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2100501353) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&menu) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()
  pinner.Pin(&menu)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
