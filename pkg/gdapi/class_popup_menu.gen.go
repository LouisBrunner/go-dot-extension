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

func  (me *PopupMenu) AddItem(label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconItem(texture Texture2D, label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddCheckItem(label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconCheckItem(texture Texture2D, label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddRadioCheckItem(label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconRadioCheckItem(texture Texture2D, label String, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddMultistateItem(label String, max_states int, default_state int, id int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddShortcut(shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddCheckShortcut(shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconCheckShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddRadioCheckShortcut(shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddIconRadioCheckShortcut(texture Texture2D, shortcut Shortcut, id int, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddSubmenuItem(label String, submenu String, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemText(index int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemTextDirection(index int, direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemLanguage(index int, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemIcon(index int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemIconMaxWidth(index int, width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemIconModulate(index int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemChecked(index int, checked bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemId(index int, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemAccelerator(index int, accel Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemMetadata(index int, metadata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemDisabled(index int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemSubmenu(index int, submenu String, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemAsSeparator(index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemAsCheckable(index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemAsRadioCheckable(index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemTooltip(index int, tooltip String, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemShortcut(index int, shortcut Shortcut, global bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemIndent(index int, indent int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemMultistate(index int, state int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemShortcutDisabled(index int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) ToggleItemChecked(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) ToggleItemMultistate(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemText(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemTextDirection(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemLanguage(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemIcon(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemIconMaxWidth(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemIconModulate(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemChecked(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemId(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemIndex(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemAccelerator(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemMetadata(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemDisabled(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemSubmenu(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemSeparator(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemCheckable(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemRadioCheckable(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsItemShortcutDisabled(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemTooltip(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemShortcut(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemIndent(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetFocusedItem(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetFocusedItem() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetItemCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetItemCount() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) ScrollToItem(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) RemoveItem(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) AddSeparator(label String, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetHideOnItemSelection(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsHideOnItemSelection() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetHideOnCheckableItemSelection(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsHideOnCheckableItemSelection() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetHideOnStateItemSelection(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) IsHideOnStateItemSelection() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetSubmenuPopupDelay(seconds float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetSubmenuPopupDelay() { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) SetAllowSearch(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PopupMenu) GetAllowSearch() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
