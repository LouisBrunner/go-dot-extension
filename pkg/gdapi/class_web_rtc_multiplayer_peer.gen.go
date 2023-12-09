// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebRTCMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebRTCMultiplayerPeer) BaseClass() string {
  return "WebRTCMultiplayerPeer"
}



// Enums

func (me *WebRTCMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebRTCMultiplayerPeer) CreateServer(channels_config Array, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) CreateClient(peer_id int, channels_config Array, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) CreateMesh(peer_id int, channels_config Array, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) AddPeer(peer WebRTCPeerConnection, peer_id int, unreliable_lifetime int, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) RemovePeer(peer_id int, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) HasPeer(peer_id int, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) GetPeer(peer_id int, )  {
  panic("TODO: implement")
}

func  (me *WebRTCMultiplayerPeer) GetPeers()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
