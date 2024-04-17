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

type ptrsForXRController3DList struct {
	fnIsButtonPressed gdc.MethodBindPtr
	fnGetInput        gdc.MethodBindPtr
	fnGetFloat        gdc.MethodBindPtr
	fnGetVector2      gdc.MethodBindPtr
	fnGetTrackerHand  gdc.MethodBindPtr
}

var ptrsForXRController3D ptrsForXRController3DList

func initXRController3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRController3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_button_pressed")
		defer methodName.Destroy()
		ptrsForXRController3D.fnIsButtonPressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_input")
		defer methodName.Destroy()
		ptrsForXRController3D.fnGetInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("get_float")
		defer methodName.Destroy()
		ptrsForXRController3D.fnGetFloat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2349060816))
	}
	{
		methodName := StringNameFromStr("get_vector2")
		defer methodName.Destroy()
		ptrsForXRController3D.fnGetVector2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3100822709))
	}
	{
		methodName := StringNameFromStr("get_tracker_hand")
		defer methodName.Destroy()
		ptrsForXRController3D.fnGetTrackerHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4181770860))
	}

}

type XRController3D struct {
	XRNode3D
}

func (me *XRController3D) BaseClass() string {
	return "XRController3D"
}

func NewXRController3D() *XRController3D {
	str := StringNameFromStr("XRController3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRController3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XRController3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRController3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRController3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRController3D) IsButtonPressed(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRController3D.fnIsButtonPressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRController3D) GetInput(name StringName) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRController3D.fnGetInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRController3D) GetFloat(name StringName) float64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRController3D.fnGetFloat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRController3D) GetVector2(name StringName) Vector2 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRController3D.fnGetVector2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRController3D) GetTrackerHand() XRPositionalTrackerTrackerHand {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRPositionalTrackerTrackerHand

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRController3D.fnGetTrackerHand), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals

type XRController3DButtonPressedSignalFn func(name String)

func (me *XRController3D) ConnectButtonPressed(subs SignalSubscribers, fn XRController3DButtonPressedSignalFn) {
	sig := StringNameFromStr("button_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRController3D) DisconnectButtonPressed(subs SignalSubscribers, fn XRController3DButtonPressedSignalFn) {
	sig := StringNameFromStr("button_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRController3DButtonReleasedSignalFn func(name String)

func (me *XRController3D) ConnectButtonReleased(subs SignalSubscribers, fn XRController3DButtonReleasedSignalFn) {
	sig := StringNameFromStr("button_released")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRController3D) DisconnectButtonReleased(subs SignalSubscribers, fn XRController3DButtonReleasedSignalFn) {
	sig := StringNameFromStr("button_released")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRController3DInputFloatChangedSignalFn func(name String, value float32)

func (me *XRController3D) ConnectInputFloatChanged(subs SignalSubscribers, fn XRController3DInputFloatChangedSignalFn) {
	sig := StringNameFromStr("input_float_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRController3D) DisconnectInputFloatChanged(subs SignalSubscribers, fn XRController3DInputFloatChangedSignalFn) {
	sig := StringNameFromStr("input_float_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRController3DInputVector2ChangedSignalFn func(name String, value Vector2)

func (me *XRController3D) ConnectInputVector2Changed(subs SignalSubscribers, fn XRController3DInputVector2ChangedSignalFn) {
	sig := StringNameFromStr("input_vector2_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRController3D) DisconnectInputVector2Changed(subs SignalSubscribers, fn XRController3DInputVector2ChangedSignalFn) {
	sig := StringNameFromStr("input_vector2_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRController3DProfileChangedSignalFn func(role String)

func (me *XRController3D) ConnectProfileChanged(subs SignalSubscribers, fn XRController3DProfileChangedSignalFn) {
	sig := StringNameFromStr("profile_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRController3D) DisconnectProfileChanged(subs SignalSubscribers, fn XRController3DProfileChangedSignalFn) {
	sig := StringNameFromStr("profile_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
