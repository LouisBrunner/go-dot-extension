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

type XRPositionalTrackerTrackerHand int
const (
  XRPositionalTrackerTrackerHandTrackerHandUnknown XRPositionalTrackerTrackerHand = 0
  XRPositionalTrackerTrackerHandTrackerHandLeft XRPositionalTrackerTrackerHand = 1
  XRPositionalTrackerTrackerHandTrackerHandRight XRPositionalTrackerTrackerHand = 2
)

func  (me *XRPositionalTracker) GetTrackerType() { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetTrackerType(type_ XRServerTrackerType, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetTrackerName() { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetTrackerName(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetTrackerDesc() { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetTrackerDesc(description String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetTrackerProfile() { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetTrackerProfile(profile String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetTrackerHand() { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetTrackerHand(hand XRPositionalTrackerTrackerHand, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) HasPose(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetPose(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) InvalidatePose(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetPose(name StringName, transform Transform3D, linear_velocity Vector3, angular_velocity Vector3, tracking_confidence XRPoseTrackingConfidence, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) GetInput(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPositionalTracker) SetInput(name StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
