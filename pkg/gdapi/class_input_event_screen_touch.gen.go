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

type ptrsForInputEventScreenTouchList struct {
	fnSetIndex     gdc.MethodBindPtr
	fnGetIndex     gdc.MethodBindPtr
	fnSetPosition  gdc.MethodBindPtr
	fnGetPosition  gdc.MethodBindPtr
	fnSetPressed   gdc.MethodBindPtr
	fnSetCanceled  gdc.MethodBindPtr
	fnSetDoubleTap gdc.MethodBindPtr
	fnIsDoubleTap  gdc.MethodBindPtr
}

var ptrsForInputEventScreenTouch ptrsForInputEventScreenTouchList

func initInputEventScreenTouchPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventScreenTouch")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_index")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnSetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_index")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnGetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_canceled")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnSetCanceled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_double_tap")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnSetDoubleTap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_double_tap")
		defer methodName.Destroy()
		ptrsForInputEventScreenTouch.fnIsDoubleTap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type InputEventScreenTouch struct {
	InputEventFromWindow
}

func (me *InputEventScreenTouch) BaseClass() string {
	return "InputEventScreenTouch"
}

func NewInputEventScreenTouch() *InputEventScreenTouch {
	str := StringNameFromStr("InputEventScreenTouch") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventScreenTouch{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventScreenTouch) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventScreenTouch) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventScreenTouch) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventScreenTouch) SetIndex(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnSetIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenTouch) GetIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnGetIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventScreenTouch) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenTouch) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventScreenTouch) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenTouch) SetCanceled(canceled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canceled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnSetCanceled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenTouch) SetDoubleTap(double_tap bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&double_tap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnSetDoubleTap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventScreenTouch) IsDoubleTap() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventScreenTouch.fnIsDoubleTap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
