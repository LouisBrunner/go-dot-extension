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

func  (me *Tree) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) CreateItem(parent TreeItem, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetRoot() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnCustomMinimumWidth(column int, min_width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnExpand(column int, expand bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnExpandRatio(column int, ratio int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnClipContent(column int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsColumnExpanding(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsColumnClippingContent(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnExpandRatio(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnWidth(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetHideRoot(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsRootHidden() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetNextSelected(from TreeItem, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetSelected() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetSelected(item TreeItem, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetSelectedColumn() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetPressedButton() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetSelectMode(mode TreeSelectMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetSelectMode() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) DeselectAll() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumns(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumns() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetEdited() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetEditedColumn() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) EditSelected(force_edit bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetCustomPopupRect() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetItemAreaRect(item TreeItem, column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetItemAtPosition(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnAtPosition(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetDropSectionAtPosition(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetButtonIdAtPosition(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) EnsureCursorIsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnTitlesVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) AreColumnTitlesVisible() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnTitle(column int, title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnTitle(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnTitleAlignment(column int, title_alignment HorizontalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnTitleAlignment(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnTitleDirection(column int, direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnTitleDirection(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetColumnTitleLanguage(column int, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetColumnTitleLanguage(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetScroll() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) ScrollToItem(item TreeItem, center_on_item bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetHScrollEnabled(h_scroll bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsHScrollEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetVScrollEnabled(h_scroll bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsVScrollEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetHideFolding(hide bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsFoldingHidden() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetEnableRecursiveFolding(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) IsRecursiveFoldingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetDropModeFlags(flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetDropModeFlags() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetAllowRmbSelect(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetAllowRmbSelect() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetAllowReselect(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetAllowReselect() { // TODO: return value
  // TODO: implement
}

func  (me *Tree) SetAllowSearch(allow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tree) GetAllowSearch() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
