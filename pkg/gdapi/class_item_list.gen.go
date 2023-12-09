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



// Enums

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

func (me *ItemList) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ItemList) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ItemList) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ItemList) AddItem(text String, icon Texture2D, selectable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) AddIconItem(icon Texture2D, selectable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemText(idx int, text String, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemText(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemIcon(idx int, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemIcon(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemTextDirection(idx int, direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemTextDirection(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemLanguage(idx int, language String, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemLanguage(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemIconTransposed(idx int, transposed bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) IsItemIconTransposed(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemIconRegion(idx int, rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemIconRegion(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemIconModulate(idx int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemIconModulate(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemSelectable(idx int, selectable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) IsItemSelectable(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemDisabled(idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) IsItemDisabled(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemMetadata(idx int, metadata Variant, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemMetadata(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemCustomBgColor(idx int, custom_bg_color Color, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemCustomBgColor(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemCustomFgColor(idx int, custom_fg_color Color, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemCustomFgColor(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemRect(idx int, expand bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemTooltipEnabled(idx int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) IsItemTooltipEnabled(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemTooltip(idx int, tooltip String, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemTooltip(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) Select(idx int, single bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) Deselect(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) DeselectAll()  {
  panic("TODO: implement")
}

func  (me *ItemList) IsSelected(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetSelectedItems()  {
  panic("TODO: implement")
}

func  (me *ItemList) MoveItem(from_idx int, to_idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) SetItemCount(count int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemCount()  {
  panic("TODO: implement")
}

func  (me *ItemList) RemoveItem(idx int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) Clear()  {
  panic("TODO: implement")
}

func  (me *ItemList) SortItemsByText()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetFixedColumnWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetFixedColumnWidth()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetSameColumnWidth(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) IsSameColumnWidth()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetMaxTextLines(lines int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetMaxTextLines()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetMaxColumns(amount int, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetMaxColumns()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetSelectMode(mode ItemListSelectMode, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetSelectMode()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetIconMode(mode ItemListIconMode, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetIconMode()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetFixedIconSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetFixedIconSize()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetIconScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetIconScale()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetAllowRmbSelect(allow bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetAllowRmbSelect()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetAllowReselect(allow bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetAllowReselect()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetAllowSearch(allow bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetAllowSearch()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetAutoHeight(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) HasAutoHeight()  {
  panic("TODO: implement")
}

func  (me *ItemList) IsAnythingSelected()  {
  panic("TODO: implement")
}

func  (me *ItemList) GetItemAtPosition(position Vector2, exact bool, )  {
  panic("TODO: implement")
}

func  (me *ItemList) EnsureCurrentIsVisible()  {
  panic("TODO: implement")
}

func  (me *ItemList) GetVScrollBar()  {
  panic("TODO: implement")
}

func  (me *ItemList) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  panic("TODO: implement")
}

func  (me *ItemList) GetTextOverrunBehavior()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
