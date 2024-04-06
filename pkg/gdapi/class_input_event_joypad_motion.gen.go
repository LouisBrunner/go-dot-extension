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

type ptrsForInputEventJoypadMotionList struct {
	fnSetAxis      gdc.MethodBindPtr
	fnGetAxis      gdc.MethodBindPtr
	fnSetAxisValue gdc.MethodBindPtr
	fnGetAxisValue gdc.MethodBindPtr
}

var ptrsForInputEventJoypadMotion ptrsForInputEventJoypadMotionList

func initInputEventJoypadMotionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventJoypadMotion")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_axis")
		defer methodName.Destroy()
		ptrsForInputEventJoypadMotion.fnSetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1332685170))
	}
	{
		methodName := StringNameFromStr("get_axis")
		defer methodName.Destroy()
		ptrsForInputEventJoypadMotion.fnGetAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4019121683))
	}
	{
		methodName := StringNameFromStr("set_axis_value")
		defer methodName.Destroy()
		ptrsForInputEventJoypadMotion.fnSetAxisValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_axis_value")
		defer methodName.Destroy()
		ptrsForInputEventJoypadMotion.fnGetAxisValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type InputEventJoypadMotion struct {
	InputEvent
}

func (me *InputEventJoypadMotion) BaseClass() string {
	return "InputEventJoypadMotion"
}

func NewInputEventJoypadMotion() *InputEventJoypadMotion {
	str := StringNameFromStr("InputEventJoypadMotion") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventJoypadMotion{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventJoypadMotion) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventJoypadMotion) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventJoypadMotion) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventJoypadMotion) SetAxis(axis JoyAxis) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadMotion.fnSetAxis), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventJoypadMotion) GetAxis() JoyAxis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret JoyAxis

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadMotion.fnGetAxis), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventJoypadMotion) SetAxisValue(axis_value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis_value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadMotion.fnSetAxisValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventJoypadMotion) GetAxisValue() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventJoypadMotion.fnGetAxisValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
