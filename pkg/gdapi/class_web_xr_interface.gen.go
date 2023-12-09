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



// Enums

type WebXRInterfaceTargetRayMode int
const (
  WebXRInterfaceTargetRayModeTargetRayModeUnknown WebXRInterfaceTargetRayMode = 0
  WebXRInterfaceTargetRayModeTargetRayModeGaze WebXRInterfaceTargetRayMode = 1
  WebXRInterfaceTargetRayModeTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
  WebXRInterfaceTargetRayModeTargetRayModeScreen WebXRInterfaceTargetRayMode = 3
)

func (me *WebXRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebXRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebXRInterface) IsSessionSupported(session_mode String, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) SetSessionMode(session_mode String, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetSessionMode()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) SetRequiredFeatures(required_features String, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetRequiredFeatures()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) SetOptionalFeatures(optional_features String, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetOptionalFeatures()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetReferenceSpaceType()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) SetRequestedReferenceSpaceTypes(requested_reference_space_types String, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetRequestedReferenceSpaceTypes()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) IsInputSourceActive(input_source_id int, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetInputSourceTracker(input_source_id int, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetInputSourceTargetRayMode(input_source_id int, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetVisibilityState()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetDisplayRefreshRate()  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) SetDisplayRefreshRate(refresh_rate float32, )  {
  panic("TODO: implement")
}

func  (me *WebXRInterface) GetAvailableDisplayRefreshRates()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
