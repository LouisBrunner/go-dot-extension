// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TreeItem struct {
  obj gdc.ObjectPtr
}

func (me *TreeItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TreeItem) BaseClass() string {
  return "TreeItem"
}

type TreeItemTreeCellMode int
const (
  TreeItemTreeCellModeCellModeString TreeItemTreeCellMode = 0
  TreeItemTreeCellModeCellModeCheck TreeItemTreeCellMode = 1
  TreeItemTreeCellModeCellModeRange TreeItemTreeCellMode = 2
  TreeItemTreeCellModeCellModeIcon TreeItemTreeCellMode = 3
  TreeItemTreeCellModeCellModeCustom TreeItemTreeCellMode = 4
)

func  (me *TreeItem) SetCellMode(column int, mode TreeItemTreeCellMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCellMode(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetEditMultiline(column int, multiline bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsEditMultiline(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetChecked(column int, checked bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetIndeterminate(column int, indeterminate bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsChecked(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsIndeterminate(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) PropagateCheck(column int, emit_signal bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetText(column int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetText(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetTextDirection(column int, direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetTextDirection(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetAutowrapMode(column int, autowrap_mode TextServerAutowrapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetAutowrapMode(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetTextOverrunBehavior(column int, overrun_behavior TextServerOverrunBehavior, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetTextOverrunBehavior(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetStructuredTextBidiOverride(column int, parser TextServerStructuredTextParser, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetStructuredTextBidiOverride(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetStructuredTextBidiOverrideOptions(column int, args Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetStructuredTextBidiOverrideOptions(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetLanguage(column int, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetLanguage(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetSuffix(column int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetSuffix(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetIcon(column int, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetIcon(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetIconRegion(column int, region Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetIconRegion(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetIconMaxWidth(column int, width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetIconMaxWidth(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetIconModulate(column int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetIconModulate(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetRange(column int, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetRange(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetRangeConfig(column int, min float32, max float32, step float32, expr bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetRangeConfig(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetMetadata(column int, meta Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetMetadata(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomDraw(column int, object Object, callback StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCollapsed(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsCollapsed() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCollapsedRecursive(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsAnyCollapsed(only_visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetVisible(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) UncollapseTree() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomMinimumHeight(height int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCustomMinimumHeight() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetSelectable(column int, selectable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsSelectable(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsSelected(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) Select(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) Deselect(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetEditable(column int, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsEditable(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomColor(column int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCustomColor(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) ClearCustomColor(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomFont(column int, font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCustomFont(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomFontSize(column int, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCustomFontSize(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomBgColor(column int, color Color, just_outline bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) ClearCustomBgColor(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetCustomBgColor(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetCustomAsButton(column int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsCustomSetAsButton(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) AddButton(column int, button Texture2D, id int, disabled bool, tooltip_text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetButtonCount(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetButtonTooltipText(column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetButtonId(column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetButtonById(column int, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetButton(column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetButton(column int, button_index int, button Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) EraseButton(column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetButtonDisabled(column int, button_index int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetButtonColor(column int, button_index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsButtonDisabled(column int, button_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetTooltipText(column int, tooltip String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetTooltipText(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetTextAlignment(column int, text_alignment HorizontalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetTextAlignment(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetExpandRight(column int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetExpandRight(column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) SetDisableFolding(disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) IsFoldingDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) CreateChild(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) AddChild(child TreeItem, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) RemoveChild(child TreeItem, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetTree() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetNext() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetPrev() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetParent() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetFirstChild() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetNextInTree(wrap bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetPrevInTree(wrap bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetNextVisible(wrap bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetPrevVisible(wrap bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetChild(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetChildCount() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetChildren() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) GetIndex() { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) MoveBefore(item TreeItem, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) MoveAfter(item TreeItem, ) { // TODO: return value
  // TODO: implement
}

func  (me *TreeItem) CallRecursive(method StringName, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
