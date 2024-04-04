// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type BaseButton struct {
  Control
}

func (me *BaseButton) BaseClass() string {
  return "BaseButton"
}

func NewBaseButton() *BaseButton {
  str := StringNameFromStr("BaseButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BaseButton{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type BaseButtonDrawMode int
const (
  BaseButtonDrawModeDrawNormal BaseButtonDrawMode = 0
  BaseButtonDrawModeDrawPressed BaseButtonDrawMode = 1
  BaseButtonDrawModeDrawHover BaseButtonDrawMode = 2
  BaseButtonDrawModeDrawDisabled BaseButtonDrawMode = 3
  BaseButtonDrawModeDrawHoverPressed BaseButtonDrawMode = 4
)

type BaseButtonActionMode int
const (
  BaseButtonActionModeActionModeButtonPress BaseButtonActionMode = 0
  BaseButtonActionModeActionModeButtonRelease BaseButtonActionMode = 1
)

func (me *BaseButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BaseButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BaseButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BaseButton) SetPressed(pressed bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsPressed() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetPressedNoSignal(pressed bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed_no_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsHovered() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hovered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetToggleMode(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsToggleMode() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetShortcutInTooltip(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_in_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsShortcutInTooltipEnabled() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_in_tooltip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetDisabled(disabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsDisabled() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetActionMode(mode BaseButtonActionMode, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1985162088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) GetActionMode() BaseButtonActionMode {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2589712189) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseButtonActionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseButton) SetButtonMask(mask MouseButtonMask, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3950145251) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) GetButtonMask() MouseButtonMask {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2512161324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MouseButtonMask

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseButton) GetDrawMode() BaseButtonDrawMode {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2492721305) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseButtonDrawMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseButton) SetKeepPressedOutside(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_pressed_outside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsKeepPressedOutside() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keep_pressed_outside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetShortcutFeedback(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) IsShortcutFeedback() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseButton) SetShortcut(shortcut Shortcut, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 857163497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) GetShortcut() Shortcut {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415666916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShortcut()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseButton) SetButtonGroup(button_group ButtonGroup, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794463739) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{button_group.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseButton) GetButtonGroup() ButtonGroup {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281644053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewButtonGroup()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type BaseButtonPressedSignalFn func()

func (me *BaseButton) ConnectPressed(subs SignalSubscribers, fn BaseButtonPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *BaseButton) DisconnectPressed(subs SignalSubscribers, fn BaseButtonPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type BaseButtonButtonUpSignalFn func()

func (me *BaseButton) ConnectButtonUp(subs SignalSubscribers, fn BaseButtonButtonUpSignalFn) {
  sig := StringNameFromStr("button_up")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *BaseButton) DisconnectButtonUp(subs SignalSubscribers, fn BaseButtonButtonUpSignalFn) {
  sig := StringNameFromStr("button_up")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type BaseButtonButtonDownSignalFn func()

func (me *BaseButton) ConnectButtonDown(subs SignalSubscribers, fn BaseButtonButtonDownSignalFn) {
  sig := StringNameFromStr("button_down")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *BaseButton) DisconnectButtonDown(subs SignalSubscribers, fn BaseButtonButtonDownSignalFn) {
  sig := StringNameFromStr("button_down")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type BaseButtonToggledSignalFn func(toggled_on bool, )

func (me *BaseButton) ConnectToggled(subs SignalSubscribers, fn BaseButtonToggledSignalFn) {
  sig := StringNameFromStr("toggled")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *BaseButton) DisconnectToggled(subs SignalSubscribers, fn BaseButtonToggledSignalFn) {
  sig := StringNameFromStr("toggled")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
