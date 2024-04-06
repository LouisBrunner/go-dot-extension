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

type ptrsForInputEventMouseList struct {
	fnSetButtonMask     gdc.MethodBindPtr
	fnGetButtonMask     gdc.MethodBindPtr
	fnSetPosition       gdc.MethodBindPtr
	fnGetPosition       gdc.MethodBindPtr
	fnSetGlobalPosition gdc.MethodBindPtr
	fnGetGlobalPosition gdc.MethodBindPtr
}

var ptrsForInputEventMouse ptrsForInputEventMouseList

func initInputEventMousePtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventMouse")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_button_mask")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnSetButtonMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3950145251))
	}
	{
		methodName := StringNameFromStr("get_button_mask")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnGetButtonMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2512161324))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_global_position")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnSetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_global_position")
		defer methodName.Destroy()
		ptrsForInputEventMouse.fnGetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type InputEventMouse struct {
	InputEventWithModifiers
}

func (me *InputEventMouse) BaseClass() string {
	return "InputEventMouse"
}

func NewInputEventMouse() *InputEventMouse {
	str := StringNameFromStr("InputEventMouse") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventMouse{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventMouse) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventMouse) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventMouse) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventMouse) SetButtonMask(button_mask MouseButtonMask) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&button_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnSetButtonMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouse) GetButtonMask() MouseButtonMask {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MouseButtonMask

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnGetButtonMask), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *InputEventMouse) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouse) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventMouse) SetGlobalPosition(global_position Vector2) {
	cargs := []gdc.ConstTypePtr{global_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnSetGlobalPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventMouse) GetGlobalPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMouse.fnGetGlobalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
