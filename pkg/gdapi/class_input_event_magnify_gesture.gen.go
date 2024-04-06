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

type ptrsForInputEventMagnifyGestureList struct {
	fnSetFactor gdc.MethodBindPtr
	fnGetFactor gdc.MethodBindPtr
}

var ptrsForInputEventMagnifyGesture ptrsForInputEventMagnifyGestureList

func initInputEventMagnifyGesturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventMagnifyGesture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_factor")
		defer methodName.Destroy()
		ptrsForInputEventMagnifyGesture.fnSetFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_factor")
		defer methodName.Destroy()
		ptrsForInputEventMagnifyGesture.fnGetFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type InputEventMagnifyGesture struct {
	InputEventGesture
}

func (me *InputEventMagnifyGesture) BaseClass() string {
	return "InputEventMagnifyGesture"
}

func NewInputEventMagnifyGesture() *InputEventMagnifyGesture {
	str := StringNameFromStr("InputEventMagnifyGesture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventMagnifyGesture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventMagnifyGesture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventMagnifyGesture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventMagnifyGesture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventMagnifyGesture) SetFactor(factor float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMagnifyGesture.fnSetFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMagnifyGesture) GetFactor() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMagnifyGesture.fnGetFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
