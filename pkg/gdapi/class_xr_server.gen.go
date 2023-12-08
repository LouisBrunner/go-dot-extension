// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRServer struct {
  obj gdc.ObjectPtr
}

func (me *XRServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRServer) BaseClass() string {
  return "XRServer"
}

type XRServerTrackerType int
const (
  XRServerTrackerTypeTrackerHead XRServerTrackerType = 1
  XRServerTrackerTypeTrackerController XRServerTrackerType = 2
  XRServerTrackerTypeTrackerBasestation XRServerTrackerType = 4
  XRServerTrackerTypeTrackerAnchor XRServerTrackerType = 8
  XRServerTrackerTypeTrackerAnyKnown XRServerTrackerType = 127
  XRServerTrackerTypeTrackerUnknown XRServerTrackerType = 128
  XRServerTrackerTypeTrackerAny XRServerTrackerType = 255
)

type XRServerRotationMode int
const (
  XRServerRotationModeResetFullRotation XRServerRotationMode = 0
  XRServerRotationModeResetButKeepTilt XRServerRotationMode = 1
  XRServerRotationModeDontResetRotation XRServerRotationMode = 2
)

func  (me *XRServer) GetWorldScale() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) SetWorldScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetWorldOrigin() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) SetWorldOrigin(world_origin Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetReferenceFrame() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) CenterOnHmd(rotation_mode XRServerRotationMode, keep_height bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetHmdTransform() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) AddInterface(interface_ XRInterface, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetInterfaceCount() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) RemoveInterface(interface_ XRInterface, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetInterface(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetInterfaces() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) FindInterface(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) AddTracker(tracker XRPositionalTracker, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) RemoveTracker(tracker XRPositionalTracker, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetTrackers(tracker_types int, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetTracker(tracker_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) GetPrimaryInterface() { // TODO: return value
  // TODO: implement
}

func  (me *XRServer) SetPrimaryInterface(interface_ XRInterface, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
