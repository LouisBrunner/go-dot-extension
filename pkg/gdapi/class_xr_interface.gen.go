// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  XRInterfaceCapabilitiesXrNone XRInterfaceCapabilities = 0
  XRInterfaceCapabilitiesXrMono XRInterfaceCapabilities = 1
  XRInterfaceCapabilitiesXrStereo XRInterfaceCapabilities = 2
  XRInterfaceCapabilitiesXrQuad XRInterfaceCapabilities = 4
  XRInterfaceCapabilitiesXrVr XRInterfaceCapabilities = 8
  XRInterfaceCapabilitiesXrAr XRInterfaceCapabilities = 16
  XRInterfaceCapabilitiesXrExternal XRInterfaceCapabilities = 32
)

type XRInterfaceTrackingStatus int
const (
  XRInterfaceTrackingStatusXrNormalTracking XRInterfaceTrackingStatus = 0
  XRInterfaceTrackingStatusXrExcessiveMotion XRInterfaceTrackingStatus = 1
  XRInterfaceTrackingStatusXrInsufficientFeatures XRInterfaceTrackingStatus = 2
  XRInterfaceTrackingStatusXrUnknownTracking XRInterfaceTrackingStatus = 3
  XRInterfaceTrackingStatusXrNotTracking XRInterfaceTrackingStatus = 4
)

type XRInterfacePlayAreaMode int
const (
  XRInterfacePlayAreaModeXrPlayAreaUnknown XRInterfacePlayAreaMode = 0
  XRInterfacePlayAreaModeXrPlayArea3Dof XRInterfacePlayAreaMode = 1
  XRInterfacePlayAreaModeXrPlayAreaSitting XRInterfacePlayAreaMode = 2
  XRInterfacePlayAreaModeXrPlayAreaRoomscale XRInterfacePlayAreaMode = 3
  XRInterfacePlayAreaModeXrPlayAreaStage XRInterfacePlayAreaMode = 4
)

type XRInterfaceEnvironmentBlendMode int
const (
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeOpaque XRInterfaceEnvironmentBlendMode = 0
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeAdditive XRInterfaceEnvironmentBlendMode = 1
  XRInterfaceEnvironmentBlendModeXrEnvBlendModeAlphaBlend XRInterfaceEnvironmentBlendMode = 2
)

func  (me *XRInterface) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetCapabilities() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) IsPrimary() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) SetPrimary(primary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) IsInitialized() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) Initialize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) Uninitialize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetSystemInfo() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetTrackingStatus() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetRenderTargetSize() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetViewCount() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) TriggerHapticPulse(action_name String, tracker_name StringName, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) SupportsPlayAreaMode(mode XRInterfacePlayAreaMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetPlayAreaMode() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) SetPlayAreaMode(mode XRInterfacePlayAreaMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetPlayArea() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetAnchorDetectionIsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) SetAnchorDetectionIsEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetCameraFeedId() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) IsPassthroughSupported() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) IsPassthroughEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) StartPassthrough() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) StopPassthrough() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetTransformForView(view int, cam_transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetProjectionForView(view int, aspect float32, near float32, far float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) GetSupportedEnvironmentBlendModes() { // TODO: return value
  // TODO: implement
}

func  (me *XRInterface) SetEnvironmentBlendMode(mode XRInterfaceEnvironmentBlendMode, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
