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

func  (me *WebSocketPeer) ConnectToUrl(url String, tls_client_options TLSOptions, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) AcceptStream(stream StreamPeer, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) Send(message PackedByteArray, write_mode WebSocketPeerWriteMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SendText(message String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) WasStringPacket() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) Close(code int, reason String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetConnectedHost() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetConnectedPort() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetSelectedProtocol() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetRequestedUrl() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetNoDelay(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetCurrentOutboundBufferedAmount() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetReadyState() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetCloseCode() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetCloseReason() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetSupportedProtocols() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetSupportedProtocols(protocols PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetHandshakeHeaders() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetHandshakeHeaders(protocols PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetInboundBufferSize() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetInboundBufferSize(buffer_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetOutboundBufferSize() { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetOutboundBufferSize(buffer_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) SetMaxQueuedPackets(buffer_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebSocketPeer) GetMaxQueuedPackets() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
