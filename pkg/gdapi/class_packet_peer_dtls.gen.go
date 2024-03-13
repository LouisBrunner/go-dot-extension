// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *PacketPeerDTLS) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PacketPeerDTLS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeerDTLS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PacketPeerDTLS) Poll()  {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PacketPeerDTLS) ConnectToPeer(packet_peer PacketPeerUDP, hostname String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2880188099) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(packet_peer.AsCTypePtr()), gdc.ConstTypePtr(hostname.AsCTypePtr()), gdc.ConstTypePtr(client_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerDTLS) GetStatus() PacketPeerDTLSStatus {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3248654679) // FIXME: should cache?
  var ret PacketPeerDTLSStatus
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerDTLS) DisconnectFromPeer()  {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_from_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
