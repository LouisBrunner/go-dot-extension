// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerAPI struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerAPI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerAPI) BaseClass() string {
  return "MultiplayerAPI"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) GetMultiplayerPeer() MultiplayerPeer {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3223692825) // FIXME: should cache?
  var ret MultiplayerPeer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) SetMultiplayerPeer(peer MultiplayerPeer, )  {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3694835298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(peer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiplayerAPI) GetUniqueId() int {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) IsServer() bool {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) GetRemoteSenderId() int {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remote_sender_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) Poll() Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) Rpc(peer int, object Object, method StringName, arguments Array, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rpc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1833408346) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer), gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(arguments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) ObjectConfigurationAdd(object Object, configuration Variant, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("object_configuration_add")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171879464) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(configuration.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) ObjectConfigurationRemove(object Object, configuration Variant, ) Error {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("object_configuration_remove")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171879464) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(configuration.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiplayerAPI) GetPeers() PackedInt32Array {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  MultiplayerAPISetDefaultInterface(interface_name StringName, )  {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

func  MultiplayerAPIGetDefaultInterface() StringName {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2737447660) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  MultiplayerAPICreateDefaultInterface() MultiplayerAPI {
  classNameV := StringNameFromStr("MultiplayerAPI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_default_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3294156723) // FIXME: should cache?
  var ret MultiplayerAPI
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerAPIPeerConnectedSignalFn func(id int, )

func (me *MultiplayerAPI) ConnectPeerConnected(subs SignalSubscribers, fn MultiplayerAPIPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectPeerConnected(subs SignalSubscribers, fn MultiplayerAPIPeerConnectedSignalFn) {
  sig := StringNameFromStr("peer_connected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerAPIPeerDisconnectedSignalFn func(id int, )

func (me *MultiplayerAPI) ConnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerAPIPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectPeerDisconnected(subs SignalSubscribers, fn MultiplayerAPIPeerDisconnectedSignalFn) {
  sig := StringNameFromStr("peer_disconnected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerAPIConnectedToServerSignalFn func()

func (me *MultiplayerAPI) ConnectConnectedToServer(subs SignalSubscribers, fn MultiplayerAPIConnectedToServerSignalFn) {
  sig := StringNameFromStr("connected_to_server")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectConnectedToServer(subs SignalSubscribers, fn MultiplayerAPIConnectedToServerSignalFn) {
  sig := StringNameFromStr("connected_to_server")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerAPIConnectionFailedSignalFn func()

func (me *MultiplayerAPI) ConnectConnectionFailed(subs SignalSubscribers, fn MultiplayerAPIConnectionFailedSignalFn) {
  sig := StringNameFromStr("connection_failed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectConnectionFailed(subs SignalSubscribers, fn MultiplayerAPIConnectionFailedSignalFn) {
  sig := StringNameFromStr("connection_failed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type MultiplayerAPIServerDisconnectedSignalFn func()

func (me *MultiplayerAPI) ConnectServerDisconnected(subs SignalSubscribers, fn MultiplayerAPIServerDisconnectedSignalFn) {
  sig := StringNameFromStr("server_disconnected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MultiplayerAPI) DisconnectServerDisconnected(subs SignalSubscribers, fn MultiplayerAPIServerDisconnectedSignalFn) {
  sig := StringNameFromStr("server_disconnected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
