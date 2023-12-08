// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  WebXRInterfaceTargetRayModeUnknown WebXRInterfaceTargetRayMode = 0
  WebXRInterfaceTargetRayModeGaze WebXRInterfaceTargetRayMode = 1
  WebXRInterfaceTargetRayModeTrackedPointer WebXRInterfaceTargetRayMode = 2
  WebXRInterfaceTargetRayModeScreen WebXRInterfaceTargetRayMode = 3
)
