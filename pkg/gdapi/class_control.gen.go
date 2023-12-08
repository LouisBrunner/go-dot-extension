// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Control struct {
  obj gdc.ObjectPtr
}

func (me *Control) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Control) BaseClass() string {
  return "Control"
}

// TODO: needed?
// const (
// // )

var (
  ControlNotificationResized = "40" // TODO: construct correctly
  ControlNotificationMouseEnter = "41" // TODO: construct correctly
  ControlNotificationMouseExit = "42" // TODO: construct correctly
  ControlNotificationFocusEnter = "43" // TODO: construct correctly
  ControlNotificationFocusExit = "44" // TODO: construct correctly
  ControlNotificationThemeChanged = "45" // TODO: construct correctly
  ControlNotificationScrollBegin = "47" // TODO: construct correctly
  ControlNotificationScrollEnd = "48" // TODO: construct correctly
  ControlNotificationLayoutDirectionChanged = "49" // TODO: construct correctly
)

type ControlFocusMode int
const (
  ControlFocusModeFocusNone ControlFocusMode = 0
  ControlFocusModeFocusClick ControlFocusMode = 1
  ControlFocusModeFocusAll ControlFocusMode = 2
)

type ControlCursorShape int
const (
  ControlCursorShapeCursorArrow ControlCursorShape = 0
  ControlCursorShapeCursorIbeam ControlCursorShape = 1
  ControlCursorShapeCursorPointingHand ControlCursorShape = 2
  ControlCursorShapeCursorCross ControlCursorShape = 3
  ControlCursorShapeCursorWait ControlCursorShape = 4
  ControlCursorShapeCursorBusy ControlCursorShape = 5
  ControlCursorShapeCursorDrag ControlCursorShape = 6
  ControlCursorShapeCursorCanDrop ControlCursorShape = 7
  ControlCursorShapeCursorForbidden ControlCursorShape = 8
  ControlCursorShapeCursorVsize ControlCursorShape = 9
  ControlCursorShapeCursorHsize ControlCursorShape = 10
  ControlCursorShapeCursorBdiagsize ControlCursorShape = 11
  ControlCursorShapeCursorFdiagsize ControlCursorShape = 12
  ControlCursorShapeCursorMove ControlCursorShape = 13
  ControlCursorShapeCursorVsplit ControlCursorShape = 14
  ControlCursorShapeCursorHsplit ControlCursorShape = 15
  ControlCursorShapeCursorHelp ControlCursorShape = 16
)

type ControlLayoutPreset int
const (
  ControlLayoutPresetPresetTopLeft ControlLayoutPreset = 0
  ControlLayoutPresetPresetTopRight ControlLayoutPreset = 1
  ControlLayoutPresetPresetBottomLeft ControlLayoutPreset = 2
  ControlLayoutPresetPresetBottomRight ControlLayoutPreset = 3
  ControlLayoutPresetPresetCenterLeft ControlLayoutPreset = 4
  ControlLayoutPresetPresetCenterTop ControlLayoutPreset = 5
  ControlLayoutPresetPresetCenterRight ControlLayoutPreset = 6
  ControlLayoutPresetPresetCenterBottom ControlLayoutPreset = 7
  ControlLayoutPresetPresetCenter ControlLayoutPreset = 8
  ControlLayoutPresetPresetLeftWide ControlLayoutPreset = 9
  ControlLayoutPresetPresetTopWide ControlLayoutPreset = 10
  ControlLayoutPresetPresetRightWide ControlLayoutPreset = 11
  ControlLayoutPresetPresetBottomWide ControlLayoutPreset = 12
  ControlLayoutPresetPresetVcenterWide ControlLayoutPreset = 13
  ControlLayoutPresetPresetHcenterWide ControlLayoutPreset = 14
  ControlLayoutPresetPresetFullRect ControlLayoutPreset = 15
)

type ControlLayoutPresetMode int
const (
  ControlLayoutPresetModePresetModeMinsize ControlLayoutPresetMode = 0
  ControlLayoutPresetModePresetModeKeepWidth ControlLayoutPresetMode = 1
  ControlLayoutPresetModePresetModeKeepHeight ControlLayoutPresetMode = 2
  ControlLayoutPresetModePresetModeKeepSize ControlLayoutPresetMode = 3
)

type ControlSizeFlags int
const (
  ControlSizeFlagsSizeShrinkBegin ControlSizeFlags = 0
  ControlSizeFlagsSizeFill ControlSizeFlags = 1
  ControlSizeFlagsSizeExpand ControlSizeFlags = 2
  ControlSizeFlagsSizeExpandFill ControlSizeFlags = 3
  ControlSizeFlagsSizeShrinkCenter ControlSizeFlags = 4
  ControlSizeFlagsSizeShrinkEnd ControlSizeFlags = 8
)

type ControlMouseFilter int
const (
  ControlMouseFilterMouseFilterStop ControlMouseFilter = 0
  ControlMouseFilterMouseFilterPass ControlMouseFilter = 1
  ControlMouseFilterMouseFilterIgnore ControlMouseFilter = 2
)

