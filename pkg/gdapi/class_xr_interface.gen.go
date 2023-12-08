// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRInterface struct {
  obj gdc.ObjectPtr
}

func (me *XRInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRInterface) BaseClass() string {
  return "XRInterface"
}

type XRInterfaceCapabilities int
const (
  XRInterfaceXrNone XRInterfaceCapabilities = 0
  XRInterfaceXrMono XRInterfaceCapabilities = 1
  XRInterfaceXrStereo XRInterfaceCapabilities = 2
  XRInterfaceXrQuad XRInterfaceCapabilities = 4
  XRInterfaceXrVr XRInterfaceCapabilities = 8
  XRInterfaceXrAr XRInterfaceCapabilities = 16
  XRInterfaceXrExternal XRInterfaceCapabilities = 32
)

type XRInterfaceTrackingStatus int
const (
  XRInterfaceXrNormalTracking XRInterfaceTrackingStatus = 0
  XRInterfaceXrExcessiveMotion XRInterfaceTrackingStatus = 1
  XRInterfaceXrInsufficientFeatures XRInterfaceTrackingStatus = 2
  XRInterfaceXrUnknownTracking XRInterfaceTrackingStatus = 3
  XRInterfaceXrNotTracking XRInterfaceTrackingStatus = 4
)

type XRInterfacePlayAreaMode int
const (
  XRInterfaceXrPlayAreaUnknown XRInterfacePlayAreaMode = 0
  XRInterfaceXrPlayArea3Dof XRInterfacePlayAreaMode = 1
  XRInterfaceXrPlayAreaSitting XRInterfacePlayAreaMode = 2
  XRInterfaceXrPlayAreaRoomscale XRInterfacePlayAreaMode = 3
  XRInterfaceXrPlayAreaStage XRInterfacePlayAreaMode = 4
)

type XRInterfaceEnvironmentBlendMode int
const (
  XRInterfaceXrEnvBlendModeOpaque XRInterfaceEnvironmentBlendMode = 0
  XRInterfaceXrEnvBlendModeAdditive XRInterfaceEnvironmentBlendMode = 1
  XRInterfaceXrEnvBlendModeAlphaBlend XRInterfaceEnvironmentBlendMode = 2
)
