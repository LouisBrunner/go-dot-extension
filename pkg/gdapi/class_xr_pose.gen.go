// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRPose struct {
  obj gdc.ObjectPtr
}

func (me *XRPose) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRPose) BaseClass() string {
  return "XRPose"
}

type XRPoseTrackingConfidence int
const (
  XRPoseTrackingConfidenceXrTrackingConfidenceNone XRPoseTrackingConfidence = 0
  XRPoseTrackingConfidenceXrTrackingConfidenceLow XRPoseTrackingConfidence = 1
  XRPoseTrackingConfidenceXrTrackingConfidenceHigh XRPoseTrackingConfidence = 2
)

func  (me *XRPose) SetHasTrackingData(has_tracking_data bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetHasTrackingData() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) SetName(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) SetTransform(transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetTransform() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetAdjustedTransform() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) SetLinearVelocity(velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetLinearVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) SetAngularVelocity(velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetAngularVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) SetTrackingConfidence(tracking_confidence XRPoseTrackingConfidence, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRPose) GetTrackingConfidence() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
