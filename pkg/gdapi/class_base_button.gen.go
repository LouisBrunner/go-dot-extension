// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BaseButton struct {
  obj gdc.ObjectPtr
}

func (me *BaseButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BaseButton) BaseClass() string {
  return "BaseButton"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsPressed() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetPressedNoSignal(pressed bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed_no_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsHovered() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hovered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetToggleMode(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsToggleMode() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetShortcutInTooltip(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_in_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsShortcutInTooltipEnabled() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_in_tooltip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetDisabled(disabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsDisabled() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetActionMode(mode BaseButtonActionMode, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1985162088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) GetActionMode() BaseButtonActionMode {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2589712189) // FIXME: should cache?
  var ret BaseButtonActionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetButtonMask(mask MouseButtonMask, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3950145251) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) GetButtonMask() MouseButtonMask {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2512161324) // FIXME: should cache?
  var ret MouseButtonMask
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) GetDrawMode() BaseButtonDrawMode {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2492721305) // FIXME: should cache?
  var ret BaseButtonDrawMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetKeepPressedOutside(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_pressed_outside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsKeepPressedOutside() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keep_pressed_outside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetShortcutFeedback(enabled bool, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) IsShortcutFeedback() bool {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetShortcut(shortcut Shortcut, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 857163497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shortcut.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) GetShortcut() Shortcut {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415666916) // FIXME: should cache?
  var ret Shortcut
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BaseButton) SetButtonGroup(button_group ButtonGroup, )  {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794463739) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(button_group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BaseButton) GetButtonGroup() ButtonGroup {
  classNameV := StringNameFromStr("BaseButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281644053) // FIXME: should cache?
  var ret ButtonGroup
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *BaseButton) GetPropDisabled() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropDisabled(value bool) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropToggleMode() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropToggleMode(value bool) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropButtonPressed() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropButtonPressed(value bool) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropActionMode() int {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropActionMode(value int) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropButtonMask() int {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropButtonMask(value int) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropKeepPressedOutside() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropKeepPressedOutside(value bool) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropButtonGroup() ButtonGroup {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropButtonGroup(value ButtonGroup) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropShortcut() Shortcut {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropShortcut(value Shortcut) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropShortcutFeedback() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropShortcutFeedback(value bool) {
  panic("TODO: implement")
}

func (me *BaseButton) GetPropShortcutInTooltip() bool {
  panic("TODO: implement")
}

func (me *BaseButton) SetPropShortcutInTooltip(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API