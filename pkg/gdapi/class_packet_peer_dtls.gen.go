// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerDTLS struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeerDTLS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PacketPeerDTLS) BaseClass() string {
  return "PacketPeerDTLS"
}



// Enums

type PacketPeerDTLSStatus int
const (
  PacketPeerDTLSStatusStatusDisconnected PacketPeerDTLSStatus = 0
  PacketPeerDTLSStatusStatusHandshaking PacketPeerDTLSStatus = 1
  PacketPeerDTLSStatusStatusConnected PacketPeerDTLSStatus = 2
  PacketPeerDTLSStatusStatusError PacketPeerDTLSStatus = 3
  PacketPeerDTLSStatusStatusErrorHostnameMismatch PacketPeerDTLSStatus = 4
)

func (me *PacketPeerDTLS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeerDTLS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PacketPeerDTLS) Poll()  {
  panic("TODO: implement")
}

func  (me *PacketPeerDTLS) ConnectToPeer(packet_peer PacketPeerUDP, hostname String, client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerDTLS) GetStatus()  {
  panic("TODO: implement")
}

func  (me *PacketPeerDTLS) DisconnectFromPeer()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
