// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCDataChannelExtension struct {
  obj gdc.ObjectPtr
}

func createWebRTCDataChannelExtension(obj gdc.ObjectPtr) *WebRTCDataChannelExtension {
  return &WebRTCDataChannelExtension{
    obj: obj,
  }
}

func (me *WebRTCDataChannelExtension) BaseClass() string {
  return "WebRTCDataChannelExtension"
}
