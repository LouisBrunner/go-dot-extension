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

type ptrsForOpenXRInterfaceList struct {
	fnGetDisplayRefreshRate           gdc.MethodBindPtr
	fnSetDisplayRefreshRate           gdc.MethodBindPtr
	fnGetRenderTargetSizeMultiplier   gdc.MethodBindPtr
	fnSetRenderTargetSizeMultiplier   gdc.MethodBindPtr
	fnIsFoveationSupported            gdc.MethodBindPtr
	fnGetFoveationLevel               gdc.MethodBindPtr
	fnSetFoveationLevel               gdc.MethodBindPtr
	fnGetFoveationDynamic             gdc.MethodBindPtr
	fnSetFoveationDynamic             gdc.MethodBindPtr
	fnIsActionSetActive               gdc.MethodBindPtr
	fnSetActionSetActive              gdc.MethodBindPtr
	fnGetActionSets                   gdc.MethodBindPtr
	fnGetAvailableDisplayRefreshRates gdc.MethodBindPtr
	fnSetMotionRange                  gdc.MethodBindPtr
	fnGetMotionRange                  gdc.MethodBindPtr
	fnGetHandJointFlags               gdc.MethodBindPtr
	fnGetHandJointRotation            gdc.MethodBindPtr
	fnGetHandJointPosition            gdc.MethodBindPtr
	fnGetHandJointRadius              gdc.MethodBindPtr
	fnGetHandJointLinearVelocity      gdc.MethodBindPtr
	fnGetHandJointAngularVelocity     gdc.MethodBindPtr
	fnIsHandTrackingSupported         gdc.MethodBindPtr
	fnIsEyeGazeInteractionSupported   gdc.MethodBindPtr
}

var ptrsForOpenXRInterface ptrsForOpenXRInterfaceList

func initOpenXRInterfacePtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRInterface")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_display_refresh_rate")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetDisplayRefreshRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_display_refresh_rate")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetDisplayRefreshRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_render_target_size_multiplier")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetRenderTargetSizeMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_render_target_size_multiplier")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetRenderTargetSizeMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("is_foveation_supported")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnIsFoveationSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_foveation_level")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetFoveationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_foveation_level")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetFoveationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_foveation_dynamic")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetFoveationDynamic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_foveation_dynamic")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetFoveationDynamic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_action_set_active")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnIsActionSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("set_action_set_active")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetActionSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("get_action_sets")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetActionSets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_available_display_refresh_rates")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetAvailableDisplayRefreshRates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_motion_range")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnSetMotionRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 855158159))
	}
	{
		methodName := StringNameFromStr("get_motion_range")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetMotionRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3955838114))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_flags")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 720567706))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_rotation")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1974618321))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_position")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3529194242))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_radius")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 901522724))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_linear_velocity")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3529194242))
	}
	{
		methodName := StringNameFromStr("get_hand_joint_angular_velocity")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnGetHandJointAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3529194242))
	}
	{
		methodName := StringNameFromStr("is_hand_tracking_supported")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnIsHandTrackingSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_eye_gaze_interaction_supported")
		defer methodName.Destroy()
		ptrsForOpenXRInterface.fnIsEyeGazeInteractionSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
}

type OpenXRInterface struct {
	XRInterface
}

func (me *OpenXRInterface) BaseClass() string {
	return "OpenXRInterface"
}

