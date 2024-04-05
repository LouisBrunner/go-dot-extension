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

type ptrsForMultiplayerPeerList struct {
  fnSetTransferChannel gdc.MethodBindPtr
  fnGetTransferChannel gdc.MethodBindPtr
  fnSetTransferMode gdc.MethodBindPtr
  fnGetTransferMode gdc.MethodBindPtr
  fnSetTargetPeer gdc.MethodBindPtr
  fnGetPacketPeer gdc.MethodBindPtr
  fnGetPacketChannel gdc.MethodBindPtr
  fnGetPacketMode gdc.MethodBindPtr
  fnPoll gdc.MethodBindPtr
  fnClose gdc.MethodBindPtr
  fnDisconnectPeer gdc.MethodBindPtr
  fnGetConnectionStatus gdc.MethodBindPtr
  fnGetUniqueId gdc.MethodBindPtr
  fnGenerateUniqueId gdc.MethodBindPtr
  fnSetRefuseNewConnections gdc.MethodBindPtr
  fnIsRefusingNewConnections gdc.MethodBindPtr
  fnIsServerRelaySupported gdc.MethodBindPtr
}

var ptrsForMultiplayerPeer ptrsForMultiplayerPeerList

func initMultiplayerPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MultiplayerPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_transfer_channel")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnSetTransferChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_transfer_channel")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetTransferChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_transfer_mode")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnSetTransferMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 950411049))
  }
  {
    methodName := StringNameFromStr("get_transfer_mode")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetTransferMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3369852622))
  }
  {
    methodName := StringNameFromStr("set_target_peer")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnSetTargetPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_packet_peer")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetPacketPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_packet_channel")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetPacketChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_packet_mode")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetPacketMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3369852622))
  }
  {
    methodName := StringNameFromStr("poll")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("close")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("disconnect_peer")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnDisconnectPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4023243586))
  }
  {
    methodName := StringNameFromStr("get_connection_status")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetConnectionStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2147374275))
  }
  {
    methodName := StringNameFromStr("get_unique_id")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGetUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("generate_unique_id")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnGenerateUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_refuse_new_connections")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnSetRefuseNewConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_refusing_new_connections")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnIsRefusingNewConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_server_relay_supported")
    defer methodName.Destroy()
    ptrsForMultiplayerPeer.fnIsServerRelaySupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type MultiplayerPeer struct {
  PacketPeer
}

func (me *MultiplayerPeer) BaseClass() string {
  return "MultiplayerPeer"
}

func NewMultiplayerPeer() *MultiplayerPeer {
  str := StringNameFromStr("MultiplayerPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerPeer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  MultiplayerPeerTargetPeerBroadcast = 0
  MultiplayerPeerTargetPeerServer = 1
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

func (me *MultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiplayerPeer) SetTransferChannel(channel int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnSetTransferChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) GetTransferChannel() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetTransferChannel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) SetTransferMode(mode MultiplayerPeerTransferMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnSetTransferMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) GetTransferMode() MultiplayerPeerTransferMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MultiplayerPeerTransferMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetTransferMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerPeer) SetTargetPeer(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnSetTargetPeer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) GetPacketPeer() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetPacketPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) GetPacketChannel() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetPacketChannel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) GetPacketMode() MultiplayerPeerTransferMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MultiplayerPeerTransferMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetPacketMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerPeer) Poll()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnPoll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) Close()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) DisconnectPeer(peer int64, force bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , gdc.ConstTypePtr(&force) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnDisconnectPeer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) GetConnectionStatus() MultiplayerPeerConnectionStatus {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MultiplayerPeerConnectionStatus

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetConnectionStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerPeer) GetUniqueId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGetUniqueId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) GenerateUniqueId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnGenerateUniqueId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) SetRefuseNewConnections(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnSetRefuseNewConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerPeer) IsRefusingNewConnections() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnIsRefusingNewConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerPeer) IsServerRelaySupported() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerPeer.fnIsServerRelaySupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerPeerPeerConnectedSignalFn func(id int, )

func (me *MultiplayerPeer) ConnectPeerConnected(subs SignalSubscribers, fn MultiplayerPeerPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerPeer) DisconnectPeerConnected(subs SignalSubscribers, fn MultiplayerPeerPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerPeerPeerDisconnectedSignalFn func(id int, )

func (me *MultiplayerPeer) ConnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerPeerPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerPeer) DisconnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerPeerPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
