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



// Enums

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

func (me *XRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRInterface) GetName()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetCapabilities()  {
  panic("TODO: implement")
}

func  (me *XRInterface) IsPrimary()  {
  panic("TODO: implement")
}

func  (me *XRInterface) SetPrimary(primary bool, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) IsInitialized()  {
  panic("TODO: implement")
}

func  (me *XRInterface) Initialize()  {
  panic("TODO: implement")
}

func  (me *XRInterface) Uninitialize()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetSystemInfo()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetTrackingStatus()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetRenderTargetSize()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetViewCount()  {
  panic("TODO: implement")
}

func  (me *XRInterface) TriggerHapticPulse(action_name String, tracker_name StringName, frequency float32, amplitude float32, duration_sec float32, delay_sec float32, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) SupportsPlayAreaMode(mode XRInterfacePlayAreaMode, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetPlayAreaMode()  {
  panic("TODO: implement")
}

func  (me *XRInterface) SetPlayAreaMode(mode XRInterfacePlayAreaMode, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetPlayArea()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetAnchorDetectionIsEnabled()  {
  panic("TODO: implement")
}

func  (me *XRInterface) SetAnchorDetectionIsEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetCameraFeedId()  {
  panic("TODO: implement")
}

func  (me *XRInterface) IsPassthroughSupported()  {
  panic("TODO: implement")
}

func  (me *XRInterface) IsPassthroughEnabled()  {
  panic("TODO: implement")
}

func  (me *XRInterface) StartPassthrough()  {
  panic("TODO: implement")
}

func  (me *XRInterface) StopPassthrough()  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetTransformForView(view int, cam_transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetProjectionForView(view int, aspect float32, near float32, far float32, )  {
  panic("TODO: implement")
}

func  (me *XRInterface) GetSupportedEnvironmentBlendModes()  {
  panic("TODO: implement")
}

func  (me *XRInterface) SetEnvironmentBlendMode(mode XRInterfaceEnvironmentBlendMode, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
