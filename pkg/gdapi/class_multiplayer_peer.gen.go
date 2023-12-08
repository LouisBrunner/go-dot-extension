// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  MultiplayerPeerTARGET_PEER_BROADCAST = 0
  MultiplayerPeerTARGET_PEER_SERVER = 1
)

type MultiplayerPeerConnectionStatus int
const (
  MultiplayerPeerConnectionDisconnected MultiplayerPeerConnectionStatus = 0
  MultiplayerPeerConnectionConnecting MultiplayerPeerConnectionStatus = 1
  MultiplayerPeerConnectionConnected MultiplayerPeerConnectionStatus = 2
)

type MultiplayerPeerTransferMode int
const (
  MultiplayerPeerTransferModeUnreliable MultiplayerPeerTransferMode = 0
  MultiplayerPeerTransferModeUnreliableOrdered MultiplayerPeerTransferMode = 1
  MultiplayerPeerTransferModeReliable MultiplayerPeerTransferMode = 2
)
