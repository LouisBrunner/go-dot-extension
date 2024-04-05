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

type ptrsForENetMultiplayerPeerList struct {
  fnCreateServer gdc.MethodBindPtr
  fnCreateClient gdc.MethodBindPtr
  fnCreateMesh gdc.MethodBindPtr
  fnAddMeshPeer gdc.MethodBindPtr
  fnSetBindIp gdc.MethodBindPtr
  fnGetHost gdc.MethodBindPtr
  fnGetPeer gdc.MethodBindPtr
}

var ptrsForENetMultiplayerPeer ptrsForENetMultiplayerPeerList

func initENetMultiplayerPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ENetMultiplayerPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_server")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnCreateServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2917761309))
  }
  {
    methodName := StringNameFromStr("create_client")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnCreateClient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2327163476))
  }
  {
    methodName := StringNameFromStr("create_mesh")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnCreateMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844576869))
  }
  {
    methodName := StringNameFromStr("add_mesh_peer")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnAddMeshPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1293458335))
  }
  {
    methodName := StringNameFromStr("set_bind_ip")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnSetBindIp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_host")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnGetHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4103238886))
  }
  {
    methodName := StringNameFromStr("get_peer")
    defer methodName.Destroy()
    ptrsForENetMultiplayerPeer.fnGetPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3793311544))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , gdc.ConstTypePtr(&max_clients) , gdc.ConstTypePtr(&max_channels) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)
  pinner.Pin(&max_clients)
  pinner.Pin(&max_channels)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnCreateServer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateClient(address String, port int64, channel_count int64, in_bandwidth int64, out_bandwidth int64, local_port int64, ) Error {
  cargs := []gdc.ConstTypePtr{address.AsCTypePtr(), gdc.ConstTypePtr(&port) , gdc.ConstTypePtr(&channel_count) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , gdc.ConstTypePtr(&local_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)
  pinner.Pin(&channel_count)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)
  pinner.Pin(&local_port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnCreateClient), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) CreateMesh(unique_id int64, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&unique_id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnCreateMesh), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) AddMeshPeer(peer_id int64, host ENetConnection, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , host.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnAddMeshPeer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetMultiplayerPeer) SetBindIp(ip String, )  {
  cargs := []gdc.ConstTypePtr{ip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnSetBindIp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetMultiplayerPeer) GetHost() ENetConnection {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewENetConnection()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnGetHost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetMultiplayerPeer) GetPeer(id int64, ) ENetPacketPeer {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewENetPacketPeer()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetMultiplayerPeer.fnGetPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
