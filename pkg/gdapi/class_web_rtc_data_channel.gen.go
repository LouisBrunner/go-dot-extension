// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCDataChannel struct {
  obj gdc.ObjectPtr
}

func createWebRTCDataChannel(obj gdc.ObjectPtr) *WebRTCDataChannel {
  return &WebRTCDataChannel{
    obj: obj,
  }
}

func (me *WebRTCDataChannel) BaseClass() string {
  return "WebRTCDataChannel"
}