func NewOpenXRInterface() *OpenXRInterface {
	str := StringNameFromStr("OpenXRInterface") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRInterface{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type OpenXRInterfaceHand int

const (
	OpenXRInterfaceHandHandLeft  OpenXRInterfaceHand = 0
	OpenXRInterfaceHandHandRight OpenXRInterfaceHand = 1
	OpenXRInterfaceHandHandMax   OpenXRInterfaceHand = 2
)

type OpenXRInterfaceHandMotionRange int

const (
	OpenXRInterfaceHandMotionRangeHandMotionRangeUnobstructed        OpenXRInterfaceHandMotionRange = 0
	OpenXRInterfaceHandMotionRangeHandMotionRangeConformToController OpenXRInterfaceHandMotionRange = 1
	OpenXRInterfaceHandMotionRangeHandMotionRangeMax                 OpenXRInterfaceHandMotionRange = 2
)

type OpenXRInterfaceHandJoints int

const (
	OpenXRInterfaceHandJointsHandJointPalm               OpenXRInterfaceHandJoints = 0
	OpenXRInterfaceHandJointsHandJointWrist              OpenXRInterfaceHandJoints = 1
	OpenXRInterfaceHandJointsHandJointThumbMetacarpal    OpenXRInterfaceHandJoints = 2
	OpenXRInterfaceHandJointsHandJointThumbProximal      OpenXRInterfaceHandJoints = 3
	OpenXRInterfaceHandJointsHandJointThumbDistal        OpenXRInterfaceHandJoints = 4
	OpenXRInterfaceHandJointsHandJointThumbTip           OpenXRInterfaceHandJoints = 5
	OpenXRInterfaceHandJointsHandJointIndexMetacarpal    OpenXRInterfaceHandJoints = 6
	OpenXRInterfaceHandJointsHandJointIndexProximal      OpenXRInterfaceHandJoints = 7
	OpenXRInterfaceHandJointsHandJointIndexIntermediate  OpenXRInterfaceHandJoints = 8
	OpenXRInterfaceHandJointsHandJointIndexDistal        OpenXRInterfaceHandJoints = 9
	OpenXRInterfaceHandJointsHandJointIndexTip           OpenXRInterfaceHandJoints = 10
	OpenXRInterfaceHandJointsHandJointMiddleMetacarpal   OpenXRInterfaceHandJoints = 11
	OpenXRInterfaceHandJointsHandJointMiddleProximal     OpenXRInterfaceHandJoints = 12
	OpenXRInterfaceHandJointsHandJointMiddleIntermediate OpenXRInterfaceHandJoints = 13
	OpenXRInterfaceHandJointsHandJointMiddleDistal       OpenXRInterfaceHandJoints = 14
	OpenXRInterfaceHandJointsHandJointMiddleTip          OpenXRInterfaceHandJoints = 15
	OpenXRInterfaceHandJointsHandJointRingMetacarpal     OpenXRInterfaceHandJoints = 16
	OpenXRInterfaceHandJointsHandJointRingProximal       OpenXRInterfaceHandJoints = 17
	OpenXRInterfaceHandJointsHandJointRingIntermediate   OpenXRInterfaceHandJoints = 18
	OpenXRInterfaceHandJointsHandJointRingDistal         OpenXRInterfaceHandJoints = 19
	OpenXRInterfaceHandJointsHandJointRingTip            OpenXRInterfaceHandJoints = 20
	OpenXRInterfaceHandJointsHandJointLittleMetacarpal   OpenXRInterfaceHandJoints = 21
	OpenXRInterfaceHandJointsHandJointLittleProximal     OpenXRInterfaceHandJoints = 22
	OpenXRInterfaceHandJointsHandJointLittleIntermediate OpenXRInterfaceHandJoints = 23
	OpenXRInterfaceHandJointsHandJointLittleDistal       OpenXRInterfaceHandJoints = 24
	OpenXRInterfaceHandJointsHandJointLittleTip          OpenXRInterfaceHandJoints = 25
	OpenXRInterfaceHandJointsHandJointMax                OpenXRInterfaceHandJoints = 26
)

type OpenXRInterfaceHandJointFlags int

const (
	OpenXRInterfaceHandJointFlagsHandJointNone                 OpenXRInterfaceHandJointFlags = 0
	OpenXRInterfaceHandJointFlagsHandJointOrientationValid     OpenXRInterfaceHandJointFlags = 1
	OpenXRInterfaceHandJointFlagsHandJointOrientationTracked   OpenXRInterfaceHandJointFlags = 2
	OpenXRInterfaceHandJointFlagsHandJointPositionValid        OpenXRInterfaceHandJointFlags = 4
	OpenXRInterfaceHandJointFlagsHandJointPositionTracked      OpenXRInterfaceHandJointFlags = 8
	OpenXRInterfaceHandJointFlagsHandJointLinearVelocityValid  OpenXRInterfaceHandJointFlags = 16
	OpenXRInterfaceHandJointFlagsHandJointAngularVelocityValid OpenXRInterfaceHandJointFlags = 32
)

func (me *OpenXRInterface) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRInterface) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRInterface) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRInterface) GetDisplayRefreshRate() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetDisplayRefreshRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) SetDisplayRefreshRate(refresh_rate float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refresh_rate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetDisplayRefreshRate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) GetRenderTargetSizeMultiplier() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetRenderTargetSizeMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) SetRenderTargetSizeMultiplier(multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetRenderTargetSizeMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) IsFoveationSupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnIsFoveationSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) GetFoveationLevel() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetFoveationLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) SetFoveationLevel(foveation_level int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&foveation_level)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetFoveationLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) GetFoveationDynamic() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetFoveationDynamic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) SetFoveationDynamic(foveation_dynamic bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&foveation_dynamic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetFoveationDynamic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) IsActionSetActive(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnIsActionSetActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) SetActionSetActive(name String, active bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetActionSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) GetActionSets() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetActionSets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) GetAvailableDisplayRefreshRates() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetAvailableDisplayRefreshRates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) SetMotionRange(hand OpenXRInterfaceHand, motion_range OpenXRInterfaceHandMotionRange) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&motion_range)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnSetMotionRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInterface) GetMotionRange(hand OpenXRInterfaceHand) OpenXRInterfaceHandMotionRange {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret OpenXRInterfaceHandMotionRange
	pinner.Pin(&hand)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetMotionRange), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OpenXRInterface) GetHandJointFlags(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) OpenXRInterfaceHandJointFlags {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret OpenXRInterfaceHandJointFlags
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OpenXRInterface) GetHandJointRotation(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) Quaternion {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewQuaternion()
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) GetHandJointPosition(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) GetHandJointRadius(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) GetHandJointLinearVelocity(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) GetHandJointAngularVelocity(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&hand)
	pinner.Pin(&joint)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnGetHandJointAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRInterface) IsHandTrackingSupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnIsHandTrackingSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRInterface) IsEyeGazeInteractionSupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInterface.fnIsEyeGazeInteractionSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type OpenXRInterfaceSessionBegunSignalFn func()

