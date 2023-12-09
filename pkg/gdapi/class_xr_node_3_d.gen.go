// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRNode3D struct {
  obj gdc.ObjectPtr
}

func (me *XRNode3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRNode3D) BaseClass() string {
  return "XRNode3D"
}



// Enums

func (me *XRNode3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRNode3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRNode3D) SetTracker(tracker_name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRNode3D) GetTracker()  {
  panic("TODO: implement")
}

func  (me *XRNode3D) SetPoseName(pose StringName, )  {
  panic("TODO: implement")
}

func  (me *XRNode3D) GetPoseName()  {
  panic("TODO: implement")
}

func  (me *XRNode3D) GetIsActive()  {
  panic("TODO: implement")
}

func  (me *XRNode3D) GetHasTrackingData()  {
  panic("TODO: implement")
}

func  (me *XRNode3D) GetPose()  {
  panic("TODO: implement")
}

func  (me *XRNode3D) TriggerHapticPulse(action_name String, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
