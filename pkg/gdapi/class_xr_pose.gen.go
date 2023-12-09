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



// Enums

type XRPoseTrackingConfidence int
const (
  XRPoseTrackingConfidenceXrTrackingConfidenceNone XRPoseTrackingConfidence = 0
  XRPoseTrackingConfidenceXrTrackingConfidenceLow XRPoseTrackingConfidence = 1
  XRPoseTrackingConfidenceXrTrackingConfidenceHigh XRPoseTrackingConfidence = 2
)

func (me *XRPose) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRPose) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRPose) SetHasTrackingData(has_tracking_data bool, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetHasTrackingData()  {
  panic("TODO: implement")
}

func  (me *XRPose) SetName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetName()  {
  panic("TODO: implement")
}

func  (me *XRPose) SetTransform(transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetTransform()  {
  panic("TODO: implement")
}

func  (me *XRPose) GetAdjustedTransform()  {
  panic("TODO: implement")
}

func  (me *XRPose) SetLinearVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *XRPose) SetAngularVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *XRPose) SetTrackingConfidence(tracking_confidence XRPoseTrackingConfidence, )  {
  panic("TODO: implement")
}

func  (me *XRPose) GetTrackingConfidence()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
