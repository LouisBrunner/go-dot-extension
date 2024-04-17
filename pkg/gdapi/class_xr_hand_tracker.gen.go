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

type ptrsForXRHandTrackerList struct {
	fnSetHand                     gdc.MethodBindPtr
	fnGetHand                     gdc.MethodBindPtr
	fnSetHasTrackingData          gdc.MethodBindPtr
	fnGetHasTrackingData          gdc.MethodBindPtr
	fnSetHandTrackingSource       gdc.MethodBindPtr
	fnGetHandTrackingSource       gdc.MethodBindPtr
	fnSetHandJointFlags           gdc.MethodBindPtr
	fnGetHandJointFlags           gdc.MethodBindPtr
	fnSetHandJointTransform       gdc.MethodBindPtr
	fnGetHandJointTransform       gdc.MethodBindPtr
	fnSetHandJointRadius          gdc.MethodBindPtr
	fnGetHandJointRadius          gdc.MethodBindPtr
	fnSetHandJointLinearVelocity  gdc.MethodBindPtr
	fnGetHandJointLinearVelocity  gdc.MethodBindPtr
	fnSetHandJointAngularVelocity gdc.MethodBindPtr
	fnGetHandJointAngularVelocity gdc.MethodBindPtr
}

var ptrsForXRHandTracker ptrsForXRHandTrackerList

func initXRHandTrackerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRHandTracker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_hand")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 89041945))
	}
	{
		methodName := StringNameFromStr("get_hand")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1790719290))
	}
	{
		methodName := StringNameFromStr("set_has_tracking_data")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_has_tracking_data")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hand_tracking_source")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandTrackingSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2958308861))
	}
	{
		methodName := StringNameFromStr("get_hand_tracking_source")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandTrackingSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2475045250))
	}
	{
		methodName := StringNameFromStr("set_hand_joint_flags")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandJointFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3028437365))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_flags")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandJointFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1730972401))
	}
	{
		methodName := StringNameFromStr("set_hand_joint_transform")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandJointTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2529959613))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_transform")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandJointTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1090840196))
	}
	{
		methodName := StringNameFromStr("set_hand_joint_radius")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandJointRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2723659615))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_radius")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandJointRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3400025734))
	}
	{
		methodName := StringNameFromStr("set_hand_joint_linear_velocity")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandJointLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1978646737))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_linear_velocity")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandJointLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 547240792))
	}
	{
		methodName := StringNameFromStr("set_hand_joint_angular_velocity")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnSetHandJointAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1978646737))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_angular_velocity")
		defer methodName.Destroy()
		ptrsForXRHandTracker.fnGetHandJointAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 547240792))
	}

}

type XRHandTracker struct {
	RefCounted
}

func (me *XRHandTracker) BaseClass() string {
	return "XRHandTracker"
}

