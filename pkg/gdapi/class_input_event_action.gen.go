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

type ptrsForInputEventActionList struct {
	fnSetAction   gdc.MethodBindPtr
	fnGetAction   gdc.MethodBindPtr
	fnSetPressed  gdc.MethodBindPtr
	fnSetStrength gdc.MethodBindPtr
	fnGetStrength gdc.MethodBindPtr
}

var ptrsForInputEventAction ptrsForInputEventActionList

func initInputEventActionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventAction")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_action")
		defer methodName.Destroy()
		ptrsForInputEventAction.fnSetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_action")
		defer methodName.Destroy()
		ptrsForInputEventAction.fnGetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_pressed")
		defer methodName.Destroy()
		ptrsForInputEventAction.fnSetPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_strength")
		defer methodName.Destroy()
		ptrsForInputEventAction.fnSetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_strength")
		defer methodName.Destroy()
		ptrsForInputEventAction.fnGetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type InputEventAction struct {
	InputEvent
}

func (me *InputEventAction) BaseClass() string {
	return "InputEventAction"
}

func NewInputEventAction() *InputEventAction {
	str := StringNameFromStr("InputEventAction") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventAction{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventAction) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventAction) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventAction) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventAction) SetAction(action StringName) {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventAction.fnSetAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventAction) GetAction() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventAction.fnGetAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEventAction) SetPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventAction.fnSetPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventAction) SetStrength(strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventAction.fnSetStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventAction) GetStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventAction.fnGetStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
