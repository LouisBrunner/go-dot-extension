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



// Constants

var (
  WindowNotificationVisibilityChanged = "30" // TODO: construct correctly
  WindowNotificationThemeChanged = "32" // TODO: construct correctly
)

// Enums

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

func (me *Window) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Window) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Window) SetTitle(title String, )  {
  panic("TODO: implement")
}

func  (me *Window) GetTitle()  {
  panic("TODO: implement")
}

func  (me *Window) GetWindowId()  {
  panic("TODO: implement")
}

func  (me *Window) SetInitialPosition(initial_position WindowWindowInitialPosition, )  {
  panic("TODO: implement")
}

func  (me *Window) GetInitialPosition()  {
  panic("TODO: implement")
}

func  (me *Window) SetCurrentScreen(index int, )  {
  panic("TODO: implement")
}

func  (me *Window) GetCurrentScreen()  {
  panic("TODO: implement")
}

func  (me *Window) SetPosition(position Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) GetPosition()  {
  panic("TODO: implement")
}

func  (me *Window) SetSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) GetSize()  {
  panic("TODO: implement")
}

func  (me *Window) ResetSize()  {
  panic("TODO: implement")
}

func  (me *Window) GetPositionWithDecorations()  {
  panic("TODO: implement")
}

func  (me *Window) GetSizeWithDecorations()  {
  panic("TODO: implement")
}

func  (me *Window) SetMaxSize(max_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) GetMaxSize()  {
  panic("TODO: implement")
}

func  (me *Window) SetMinSize(min_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) GetMinSize()  {
  panic("TODO: implement")
}

func  (me *Window) SetMode(mode WindowMode, )  {
  panic("TODO: implement")
}

func  (me *Window) GetMode()  {
  panic("TODO: implement")
}

func  (me *Window) SetFlag(flag WindowFlags, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Window) GetFlag(flag WindowFlags, )  {
  panic("TODO: implement")
}

func  (me *Window) IsMaximizeAllowed()  {
  panic("TODO: implement")
}

func  (me *Window) RequestAttention()  {
  panic("TODO: implement")
}

func  (me *Window) MoveToForeground()  {
  panic("TODO: implement")
}

func  (me *Window) SetVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsVisible()  {
  panic("TODO: implement")
}

func  (me *Window) Hide()  {
  panic("TODO: implement")
}

func  (me *Window) Show()  {
  panic("TODO: implement")
}

func  (me *Window) SetTransient(transient bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsTransient()  {
  panic("TODO: implement")
}

func  (me *Window) SetExclusive(exclusive bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsExclusive()  {
  panic("TODO: implement")
}

func  (me *Window) SetUnparentWhenInvisible(unparent bool, )  {
  panic("TODO: implement")
}

func  (me *Window) CanDraw()  {
  panic("TODO: implement")
}

func  (me *Window) HasFocus()  {
  panic("TODO: implement")
}

func  (me *Window) GrabFocus()  {
  panic("TODO: implement")
}

func  (me *Window) SetImeActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *Window) SetImePosition(position Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) IsEmbedded()  {
  panic("TODO: implement")
}

func  (me *Window) GetContentsMinimumSize()  {
  panic("TODO: implement")
}

func  (me *Window) SetContentScaleSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) GetContentScaleSize()  {
  panic("TODO: implement")
}

func  (me *Window) SetContentScaleMode(mode WindowContentScaleMode, )  {
  panic("TODO: implement")
}

func  (me *Window) GetContentScaleMode()  {
  panic("TODO: implement")
}

func  (me *Window) SetContentScaleAspect(aspect WindowContentScaleAspect, )  {
  panic("TODO: implement")
}

func  (me *Window) GetContentScaleAspect()  {
  panic("TODO: implement")
}

func  (me *Window) SetContentScaleFactor(factor float32, )  {
  panic("TODO: implement")
}

func  (me *Window) GetContentScaleFactor()  {
  panic("TODO: implement")
}

func  (me *Window) SetUseFontOversampling(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsUsingFontOversampling()  {
  panic("TODO: implement")
}

func  (me *Window) SetMousePassthroughPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *Window) GetMousePassthroughPolygon()  {
  panic("TODO: implement")
}

func  (me *Window) SetWrapControls(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsWrappingControls()  {
  panic("TODO: implement")
}

func  (me *Window) ChildControlsChanged()  {
  panic("TODO: implement")
}

func  (me *Window) SetTheme(theme Theme, )  {
  panic("TODO: implement")
}

func  (me *Window) GetTheme()  {
  panic("TODO: implement")
}

func  (me *Window) SetThemeTypeVariation(theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeTypeVariation()  {
  panic("TODO: implement")
}

func  (me *Window) BeginBulkThemeOverride()  {
  panic("TODO: implement")
}

func  (me *Window) EndBulkThemeOverride()  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeIconOverride(name StringName, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, )  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeFontOverride(name StringName, font Font, )  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeFontSizeOverride(name StringName, font_size int, )  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeColorOverride(name StringName, color Color, )  {
  panic("TODO: implement")
}

func  (me *Window) AddThemeConstantOverride(name StringName, constant int, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeIconOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeStyleboxOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeFontOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeFontSizeOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeColorOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) RemoveThemeConstantOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeIconOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeStyleboxOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeFontOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeFontSizeOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeColorOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeConstantOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeIcon(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeStylebox(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeFont(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeFontSize(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeColor(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) HasThemeConstant(name StringName, theme_type StringName, )  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeDefaultBaseScale()  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeDefaultFont()  {
  panic("TODO: implement")
}

func  (me *Window) GetThemeDefaultFontSize()  {
  panic("TODO: implement")
}

func  (me *Window) SetLayoutDirection(direction WindowLayoutDirection, )  {
  panic("TODO: implement")
}

func  (me *Window) GetLayoutDirection()  {
  panic("TODO: implement")
}

func  (me *Window) IsLayoutRtl()  {
  panic("TODO: implement")
}

func  (me *Window) SetAutoTranslate(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Window) IsAutoTranslating()  {
  panic("TODO: implement")
}

func  (me *Window) Popup(rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupOnParent(parent_rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupCentered(minsize Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupCenteredRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupCenteredClamped(minsize Vector2i, fallback_ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupExclusive(from_node Node, rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupExclusiveOnParent(from_node Node, parent_rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupExclusiveCentered(from_node Node, minsize Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupExclusiveCenteredRatio(from_node Node, ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Window) PopupExclusiveCenteredClamped(from_node Node, minsize Vector2i, fallback_ratio float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
