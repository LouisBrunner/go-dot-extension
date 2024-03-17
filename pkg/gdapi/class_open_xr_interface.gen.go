// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRInterface struct {
  XRInterface
}

func (me *OpenXRInterface) BaseClass() string {
  return "OpenXRInterface"
}



// Enums

type OpenXRInterfaceHand int
const (
  OpenXRInterfaceHandHandLeft OpenXRInterfaceHand = 0
  OpenXRInterfaceHandHandRight OpenXRInterfaceHand = 1
  OpenXRInterfaceHandHandMax OpenXRInterfaceHand = 2
)

type OpenXRInterfaceHandMotionRange int
const (
  OpenXRInterfaceHandMotionRangeHandMotionRangeUnobstructed OpenXRInterfaceHandMotionRange = 0
  OpenXRInterfaceHandMotionRangeHandMotionRangeConformToController OpenXRInterfaceHandMotionRange = 1
  OpenXRInterfaceHandMotionRangeHandMotionRangeMax OpenXRInterfaceHandMotionRange = 2
)

type OpenXRInterfaceHandJoints int
const (
  OpenXRInterfaceHandJointsHandJointPalm OpenXRInterfaceHandJoints = 0
  OpenXRInterfaceHandJointsHandJointWrist OpenXRInterfaceHandJoints = 1
  OpenXRInterfaceHandJointsHandJointThumbMetacarpal OpenXRInterfaceHandJoints = 2
  OpenXRInterfaceHandJointsHandJointThumbProximal OpenXRInterfaceHandJoints = 3
  OpenXRInterfaceHandJointsHandJointThumbDistal OpenXRInterfaceHandJoints = 4
  OpenXRInterfaceHandJointsHandJointThumbTip OpenXRInterfaceHandJoints = 5
  OpenXRInterfaceHandJointsHandJointIndexMetacarpal OpenXRInterfaceHandJoints = 6
  OpenXRInterfaceHandJointsHandJointIndexProximal OpenXRInterfaceHandJoints = 7
  OpenXRInterfaceHandJointsHandJointIndexIntermediate OpenXRInterfaceHandJoints = 8
  OpenXRInterfaceHandJointsHandJointIndexDistal OpenXRInterfaceHandJoints = 9
  OpenXRInterfaceHandJointsHandJointIndexTip OpenXRInterfaceHandJoints = 10
  OpenXRInterfaceHandJointsHandJointMiddleMetacarpal OpenXRInterfaceHandJoints = 11
  OpenXRInterfaceHandJointsHandJointMiddleProximal OpenXRInterfaceHandJoints = 12
  OpenXRInterfaceHandJointsHandJointMiddleIntermediate OpenXRInterfaceHandJoints = 13
  OpenXRInterfaceHandJointsHandJointMiddleDistal OpenXRInterfaceHandJoints = 14
  OpenXRInterfaceHandJointsHandJointMiddleTip OpenXRInterfaceHandJoints = 15
  OpenXRInterfaceHandJointsHandJointRingMetacarpal OpenXRInterfaceHandJoints = 16
  OpenXRInterfaceHandJointsHandJointRingProximal OpenXRInterfaceHandJoints = 17
  OpenXRInterfaceHandJointsHandJointRingIntermediate OpenXRInterfaceHandJoints = 18
  OpenXRInterfaceHandJointsHandJointRingDistal OpenXRInterfaceHandJoints = 19
  OpenXRInterfaceHandJointsHandJointRingTip OpenXRInterfaceHandJoints = 20
  OpenXRInterfaceHandJointsHandJointLittleMetacarpal OpenXRInterfaceHandJoints = 21
  OpenXRInterfaceHandJointsHandJointLittleProximal OpenXRInterfaceHandJoints = 22
  OpenXRInterfaceHandJointsHandJointLittleIntermediate OpenXRInterfaceHandJoints = 23
  OpenXRInterfaceHandJointsHandJointLittleDistal OpenXRInterfaceHandJoints = 24
  OpenXRInterfaceHandJointsHandJointLittleTip OpenXRInterfaceHandJoints = 25
  OpenXRInterfaceHandJointsHandJointMax OpenXRInterfaceHandJoints = 26
)

type OpenXRInterfaceHandJointFlags int
const (
  OpenXRInterfaceHandJointFlagsHandJointNone OpenXRInterfaceHandJointFlags = 0
  OpenXRInterfaceHandJointFlagsHandJointOrientationValid OpenXRInterfaceHandJointFlags = 1
  OpenXRInterfaceHandJointFlagsHandJointOrientationTracked OpenXRInterfaceHandJointFlags = 2
  OpenXRInterfaceHandJointFlagsHandJointPositionValid OpenXRInterfaceHandJointFlags = 4
  OpenXRInterfaceHandJointFlagsHandJointPositionTracked OpenXRInterfaceHandJointFlags = 8
  OpenXRInterfaceHandJointFlagsHandJointLinearVelocityValid OpenXRInterfaceHandJointFlags = 16
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

func  (me *OpenXRInterface) GetDisplayRefreshRate() float32 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetDisplayRefreshRate(refresh_rate float32, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refresh_rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) GetRenderTargetSizeMultiplier() float32 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target_size_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetRenderTargetSizeMultiplier(multiplier float32, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_render_target_size_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) IsFoveationSupported() bool {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_foveation_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetFoveationLevel() int {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_foveation_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetFoveationLevel(foveation_level int, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_foveation_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&foveation_level), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) GetFoveationDynamic() bool {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_foveation_dynamic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetFoveationDynamic(foveation_dynamic bool, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_foveation_dynamic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&foveation_dynamic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) IsActionSetActive(name String, ) bool {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetActionSetActive(name String, active bool, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) GetActionSets() Array {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_sets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetAvailableDisplayRefreshRates() Array {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_display_refresh_rates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) SetMotionRange(hand OpenXRInterfaceHand, motion_range OpenXRInterfaceHandMotionRange, )  {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 855158159) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&motion_range), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInterface) GetMotionRange(hand OpenXRInterfaceHand, ) OpenXRInterfaceHandMotionRange {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3955838114) // FIXME: should cache?
  var ret OpenXRInterfaceHandMotionRange
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointFlags(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) OpenXRInterfaceHandJointFlags {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 720567706) // FIXME: should cache?
  var ret OpenXRInterfaceHandJointFlags
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointRotation(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) Quaternion {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1974618321) // FIXME: should cache?
  var ret Quaternion
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointPosition(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) Vector3 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3529194242) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointRadius(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) float32 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 901522724) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointLinearVelocity(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) Vector3 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3529194242) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) GetHandJointAngularVelocity(hand OpenXRInterfaceHand, joint OpenXRInterfaceHandJoints, ) Vector3 {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hand_joint_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3529194242) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), gdc.ConstTypePtr(&joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) IsHandTrackingSupported() bool {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hand_tracking_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRInterface) IsEyeGazeInteractionSupported() bool {
  classNameV := StringNameFromStr("OpenXRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_eye_gaze_interaction_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
