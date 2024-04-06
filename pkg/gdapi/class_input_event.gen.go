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

type ptrsForInputEventList struct {
	fnSetDevice         gdc.MethodBindPtr
	fnGetDevice         gdc.MethodBindPtr
	fnIsAction          gdc.MethodBindPtr
	fnIsActionPressed   gdc.MethodBindPtr
	fnIsActionReleased  gdc.MethodBindPtr
	fnGetActionStrength gdc.MethodBindPtr
	fnIsCanceled        gdc.MethodBindPtr
	fnIsPressed         gdc.MethodBindPtr
	fnIsReleased        gdc.MethodBindPtr
	fnIsEcho            gdc.MethodBindPtr
	fnAsText            gdc.MethodBindPtr
	fnIsMatch           gdc.MethodBindPtr
	fnIsActionType      gdc.MethodBindPtr
	fnAccumulate        gdc.MethodBindPtr
	fnXformedBy         gdc.MethodBindPtr
}

var ptrsForInputEvent ptrsForInputEventList

func initInputEventPtrs(iface gdc.Interface) {

	className := StringNameFromStr("InputEvent")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_device")
		defer methodName.Destroy()
		ptrsForInputEvent.fnSetDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_device")
		defer methodName.Destroy()
		ptrsForInputEvent.fnGetDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_action")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558498928))
	}
	{
		methodName := StringNameFromStr("is_action_pressed")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsActionPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1631499404))
	}
	{
		methodName := StringNameFromStr("is_action_released")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsActionReleased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558498928))
	}
	{
		methodName := StringNameFromStr("get_action_strength")
		defer methodName.Destroy()
		ptrsForInputEvent.fnGetActionStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 801543509))
	}
	{
		methodName := StringNameFromStr("is_canceled")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsCanceled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_pressed")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_released")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsReleased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_echo")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsEcho = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("as_text")
		defer methodName.Destroy()
		ptrsForInputEvent.fnAsText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_match")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsMatch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1754951977))
	}
	{
		methodName := StringNameFromStr("is_action_type")
		defer methodName.Destroy()
		ptrsForInputEvent.fnIsActionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("accumulate")
		defer methodName.Destroy()
		ptrsForInputEvent.fnAccumulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1062211774))
	}
	{
		methodName := StringNameFromStr("xformed_by")
		defer methodName.Destroy()
		ptrsForInputEvent.fnXformedBy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1282766827))
	}
}

type InputEvent struct {
	Resource
}

func (me *InputEvent) BaseClass() string {
	return "InputEvent"
}

func NewInputEvent() *InputEvent {
	str := StringNameFromStr("InputEvent") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &InputEvent{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *InputEvent) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *InputEvent) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *InputEvent) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *InputEvent) SetDevice(device int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnSetDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *InputEvent) GetDevice() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnGetDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsAction(action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsActionPressed(action StringName, allow_echo bool, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&allow_echo), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&allow_echo)
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsActionPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsActionReleased(action StringName, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsActionReleased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) GetActionStrength(action StringName, exact_match bool) float64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnGetActionStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsCanceled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsCanceled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsPressed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsReleased() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsReleased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsEcho() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsEcho), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) AsText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnAsText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *InputEvent) IsMatch(event InputEvent, exact_match bool) bool {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&exact_match)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&exact_match)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsMatch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) IsActionType() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnIsActionType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) Accumulate(with_event InputEvent) bool {
	cargs := []gdc.ConstTypePtr{with_event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnAccumulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *InputEvent) XformedBy(xform Transform2D, local_ofs Vector2) InputEvent {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), local_ofs.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInputEvent()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEvent.fnXformedBy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
