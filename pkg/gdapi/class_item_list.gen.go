// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ItemList struct {
  obj gdc.ObjectPtr
}

func (me *ItemList) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ItemList) BaseClass() string {
  return "ItemList"
}

type ItemListIconMode int
const (
  ItemListIconModeIconModeTop ItemListIconMode = 0
  ItemListIconModeIconModeLeft ItemListIconMode = 1
)

type ItemListSelectMode int
const (
  ItemListSelectModeSelectSingle ItemListSelectMode = 0
  ItemListSelectModeSelectMulti ItemListSelectMode = 1
)

func  (me *ItemList) AddItem(text String, icon Texture2D, selectable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) AddIconItem(icon Texture2D, selectable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemText(idx int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemText(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemIcon(idx int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemIcon(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemTextDirection(idx int, direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemTextDirection(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemLanguage(idx int, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemLanguage(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemIconTransposed(idx int, transposed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsItemIconTransposed(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemIconRegion(idx int, rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemIconRegion(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemIconModulate(idx int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemIconModulate(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemSelectable(idx int, selectable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsItemSelectable(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemDisabled(idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsItemDisabled(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemMetadata(idx int, metadata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemMetadata(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemCustomBgColor(idx int, custom_bg_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemCustomBgColor(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemCustomFgColor(idx int, custom_fg_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemCustomFgColor(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemRect(idx int, expand bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemTooltipEnabled(idx int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsItemTooltipEnabled(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemTooltip(idx int, tooltip String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemTooltip(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) Select(idx int, single bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) Deselect(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) DeselectAll() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsSelected(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetSelectedItems() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) MoveItem(from_idx int, to_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetItemCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemCount() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) RemoveItem(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SortItemsByText() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetFixedColumnWidth(width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetFixedColumnWidth() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetSameColumnWidth(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsSameColumnWidth() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetMaxTextLines(lines int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetMaxTextLines() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetMaxColumns(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetMaxColumns() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetSelectMode(mode ItemListSelectMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetSelectMode() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetIconMode(mode ItemListIconMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetIconMode() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetFixedIconSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetFixedIconSize() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetIconScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetIconScale() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetAllowRmbSelect(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetAllowRmbSelect() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetAllowReselect(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetAllowReselect() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetAllowSearch(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetAllowSearch() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetAutoHeight(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) HasAutoHeight() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) IsAnythingSelected() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetItemAtPosition(position Vector2, exact bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) EnsureCurrentIsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetVScrollBar() { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, ) { // TODO: return value
  // TODO: implement
}

func  (me *ItemList) GetTextOverrunBehavior() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
