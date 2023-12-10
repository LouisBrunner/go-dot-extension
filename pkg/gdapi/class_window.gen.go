// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *Window) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Window) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Window) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Window) SetTitle(title String, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetTitle() String {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetWindowId() int {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetInitialPosition(initial_position WindowWindowInitialPosition, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_initial_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4084468099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&initial_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetInitialPosition() WindowWindowInitialPosition {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_initial_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4294066647) // FIXME: should cache?
  var ret WindowWindowInitialPosition
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetCurrentScreen(index int, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetCurrentScreen() int {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetPosition(position Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetPosition() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetSize(size Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetSize() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) ResetSize()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetPositionWithDecorations() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_with_decorations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetSizeWithDecorations() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_with_decorations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetMaxSize(max_size Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(max_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetMaxSize() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetMinSize(min_size Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(min_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetMinSize() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetMode(mode WindowMode, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3095236531) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetMode() WindowMode {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2566346114) // FIXME: should cache?
  var ret WindowMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetFlag(flag WindowFlags, enabled bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426449779) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetFlag(flag WindowFlags, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3062752289) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) IsMaximizeAllowed() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_maximize_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) RequestAttention()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_attention")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) MoveToForeground()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_to_foreground")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) SetVisible(visible bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsVisible() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) Hide()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) Show()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) SetTransient(transient bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transient), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsTransient() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_transient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetExclusive(exclusive bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclusive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclusive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsExclusive() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_exclusive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetUnparentWhenInvisible(unparent bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unparent_when_invisible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unparent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) CanDraw() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasFocus() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GrabFocus()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("grab_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) SetImeActive(active bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ime_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) SetImePosition(position Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ime_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsEmbedded() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_embedded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetContentsMinimumSize() Vector2 {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contents_minimum_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetContentScaleSize(size Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_content_scale_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetContentScaleSize() Vector2i {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_scale_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetContentScaleMode(mode WindowContentScaleMode, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_content_scale_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2937716473) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetContentScaleMode() WindowContentScaleMode {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_scale_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 161585230) // FIXME: should cache?
  var ret WindowContentScaleMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetContentScaleAspect(aspect WindowContentScaleAspect, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_content_scale_aspect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2370399418) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aspect), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetContentScaleAspect() WindowContentScaleAspect {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_scale_aspect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4158790715) // FIXME: should cache?
  var ret WindowContentScaleAspect
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetContentScaleFactor(factor float32, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_content_scale_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetContentScaleFactor() float32 {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_scale_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetUseFontOversampling(enable bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_font_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsUsingFontOversampling() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_font_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetMousePassthroughPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mouse_passthrough_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetMousePassthroughPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mouse_passthrough_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetWrapControls(enable bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wrap_controls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsWrappingControls() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_wrapping_controls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) ChildControlsChanged()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("child_controls_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) SetTheme(theme Theme, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2326690814) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(theme.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetTheme() Theme {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3846893731) // FIXME: should cache?
  var ret Theme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetThemeTypeVariation(theme_type StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_theme_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetThemeTypeVariation() StringName {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_type_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) BeginBulkThemeOverride()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("begin_bulk_theme_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) EndBulkThemeOverride()  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("end_bulk_theme_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeIconOverride(name StringName, texture Texture2D, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1373065600) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeStyleboxOverride(name StringName, stylebox StyleBox, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4188838905) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(stylebox.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeFontOverride(name StringName, font Font, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3518018674) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeFontSizeOverride(name StringName, font_size int, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeColorOverride(name StringName, color Color, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4260178595) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) AddThemeConstantOverride(name StringName, constant int, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&constant), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeIconOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeStyleboxOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeFontOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeFontSizeOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeColorOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) RemoveThemeConstantOverride(name StringName, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetThemeIcon(name StringName, theme_type StringName, ) Texture2D {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2336455395) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeStylebox(name StringName, theme_type StringName, ) StyleBox {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2759935355) // FIXME: should cache?
  var ret StyleBox
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeFont(name StringName, theme_type StringName, ) Font {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 387378635) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeFontSize(name StringName, theme_type StringName, ) int {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229578101) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeColor(name StringName, theme_type StringName, ) Color {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2377051548) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeConstant(name StringName, theme_type StringName, ) int {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229578101) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeIconOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_icon_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeStyleboxOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_stylebox_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeFontOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeFontSizeOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeColorOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_color_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeConstantOverride(name StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_constant_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeIcon(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeStylebox(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeFont(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeFontSize(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeColor(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) HasThemeConstant(name StringName, theme_type StringName, ) bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_theme_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187511791) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(theme_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeDefaultBaseScale() float32 {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeDefaultFont() Font {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) GetThemeDefaultFontSize() int {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_theme_default_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetLayoutDirection(direction WindowLayoutDirection, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layout_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3094704184) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) GetLayoutDirection() WindowLayoutDirection {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layout_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3909617982) // FIXME: should cache?
  var ret WindowLayoutDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) IsLayoutRtl() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layout_rtl")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) SetAutoTranslate(enable bool, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) IsAutoTranslating() bool {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_translating")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Window) Popup(rect Rect2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1680304321) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupOnParent(parent_rect Rect2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_on_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1763793166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parent_rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupCentered(minsize Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3447975422) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(minsize.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupCenteredRatio(ratio float32, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_centered_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1014814997) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupCenteredClamped(minsize Vector2i, fallback_ratio float32, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_centered_clamped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2613752477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(minsize.AsCTypePtr()), gdc.ConstTypePtr(&fallback_ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupExclusive(from_node Node, rect Rect2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_exclusive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1728044812) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupExclusiveOnParent(from_node Node, parent_rect Rect2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_exclusive_on_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2344671043) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(parent_rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupExclusiveCentered(from_node Node, minsize Vector2i, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_exclusive_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2561668109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(minsize.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupExclusiveCenteredRatio(from_node Node, ratio float32, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_exclusive_centered_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4257659513) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Window) PopupExclusiveCenteredClamped(from_node Node, minsize Vector2i, fallback_ratio float32, )  {
  classNameV := StringNameFromStr("Window")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_exclusive_centered_clamped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 224798062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(minsize.AsCTypePtr()), gdc.ConstTypePtr(&fallback_ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *Window) GetPropMode() int {
  panic("TODO: implement")
}

func (me *Window) SetPropMode(value int) {
  panic("TODO: implement")
}

func (me *Window) GetPropTitle() String {
  panic("TODO: implement")
}

func (me *Window) SetPropTitle(value String) {
  panic("TODO: implement")
}

func (me *Window) GetPropInitialPosition() int {
  panic("TODO: implement")
}

func (me *Window) SetPropInitialPosition(value int) {
  panic("TODO: implement")
}

func (me *Window) GetPropPosition() Vector2i {
  panic("TODO: implement")
}

func (me *Window) SetPropPosition(value Vector2i) {
  panic("TODO: implement")
}

func (me *Window) GetPropSize() Vector2i {
  panic("TODO: implement")
}

func (me *Window) SetPropSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *Window) GetPropCurrentScreen() int {
  panic("TODO: implement")
}

func (me *Window) SetPropCurrentScreen(value int) {
  panic("TODO: implement")
}

func (me *Window) GetPropMousePassthroughPolygon() PackedVector2Array {
  panic("TODO: implement")
}

func (me *Window) SetPropMousePassthroughPolygon(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *Window) GetPropVisible() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropVisible(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropWrapControls() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropWrapControls(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropTransient() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropTransient(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropExclusive() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropExclusive(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropUnresizable() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropUnresizable(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropBorderless() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropBorderless(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropAlwaysOnTop() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropAlwaysOnTop(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropTransparent() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropTransparent(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropUnfocusable() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropUnfocusable(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropPopupWindow() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropPopupWindow(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropExtendToTitle() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropExtendToTitle(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropMousePassthrough() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropMousePassthrough(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropMinSize() Vector2i {
  panic("TODO: implement")
}

func (me *Window) SetPropMinSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *Window) GetPropMaxSize() Vector2i {
  panic("TODO: implement")
}

func (me *Window) SetPropMaxSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *Window) GetPropContentScaleSize() Vector2i {
  panic("TODO: implement")
}

func (me *Window) SetPropContentScaleSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *Window) GetPropContentScaleMode() int {
  panic("TODO: implement")
}

func (me *Window) SetPropContentScaleMode(value int) {
  panic("TODO: implement")
}

func (me *Window) GetPropContentScaleAspect() int {
  panic("TODO: implement")
}

func (me *Window) SetPropContentScaleAspect(value int) {
  panic("TODO: implement")
}

func (me *Window) GetPropContentScaleFactor() float32 {
  panic("TODO: implement")
}

func (me *Window) SetPropContentScaleFactor(value float32) {
  panic("TODO: implement")
}

func (me *Window) GetPropAutoTranslate() bool {
  panic("TODO: implement")
}

func (me *Window) SetPropAutoTranslate(value bool) {
  panic("TODO: implement")
}

func (me *Window) GetPropTheme() Theme {
  panic("TODO: implement")
}

func (me *Window) SetPropTheme(value Theme) {
  panic("TODO: implement")
}

func (me *Window) GetPropThemeTypeVariation() String {
  panic("TODO: implement")
}

func (me *Window) SetPropThemeTypeVariation(value String) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API