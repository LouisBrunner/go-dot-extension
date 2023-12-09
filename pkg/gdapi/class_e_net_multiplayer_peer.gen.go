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



// Enums

func (me *ENetMultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ENetMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ENetMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ENetMultiplayerPeer) CreateServer(port int, max_clients int, max_channels int, in_bandwidth int, out_bandwidth int, )  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) CreateClient(address String, port int, channel_count int, in_bandwidth int, out_bandwidth int, local_port int, )  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) CreateMesh(unique_id int, )  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) AddMeshPeer(peer_id int, host ENetConnection, )  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) SetBindIp(ip String, )  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) GetHost()  {
  panic("TODO: implement")
}

func  (me *ENetMultiplayerPeer) GetPeer(id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
