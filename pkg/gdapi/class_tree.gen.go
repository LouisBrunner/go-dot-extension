// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Tree struct {
  obj gdc.ObjectPtr
}

func (me *Tree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Tree) BaseClass() string {
  return "Tree"
}



// Enums

type TreeSelectMode int
const (
  TreeSelectModeSelectSingle TreeSelectMode = 0
  TreeSelectModeSelectRow TreeSelectMode = 1
  TreeSelectModeSelectMulti TreeSelectMode = 2
)

type TreeDropModeFlags int
const (
  TreeDropModeFlagsDropModeDisabled TreeDropModeFlags = 0
  TreeDropModeFlagsDropModeOnItem TreeDropModeFlags = 1
  TreeDropModeFlagsDropModeInbetween TreeDropModeFlags = 2
)

func (me *Tree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Tree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Tree) Clear()  {
  panic("TODO: implement")
}

func  (me *Tree) CreateItem(parent TreeItem, index int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetRoot()  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnCustomMinimumWidth(column int, min_width int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnExpand(column int, expand bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnExpandRatio(column int, ratio int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnClipContent(column int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsColumnExpanding(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsColumnClippingContent(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnExpandRatio(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnWidth(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetHideRoot(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsRootHidden()  {
  panic("TODO: implement")
}

func  (me *Tree) GetNextSelected(from TreeItem, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetSelected()  {
  panic("TODO: implement")
}

func  (me *Tree) SetSelected(item TreeItem, column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetSelectedColumn()  {
  panic("TODO: implement")
}

func  (me *Tree) GetPressedButton()  {
  panic("TODO: implement")
}

func  (me *Tree) SetSelectMode(mode TreeSelectMode, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetSelectMode()  {
  panic("TODO: implement")
}

func  (me *Tree) DeselectAll()  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumns(amount int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumns()  {
  panic("TODO: implement")
}

func  (me *Tree) GetEdited()  {
  panic("TODO: implement")
}

func  (me *Tree) GetEditedColumn()  {
  panic("TODO: implement")
}

func  (me *Tree) EditSelected(force_edit bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetCustomPopupRect()  {
  panic("TODO: implement")
}

func  (me *Tree) GetItemAreaRect(item TreeItem, column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetItemAtPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnAtPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetDropSectionAtPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetButtonIdAtPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Tree) EnsureCursorIsVisible()  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnTitlesVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) AreColumnTitlesVisible()  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnTitle(column int, title String, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnTitle(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnTitleAlignment(column int, title_alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnTitleAlignment(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnTitleDirection(column int, direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnTitleDirection(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetColumnTitleLanguage(column int, language String, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetColumnTitleLanguage(column int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetScroll()  {
  panic("TODO: implement")
}

func  (me *Tree) ScrollToItem(item TreeItem, center_on_item bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) SetHScrollEnabled(h_scroll bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsHScrollEnabled()  {
  panic("TODO: implement")
}

func  (me *Tree) SetVScrollEnabled(h_scroll bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsVScrollEnabled()  {
  panic("TODO: implement")
}

func  (me *Tree) SetHideFolding(hide bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsFoldingHidden()  {
  panic("TODO: implement")
}

func  (me *Tree) SetEnableRecursiveFolding(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) IsRecursiveFoldingEnabled()  {
  panic("TODO: implement")
}

func  (me *Tree) SetDropModeFlags(flags int, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetDropModeFlags()  {
  panic("TODO: implement")
}

func  (me *Tree) SetAllowRmbSelect(allow bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetAllowRmbSelect()  {
  panic("TODO: implement")
}

func  (me *Tree) SetAllowReselect(allow bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetAllowReselect()  {
  panic("TODO: implement")
}

func  (me *Tree) SetAllowSearch(allow bool, )  {
  panic("TODO: implement")
}

func  (me *Tree) GetAllowSearch()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
