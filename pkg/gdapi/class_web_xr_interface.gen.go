// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebXRInterface struct {
  obj gdc.ObjectPtr
}

func (me *WebXRInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebXRInterface) BaseClass() string {
  return "WebXRInterface"
}

type WebXRInterfaceTargetRayMode int
const (
  WebXRInterfaceTargetRayModeTargetRayModeUnknown WebXRInterfaceTargetRayMode = 0
  WebXRInterfaceTargetRayModeTargetRayModeGaze WebXRInterfaceTargetRayMode = 1
  WebXRInterfaceTargetRayModeTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
  WebXRInterfaceTargetRayModeTargetRayModeScreen WebXRInterfaceTargetRayMode = 3
)

func  (me *WebXRInterface) IsSessionSupported(session_mode String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) SetSessionMode(session_mode String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetSessionMode() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) SetRequiredFeatures(required_features String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetRequiredFeatures() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) SetOptionalFeatures(optional_features String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetOptionalFeatures() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetReferenceSpaceType() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) SetRequestedReferenceSpaceTypes(requested_reference_space_types String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetRequestedReferenceSpaceTypes() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) IsInputSourceActive(input_source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetInputSourceTracker(input_source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetInputSourceTargetRayMode(input_source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetVisibilityState() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetDisplayRefreshRate() { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) SetDisplayRefreshRate(refresh_rate float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebXRInterface) GetAvailableDisplayRefreshRates() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
