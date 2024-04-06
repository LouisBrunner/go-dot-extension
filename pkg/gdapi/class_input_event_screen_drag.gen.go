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

type ptrsForInputEventScreenDragList struct {
	fnSetIndex       gdc.MethodBindPtr
	fnGetIndex       gdc.MethodBindPtr
	fnSetTilt        gdc.MethodBindPtr
	fnGetTilt        gdc.MethodBindPtr
	fnSetPressure    gdc.MethodBindPtr
	fnGetPressure    gdc.MethodBindPtr
	fnSetPenInverted gdc.MethodBindPtr
	fnGetPenInverted gdc.MethodBindPtr
	fnSetPosition    gdc.MethodBindPtr
	fnGetPosition    gdc.MethodBindPtr
	fnSetRelative    gdc.MethodBindPtr
	fnGetRelative    gdc.MethodBindPtr
	fnSetVelocity    gdc.MethodBindPtr
	fnGetVelocity    gdc.MethodBindPtr
}

var ptrsForInputEventScreenDrag ptrsForInputEventScreenDragList

func initInputEventScreenDragPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventScreenDrag")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_index")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_index")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_tilt")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_tilt")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_pressure")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pressure")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pen_inverted")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetPenInverted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_pen_inverted")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetPenInverted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_relative")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_relative")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForInputEventScreenDrag.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type InputEventScreenDrag struct {
	InputEventFromWindow
}

func (me *InputEventScreenDrag) BaseClass() string {
	return "InputEventScreenDrag"
}

func NewInputEventScreenDrag() *InputEventScreenDrag {
	str := StringNameFromStr("InputEventScreenDrag") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventScreenDrag{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventScreenDrag) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventScreenDrag) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventScreenDrag) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventScreenDrag) SetIndex(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventScreenDrag) SetTilt(tilt Vector2) {
	cargs := []gdc.ConstTypePtr{tilt.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetTilt), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetTilt() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetTilt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventScreenDrag) SetPressure(pressure float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetPressure), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetPressure() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetPressure), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventScreenDrag) SetPenInverted(pen_inverted bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pen_inverted)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetPenInverted), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetPenInverted() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetPenInverted), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventScreenDrag) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventScreenDrag) SetRelative(relative Vector2) {
	cargs := []gdc.ConstTypePtr{relative.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetRelative), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetRelative() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetRelative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventScreenDrag) SetVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenDrag) GetVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenDrag.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
