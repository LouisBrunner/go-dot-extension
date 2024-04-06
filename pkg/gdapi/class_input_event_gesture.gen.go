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

type ptrsForInputEventGestureList struct {
	fnSetPosition gdc.MethodBindPtr
	fnGetPosition gdc.MethodBindPtr
}

var ptrsForInputEventGesture ptrsForInputEventGestureList

func initInputEventGesturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventGesture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForInputEventGesture.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForInputEventGesture.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type InputEventGesture struct {
	InputEventWithModifiers
}

func (me *InputEventGesture) BaseClass() string {
	return "InputEventGesture"
}

func NewInputEventGesture() *InputEventGesture {
	str := StringNameFromStr("InputEventGesture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventGesture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventGesture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventGesture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventGesture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventGesture) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventGesture.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventGesture) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventGesture.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
