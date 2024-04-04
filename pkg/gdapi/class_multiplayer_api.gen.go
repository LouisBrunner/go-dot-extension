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

type MultiplayerAPI struct {
  RefCounted
}

func (me *MultiplayerAPI) BaseClass() string {
  return "MultiplayerAPI"
}

func NewMultiplayerAPI() *MultiplayerAPI {
  str := StringNameFromStr("MultiplayerAPI") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerAPI{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type MultiplayerAPIRPCMode int
const (
  MultiplayerAPIRPCModeRpcModeDisabled MultiplayerAPIRPCMode = 0
  MultiplayerAPIRPCModeRpcModeAnyPeer MultiplayerAPIRPCMode = 1
  MultiplayerAPIRPCModeRpcModeAuthority MultiplayerAPIRPCMode = 2
)

func (me *MultiplayerAPI) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerAPI) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerAPI) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiplayerAPI) HasMultiplayerPeer() bool {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_multiplayer_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerAPI) GetMultiplayerPeer() MultiplayerPeer {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3223692825) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMultiplayerPeer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiplayerAPI) SetMultiplayerPeer(peer MultiplayerPeer, )  {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3694835298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{peer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiplayerAPI) GetUniqueId() int64 {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerAPI) IsServer() bool {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerAPI) GetRemoteSenderId() int64 {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remote_sender_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiplayerAPI) Poll() Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerAPI) Rpc(peer int64, object Object, method StringName, arguments Array, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rpc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2077486355) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer) , object.AsCTypePtr(), method.AsCTypePtr(), arguments.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&peer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerAPI) ObjectConfigurationAdd(object Object, configuration Variant, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("object_configuration_add")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171879464) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), configuration.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerAPI) ObjectConfigurationRemove(object Object, configuration Variant, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("object_configuration_remove")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171879464) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), configuration.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiplayerAPI) GetPeers() PackedInt32Array {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  MultiplayerAPISetDefaultInterface(interface_name StringName, )  {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interface_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}

func  MultiplayerAPIGetDefaultInterface() StringName {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2737447660) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  MultiplayerAPICreateDefaultInterface() MultiplayerAPI {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3294156723) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMultiplayerAPI()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerAPIPeerConnectedSignalFn func(id int, )

func (me *MultiplayerAPI) ConnectPeerConnected(subs SignalSubscribers, fn MultiplayerAPIPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectPeerConnected(subs SignalSubscribers, fn MultiplayerAPIPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerAPIPeerDisconnectedSignalFn func(id int, )

func (me *MultiplayerAPI) ConnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerAPIPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerAPIPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerAPIConnectedToServerSignalFn func()

func (me *MultiplayerAPI) ConnectConnectedToServer(subs SignalSubscribers, fn MultiplayerAPIConnectedToServerSignalFn) {
  sig := StringNameFromStr("connected_to_server")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectConnectedToServer(subs SignalSubscribers, fn MultiplayerAPIConnectedToServerSignalFn) {
  sig := StringNameFromStr("connected_to_server")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerAPIConnectionFailedSignalFn func()

func (me *MultiplayerAPI) ConnectConnectionFailed(subs SignalSubscribers, fn MultiplayerAPIConnectionFailedSignalFn) {
  sig := StringNameFromStr("connection_failed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectConnectionFailed(subs SignalSubscribers, fn MultiplayerAPIConnectionFailedSignalFn) {
  sig := StringNameFromStr("connection_failed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type MultiplayerAPIServerDisconnectedSignalFn func()

func (me *MultiplayerAPI) ConnectServerDisconnected(subs SignalSubscribers, fn MultiplayerAPIServerDisconnectedSignalFn) {
  sig := StringNameFromStr("server_disconnected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectServerDisconnected(subs SignalSubscribers, fn MultiplayerAPIServerDisconnectedSignalFn) {
  sig := StringNameFromStr("server_disconnected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
