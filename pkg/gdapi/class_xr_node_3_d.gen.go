// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRNode3D struct {
  Node3D
}

func (me *XRNode3D) BaseClass() string {
  return "XRNode3D"
}



// Enums

func (me *XRNode3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRNode3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRNode3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRNode3D) SetTracker(tracker_name StringName, )  {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tracker_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRNode3D) GetTracker() StringName {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRNode3D) SetPoseName(pose StringName, )  {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pose_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pose.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XRNode3D) GetPoseName() StringName {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRNode3D) GetIsActive() bool {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRNode3D) GetHasTrackingData() bool {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_has_tracking_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRNode3D) GetPose() XRPose {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2806551826) // FIXME: should cache?
  var ret XRPose
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRNode3D) TriggerHapticPulse(action_name String, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, )  {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("trigger_haptic_pulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 508576839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action_name.AsCTypePtr()), gdc.ConstTypePtr(&frequency), gdc.ConstTypePtr(&amplitude), gdc.ConstTypePtr(&duration_sec), gdc.ConstTypePtr(&delay_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRNode3DTrackingChangedSignalFn func(tracking bool, )

func (me *XRNode3D) ConnectTrackingChanged(subs SignalSubscribers, fn XRNode3DTrackingChangedSignalFn) {
  sig := StringNameFromStr("tracking_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *XRNode3D) DisconnectTrackingChanged(subs SignalSubscribers, fn XRNode3DTrackingChangedSignalFn) {
  sig := StringNameFromStr("tracking_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
