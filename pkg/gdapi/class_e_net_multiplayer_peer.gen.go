// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *ENetMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ENetMultiplayerPeer) BaseClass() string {
  return "ENetMultiplayerPeer"
}

func  (me *ENetMultiplayerPeer) CreateServer(port int, max_clients int, max_channels int, in_bandwidth int, out_bandwidth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) CreateClient(address String, port int, channel_count int, in_bandwidth int, out_bandwidth int, local_port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) CreateMesh(unique_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) AddMeshPeer(peer_id int, host ENetConnection, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) SetBindIp(ip String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) GetHost() { // TODO: return value
  // TODO: implement
}

func  (me *ENetMultiplayerPeer) GetPeer(id int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
