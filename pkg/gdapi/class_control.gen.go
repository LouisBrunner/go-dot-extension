// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  ControlNOTIFICATION_RESIZED = 40
  ControlNOTIFICATION_MOUSE_ENTER = 41
  ControlNOTIFICATION_MOUSE_EXIT = 42
  ControlNOTIFICATION_FOCUS_ENTER = 43
  ControlNOTIFICATION_FOCUS_EXIT = 44
  ControlNOTIFICATION_THEME_CHANGED = 45
  ControlNOTIFICATION_SCROLL_BEGIN = 47
  ControlNOTIFICATION_SCROLL_END = 48
  ControlNOTIFICATION_LAYOUT_DIRECTION_CHANGED = 49
)

type ControlFocusMode int
const (
  ControlFocusNone ControlFocusMode = 0
  ControlFocusClick ControlFocusMode = 1
  ControlFocusAll ControlFocusMode = 2
)

type ControlCursorShape int
const (
  ControlCursorArrow ControlCursorShape = 0
  ControlCursorIbeam ControlCursorShape = 1
  ControlCursorPointingHand ControlCursorShape = 2
  ControlCursorCross ControlCursorShape = 3
  ControlCursorWait ControlCursorShape = 4
  ControlCursorBusy ControlCursorShape = 5
  ControlCursorDrag ControlCursorShape = 6
  ControlCursorCanDrop ControlCursorShape = 7
  ControlCursorForbidden ControlCursorShape = 8
  ControlCursorVsize ControlCursorShape = 9
  ControlCursorHsize ControlCursorShape = 10
  ControlCursorBdiagsize ControlCursorShape = 11
  ControlCursorFdiagsize ControlCursorShape = 12
  ControlCursorMove ControlCursorShape = 13
  ControlCursorVsplit ControlCursorShape = 14
  ControlCursorHsplit ControlCursorShape = 15
  ControlCursorHelp ControlCursorShape = 16
)

type ControlLayoutPreset int
const (
  ControlPresetTopLeft ControlLayoutPreset = 0
  ControlPresetTopRight ControlLayoutPreset = 1
  ControlPresetBottomLeft ControlLayoutPreset = 2
  ControlPresetBottomRight ControlLayoutPreset = 3
  ControlPresetCenterLeft ControlLayoutPreset = 4
  ControlPresetCenterTop ControlLayoutPreset = 5
  ControlPresetCenterRight ControlLayoutPreset = 6
  ControlPresetCenterBottom ControlLayoutPreset = 7
  ControlPresetCenter ControlLayoutPreset = 8
  ControlPresetLeftWide ControlLayoutPreset = 9
  ControlPresetTopWide ControlLayoutPreset = 10
  ControlPresetRightWide ControlLayoutPreset = 11
  ControlPresetBottomWide ControlLayoutPreset = 12
  ControlPresetVcenterWide ControlLayoutPreset = 13
  ControlPresetHcenterWide ControlLayoutPreset = 14
  ControlPresetFullRect ControlLayoutPreset = 15
)

type ControlLayoutPresetMode int
const (
  ControlPresetModeMinsize ControlLayoutPresetMode = 0
  ControlPresetModeKeepWidth ControlLayoutPresetMode = 1
  ControlPresetModeKeepHeight ControlLayoutPresetMode = 2
  ControlPresetModeKeepSize ControlLayoutPresetMode = 3
)

type ControlSizeFlags int
const (
  ControlSizeShrinkBegin ControlSizeFlags = 0
  ControlSizeFill ControlSizeFlags = 1
  ControlSizeExpand ControlSizeFlags = 2
  ControlSizeExpandFill ControlSizeFlags = 3
  ControlSizeShrinkCenter ControlSizeFlags = 4
  ControlSizeShrinkEnd ControlSizeFlags = 8
)

type ControlMouseFilter int
const (
  ControlMouseFilterStop ControlMouseFilter = 0
  ControlMouseFilterPass ControlMouseFilter = 1
  ControlMouseFilterIgnore ControlMouseFilter = 2
)

type ControlGrowDirection int
const (
  ControlGrowDirectionBegin ControlGrowDirection = 0
  ControlGrowDirectionEnd ControlGrowDirection = 1
  ControlGrowDirectionBoth ControlGrowDirection = 2
)

type ControlAnchor int
const (
  ControlAnchorBegin ControlAnchor = 0
  ControlAnchorEnd ControlAnchor = 1
)

type ControlLayoutDirection int
const (
  ControlLayoutDirectionInherited ControlLayoutDirection = 0
  ControlLayoutDirectionLocale ControlLayoutDirection = 1
  ControlLayoutDirectionLtr ControlLayoutDirection = 2
  ControlLayoutDirectionRtl ControlLayoutDirection = 3
)

type ControlTextDirection int
const (
  ControlTextDirectionInherited ControlTextDirection = 3
  ControlTextDirectionAuto ControlTextDirection = 0
  ControlTextDirectionLtr ControlTextDirection = 1
  ControlTextDirectionRtl ControlTextDirection = 2
)
