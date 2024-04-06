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

type ptrsForBaseButtonList struct {
	fnXPressed                   gdc.MethodBindPtr
	fnXToggled                   gdc.MethodBindPtr
	fnSetPressed                 gdc.MethodBindPtr
	fnIsPressed                  gdc.MethodBindPtr
	fnSetPressedNoSignal         gdc.MethodBindPtr
	fnIsHovered                  gdc.MethodBindPtr
	fnSetToggleMode              gdc.MethodBindPtr
	fnIsToggleMode               gdc.MethodBindPtr
	fnSetShortcutInTooltip       gdc.MethodBindPtr
	fnIsShortcutInTooltipEnabled gdc.MethodBindPtr
	fnSetDisabled                gdc.MethodBindPtr
	fnIsDisabled                 gdc.MethodBindPtr
	fnSetActionMode              gdc.MethodBindPtr
	fnGetActionMode              gdc.MethodBindPtr
	fnSetButtonMask              gdc.MethodBindPtr
	fnGetButtonMask              gdc.MethodBindPtr
	fnGetDrawMode                gdc.MethodBindPtr
	fnSetKeepPressedOutside      gdc.MethodBindPtr
	fnIsKeepPressedOutside       gdc.MethodBindPtr
	fnSetShortcutFeedback        gdc.MethodBindPtr
	fnIsShortcutFeedback         gdc.MethodBindPtr
	fnSetShortcut                gdc.MethodBindPtr
	fnGetShortcut                gdc.MethodBindPtr
	fnSetButtonGroup             gdc.MethodBindPtr
	fnGetButtonGroup             gdc.MethodBindPtr
}

var ptrsForBaseButton ptrsForBaseButtonList

func initBaseButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BaseButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_pressed")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_pressed_no_signal")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetPressedNoSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hovered")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsHovered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_toggle_mode")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetToggleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_toggle_mode")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsToggleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shortcut_in_tooltip")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetShortcutInTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shortcut_in_tooltip_enabled")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsShortcutInTooltipEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_disabled")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_disabled")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_action_mode")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetActionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1985162088))
	}
	{
		methodName := StringNameFromStr("get_action_mode")
		defer methodName.Destroy()
		ptrsForBaseButton.fnGetActionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2589712189))
	}
	{
		methodName := StringNameFromStr("set_button_mask")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetButtonMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3950145251))
	}
	{
		methodName := StringNameFromStr("get_button_mask")
		defer methodName.Destroy()
		ptrsForBaseButton.fnGetButtonMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2512161324))
	}
	{
		methodName := StringNameFromStr("get_draw_mode")
		defer methodName.Destroy()
		ptrsForBaseButton.fnGetDrawMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2492721305))
	}
	{
		methodName := StringNameFromStr("set_keep_pressed_outside")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetKeepPressedOutside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_keep_pressed_outside")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsKeepPressedOutside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shortcut_feedback")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetShortcutFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shortcut_feedback")
		defer methodName.Destroy()
		ptrsForBaseButton.fnIsShortcutFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shortcut")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 857163497))
	}
	{
		methodName := StringNameFromStr("get_shortcut")
		defer methodName.Destroy()
		ptrsForBaseButton.fnGetShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3415666916))
	}
	{
		methodName := StringNameFromStr("set_button_group")
		defer methodName.Destroy()
		ptrsForBaseButton.fnSetButtonGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794463739))
	}
	{
		methodName := StringNameFromStr("get_button_group")
		defer methodName.Destroy()
		ptrsForBaseButton.fnGetButtonGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 281644053))
	}
}

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
	BaseButtonDrawModeDrawNormal       BaseButtonDrawMode = 0
	BaseButtonDrawModeDrawPressed      BaseButtonDrawMode = 1
	BaseButtonDrawModeDrawHover        BaseButtonDrawMode = 2
	BaseButtonDrawModeDrawDisabled     BaseButtonDrawMode = 3
	BaseButtonDrawModeDrawHoverPressed BaseButtonDrawMode = 4
)

type BaseButtonActionMode int

const (
	BaseButtonActionModeActionModeButtonPress   BaseButtonActionMode = 0
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

func (me *BaseButton) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetPressedNoSignal(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetPressedNoSignal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsHovered() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsHovered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetToggleMode(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetToggleMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsToggleMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsToggleMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetShortcutInTooltip(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetShortcutInTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsShortcutInTooltipEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsShortcutInTooltipEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetDisabled(disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetActionMode(mode BaseButtonActionMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetActionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) GetActionMode() BaseButtonActionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseButtonActionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnGetActionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseButton) SetButtonMask(mask MouseButtonMask) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetButtonMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) GetButtonMask() MouseButtonMask {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MouseButtonMask

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnGetButtonMask), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseButton) GetDrawMode() BaseButtonDrawMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseButtonDrawMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnGetDrawMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseButton) SetKeepPressedOutside(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetKeepPressedOutside), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsKeepPressedOutside() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsKeepPressedOutside), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetShortcutFeedback(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetShortcutFeedback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) IsShortcutFeedback() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnIsShortcutFeedback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseButton) SetShortcut(shortcut Shortcut) {
	cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) GetShortcut() Shortcut {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShortcut()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnGetShortcut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseButton) SetButtonGroup(button_group ButtonGroup) {
	cargs := []gdc.ConstTypePtr{button_group.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnSetButtonGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseButton) GetButtonGroup() ButtonGroup {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewButtonGroup()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseButton.fnGetButtonGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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

type BaseButtonToggledSignalFn func(toggled_on bool)

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
