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

type ptrsForXRBodyTrackerList struct {
	fnSetHasTrackingData gdc.MethodBindPtr
	fnGetHasTrackingData gdc.MethodBindPtr
	fnSetBodyFlags       gdc.MethodBindPtr
	fnGetBodyFlags       gdc.MethodBindPtr
	fnSetJointFlags      gdc.MethodBindPtr
	fnGetJointFlags      gdc.MethodBindPtr
	fnSetJointTransform  gdc.MethodBindPtr
	fnGetJointTransform  gdc.MethodBindPtr
}

var ptrsForXRBodyTracker ptrsForXRBodyTrackerList

func initXRBodyTrackerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRBodyTracker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_has_tracking_data")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnSetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_has_tracking_data")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnGetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_body_flags")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnSetBodyFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2103235750))
	}
	{
		methodName := StringNameFromStr("get_body_flags")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnGetBodyFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3543166366))
	}
	{
		methodName := StringNameFromStr("set_joint_flags")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnSetJointFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 592144999))
	}
	{
		methodName := StringNameFromStr("get_joint_flags")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnGetJointFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1030162609))
	}
	{
		methodName := StringNameFromStr("set_joint_transform")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnSetJointTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2635424328))
	}
	{
		methodName := StringNameFromStr("get_joint_transform")
		defer methodName.Destroy()
		ptrsForXRBodyTracker.fnGetJointTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3474811534))
	}

}

type XRBodyTracker struct {
	RefCounted
}

func (me *XRBodyTracker) BaseClass() string {
	return "XRBodyTracker"
}