func (me *OpenXRInterface) ConnectSessionBegun(subs SignalSubscribers, fn OpenXRInterfaceSessionBegunSignalFn) {
	sig := StringNameFromStr("session_begun")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *OpenXRInterface) DisconnectSessionBegun(subs SignalSubscribers, fn OpenXRInterfaceSessionBegunSignalFn) {
	sig := StringNameFromStr("session_begun")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type OpenXRInterfaceSessionStoppingSignalFn func()

func (me *OpenXRInterface) ConnectSessionStopping(subs SignalSubscribers, fn OpenXRInterfaceSessionStoppingSignalFn) {
	sig := StringNameFromStr("session_stopping")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *OpenXRInterface) DisconnectSessionStopping(subs SignalSubscribers, fn OpenXRInterfaceSessionStoppingSignalFn) {
	sig := StringNameFromStr("session_stopping")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type OpenXRInterfaceSessionFocussedSignalFn func()

func (me *OpenXRInterface) ConnectSessionFocussed(subs SignalSubscribers, fn OpenXRInterfaceSessionFocussedSignalFn) {
	sig := StringNameFromStr("session_focussed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *OpenXRInterface) DisconnectSessionFocussed(subs SignalSubscribers, fn OpenXRInterfaceSessionFocussedSignalFn) {
	sig := StringNameFromStr("session_focussed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type OpenXRInterfaceSessionVisibleSignalFn func()

func (me *OpenXRInterface) ConnectSessionVisible(subs SignalSubscribers, fn OpenXRInterfaceSessionVisibleSignalFn) {
	sig := StringNameFromStr("session_visible")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *OpenXRInterface) DisconnectSessionVisible(subs SignalSubscribers, fn OpenXRInterfaceSessionVisibleSignalFn) {
	sig := StringNameFromStr("session_visible")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type OpenXRInterfacePoseRecenteredSignalFn func()

func (me *OpenXRInterface) ConnectPoseRecentered(subs SignalSubscribers, fn OpenXRInterfacePoseRecenteredSignalFn) {
	sig := StringNameFromStr("pose_recentered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *OpenXRInterface) DisconnectPoseRecentered(subs SignalSubscribers, fn OpenXRInterfacePoseRecenteredSignalFn) {
	sig := StringNameFromStr("pose_recentered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
