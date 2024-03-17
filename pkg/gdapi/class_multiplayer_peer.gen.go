// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerPeer struct {
  PacketPeer
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

func  (me *MultiplayerPeer) SetTransferChannel(channel int, )  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transfer_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) GetTransferChannel() int {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transfer_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) SetTransferMode(mode MultiplayerPeerTransferMode, )  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transfer_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 950411049) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) GetTransferMode() MultiplayerPeerTransferMode {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transfer_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3369852622) // FIXME: should cache?
  var ret MultiplayerPeerTransferMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) SetTargetPeer(id int, )  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) GetPacketPeer() int {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) GetPacketChannel() int {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) GetPacketMode() MultiplayerPeerTransferMode {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3369852622) // FIXME: should cache?
  var ret MultiplayerPeerTransferMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) Poll()  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) Close()  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) DisconnectPeer(peer int, force bool, )  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4023243586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer), gdc.ConstTypePtr(&force), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) GetConnectionStatus() MultiplayerPeerConnectionStatus {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147374275) // FIXME: should cache?
  var ret MultiplayerPeerConnectionStatus
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) GetUniqueId() int {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) GenerateUniqueId() int {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_unique_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) SetRefuseNewConnections(enable bool, )  {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_refuse_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerPeer) IsRefusingNewConnections() bool {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_refusing_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerPeer) IsServerRelaySupported() bool {
  classNameV := StringNameFromStr("MultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_server_relay_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