func NewXRHandTracker() *XRHandTracker {
	str := StringNameFromStr("XRHandTracker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRHandTracker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRHandTrackerHand int

const (
	XRHandTrackerHandHandLeft  XRHandTrackerHand = 0
	XRHandTrackerHandHandRight XRHandTrackerHand = 1
	XRHandTrackerHandHandMax   XRHandTrackerHand = 2
)

type XRHandTrackerHandTrackingSource int

const (
	XRHandTrackerHandTrackingSourceHandTrackingSourceUnknown      XRHandTrackerHandTrackingSource = 0
	XRHandTrackerHandTrackingSourceHandTrackingSourceUnobstructed XRHandTrackerHandTrackingSource = 1
	XRHandTrackerHandTrackingSourceHandTrackingSourceController   XRHandTrackerHandTrackingSource = 2
	XRHandTrackerHandTrackingSourceHandTrackingSourceMax          XRHandTrackerHandTrackingSource = 3
)

type XRHandTrackerHandJoint int

const (
	XRHandTrackerHandJointHandJointPalm                            XRHandTrackerHandJoint = 0
	XRHandTrackerHandJointHandJointWrist                           XRHandTrackerHandJoint = 1
	XRHandTrackerHandJointHandJointThumbMetacarpal                 XRHandTrackerHandJoint = 2
	XRHandTrackerHandJointHandJointThumbPhalanxProximal            XRHandTrackerHandJoint = 3
	XRHandTrackerHandJointHandJointThumbPhalanxDistal              XRHandTrackerHandJoint = 4
	XRHandTrackerHandJointHandJointThumbTip                        XRHandTrackerHandJoint = 5
	XRHandTrackerHandJointHandJointIndexFingerMetacarpal           XRHandTrackerHandJoint = 6
	XRHandTrackerHandJointHandJointIndexFingerPhalanxProximal      XRHandTrackerHandJoint = 7
	XRHandTrackerHandJointHandJointIndexFingerPhalanxIntermediate  XRHandTrackerHandJoint = 8
	XRHandTrackerHandJointHandJointIndexFingerPhalanxDistal        XRHandTrackerHandJoint = 9
	XRHandTrackerHandJointHandJointIndexFingerTip                  XRHandTrackerHandJoint = 10
	XRHandTrackerHandJointHandJointMiddleFingerMetacarpal          XRHandTrackerHandJoint = 11
	XRHandTrackerHandJointHandJointMiddleFingerPhalanxProximal     XRHandTrackerHandJoint = 12
	XRHandTrackerHandJointHandJointMiddleFingerPhalanxIntermediate XRHandTrackerHandJoint = 13
	XRHandTrackerHandJointHandJointMiddleFingerPhalanxDistal       XRHandTrackerHandJoint = 14
	XRHandTrackerHandJointHandJointMiddleFingerTip                 XRHandTrackerHandJoint = 15
	XRHandTrackerHandJointHandJointRingFingerMetacarpal            XRHandTrackerHandJoint = 16
	XRHandTrackerHandJointHandJointRingFingerPhalanxProximal       XRHandTrackerHandJoint = 17
	XRHandTrackerHandJointHandJointRingFingerPhalanxIntermediate   XRHandTrackerHandJoint = 18
	XRHandTrackerHandJointHandJointRingFingerPhalanxDistal         XRHandTrackerHandJoint = 19
	XRHandTrackerHandJointHandJointRingFingerTip                   XRHandTrackerHandJoint = 20
	XRHandTrackerHandJointHandJointPinkyFingerMetacarpal           XRHandTrackerHandJoint = 21
	XRHandTrackerHandJointHandJointPinkyFingerPhalanxProximal      XRHandTrackerHandJoint = 22
	XRHandTrackerHandJointHandJointPinkyFingerPhalanxIntermediate  XRHandTrackerHandJoint = 23
	XRHandTrackerHandJointHandJointPinkyFingerPhalanxDistal        XRHandTrackerHandJoint = 24
	XRHandTrackerHandJointHandJointPinkyFingerTip                  XRHandTrackerHandJoint = 25
	XRHandTrackerHandJointHandJointMax                             XRHandTrackerHandJoint = 26
)

type XRHandTrackerHandJointFlags int

const (
	XRHandTrackerHandJointFlagsHandJointFlagOrientationValid     XRHandTrackerHandJointFlags = 1
	XRHandTrackerHandJointFlagsHandJointFlagOrientationTracked   XRHandTrackerHandJointFlags = 2
	XRHandTrackerHandJointFlagsHandJointFlagPositionValid        XRHandTrackerHandJointFlags = 4
	XRHandTrackerHandJointFlagsHandJointFlagPositionTracked      XRHandTrackerHandJointFlags = 8
	XRHandTrackerHandJointFlagsHandJointFlagLinearVelocityValid  XRHandTrackerHandJointFlags = 16
	XRHandTrackerHandJointFlagsHandJointFlagAngularVelocityValid XRHandTrackerHandJointFlags = 32
)

func (me *XRHandTracker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRHandTracker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRHandTracker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRHandTracker) SetHand(hand XRHandTrackerHand) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHand), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHand() XRHandTrackerHand {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRHandTrackerHand

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHand), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRHandTracker) SetHasTrackingData(has_data bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&has_data)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHasTrackingData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHasTrackingData() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHasTrackingData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRHandTracker) SetHandTrackingSource(source XRHandTrackerHandTrackingSource) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandTrackingSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandTrackingSource() XRHandTrackerHandTrackingSource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRHandTrackerHandTrackingSource

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandTrackingSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRHandTracker) SetHandJointFlags(joint XRHandTrackerHandJoint, flags XRHandTrackerHandJointFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandJointFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandJointFlags(joint XRHandTrackerHandJoint) XRHandTrackerHandJointFlags {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRHandTrackerHandJointFlags
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandJointFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRHandTracker) SetHandJointTransform(joint XRHandTrackerHandJoint, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandJointTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandJointTransform(joint XRHandTrackerHandJoint) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandJointTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRHandTracker) SetHandJointRadius(joint XRHandTrackerHandJoint, radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandJointRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandJointRadius(joint XRHandTrackerHandJoint) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandJointRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRHandTracker) SetHandJointLinearVelocity(joint XRHandTrackerHandJoint, linear_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), linear_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandJointLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandJointLinearVelocity(joint XRHandTrackerHandJoint) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandJointLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRHandTracker) SetHandJointAngularVelocity(joint XRHandTrackerHandJoint, angular_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), angular_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnSetHandJointAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandTracker) GetHandJointAngularVelocity(joint XRHandTrackerHandJoint) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandTracker.fnGetHandJointAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
