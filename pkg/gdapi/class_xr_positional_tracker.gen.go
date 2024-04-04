// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  XRPositionalTrackerTrackerHandTrackerHandLeft XRPositionalTrackerTrackerHand = 1
  XRPositionalTrackerTrackerHandTrackerHandRight XRPositionalTrackerTrackerHand = 2
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

func  (me *XRPositionalTracker) GetTrackerType() XRServerTrackerType {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2784508102) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRServerTrackerType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerType(type_ XRServerTrackerType, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3055763575) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) GetTrackerName() StringName {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPositionalTracker) SetTrackerName(name StringName, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) GetTrackerDesc() String {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_desc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPositionalTracker) SetTrackerDesc(description String, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_desc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) GetTrackerProfile() String {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPositionalTracker) SetTrackerProfile(profile String, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{profile.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) GetTrackerHand() XRPositionalTrackerTrackerHand {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4181770860) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret XRPositionalTrackerTrackerHand

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerHand(hand XRPositionalTrackerTrackerHand, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3904108980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) HasPose(name StringName, ) bool {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRPositionalTracker) GetPose(name StringName, ) XRPose {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4099720006) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRPose()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPositionalTracker) InvalidatePose(name StringName, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("invalidate_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) SetPose(name StringName, transform Transform3D, linear_velocity Vector3, angular_velocity Vector3, tracking_confidence XRPoseTrackingConfidence, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3451230163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), transform.AsCTypePtr(), linear_velocity.AsCTypePtr(), angular_velocity.AsCTypePtr(), gdc.ConstTypePtr(&tracking_confidence) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRPositionalTracker) GetInput(name StringName, ) Variant {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRPositionalTracker) SetInput(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRPositionalTrackerPoseChangedSignalFn func(pose XRPose, )

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

type XRPositionalTrackerPoseLostTrackingSignalFn func(pose XRPose, )

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

type XRPositionalTrackerButtonPressedSignalFn func(name String, )

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

type XRPositionalTrackerButtonReleasedSignalFn func(name String, )

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

type XRPositionalTrackerInputFloatChangedSignalFn func(name String, value float32, )

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

type XRPositionalTrackerInputVector2ChangedSignalFn func(name String, vector Vector2, )

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

type XRPositionalTrackerProfileChangedSignalFn func(role String, )

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