func NewXRBodyTracker() *XRBodyTracker {
	str := StringNameFromStr("XRBodyTracker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRBodyTracker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRBodyTrackerBodyFlags int

const (
	XRBodyTrackerBodyFlagsBodyFlagUpperBodySupported XRBodyTrackerBodyFlags = 1
	XRBodyTrackerBodyFlagsBodyFlagLowerBodySupported XRBodyTrackerBodyFlags = 2
	XRBodyTrackerBodyFlagsBodyFlagHandsSupported     XRBodyTrackerBodyFlags = 4
)

type XRBodyTrackerJoint int

const (
	XRBodyTrackerJointJointRoot                                 XRBodyTrackerJoint = 0
	XRBodyTrackerJointJointHips                                 XRBodyTrackerJoint = 1
	XRBodyTrackerJointJointSpine                                XRBodyTrackerJoint = 2
	XRBodyTrackerJointJointChest                                XRBodyTrackerJoint = 3
	XRBodyTrackerJointJointUpperChest                           XRBodyTrackerJoint = 4
	XRBodyTrackerJointJointNeck                                 XRBodyTrackerJoint = 5
	XRBodyTrackerJointJointHead                                 XRBodyTrackerJoint = 6
	XRBodyTrackerJointJointHeadTip                              XRBodyTrackerJoint = 7
	XRBodyTrackerJointJointLeftShoulder                         XRBodyTrackerJoint = 8
	XRBodyTrackerJointJointLeftUpperArm                         XRBodyTrackerJoint = 9
	XRBodyTrackerJointJointLeftLowerArm                         XRBodyTrackerJoint = 10
	XRBodyTrackerJointJointRightShoulder                        XRBodyTrackerJoint = 11
	XRBodyTrackerJointJointRightUpperArm                        XRBodyTrackerJoint = 12
	XRBodyTrackerJointJointRightLowerArm                        XRBodyTrackerJoint = 13
	XRBodyTrackerJointJointLeftUpperLeg                         XRBodyTrackerJoint = 14
	XRBodyTrackerJointJointLeftLowerLeg                         XRBodyTrackerJoint = 15
	XRBodyTrackerJointJointLeftFoot                             XRBodyTrackerJoint = 16
	XRBodyTrackerJointJointLeftToes                             XRBodyTrackerJoint = 17
	XRBodyTrackerJointJointRightUpperLeg                        XRBodyTrackerJoint = 18
	XRBodyTrackerJointJointRightLowerLeg                        XRBodyTrackerJoint = 19
	XRBodyTrackerJointJointRightFoot                            XRBodyTrackerJoint = 20
	XRBodyTrackerJointJointRightToes                            XRBodyTrackerJoint = 21
	XRBodyTrackerJointJointLeftHand                             XRBodyTrackerJoint = 22
	XRBodyTrackerJointJointLeftPalm                             XRBodyTrackerJoint = 23
	XRBodyTrackerJointJointLeftWrist                            XRBodyTrackerJoint = 24
	XRBodyTrackerJointJointLeftThumbMetacarpal                  XRBodyTrackerJoint = 25
	XRBodyTrackerJointJointLeftThumbPhalanxProximal             XRBodyTrackerJoint = 26
	XRBodyTrackerJointJointLeftThumbPhalanxDistal               XRBodyTrackerJoint = 27
	XRBodyTrackerJointJointLeftThumbTip                         XRBodyTrackerJoint = 28
	XRBodyTrackerJointJointLeftIndexFingerMetacarpal            XRBodyTrackerJoint = 29
	XRBodyTrackerJointJointLeftIndexFingerPhalanxProximal       XRBodyTrackerJoint = 30
	XRBodyTrackerJointJointLeftIndexFingerPhalanxIntermediate   XRBodyTrackerJoint = 31
	XRBodyTrackerJointJointLeftIndexFingerPhalanxDistal         XRBodyTrackerJoint = 32
	XRBodyTrackerJointJointLeftIndexFingerTip                   XRBodyTrackerJoint = 33
	XRBodyTrackerJointJointLeftMiddleFingerMetacarpal           XRBodyTrackerJoint = 34
	XRBodyTrackerJointJointLeftMiddleFingerPhalanxProximal      XRBodyTrackerJoint = 35
	XRBodyTrackerJointJointLeftMiddleFingerPhalanxIntermediate  XRBodyTrackerJoint = 36
	XRBodyTrackerJointJointLeftMiddleFingerPhalanxDistal        XRBodyTrackerJoint = 37
	XRBodyTrackerJointJointLeftMiddleFingerTip                  XRBodyTrackerJoint = 38
	XRBodyTrackerJointJointLeftRingFingerMetacarpal             XRBodyTrackerJoint = 39
	XRBodyTrackerJointJointLeftRingFingerPhalanxProximal        XRBodyTrackerJoint = 40
	XRBodyTrackerJointJointLeftRingFingerPhalanxIntermediate    XRBodyTrackerJoint = 41
	XRBodyTrackerJointJointLeftRingFingerPhalanxDistal          XRBodyTrackerJoint = 42
	XRBodyTrackerJointJointLeftRingFingerTip                    XRBodyTrackerJoint = 43
	XRBodyTrackerJointJointLeftPinkyFingerMetacarpal            XRBodyTrackerJoint = 44
	XRBodyTrackerJointJointLeftPinkyFingerPhalanxProximal       XRBodyTrackerJoint = 45
	XRBodyTrackerJointJointLeftPinkyFingerPhalanxIntermediate   XRBodyTrackerJoint = 46
	XRBodyTrackerJointJointLeftPinkyFingerPhalanxDistal         XRBodyTrackerJoint = 47
	XRBodyTrackerJointJointLeftPinkyFingerTip                   XRBodyTrackerJoint = 48
	XRBodyTrackerJointJointRightHand                            XRBodyTrackerJoint = 49
	XRBodyTrackerJointJointRightPalm                            XRBodyTrackerJoint = 50
	XRBodyTrackerJointJointRightWrist                           XRBodyTrackerJoint = 51
	XRBodyTrackerJointJointRightThumbMetacarpal                 XRBodyTrackerJoint = 52
	XRBodyTrackerJointJointRightThumbPhalanxProximal            XRBodyTrackerJoint = 53
	XRBodyTrackerJointJointRightThumbPhalanxDistal              XRBodyTrackerJoint = 54
	XRBodyTrackerJointJointRightThumbTip                        XRBodyTrackerJoint = 55
	XRBodyTrackerJointJointRightIndexFingerMetacarpal           XRBodyTrackerJoint = 56
	XRBodyTrackerJointJointRightIndexFingerPhalanxProximal      XRBodyTrackerJoint = 57
	XRBodyTrackerJointJointRightIndexFingerPhalanxIntermediate  XRBodyTrackerJoint = 58
	XRBodyTrackerJointJointRightIndexFingerPhalanxDistal        XRBodyTrackerJoint = 59
	XRBodyTrackerJointJointRightIndexFingerTip                  XRBodyTrackerJoint = 60
	XRBodyTrackerJointJointRightMiddleFingerMetacarpal          XRBodyTrackerJoint = 61
	XRBodyTrackerJointJointRightMiddleFingerPhalanxProximal     XRBodyTrackerJoint = 62
	XRBodyTrackerJointJointRightMiddleFingerPhalanxIntermediate XRBodyTrackerJoint = 63
	XRBodyTrackerJointJointRightMiddleFingerPhalanxDistal       XRBodyTrackerJoint = 64
	XRBodyTrackerJointJointRightMiddleFingerTip                 XRBodyTrackerJoint = 65
	XRBodyTrackerJointJointRightRingFingerMetacarpal            XRBodyTrackerJoint = 66
	XRBodyTrackerJointJointRightRingFingerPhalanxProximal       XRBodyTrackerJoint = 67
	XRBodyTrackerJointJointRightRingFingerPhalanxIntermediate   XRBodyTrackerJoint = 68
	XRBodyTrackerJointJointRightRingFingerPhalanxDistal         XRBodyTrackerJoint = 69
	XRBodyTrackerJointJointRightRingFingerTip                   XRBodyTrackerJoint = 70
	XRBodyTrackerJointJointRightPinkyFingerMetacarpal           XRBodyTrackerJoint = 71
	XRBodyTrackerJointJointRightPinkyFingerPhalanxProximal      XRBodyTrackerJoint = 72
	XRBodyTrackerJointJointRightPinkyFingerPhalanxIntermediate  XRBodyTrackerJoint = 73
	XRBodyTrackerJointJointRightPinkyFingerPhalanxDistal        XRBodyTrackerJoint = 74
	XRBodyTrackerJointJointRightPinkyFingerTip                  XRBodyTrackerJoint = 75
	XRBodyTrackerJointJointMax                                  XRBodyTrackerJoint = 76
)

type XRBodyTrackerJointFlags int

const (
	XRBodyTrackerJointFlagsJointFlagOrientationValid   XRBodyTrackerJointFlags = 1
	XRBodyTrackerJointFlagsJointFlagOrientationTracked XRBodyTrackerJointFlags = 2
	XRBodyTrackerJointFlagsJointFlagPositionValid      XRBodyTrackerJointFlags = 4
	XRBodyTrackerJointFlagsJointFlagPositionTracked    XRBodyTrackerJointFlags = 8
)

func (me *XRBodyTracker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRBodyTracker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRBodyTracker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRBodyTracker) SetHasTrackingData(has_data bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&has_data)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnSetHasTrackingData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyTracker) GetHasTrackingData() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnGetHasTrackingData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRBodyTracker) SetBodyFlags(flags XRBodyTrackerBodyFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnSetBodyFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyTracker) GetBodyFlags() XRBodyTrackerBodyFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRBodyTrackerBodyFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnGetBodyFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRBodyTracker) SetJointFlags(joint XRBodyTrackerJoint, flags XRBodyTrackerJointFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnSetJointFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyTracker) GetJointFlags(joint XRBodyTrackerJoint) XRBodyTrackerJointFlags {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRBodyTrackerJointFlags
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnGetJointFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRBodyTracker) SetJointTransform(joint XRBodyTrackerJoint, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnSetJointTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyTracker) GetJointTransform(joint XRBodyTrackerJoint) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyTracker.fnGetJointTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
