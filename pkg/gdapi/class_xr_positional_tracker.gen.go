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

type ptrsForXRPositionalTrackerList struct {
	fnGetTrackerType    gdc.MethodBindPtr
	fnSetTrackerType    gdc.MethodBindPtr
	fnGetTrackerName    gdc.MethodBindPtr
	fnSetTrackerName    gdc.MethodBindPtr
	fnGetTrackerDesc    gdc.MethodBindPtr
	fnSetTrackerDesc    gdc.MethodBindPtr
	fnGetTrackerProfile gdc.MethodBindPtr
	fnSetTrackerProfile gdc.MethodBindPtr
	fnGetTrackerHand    gdc.MethodBindPtr
	fnSetTrackerHand    gdc.MethodBindPtr
	fnHasPose           gdc.MethodBindPtr
	fnGetPose           gdc.MethodBindPtr
	fnInvalidatePose    gdc.MethodBindPtr
	fnSetPose           gdc.MethodBindPtr
	fnGetInput          gdc.MethodBindPtr
	fnSetInput          gdc.MethodBindPtr
}

var ptrsForXRPositionalTracker ptrsForXRPositionalTrackerList

func initXRPositionalTrackerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRPositionalTracker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_tracker_type")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetTrackerType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2784508102))
	}
	{
		methodName := StringNameFromStr("set_tracker_type")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetTrackerType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3055763575))
	}
	{
		methodName := StringNameFromStr("get_tracker_name")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetTrackerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_tracker_name")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetTrackerName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_tracker_desc")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetTrackerDesc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_tracker_desc")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetTrackerDesc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_tracker_profile")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetTrackerProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_tracker_profile")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetTrackerProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_tracker_hand")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetTrackerHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4181770860))
	}
	{
		methodName := StringNameFromStr("set_tracker_hand")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetTrackerHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3904108980))
	}
	{
		methodName := StringNameFromStr("has_pose")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnHasPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_pose")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4099720006))
	}
	{
		methodName := StringNameFromStr("invalidate_pose")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnInvalidatePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("set_pose")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3451230163))
	}
	{
		methodName := StringNameFromStr("get_input")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnGetInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("set_input")
		defer methodName.Destroy()
		ptrsForXRPositionalTracker.fnSetInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}

}

type XRPositionalTracker struct {
	RefCounted
}

func (me *XRPositionalTracker) BaseClass() string {
	return "XRPositionalTracker"
}

func NewXRPositionalTracker() *XRPositionalTracker {
	str := StringNameFromStr("XRPositionalTracker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRPositionalTracker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRPositionalTrackerTrackerHand int

const (
	XRPositionalTrackerTrackerHandTrackerHandUnknown XRPositionalTrackerTrackerHand = 0
	XRPositionalTrackerTrackerHandTrackerHandLeft    XRPositionalTrackerTrackerHand = 1
	XRPositionalTrackerTrackerHandTrackerHandRight   XRPositionalTrackerTrackerHand = 2
)

func (me *XRPositionalTracker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRPositionalTracker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRPositionalTracker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRPositionalTracker) GetTrackerType() XRServerTrackerType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRServerTrackerType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetTrackerType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRPositionalTracker) SetTrackerType(type_ XRServerTrackerType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetTrackerType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) GetTrackerName() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetTrackerName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRPositionalTracker) SetTrackerName(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetTrackerName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) GetTrackerDesc() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetTrackerDesc), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRPositionalTracker) SetTrackerDesc(description String) {
	cargs := []gdc.ConstTypePtr{description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetTrackerDesc), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) GetTrackerProfile() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetTrackerProfile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRPositionalTracker) SetTrackerProfile(profile String) {
	cargs := []gdc.ConstTypePtr{profile.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetTrackerProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) GetTrackerHand() XRPositionalTrackerTrackerHand {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRPositionalTrackerTrackerHand

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetTrackerHand), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRPositionalTracker) SetTrackerHand(hand XRPositionalTrackerTrackerHand) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetTrackerHand), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) HasPose(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnHasPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRPositionalTracker) GetPose(name StringName) XRPose {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRPose()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRPositionalTracker) InvalidatePose(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnInvalidatePose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) SetPose(name StringName, transform Transform3D, linear_velocity Vector3, angular_velocity Vector3, tracking_confidence XRPoseTrackingConfidence) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), transform.AsCTypePtr(), linear_velocity.AsCTypePtr(), angular_velocity.AsCTypePtr(), gdc.ConstTypePtr(&tracking_confidence)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetPose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRPositionalTracker) GetInput(name StringName) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnGetInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRPositionalTracker) SetInput(name StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRPositionalTracker.fnSetInput), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRPositionalTrackerPoseChangedSignalFn func(pose XRPose)

func (me *XRPositionalTracker) ConnectPoseChanged(subs SignalSubscribers, fn XRPositionalTrackerPoseChangedSignalFn) {
	sig := StringNameFromStr("pose_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectPoseChanged(subs SignalSubscribers, fn XRPositionalTrackerPoseChangedSignalFn) {
	sig := StringNameFromStr("pose_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerPoseLostTrackingSignalFn func(pose XRPose)

func (me *XRPositionalTracker) ConnectPoseLostTracking(subs SignalSubscribers, fn XRPositionalTrackerPoseLostTrackingSignalFn) {
	sig := StringNameFromStr("pose_lost_tracking")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectPoseLostTracking(subs SignalSubscribers, fn XRPositionalTrackerPoseLostTrackingSignalFn) {
	sig := StringNameFromStr("pose_lost_tracking")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerButtonPressedSignalFn func(name String)

func (me *XRPositionalTracker) ConnectButtonPressed(subs SignalSubscribers, fn XRPositionalTrackerButtonPressedSignalFn) {
	sig := StringNameFromStr("button_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectButtonPressed(subs SignalSubscribers, fn XRPositionalTrackerButtonPressedSignalFn) {
	sig := StringNameFromStr("button_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerButtonReleasedSignalFn func(name String)

func (me *XRPositionalTracker) ConnectButtonReleased(subs SignalSubscribers, fn XRPositionalTrackerButtonReleasedSignalFn) {
	sig := StringNameFromStr("button_released")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectButtonReleased(subs SignalSubscribers, fn XRPositionalTrackerButtonReleasedSignalFn) {
	sig := StringNameFromStr("button_released")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerInputFloatChangedSignalFn func(name String, value float32)

func (me *XRPositionalTracker) ConnectInputFloatChanged(subs SignalSubscribers, fn XRPositionalTrackerInputFloatChangedSignalFn) {
	sig := StringNameFromStr("input_float_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectInputFloatChanged(subs SignalSubscribers, fn XRPositionalTrackerInputFloatChangedSignalFn) {
	sig := StringNameFromStr("input_float_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerInputVector2ChangedSignalFn func(name String, vector Vector2)

func (me *XRPositionalTracker) ConnectInputVector2Changed(subs SignalSubscribers, fn XRPositionalTrackerInputVector2ChangedSignalFn) {
	sig := StringNameFromStr("input_vector2_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectInputVector2Changed(subs SignalSubscribers, fn XRPositionalTrackerInputVector2ChangedSignalFn) {
	sig := StringNameFromStr("input_vector2_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRPositionalTrackerProfileChangedSignalFn func(role String)

func (me *XRPositionalTracker) ConnectProfileChanged(subs SignalSubscribers, fn XRPositionalTrackerProfileChangedSignalFn) {
	sig := StringNameFromStr("profile_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectProfileChanged(subs SignalSubscribers, fn XRPositionalTrackerProfileChangedSignalFn) {
	sig := StringNameFromStr("profile_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
