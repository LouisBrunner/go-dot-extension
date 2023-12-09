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



// Enums

type TreeItemTreeCellMode int
const (
  TreeItemTreeCellModeCellModeString TreeItemTreeCellMode = 0
  TreeItemTreeCellModeCellModeCheck TreeItemTreeCellMode = 1
  TreeItemTreeCellModeCellModeRange TreeItemTreeCellMode = 2
  TreeItemTreeCellModeCellModeIcon TreeItemTreeCellMode = 3
  TreeItemTreeCellModeCellModeCustom TreeItemTreeCellMode = 4
)

func (me *TreeItem) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TreeItem) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TreeItem) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TreeItem) SetCellMode(column int, mode TreeItemTreeCellMode, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCellMode(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetEditMultiline(column int, multiline bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsEditMultiline(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetChecked(column int, checked bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetIndeterminate(column int, indeterminate bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsChecked(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsIndeterminate(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) PropagateCheck(column int, emit_signal bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetText(column int, text String, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetText(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetTextDirection(column int, direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetTextDirection(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetAutowrapMode(column int, autowrap_mode TextServerAutowrapMode, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetAutowrapMode(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetTextOverrunBehavior(column int, overrun_behavior TextServerOverrunBehavior, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetTextOverrunBehavior(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetStructuredTextBidiOverride(column int, parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetStructuredTextBidiOverride(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetStructuredTextBidiOverrideOptions(column int, args Array, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetStructuredTextBidiOverrideOptions(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetLanguage(column int, language String, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetLanguage(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetSuffix(column int, text String, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetSuffix(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetIcon(column int, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetIcon(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetIconRegion(column int, region Rect2, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetIconRegion(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetIconMaxWidth(column int, width int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetIconMaxWidth(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetIconModulate(column int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetIconModulate(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetRange(column int, value float32, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetRange(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetRangeConfig(column int, min float32, max float32, step float32, expr bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetRangeConfig(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetMetadata(column int, meta Variant, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetMetadata(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomDraw(column int, object Object, callback StringName, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCollapsed(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsCollapsed()  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCollapsedRecursive(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsAnyCollapsed(only_visible bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetVisible(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsVisible()  {
  panic("TODO: implement")
}

func  (me *TreeItem) UncollapseTree()  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomMinimumHeight(height int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCustomMinimumHeight()  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetSelectable(column int, selectable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsSelectable(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsSelected(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) Select(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) Deselect(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetEditable(column int, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsEditable(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomColor(column int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCustomColor(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) ClearCustomColor(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomFont(column int, font Font, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCustomFont(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomFontSize(column int, font_size int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCustomFontSize(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomBgColor(column int, color Color, just_outline bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) ClearCustomBgColor(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetCustomBgColor(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetCustomAsButton(column int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsCustomSetAsButton(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) AddButton(column int, button Texture2D, id int, disabled bool, tooltip_text String, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetButtonCount(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetButtonTooltipText(column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetButtonId(column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetButtonById(column int, id int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetButton(column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetButton(column int, button_index int, button Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) EraseButton(column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetButtonDisabled(column int, button_index int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetButtonColor(column int, button_index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsButtonDisabled(column int, button_index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetTooltipText(column int, tooltip String, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetTooltipText(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetTextAlignment(column int, text_alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetTextAlignment(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetExpandRight(column int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetExpandRight(column int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) SetDisableFolding(disable bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) IsFoldingDisabled()  {
  panic("TODO: implement")
}

func  (me *TreeItem) CreateChild(index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) AddChild(child TreeItem, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) RemoveChild(child TreeItem, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetTree()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetNext()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetPrev()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetParent()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetFirstChild()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetNextInTree(wrap bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetPrevInTree(wrap bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetNextVisible(wrap bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetPrevVisible(wrap bool, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetChild(index int, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetChildCount()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetChildren()  {
  panic("TODO: implement")
}

func  (me *TreeItem) GetIndex()  {
  panic("TODO: implement")
}

func  (me *TreeItem) MoveBefore(item TreeItem, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) MoveAfter(item TreeItem, )  {
  panic("TODO: implement")
}

func  (me *TreeItem) CallRecursive(method StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
