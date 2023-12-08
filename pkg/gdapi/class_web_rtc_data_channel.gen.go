// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  WebRTCDataChannelWriteModeWriteModeText WebRTCDataChannelWriteMode = 0
  WebRTCDataChannelWriteModeWriteModeBinary WebRTCDataChannelWriteMode = 1
)

type WebRTCDataChannelChannelState int
const (
  WebRTCDataChannelChannelStateStateConnecting WebRTCDataChannelChannelState = 0
  WebRTCDataChannelChannelStateStateOpen WebRTCDataChannelChannelState = 1
  WebRTCDataChannelChannelStateStateClosing WebRTCDataChannelChannelState = 2
  WebRTCDataChannelChannelStateStateClosed WebRTCDataChannelChannelState = 3
)

func  (me *WebRTCDataChannel) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) Close() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) WasStringPacket() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) SetWriteMode(write_mode WebRTCDataChannelWriteMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetWriteMode() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetReadyState() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetLabel() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) IsOrdered() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetId() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetMaxPacketLifeTime() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetMaxRetransmits() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetProtocol() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) IsNegotiated() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCDataChannel) GetBufferedAmount() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
