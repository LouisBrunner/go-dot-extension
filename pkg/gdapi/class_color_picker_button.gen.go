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

type ptrsForColorPickerButtonList struct {
	fnSetPickColor   gdc.MethodBindPtr
	fnGetPickColor   gdc.MethodBindPtr
	fnGetPicker      gdc.MethodBindPtr
	fnGetPopup       gdc.MethodBindPtr
	fnSetEditAlpha   gdc.MethodBindPtr
	fnIsEditingAlpha gdc.MethodBindPtr
}

var ptrsForColorPickerButton ptrsForColorPickerButtonList

func initColorPickerButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ColorPickerButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pick_color")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnSetPickColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_pick_color")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnGetPickColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("get_picker")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnGetPicker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 331835996))
	}
	{
		methodName := StringNameFromStr("get_popup")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnGetPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1322440207))
	}
	{
		methodName := StringNameFromStr("set_edit_alpha")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnSetEditAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_editing_alpha")
		defer methodName.Destroy()
		ptrsForColorPickerButton.fnIsEditingAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type ColorPickerButton struct {
	Button
}

func (me *ColorPickerButton) BaseClass() string {
	return "ColorPickerButton"
}

func NewColorPickerButton() *ColorPickerButton {
	str := StringNameFromStr("ColorPickerButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ColorPickerButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ColorPickerButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ColorPickerButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ColorPickerButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ColorPickerButton) SetPickColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnSetPickColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPickerButton) GetPickColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnGetPickColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPickerButton) GetPicker() ColorPicker {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColorPicker()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnGetPicker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPickerButton) GetPopup() PopupPanel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopupPanel()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnGetPopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPickerButton) SetEditAlpha(show bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnSetEditAlpha), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPickerButton) IsEditingAlpha() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPickerButton.fnIsEditingAlpha), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ColorPickerButtonColorChangedSignalFn func(color Color)

func (me *ColorPickerButton) ConnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
	sig := StringNameFromStr("color_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
	sig := StringNameFromStr("color_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerButtonPopupClosedSignalFn func()

func (me *ColorPickerButton) ConnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
	sig := StringNameFromStr("popup_closed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
	sig := StringNameFromStr("popup_closed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerButtonPickerCreatedSignalFn func()

func (me *ColorPickerButton) ConnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
	sig := StringNameFromStr("picker_created")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
	sig := StringNameFromStr("picker_created")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
