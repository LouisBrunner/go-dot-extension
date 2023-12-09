// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebSocketMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebSocketMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebSocketMultiplayerPeer) BaseClass() string {
  return "WebSocketMultiplayerPeer"
}



// Enums

func (me *WebSocketMultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebSocketMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebSocketMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebSocketMultiplayerPeer) CreateClient(url String, tls_client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) CreateServer(port int, bind_address String, tls_server_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetPeer(peer_id int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetPeerAddress(id int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetPeerPort(id int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetSupportedProtocols()  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetHandshakeHeaders()  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetInboundBufferSize()  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetInboundBufferSize(buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetOutboundBufferSize()  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetOutboundBufferSize(buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetHandshakeTimeout()  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeTimeout(timeout float32, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) SetMaxQueuedPackets(max_queued_packets int, )  {
  panic("TODO: implement")
}

func  (me *WebSocketMultiplayerPeer) GetMaxQueuedPackets()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
