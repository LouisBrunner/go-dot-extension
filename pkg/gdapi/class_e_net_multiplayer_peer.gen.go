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

func NewENetMultiplayerPeer() *ENetMultiplayerPeer {
  str := StringNameFromStr("ENetMultiplayerPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ENetMultiplayerPeer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *ENetMultiplayerPeer) CreateServer(port int64, max_clients int64, max_channels int64, in_bandwidth int64, out_bandwidth int64, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2917761309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&max_clients), gdc.ConstTypePtr(&max_channels), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateClient(address String, port int64, channel_count int64, in_bandwidth int64, out_bandwidth int64, local_port int64, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2327163476) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(address.AsCTypePtr()), gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&channel_count), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), gdc.ConstTypePtr(&local_port), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateMesh(unique_id int64, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_id), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) AddMeshPeer(peer_id int64, host ENetConnection, ) Error {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1293458335) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(host.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewENetConnection()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetMultiplayerPeer) GetPeer(id int64, ) ENetPacketPeer {
  classNameV := StringNameFromStr("ENetMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3793311544) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  ret := NewENetPacketPeer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
