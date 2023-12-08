// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Window struct {
  obj gdc.ObjectPtr
}

func (me *Window) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Window) BaseClass() string {
  return "Window"
}

const (
  WindowNOTIFICATION_VISIBILITY_CHANGED = 30
  WindowNOTIFICATION_THEME_CHANGED = 32
)

type WindowMode int
const (
  WindowModeWindowed WindowMode = 0
  WindowModeMinimized WindowMode = 1
  WindowModeMaximized WindowMode = 2
  WindowModeFullscreen WindowMode = 3
  WindowModeExclusiveFullscreen WindowMode = 4
)

type WindowFlags int
const (
  WindowFlagResizeDisabled WindowFlags = 0
  WindowFlagBorderless WindowFlags = 1
  WindowFlagAlwaysOnTop WindowFlags = 2
  WindowFlagTransparent WindowFlags = 3
  WindowFlagNoFocus WindowFlags = 4
  WindowFlagPopup WindowFlags = 5
  WindowFlagExtendToTitle WindowFlags = 6
  WindowFlagMousePassthrough WindowFlags = 7
  WindowFlagMax WindowFlags = 8
)

type WindowContentScaleMode int
const (
  WindowContentScaleModeDisabled WindowContentScaleMode = 0
  WindowContentScaleModeCanvasItems WindowContentScaleMode = 1
  WindowContentScaleModeViewport WindowContentScaleMode = 2
)

type WindowContentScaleAspect int
const (
  WindowContentScaleAspectIgnore WindowContentScaleAspect = 0
  WindowContentScaleAspectKeep WindowContentScaleAspect = 1
  WindowContentScaleAspectKeepWidth WindowContentScaleAspect = 2
  WindowContentScaleAspectKeepHeight WindowContentScaleAspect = 3
  WindowContentScaleAspectExpand WindowContentScaleAspect = 4
)

type WindowLayoutDirection int
const (
  WindowLayoutDirectionInherited WindowLayoutDirection = 0
  WindowLayoutDirectionLocale WindowLayoutDirection = 1
  WindowLayoutDirectionLtr WindowLayoutDirection = 2
  WindowLayoutDirectionRtl WindowLayoutDirection = 3
)

type WindowWindowInitialPosition int
const (
  WindowWindowInitialPositionAbsolute WindowWindowInitialPosition = 0
  WindowWindowInitialPositionCenterPrimaryScreen WindowWindowInitialPosition = 1
  WindowWindowInitialPositionCenterMainWindowScreen WindowWindowInitialPosition = 2
  WindowWindowInitialPositionCenterOtherScreen WindowWindowInitialPosition = 3
  WindowWindowInitialPositionCenterScreenWithMouseFocus WindowWindowInitialPosition = 4
  WindowWindowInitialPositionCenterScreenWithKeyboardFocus WindowWindowInitialPosition = 5
)
