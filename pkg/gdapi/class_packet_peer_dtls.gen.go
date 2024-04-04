// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type PacketPeerDTLS struct {
  PacketPeer
}

func (me *PacketPeerDTLS) BaseClass() string {
  return "PacketPeerDTLS"
}

func NewPacketPeerDTLS() *PacketPeerDTLS {
  str := StringNameFromStr("PacketPeerDTLS") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PacketPeerDTLS{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerDTLS) ConnectToPeer(packet_peer PacketPeerUDP, hostname String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2880188099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{packet_peer.AsCTypePtr(), hostname.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerDTLS) GetStatus() PacketPeerDTLSStatus {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3248654679) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PacketPeerDTLSStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerDTLS) DisconnectFromPeer()  {
  classNameV := StringNameFromStr("PacketPeerDTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_from_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
