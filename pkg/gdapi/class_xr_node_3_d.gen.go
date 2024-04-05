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

type ptrsForXRNode3DList struct {
  fnSetTracker gdc.MethodBindPtr
  fnGetTracker gdc.MethodBindPtr
  fnSetPoseName gdc.MethodBindPtr
  fnGetPoseName gdc.MethodBindPtr
  fnGetIsActive gdc.MethodBindPtr
  fnGetHasTrackingData gdc.MethodBindPtr
  fnGetPose gdc.MethodBindPtr
  fnTriggerHapticPulse gdc.MethodBindPtr
}

var ptrsForXRNode3D ptrsForXRNode3DList

func initXRNode3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("XRNode3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_tracker")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnSetTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_tracker")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnGetTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_pose_name")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnSetPoseName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_pose_name")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnGetPoseName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("get_is_active")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnGetIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_has_tracking_data")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnGetHasTrackingData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_pose")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnGetPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2806551826))
  }
  {
    methodName := StringNameFromStr("trigger_haptic_pulse")
    defer methodName.Destroy()
    ptrsForXRNode3D.fnTriggerHapticPulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 508576839))
  }
}

type XRNode3D struct {
  Node3D
}

func (me *XRNode3D) BaseClass() string {
  return "XRNode3D"
}

func NewXRNode3D() *XRNode3D {
  str := StringNameFromStr("XRNode3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRNode3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnSetTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRNode3D) GetTracker() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnGetTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) SetPoseName(pose StringName, )  {
  cargs := []gdc.ConstTypePtr{pose.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnSetPoseName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRNode3D) GetPoseName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnGetPoseName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) GetIsActive() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnGetIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRNode3D) GetHasTrackingData() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnGetHasTrackingData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRNode3D) GetPose() XRPose {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRPose()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnGetPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) TriggerHapticPulse(action_name String, frequency float64, amplitude float64, duration_sec float64, delay_sec float64, )  {
  cargs := []gdc.ConstTypePtr{action_name.AsCTypePtr(), gdc.ConstTypePtr(&frequency) , gdc.ConstTypePtr(&amplitude) , gdc.ConstTypePtr(&duration_sec) , gdc.ConstTypePtr(&delay_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRNode3D.fnTriggerHapticPulse), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRNode3DTrackingChangedSignalFn func(tracking bool, )

func (me *XRNode3D) ConnectTrackingChanged(subs SignalSubscribers, fn XRNode3DTrackingChangedSignalFn) {
  sig := StringNameFromStr("tracking_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRNode3D) DisconnectTrackingChanged(subs SignalSubscribers, fn XRNode3DTrackingChangedSignalFn) {
  sig := StringNameFromStr("tracking_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
