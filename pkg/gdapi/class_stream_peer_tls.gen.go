// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTLS struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerTLS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerTLS) BaseClass() string {
  return "StreamPeerTLS"
}



// Enums

type StreamPeerTLSStatus int
const (
  StreamPeerTLSStatusStatusDisconnected StreamPeerTLSStatus = 0
  StreamPeerTLSStatusStatusHandshaking StreamPeerTLSStatus = 1
  StreamPeerTLSStatusStatusConnected StreamPeerTLSStatus = 2
  StreamPeerTLSStatusStatusError StreamPeerTLSStatus = 3
  StreamPeerTLSStatusStatusErrorHostnameMismatch StreamPeerTLSStatus = 4
)

func (me *StreamPeerTLS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerTLS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StreamPeerTLS) Poll()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTLS) AcceptStream(stream StreamPeer, server_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerTLS) ConnectToStream(stream StreamPeer, common_name String, client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerTLS) GetStatus()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTLS) GetStream()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTLS) DisconnectFromStream()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
