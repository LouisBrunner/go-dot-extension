// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ENetMultiplayerPeer struct {
  MultiplayerPeer
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

func  (me *ENetMultiplayerPeer) CreateServer(port int, max_clients int, max_channels int, in_bandwidth int, out_bandwidth int, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2917761309) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&max_clients), gdc.ConstTypePtr(&max_channels), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateClient(address String, port int, channel_count int, in_bandwidth int, out_bandwidth int, local_port int, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2327163476) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(address.AsCTypePtr()), gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&channel_count), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), gdc.ConstTypePtr(&local_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateMesh(unique_id int, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetMultiplayerPeer) AddMeshPeer(peer_id int, host ENetConnection, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1293458335) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(host.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetMultiplayerPeer) SetBindIp(ip String, )  {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_ip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ip.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetMultiplayerPeer) GetHost() ENetConnection {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4103238886) // FIXME: should cache?
  var ret ENetConnection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetMultiplayerPeer) GetPeer(id int, ) ENetPacketPeer {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3793311544) // FIXME: should cache?
  var ret ENetPacketPeer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
