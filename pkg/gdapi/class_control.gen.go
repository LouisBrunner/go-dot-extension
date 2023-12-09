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



// Constants

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

// Enums

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

func (me *Control) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Control) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Control) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Control) XHasPoint(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) XStructuredTextParser(args Array, text String, )  {
  panic("TODO: implement")
}

func  (me *Control) XGetMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Control) XGetTooltip(at_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) XGetDragData(at_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) XCanDropData(at_position Vector2, data Variant, )  {
  panic("TODO: implement")
}

func  (me *Control) XDropData(at_position Vector2, data Variant, )  {
  panic("TODO: implement")
}

func  (me *Control) XMakeCustomTooltip(for_text String, )  {
  panic("TODO: implement")
}

func  (me *Control) XGuiInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Control) AcceptEvent()  {
  panic("TODO: implement")
}

func  (me *Control) GetMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Control) GetCombinedMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Control) SetAnchorsPreset(preset ControlLayoutPreset, keep_offsets bool, )  {
  panic("TODO: implement")
}

func  (me *Control) SetOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, )  {
  panic("TODO: implement")
}

func  (me *Control) SetAnchorsAndOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int, )  {
  panic("TODO: implement")
}

func  (me *Control) SetAnchor(side Side, anchor float32, keep_offset bool, push_opposite_anchor bool, )  {
  panic("TODO: implement")
}

func  (me *Control) GetAnchor(side Side, )  {
  panic("TODO: implement")
}

func  (me *Control) SetOffset(side Side, offset float32, )  {
  panic("TODO: implement")
}

func  (me *Control) GetOffset(offset Side, )  {
  panic("TODO: implement")
}

func  (me *Control) SetAnchorAndOffset(side Side, anchor float32, offset float32, push_opposite_anchor bool, )  {
  panic("TODO: implement")
}

