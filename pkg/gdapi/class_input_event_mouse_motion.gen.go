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

type ptrsForInputEventMouseMotionList struct {
	fnSetTilt           gdc.MethodBindPtr
	fnGetTilt           gdc.MethodBindPtr
	fnSetPressure       gdc.MethodBindPtr
	fnGetPressure       gdc.MethodBindPtr
	fnSetPenInverted    gdc.MethodBindPtr
	fnGetPenInverted    gdc.MethodBindPtr
	fnSetRelative       gdc.MethodBindPtr
	fnGetRelative       gdc.MethodBindPtr
	fnSetScreenRelative gdc.MethodBindPtr
	fnGetScreenRelative gdc.MethodBindPtr
	fnSetVelocity       gdc.MethodBindPtr
	fnGetVelocity       gdc.MethodBindPtr
	fnSetScreenVelocity gdc.MethodBindPtr
	fnGetScreenVelocity gdc.MethodBindPtr
}

var ptrsForInputEventMouseMotion ptrsForInputEventMouseMotionList

func initInputEventMouseMotionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventMouseMotion")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_tilt")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_tilt")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_pressure")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pressure")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pen_inverted")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetPenInverted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_pen_inverted")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetPenInverted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_relative")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_relative")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_screen_relative")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetScreenRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_screen_relative")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetScreenRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_screen_velocity")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnSetScreenVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_screen_velocity")
		defer methodName.Destroy()
		ptrsForInputEventMouseMotion.fnGetScreenVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type InputEventMouseMotion struct {
	InputEventMouse
}

func (me *InputEventMouseMotion) BaseClass() string {
	return "InputEventMouseMotion"
}

func NewInputEventMouseMotion() *InputEventMouseMotion {
	str := StringNameFromStr("InputEventMouseMotion") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventMouseMotion{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventMouseMotion) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventMouseMotion) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventMouseMotion) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventMouseMotion) SetTilt(tilt Vector2) {
	cargs := []gdc.ConstTypePtr{tilt.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetTilt), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetTilt() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetTilt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventMouseMotion) SetPressure(pressure float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetPressure), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetPressure() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetPressure), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventMouseMotion) SetPenInverted(pen_inverted bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pen_inverted)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetPenInverted), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetPenInverted() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetPenInverted), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventMouseMotion) SetRelative(relative Vector2) {
	cargs := []gdc.ConstTypePtr{relative.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetRelative), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetRelative() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetRelative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventMouseMotion) SetScreenRelative(relative Vector2) {
	cargs := []gdc.ConstTypePtr{relative.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetScreenRelative), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetScreenRelative() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetScreenRelative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventMouseMotion) SetVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventMouseMotion) SetScreenVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnSetScreenVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouseMotion) GetScreenVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouseMotion.fnGetScreenVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