type ControlGrowDirection int
const (
  ControlGrowDirectionGrowDirectionBegin ControlGrowDirection = 0
  ControlGrowDirectionGrowDirectionEnd ControlGrowDirection = 1
  ControlGrowDirectionGrowDirectionBoth ControlGrowDirection = 2
)

type ControlAnchor int
const (
  ControlAnchorAnchorBegin ControlAnchor = 0
  ControlAnchorAnchorEnd ControlAnchor = 1
)

type ControlLayoutDirection int
const (
  ControlLayoutDirectionLayoutDirectionInherited ControlLayoutDirection = 0
  ControlLayoutDirectionLayoutDirectionLocale ControlLayoutDirection = 1
  ControlLayoutDirectionLayoutDirectionLtr ControlLayoutDirection = 2
  ControlLayoutDirectionLayoutDirectionRtl ControlLayoutDirection = 3
)

type ControlTextDirection int
const (
  ControlTextDirectionTextDirectionInherited ControlTextDirection = 3
  ControlTextDirectionTextDirectionAuto ControlTextDirection = 0
  ControlTextDirectionTextDirectionLtr ControlTextDirection = 1
  ControlTextDirectionTextDirectionRtl ControlTextDirection = 2
)

func  (me *Control) XHasPoint(point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XStructuredTextParser(args Array, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XGetMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) XGetTooltip(at_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XGetDragData(at_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XCanDropData(at_position Vector2, data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XDropData(at_position Vector2, data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XMakeCustomTooltip(for_text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) XGuiInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AcceptEvent() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetCombinedMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetAnchorsPreset(preset ControlLayoutPreset, keep_offsets bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetAnchorsAndOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetAnchor(side Side, anchor float32, keep_offset bool, push_opposite_anchor bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetAnchor(side Side, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetOffset(side Side, offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetOffset(offset Side, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetAnchorAndOffset(side Side, anchor float32, offset float32, push_opposite_anchor bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetBegin(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetEnd(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetPosition(position Vector2, keep_offsets bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetSize(size Vector2, keep_offsets bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) ResetSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetCustomMinimumSize(size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetGlobalPosition(position Vector2, keep_offsets bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetRotation(radians float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetRotationDegrees(degrees float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetScale(scale Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetPivotOffset(pivot_offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetBegin() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetEnd() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetRotationDegrees() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetScale() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetPivotOffset() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetCustomMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetParentAreaSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetGlobalPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetScreenPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetRect() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetGlobalRect() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetFocusMode(mode ControlFocusMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetFocusMode() { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GrabFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) ReleaseFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) FindPrevValidFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) FindNextValidFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetHSizeFlags(flags ControlSizeFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetHSizeFlags() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetStretchRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetStretchRatio() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetVSizeFlags(flags ControlSizeFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetVSizeFlags() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetTheme(theme Theme, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetTheme() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetThemeTypeVariation(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeTypeVariation() { // TODO: return value
  // TODO: implement
}

func  (me *Control) BeginBulkThemeOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Control) EndBulkThemeOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeIconOverride(name StringName, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeFontOverride(name StringName, font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeFontSizeOverride(name StringName, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeColorOverride(name StringName, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) AddThemeConstantOverride(name StringName, constant int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeIconOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeStyleboxOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeFontOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeFontSizeOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeColorOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) RemoveThemeConstantOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeIconOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeStyleboxOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeFontOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeFontSizeOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeColorOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeConstantOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) HasThemeConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeDefaultBaseScale() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeDefaultFont() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetThemeDefaultFontSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetParentControl() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetHGrowDirection(direction ControlGrowDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetHGrowDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetVGrowDirection(direction ControlGrowDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetVGrowDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetTooltipText(hint String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetTooltipText() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetTooltip(at_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetDefaultCursorShape(shape ControlCursorShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetDefaultCursorShape() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetCursorShape(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetFocusNeighbor(side Side, neighbor NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetFocusNeighbor(side Side, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetFocusNext(next NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetFocusNext() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetFocusPrevious(previous NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetFocusPrevious() { // TODO: return value
  // TODO: implement
}

func  (me *Control) ForceDrag(data Variant, preview Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetMouseFilter(filter ControlMouseFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetMouseFilter() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetForcePassScrollEvents(force_pass_scroll_events bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsForcePassScrollEvents() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetClipContents(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsClippingContents() { // TODO: return value
  // TODO: implement
}

func  (me *Control) GrabClickFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetDragForwarding(drag_func Callable, can_drop_func Callable, drop_func Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetDragPreview(control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsDragSuccessful() { // TODO: return value
  // TODO: implement
}

func  (me *Control) WarpMouse(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetShortcutContext(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetShortcutContext() { // TODO: return value
  // TODO: implement
}

func  (me *Control) UpdateMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetLayoutDirection(direction ControlLayoutDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) GetLayoutDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsLayoutRtl() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetAutoTranslate(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsAutoTranslating() { // TODO: return value
  // TODO: implement
}

func  (me *Control) SetLocalizeNumeralSystem(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Control) IsLocalizingNumeralSystem() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
