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
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRNode3D) GetTracker() StringName {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) SetPoseName(pose StringName, )  {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pose_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{pose.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRNode3D) GetPoseName() StringName {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) GetIsActive() bool {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRNode3D) GetHasTrackingData() bool {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_has_tracking_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *XRNode3D) GetPose() XRPose {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2806551826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewXRPose()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRNode3D) TriggerHapticPulse(action_name String, frequency float64, amplitude float64, duration_sec float64, delay_sec float64, )  {
  classNameV := StringNameFromStr("XRNode3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("trigger_haptic_pulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 508576839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action_name.AsCTypePtr(), gdc.ConstTypePtr(&frequency) , gdc.ConstTypePtr(&amplitude) , gdc.ConstTypePtr(&duration_sec) , gdc.ConstTypePtr(&delay_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

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
