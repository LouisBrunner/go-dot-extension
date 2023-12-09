// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerPeerExtension struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerPeerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerPeerExtension) BaseClass() string {
  return "MultiplayerPeerExtension"
}



// Enums

func (me *MultiplayerPeerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerPeerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerPeerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiplayerPeerExtension) XGetPacket(r_buffer **uint8, r_buffer_size *int32, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XPutPacket(p_buffer *uint8, p_buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetAvailablePacketCount()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetMaxPacketSize()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetPacketScript()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XPutPacketScript(p_buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetPacketChannel()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetPacketMode()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XSetTransferChannel(p_channel int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetTransferChannel()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XSetTransferMode(p_mode MultiplayerPeerTransferMode, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetTransferMode()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XSetTargetPeer(p_peer int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetPacketPeer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XIsServer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XPoll()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XClose()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XDisconnectPeer(p_peer int, p_force bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetUniqueId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XSetRefuseNewConnections(p_enable bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XIsRefusingNewConnections()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XIsServerRelaySupported()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeerExtension) XGetConnectionStatus()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
