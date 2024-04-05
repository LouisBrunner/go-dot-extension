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

type ptrsForPacketPeerDTLSList struct {
  fnPoll gdc.MethodBindPtr
  fnConnectToPeer gdc.MethodBindPtr
  fnGetStatus gdc.MethodBindPtr
  fnDisconnectFromPeer gdc.MethodBindPtr
}

var ptrsForPacketPeerDTLS ptrsForPacketPeerDTLSList

func initPacketPeerDTLSPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PacketPeerDTLS")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("poll")
    defer methodName.Destroy()
    ptrsForPacketPeerDTLS.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("connect_to_peer")
    defer methodName.Destroy()
    ptrsForPacketPeerDTLS.fnConnectToPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2880188099))
  }
  {
    methodName := StringNameFromStr("get_status")
    defer methodName.Destroy()
    ptrsForPacketPeerDTLS.fnGetStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3248654679))
  }
  {
    methodName := StringNameFromStr("disconnect_from_peer")
    defer methodName.Destroy()
    ptrsForPacketPeerDTLS.fnDisconnectFromPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerDTLS.fnPoll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerDTLS) ConnectToPeer(packet_peer PacketPeerUDP, hostname String, client_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{packet_peer.AsCTypePtr(), hostname.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerDTLS.fnConnectToPeer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerDTLS) GetStatus() PacketPeerDTLSStatus {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PacketPeerDTLSStatus

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerDTLS.fnGetStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerDTLS) DisconnectFromPeer()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerDTLS.fnDisconnectFromPeer), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
