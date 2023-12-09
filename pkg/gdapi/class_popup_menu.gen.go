// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PopupMenu struct {
  obj gdc.ObjectPtr
}

func (me *PopupMenu) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PopupMenu) BaseClass() string {
  return "PopupMenu"
}



// Enums

func (me *PopupMenu) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PopupMenu) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PopupMenu) AddItem(label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconItem(texture Texture2D, label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddCheckItem(label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconCheckItem(texture Texture2D, label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddRadioCheckItem(label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconRadioCheckItem(texture Texture2D, label String, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddMultistateItem(label String, max_states int, default_state int, id int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddShortcut(shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddCheckShortcut(shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconCheckShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddRadioCheckShortcut(shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddIconRadioCheckShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddSubmenuItem(label String, submenu String, id int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemText(index int, text String, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemTextDirection(index int, direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemLanguage(index int, language String, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemIcon(index int, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemIconMaxWidth(index int, width int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemIconModulate(index int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemChecked(index int, checked bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemId(index int, id int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemAccelerator(index int, accel Key, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemMetadata(index int, metadata Variant, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemDisabled(index int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemSubmenu(index int, submenu String, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemAsSeparator(index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemAsCheckable(index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemAsRadioCheckable(index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemTooltip(index int, tooltip String, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemShortcut(index int, shortcut Shortcut, global bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemIndent(index int, indent int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemMultistate(index int, state int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemShortcutDisabled(index int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) ToggleItemChecked(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) ToggleItemMultistate(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemText(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemTextDirection(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemLanguage(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemIcon(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemIconMaxWidth(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemIconModulate(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemChecked(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemId(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemIndex(id int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemAccelerator(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemMetadata(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemDisabled(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemSubmenu(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemSeparator(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemCheckable(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemRadioCheckable(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsItemShortcutDisabled(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemTooltip(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemShortcut(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemIndent(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetFocusedItem(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetFocusedItem()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetItemCount(count int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetItemCount()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) ScrollToItem(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) RemoveItem(index int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) AddSeparator(label String, id int, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) Clear()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetHideOnItemSelection(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsHideOnItemSelection()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetHideOnCheckableItemSelection(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsHideOnCheckableItemSelection()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetHideOnStateItemSelection(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) IsHideOnStateItemSelection()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetSubmenuPopupDelay(seconds float32, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetSubmenuPopupDelay()  {
  panic("TODO: implement")
}

func  (me *PopupMenu) SetAllowSearch(allow bool, )  {
  panic("TODO: implement")
}

func  (me *PopupMenu) GetAllowSearch()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