func  (me *Control) SetBegin(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetEnd(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetPosition(position Vector2, keep_offsets bool, )  {
  panic("TODO: implement")
}

func  (me *Control) SetSize(size Vector2, keep_offsets bool, )  {
  panic("TODO: implement")
}

func  (me *Control) ResetSize()  {
  panic("TODO: implement")
}

func  (me *Control) SetCustomMinimumSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetGlobalPosition(position Vector2, keep_offsets bool, )  {
  panic("TODO: implement")
}

func  (me *Control) SetRotation(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Control) SetRotationDegrees(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *Control) SetScale(scale Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetPivotOffset(pivot_offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) GetBegin()  {
  panic("TODO: implement")
}

func  (me *Control) GetEnd()  {
  panic("TODO: implement")
}

func  (me *Control) GetPosition()  {
  panic("TODO: implement")
}

func  (me *Control) GetSize()  {
  panic("TODO: implement")
}

func  (me *Control) GetRotation()  {
  panic("TODO: implement")
}

func  (me *Control) GetRotationDegrees()  {
  panic("TODO: implement")
}

func  (me *Control) GetScale()  {
  panic("TODO: implement")
}

func  (me *Control) GetPivotOffset()  {
  panic("TODO: implement")
}

func  (me *Control) GetCustomMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Control) GetParentAreaSize()  {
  panic("TODO: implement")
}

func  (me *Control) GetGlobalPosition()  {
  panic("TODO: implement")
}

func  (me *Control) GetScreenPosition()  {
  panic("TODO: implement")
}

func  (me *Control) GetRect()  {
  panic("TODO: implement")
}

func  (me *Control) GetGlobalRect()  {
  panic("TODO: implement")
}

func  (me *Control) SetFocusMode(mode ControlFocusMode, )  {
  panic("TODO: implement")
}

func  (me *Control) GetFocusMode()  {
  panic("TODO: implement")
}

func  (me *Control) HasFocus()  {
  panic("TODO: implement")
}

func  (me *Control) GrabFocus()  {
  panic("TODO: implement")
}

func  (me *Control) ReleaseFocus()  {
  panic("TODO: implement")
}

func  (me *Control) FindPrevValidFocus()  {
  panic("TODO: implement")
}

func  (me *Control) FindNextValidFocus()  {
  panic("TODO: implement")
}

func  (me *Control) SetHSizeFlags(flags ControlSizeFlags, )  {
  panic("TODO: implement")
}

func  (me *Control) GetHSizeFlags()  {
  panic("TODO: implement")
}

func  (me *Control) SetStretchRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Control) GetStretchRatio()  {
  panic("TODO: implement")
}

func  (me *Control) SetVSizeFlags(flags ControlSizeFlags, )  {
  panic("TODO: implement")
}

func  (me *Control) GetVSizeFlags()  {
  panic("TODO: implement")
}

func  (me *Control) SetTheme(theme Theme, )  {
  panic("TODO: implement")
}

func  (me *Control) GetTheme()  {
  panic("TODO: implement")
}

func  (me *Control) SetThemeTypeVariation(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeTypeVariation()  {
  panic("TODO: implement")
}

func  (me *Control) BeginBulkThemeOverride()  {
  panic("TODO: implement")
}

func  (me *Control) EndBulkThemeOverride()  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeIconOverride(name StringName, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, )  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeFontOverride(name StringName, font Font, )  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeFontSizeOverride(name StringName, font_size int, )  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeColorOverride(name StringName, color Color, )  {
  panic("TODO: implement")
}

func  (me *Control) AddThemeConstantOverride(name StringName, constant int, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeIconOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeStyleboxOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeFontOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeFontSizeOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeColorOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) RemoveThemeConstantOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeIconOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeStyleboxOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeFontOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeFontSizeOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeColorOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeConstantOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) HasThemeConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeDefaultBaseScale()  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeDefaultFont()  {
  panic("TODO: implement")
}

func  (me *Control) GetThemeDefaultFontSize()  {
  panic("TODO: implement")
}

func  (me *Control) GetParentControl()  {
  panic("TODO: implement")
}

func  (me *Control) SetHGrowDirection(direction ControlGrowDirection, )  {
  panic("TODO: implement")
}

func  (me *Control) GetHGrowDirection()  {
  panic("TODO: implement")
}

func  (me *Control) SetVGrowDirection(direction ControlGrowDirection, )  {
  panic("TODO: implement")
}

func  (me *Control) GetVGrowDirection()  {
  panic("TODO: implement")
}

func  (me *Control) SetTooltipText(hint String, )  {
  panic("TODO: implement")
}

func  (me *Control) GetTooltipText()  {
  panic("TODO: implement")
}

func  (me *Control) GetTooltip(at_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetDefaultCursorShape(shape ControlCursorShape, )  {
  panic("TODO: implement")
}

func  (me *Control) GetDefaultCursorShape()  {
  panic("TODO: implement")
}

func  (me *Control) GetCursorShape(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetFocusNeighbor(side Side, neighbor NodePath, )  {
  panic("TODO: implement")
}

func  (me *Control) GetFocusNeighbor(side Side, )  {
  panic("TODO: implement")
}

func  (me *Control) SetFocusNext(next NodePath, )  {
  panic("TODO: implement")
}

func  (me *Control) GetFocusNext()  {
  panic("TODO: implement")
}

func  (me *Control) SetFocusPrevious(previous NodePath, )  {
  panic("TODO: implement")
}

func  (me *Control) GetFocusPrevious()  {
  panic("TODO: implement")
}

func  (me *Control) ForceDrag(data Variant, preview Control, )  {
  panic("TODO: implement")
}

func  (me *Control) SetMouseFilter(filter ControlMouseFilter, )  {
  panic("TODO: implement")
}

func  (me *Control) GetMouseFilter()  {
  panic("TODO: implement")
}

func  (me *Control) SetForcePassScrollEvents(force_pass_scroll_events bool, )  {
  panic("TODO: implement")
}

func  (me *Control) IsForcePassScrollEvents()  {
  panic("TODO: implement")
}

func  (me *Control) SetClipContents(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Control) IsClippingContents()  {
  panic("TODO: implement")
}

func  (me *Control) GrabClickFocus()  {
  panic("TODO: implement")
}

func  (me *Control) SetDragForwarding(drag_func Callable, can_drop_func Callable, drop_func Callable, )  {
  panic("TODO: implement")
}

func  (me *Control) SetDragPreview(control Control, )  {
  panic("TODO: implement")
}

func  (me *Control) IsDragSuccessful()  {
  panic("TODO: implement")
}

func  (me *Control) WarpMouse(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Control) SetShortcutContext(node Node, )  {
  panic("TODO: implement")
}

func  (me *Control) GetShortcutContext()  {
  panic("TODO: implement")
}

func  (me *Control) UpdateMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Control) SetLayoutDirection(direction ControlLayoutDirection, )  {
  panic("TODO: implement")
}

func  (me *Control) GetLayoutDirection()  {
  panic("TODO: implement")
}

func  (me *Control) IsLayoutRtl()  {
  panic("TODO: implement")
}

func  (me *Control) SetAutoTranslate(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Control) IsAutoTranslating()  {
  panic("TODO: implement")
}

func  (me *Control) SetLocalizeNumeralSystem(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Control) IsLocalizingNumeralSystem()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
