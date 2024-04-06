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

type ptrsForInputEventJoypadButtonList struct {
	fnSetButtonIndex gdc.MethodBindPtr
	fnGetButtonIndex gdc.MethodBindPtr
	fnSetPressure    gdc.MethodBindPtr
	fnGetPressure    gdc.MethodBindPtr
	fnSetPressed     gdc.MethodBindPtr
}

var ptrsForInputEventJoypadButton ptrsForInputEventJoypadButtonList

func initInputEventJoypadButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventJoypadButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_button_index")
		defer methodName.Destroy()
		ptrsForInputEventJoypadButton.fnSetButtonIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1466368136))
	}
	{
		methodName := StringNameFromStr("get_button_index")
		defer methodName.Destroy()
		ptrsForInputEventJoypadButton.fnGetButtonIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 595588182))
	}
	{
		methodName := StringNameFromStr("set_pressure")
		defer methodName.Destroy()
		ptrsForInputEventJoypadButton.fnSetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pressure")
		defer methodName.Destroy()
		ptrsForInputEventJoypadButton.fnGetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForInputEventJoypadButton.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}

}

type InputEventJoypadButton struct {
	InputEvent
}

func (me *InputEventJoypadButton) BaseClass() string {
	return "InputEventJoypadButton"
}

func NewInputEventJoypadButton() *InputEventJoypadButton {
	str := StringNameFromStr("InputEventJoypadButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventJoypadButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventJoypadButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventJoypadButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventJoypadButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventJoypadButton) SetButtonIndex(button_index JoyButton) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadButton.fnSetButtonIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventJoypadButton) GetButtonIndex() JoyButton {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret JoyButton

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadButton.fnGetButtonIndex), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventJoypadButton) SetPressure(pressure float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadButton.fnSetPressure), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventJoypadButton) GetPressure() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadButton.fnGetPressure), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventJoypadButton) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadButton.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
