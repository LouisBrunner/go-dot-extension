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



// Enums

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

func (me *XRServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRServer) GetWorldScale()  {
  panic("TODO: implement")
}

func  (me *XRServer) SetWorldScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetWorldOrigin()  {
  panic("TODO: implement")
}

func  (me *XRServer) SetWorldOrigin(world_origin Transform3D, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetReferenceFrame()  {
  panic("TODO: implement")
}

func  (me *XRServer) CenterOnHmd(rotation_mode XRServerRotationMode, keep_height bool, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetHmdTransform()  {
  panic("TODO: implement")
}

func  (me *XRServer) AddInterface(interface_ XRInterface, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetInterfaceCount()  {
  panic("TODO: implement")
}

func  (me *XRServer) RemoveInterface(interface_ XRInterface, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetInterface(idx int, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetInterfaces()  {
  panic("TODO: implement")
}

func  (me *XRServer) FindInterface(name String, )  {
  panic("TODO: implement")
}

func  (me *XRServer) AddTracker(tracker XRPositionalTracker, )  {
  panic("TODO: implement")
}

func  (me *XRServer) RemoveTracker(tracker XRPositionalTracker, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetTrackers(tracker_types int, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetTracker(tracker_name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRServer) GetPrimaryInterface()  {
  panic("TODO: implement")
}

func  (me *XRServer) SetPrimaryInterface(interface_ XRInterface, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
