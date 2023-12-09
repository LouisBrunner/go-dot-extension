// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerPeer) BaseClass() string {
  return "MultiplayerPeer"
}



// Constants

var (
  MultiplayerPeerTargetPeerBroadcast = "0" // TODO: construct correctly
  MultiplayerPeerTargetPeerServer = "1" // TODO: construct correctly
)

// Enums

type MultiplayerPeerConnectionStatus int
const (
  MultiplayerPeerConnectionStatusConnectionDisconnected MultiplayerPeerConnectionStatus = 0
  MultiplayerPeerConnectionStatusConnectionConnecting MultiplayerPeerConnectionStatus = 1
  MultiplayerPeerConnectionStatusConnectionConnected MultiplayerPeerConnectionStatus = 2
)

type MultiplayerPeerTransferMode int
const (
  MultiplayerPeerTransferModeTransferModeUnreliable MultiplayerPeerTransferMode = 0
  MultiplayerPeerTransferModeTransferModeUnreliableOrdered MultiplayerPeerTransferMode = 1
  MultiplayerPeerTransferModeTransferModeReliable MultiplayerPeerTransferMode = 2
)

func (me *MultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiplayerPeer) SetTransferChannel(channel int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetTransferChannel()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) SetTransferMode(mode MultiplayerPeerTransferMode, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetTransferMode()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) SetTargetPeer(id int, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetPacketPeer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetPacketChannel()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetPacketMode()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) Poll()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) Close()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) DisconnectPeer(peer int, force bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetConnectionStatus()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GetUniqueId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) GenerateUniqueId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) SetRefuseNewConnections(enable bool, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) IsRefusingNewConnections()  {
  panic("TODO: implement")
}

func  (me *MultiplayerPeer) IsServerRelaySupported()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
