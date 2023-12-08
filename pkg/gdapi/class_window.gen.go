// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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

// TODO: needed?
// const (
// // )

var (
  WindowNotificationVisibilityChanged = "30" // TODO: construct correctly
  WindowNotificationThemeChanged = "32" // TODO: construct correctly
)

type WindowMode int
const (
  WindowModeModeWindowed WindowMode = 0
  WindowModeModeMinimized WindowMode = 1
  WindowModeModeMaximized WindowMode = 2
  WindowModeModeFullscreen WindowMode = 3
  WindowModeModeExclusiveFullscreen WindowMode = 4
)

type WindowFlags int
const (
  WindowFlagsFlagResizeDisabled WindowFlags = 0
  WindowFlagsFlagBorderless WindowFlags = 1
  WindowFlagsFlagAlwaysOnTop WindowFlags = 2
  WindowFlagsFlagTransparent WindowFlags = 3
  WindowFlagsFlagNoFocus WindowFlags = 4
  WindowFlagsFlagPopup WindowFlags = 5
  WindowFlagsFlagExtendToTitle WindowFlags = 6
  WindowFlagsFlagMousePassthrough WindowFlags = 7
  WindowFlagsFlagMax WindowFlags = 8
)

type WindowContentScaleMode int
const (
  WindowContentScaleModeContentScaleModeDisabled WindowContentScaleMode = 0
  WindowContentScaleModeContentScaleModeCanvasItems WindowContentScaleMode = 1
  WindowContentScaleModeContentScaleModeViewport WindowContentScaleMode = 2
)

type WindowContentScaleAspect int
const (
  WindowContentScaleAspectContentScaleAspectIgnore WindowContentScaleAspect = 0
  WindowContentScaleAspectContentScaleAspectKeep WindowContentScaleAspect = 1
  WindowContentScaleAspectContentScaleAspectKeepWidth WindowContentScaleAspect = 2
  WindowContentScaleAspectContentScaleAspectKeepHeight WindowContentScaleAspect = 3
  WindowContentScaleAspectContentScaleAspectExpand WindowContentScaleAspect = 4
)

type WindowLayoutDirection int
const (
  WindowLayoutDirectionLayoutDirectionInherited WindowLayoutDirection = 0
  WindowLayoutDirectionLayoutDirectionLocale WindowLayoutDirection = 1
  WindowLayoutDirectionLayoutDirectionLtr WindowLayoutDirection = 2
  WindowLayoutDirectionLayoutDirectionRtl WindowLayoutDirection = 3
)

type WindowWindowInitialPosition int
const (
  WindowWindowInitialPositionWindowInitialPositionAbsolute WindowWindowInitialPosition = 0
  WindowWindowInitialPositionWindowInitialPositionCenterPrimaryScreen WindowWindowInitialPosition = 1
  WindowWindowInitialPositionWindowInitialPositionCenterMainWindowScreen WindowWindowInitialPosition = 2
  WindowWindowInitialPositionWindowInitialPositionCenterOtherScreen WindowWindowInitialPosition = 3
  WindowWindowInitialPositionWindowInitialPositionCenterScreenWithMouseFocus WindowWindowInitialPosition = 4
  WindowWindowInitialPositionWindowInitialPositionCenterScreenWithKeyboardFocus WindowWindowInitialPosition = 5
)

func  (me *Window) SetTitle(title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetTitle() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetWindowId() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetInitialPosition(initial_position WindowWindowInitialPosition, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetInitialPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetCurrentScreen(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetCurrentScreen() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetPosition(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) ResetSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetPositionWithDecorations() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetSizeWithDecorations() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetMaxSize(max_size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetMaxSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetMinSize(min_size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetMinSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetMode(mode WindowMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetMode() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetFlag(flag WindowFlags, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetFlag(flag WindowFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsMaximizeAllowed() { // TODO: return value
  // TODO: implement
}

func  (me *Window) RequestAttention() { // TODO: return value
  // TODO: implement
}

func  (me *Window) MoveToForeground() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *Window) Hide() { // TODO: return value
  // TODO: implement
}

func  (me *Window) Show() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetTransient(transient bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsTransient() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetExclusive(exclusive bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsExclusive() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetUnparentWhenInvisible(unparent bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) CanDraw() { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GrabFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetImeActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetImePosition(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsEmbedded() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetContentsMinimumSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetContentScaleSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetContentScaleSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetContentScaleMode(mode WindowContentScaleMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetContentScaleMode() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetContentScaleAspect(aspect WindowContentScaleAspect, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetContentScaleAspect() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetContentScaleFactor(factor float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetContentScaleFactor() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetUseFontOversampling(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsUsingFontOversampling() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetMousePassthroughPolygon(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetMousePassthroughPolygon() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetWrapControls(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsWrappingControls() { // TODO: return value
  // TODO: implement
}

func  (me *Window) ChildControlsChanged() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetTheme(theme Theme, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetTheme() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetThemeTypeVariation(theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeTypeVariation() { // TODO: return value
  // TODO: implement
}

func  (me *Window) BeginBulkThemeOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Window) EndBulkThemeOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeIconOverride(name StringName, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeFontOverride(name StringName, font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeFontSizeOverride(name StringName, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeColorOverride(name StringName, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) AddThemeConstantOverride(name StringName, constant int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeIconOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeStyleboxOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeFontOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeFontSizeOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeColorOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) RemoveThemeConstantOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeIconOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeStyleboxOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeFontOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeFontSizeOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeColorOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeConstantOverride(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeIcon(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeStylebox(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeFont(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeFontSize(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeColor(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) HasThemeConstant(name StringName, theme_type StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeDefaultBaseScale() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeDefaultFont() { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetThemeDefaultFontSize() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetLayoutDirection(direction WindowLayoutDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) GetLayoutDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsLayoutRtl() { // TODO: return value
  // TODO: implement
}

func  (me *Window) SetAutoTranslate(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) IsAutoTranslating() { // TODO: return value
  // TODO: implement
}

func  (me *Window) Popup(rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupOnParent(parent_rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupCentered(minsize Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupCenteredRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupCenteredClamped(minsize Vector2i, fallback_ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupExclusive(from_node Node, rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupExclusiveOnParent(from_node Node, parent_rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupExclusiveCentered(from_node Node, minsize Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupExclusiveCenteredRatio(from_node Node, ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Window) PopupExclusiveCenteredClamped(from_node Node, minsize Vector2i, fallback_ratio float32, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
