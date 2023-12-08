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

// TODO: needed?
// const (
// // )

var (
  MultiplayerPeerTargetPeerBroadcast = "0" // TODO: construct correctly
  MultiplayerPeerTargetPeerServer = "1" // TODO: construct correctly
)

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

func  (me *MultiplayerPeer) SetTransferChannel(channel int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetTransferChannel() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) SetTransferMode(mode MultiplayerPeerTransferMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetTransferMode() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) SetTargetPeer(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetPacketPeer() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetPacketChannel() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetPacketMode() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) Close() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) DisconnectPeer(peer int, force bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetConnectionStatus() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GetUniqueId() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) GenerateUniqueId() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) SetRefuseNewConnections(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) IsRefusingNewConnections() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerPeer) IsServerRelaySupported() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
