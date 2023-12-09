// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebSocketPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebSocketPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebSocketPeer) BaseClass() string {
  return "WebSocketPeer"
}



// Enums

type WebSocketPeerWriteMode int
const (
  WebSocketPeerWriteModeWriteModeText WebSocketPeerWriteMode = 0
  WebSocketPeerWriteModeWriteModeBinary WebSocketPeerWriteMode = 1
)

type WebSocketPeerState int
const (
  WebSocketPeerStateStateConnecting WebSocketPeerState = 0
  WebSocketPeerStateStateOpen WebSocketPeerState = 1
  WebSocketPeerStateStateClosing WebSocketPeerState = 2
  WebSocketPeerStateStateClosed WebSocketPeerState = 3
)

func (me *WebSocketPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebSocketPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebSocketPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebSocketPeer) ConnectToUrl(url String, tls_client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) AcceptStream(stream StreamPeer, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) Send(message PackedByteArray, write_mode WebSocketPeerWriteMode, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SendText(message String, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) WasStringPacket()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) Poll()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) Close(code int, reason String, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetConnectedHost()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetConnectedPort()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetSelectedProtocol()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetRequestedUrl()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetNoDelay(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetCurrentOutboundBufferedAmount()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetReadyState()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetCloseCode()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetCloseReason()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetSupportedProtocols()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetHandshakeHeaders()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetInboundBufferSize()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetInboundBufferSize(buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetOutboundBufferSize()  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetOutboundBufferSize(buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) SetMaxQueuedPackets(buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketPeer) GetMaxQueuedPackets()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
