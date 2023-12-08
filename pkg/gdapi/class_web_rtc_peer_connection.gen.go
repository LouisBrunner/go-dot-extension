// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCPeerConnection struct {
  obj gdc.ObjectPtr
}

func createWebRTCPeerConnection(obj gdc.ObjectPtr) *WebRTCPeerConnection {
  return &WebRTCPeerConnection{
    obj: obj,
  }
}

func (me *WebRTCPeerConnection) BaseClass() string {
  return "WebRTCPeerConnection"
}
