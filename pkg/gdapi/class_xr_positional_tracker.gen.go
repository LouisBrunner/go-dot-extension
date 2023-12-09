// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *XRPositionalTracker) GetTrackerType()  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetTrackerType(type_ XRServerTrackerType, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetTrackerName()  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetTrackerName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetTrackerDesc()  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetTrackerDesc(description String, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetTrackerProfile()  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetTrackerProfile(profile String, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetTrackerHand()  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetTrackerHand(hand XRPositionalTrackerTrackerHand, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) HasPose(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetPose(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) InvalidatePose(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetPose(name StringName, transform Transform3D, linear_velocity Vector3, angular_velocity Vector3, tracking_confidence XRPoseTrackingConfidence, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) GetInput(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPositionalTracker) SetInput(name StringName, value Variant, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
