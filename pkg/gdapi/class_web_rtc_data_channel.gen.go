// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCDataChannel struct {
  obj gdc.ObjectPtr
}

func (me *WebRTCDataChannel) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebRTCDataChannel) BaseClass() string {
  return "WebRTCDataChannel"
}

type WebRTCDataChannelWriteMode int
const (
  WebRTCDataChannelWriteModeText WebRTCDataChannelWriteMode = 0
  WebRTCDataChannelWriteModeBinary WebRTCDataChannelWriteMode = 1
)

type WebRTCDataChannelChannelState int
const (
  WebRTCDataChannelStateConnecting WebRTCDataChannelChannelState = 0
  WebRTCDataChannelStateOpen WebRTCDataChannelChannelState = 1
  WebRTCDataChannelStateClosing WebRTCDataChannelChannelState = 2
  WebRTCDataChannelStateClosed WebRTCDataChannelChannelState = 3
)
