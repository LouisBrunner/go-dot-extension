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



// Enums

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

func (me *WebRTCDataChannel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCDataChannel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebRTCDataChannel) Poll()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) Close()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) WasStringPacket()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) SetWriteMode(write_mode WebRTCDataChannelWriteMode, )  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetWriteMode()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetReadyState()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetLabel()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) IsOrdered()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetId()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetMaxPacketLifeTime()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetMaxRetransmits()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetProtocol()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) IsNegotiated()  {
  panic("TODO: implement")
}

func  (me *WebRTCDataChannel) GetBufferedAmount()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
