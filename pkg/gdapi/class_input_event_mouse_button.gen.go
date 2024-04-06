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

type ptrsForInputEventMouseButtonList struct {
	fnSetFactor      gdc.MethodBindPtr
	fnGetFactor      gdc.MethodBindPtr
	fnSetButtonIndex gdc.MethodBindPtr
	fnGetButtonIndex gdc.MethodBindPtr
	fnSetPressed     gdc.MethodBindPtr
	fnSetCanceled    gdc.MethodBindPtr
	fnSetDoubleClick gdc.MethodBindPtr
	fnIsDoubleClick  gdc.MethodBindPtr
}

var ptrsForInputEventMouseButton ptrsForInputEventMouseButtonList

func initInputEventMouseButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventMouseButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_factor")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnSetFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_factor")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnGetFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_button_index")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnSetButtonIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3624991109))
	}
	{
		methodName := StringNameFromStr("get_button_index")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnGetButtonIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1132662608))
	}
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_canceled")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnSetCanceled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_double_click")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnSetDoubleClick = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_double_click")
		defer methodName.Destroy()
		ptrsForInputEventMouseButton.fnIsDoubleClick = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type InputEventMouseButton struct {
	InputEventMouse
}

func (me *InputEventMouseButton) BaseClass() string {
	return "InputEventMouseButton"
}

func NewInputEventMouseButton() *InputEventMouseButton {
	str := StringNameFromStr("InputEventMouseButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventMouseButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventMouseButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventMouseButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventMouseButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventMouseButton) SetFactor(factor float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnSetFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseButton) GetFactor() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnGetFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventMouseButton) SetButtonIndex(button_index MouseButton) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnSetButtonIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseButton) GetButtonIndex() MouseButton {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MouseButton

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnGetButtonIndex), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventMouseButton) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseButton) SetCanceled(canceled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canceled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnSetCanceled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseButton) SetDoubleClick(double_click bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&double_click)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnSetDoubleClick), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseButton) IsDoubleClick() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseButton.fnIsDoubleClick), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
