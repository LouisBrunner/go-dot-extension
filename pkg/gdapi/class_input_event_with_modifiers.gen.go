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

type ptrsForInputEventWithModifiersList struct {
	fnSetCommandOrControlAutoremap gdc.MethodBindPtr
	fnIsCommandOrControlAutoremap  gdc.MethodBindPtr
	fnIsCommandOrControlPressed    gdc.MethodBindPtr
	fnSetAltPressed                gdc.MethodBindPtr
	fnIsAltPressed                 gdc.MethodBindPtr
	fnSetShiftPressed              gdc.MethodBindPtr
	fnIsShiftPressed               gdc.MethodBindPtr
	fnSetCtrlPressed               gdc.MethodBindPtr
	fnIsCtrlPressed                gdc.MethodBindPtr
	fnSetMetaPressed               gdc.MethodBindPtr
	fnIsMetaPressed                gdc.MethodBindPtr
	fnGetModifiersMask             gdc.MethodBindPtr
}

var ptrsForInputEventWithModifiers ptrsForInputEventWithModifiersList

func initInputEventWithModifiersPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEventWithModifiers")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_command_or_control_autoremap")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnSetCommandOrControlAutoremap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_command_or_control_autoremap")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsCommandOrControlAutoremap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_command_or_control_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsCommandOrControlPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_alt_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnSetAltPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_alt_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsAltPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shift_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnSetShiftPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shift_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsShiftPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_ctrl_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnSetCtrlPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ctrl_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsCtrlPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_meta_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnSetMetaPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_meta_pressed")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnIsMetaPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_modifiers_mask")
		defer methodName.Destroy()
		ptrsForInputEventWithModifiers.fnGetModifiersMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1258259499))
	}

}

type InputEventWithModifiers struct {
	InputEventFromWindow
}

func (me *InputEventWithModifiers) BaseClass() string {
	return "InputEventWithModifiers"
}

func NewInputEventWithModifiers() *InputEventWithModifiers {
	str := StringNameFromStr("InputEventWithModifiers") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEventWithModifiers{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEventWithModifiers) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEventWithModifiers) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEventWithModifiers) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEventWithModifiers) SetCommandOrControlAutoremap(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnSetCommandOrControlAutoremap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventWithModifiers) IsCommandOrControlAutoremap() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsCommandOrControlAutoremap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) IsCommandOrControlPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsCommandOrControlPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) SetAltPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnSetAltPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventWithModifiers) IsAltPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsAltPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) SetShiftPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnSetShiftPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventWithModifiers) IsShiftPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsShiftPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) SetCtrlPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnSetCtrlPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventWithModifiers) IsCtrlPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsCtrlPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) SetMetaPressed(pressed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnSetMetaPressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEventWithModifiers) IsMetaPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnIsMetaPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEventWithModifiers) GetModifiersMask() KeyModifierMask {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret KeyModifierMask

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventWithModifiers.fnGetModifiersMask), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
