// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCPeerConnectionExtension struct {
  obj gdc.ObjectPtr
}

func createWebRTCPeerConnectionExtension(obj gdc.ObjectPtr) *WebRTCPeerConnectionExtension {
  return &WebRTCPeerConnectionExtension{
    obj: obj,
  }
}

func (me *WebRTCPeerConnectionExtension) BaseClass() string {
  return "WebRTCPeerConnectionExtension"
}
