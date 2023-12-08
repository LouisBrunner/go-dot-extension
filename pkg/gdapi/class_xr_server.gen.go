// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  XRServerTrackerHead XRServerTrackerType = 1
  XRServerTrackerController XRServerTrackerType = 2
  XRServerTrackerBasestation XRServerTrackerType = 4
  XRServerTrackerAnchor XRServerTrackerType = 8
  XRServerTrackerAnyKnown XRServerTrackerType = 127
  XRServerTrackerUnknown XRServerTrackerType = 128
  XRServerTrackerAny XRServerTrackerType = 255
)

type XRServerRotationMode int
const (
  XRServerResetFullRotation XRServerRotationMode = 0
  XRServerResetButKeepTilt XRServerRotationMode = 1
  XRServerDontResetRotation XRServerRotationMode = 2
)
