// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRPositionalTracker struct {
  obj gdc.ObjectPtr
}

func (me *XRPositionalTracker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRPositionalTracker) BaseClass() string {
  return "XRPositionalTracker"
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
  var ret XRServerTrackerType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerType(type_ XRServerTrackerType, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3055763575) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) GetTrackerName() StringName {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerName(name StringName, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) GetTrackerDesc() String {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_desc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerDesc(description String, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_desc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) GetTrackerProfile() String {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerProfile(profile String, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profile.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) GetTrackerHand() XRPositionalTrackerTrackerHand {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4181770860) // FIXME: should cache?
  var ret XRPositionalTrackerTrackerHand
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetTrackerHand(hand XRPositionalTrackerTrackerHand, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3904108980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hand), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) HasPose(name StringName, ) bool {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) GetPose(name StringName, ) XRPose {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4099720006) // FIXME: should cache?
  var ret XRPose
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) InvalidatePose(name StringName, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("invalidate_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) SetPose(name StringName, transform Transform3D, linear_velocity Vector3, angular_velocity Vector3, tracking_confidence XRPoseTrackingConfidence, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3451230163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(linear_velocity.AsCTypePtr()), gdc.ConstTypePtr(angular_velocity.AsCTypePtr()), gdc.ConstTypePtr(&tracking_confidence), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRPositionalTracker) GetInput(name StringName, ) Variant {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRPositionalTracker) SetInput(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("XRPositionalTracker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRPositionalTrackerPoseChangedSignalFn func(pose XRPose, )

func (me *XRPositionalTracker) ConnectPoseChanged(subs SignalSubscribers, fn XRPositionalTrackerPoseChangedSignalFn) {
  sig := StringNameFromStr("pose_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectPoseChanged(subs SignalSubscribers, fn XRPositionalTrackerPoseChangedSignalFn) {
  sig := StringNameFromStr("pose_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type XRPositionalTrackerButtonPressedSignalFn func(name String, )

func (me *XRPositionalTracker) ConnectButtonPressed(subs SignalSubscribers, fn XRPositionalTrackerButtonPressedSignalFn) {
  sig := StringNameFromStr("button_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectButtonPressed(subs SignalSubscribers, fn XRPositionalTrackerButtonPressedSignalFn) {
  sig := StringNameFromStr("button_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type XRPositionalTrackerButtonReleasedSignalFn func(name String, )

func (me *XRPositionalTracker) ConnectButtonReleased(subs SignalSubscribers, fn XRPositionalTrackerButtonReleasedSignalFn) {
  sig := StringNameFromStr("button_released")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectButtonReleased(subs SignalSubscribers, fn XRPositionalTrackerButtonReleasedSignalFn) {
  sig := StringNameFromStr("button_released")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type XRPositionalTrackerInputFloatChangedSignalFn func(name String, value float32, )

func (me *XRPositionalTracker) ConnectInputFloatChanged(subs SignalSubscribers, fn XRPositionalTrackerInputFloatChangedSignalFn) {
  sig := StringNameFromStr("input_float_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectInputFloatChanged(subs SignalSubscribers, fn XRPositionalTrackerInputFloatChangedSignalFn) {
  sig := StringNameFromStr("input_float_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type XRPositionalTrackerInputVector2ChangedSignalFn func(name String, vector Vector2, )

func (me *XRPositionalTracker) ConnectInputVector2Changed(subs SignalSubscribers, fn XRPositionalTrackerInputVector2ChangedSignalFn) {
  sig := StringNameFromStr("input_vector2_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectInputVector2Changed(subs SignalSubscribers, fn XRPositionalTrackerInputVector2ChangedSignalFn) {
  sig := StringNameFromStr("input_vector2_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type XRPositionalTrackerProfileChangedSignalFn func(role String, )

func (me *XRPositionalTracker) ConnectProfileChanged(subs SignalSubscribers, fn XRPositionalTrackerProfileChangedSignalFn) {
  sig := StringNameFromStr("profile_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRPositionalTracker) DisconnectProfileChanged(subs SignalSubscribers, fn XRPositionalTrackerProfileChangedSignalFn) {
  sig := StringNameFromStr("profile_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
