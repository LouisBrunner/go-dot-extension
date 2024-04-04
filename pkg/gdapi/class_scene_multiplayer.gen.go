// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SceneMultiplayer struct {
  MultiplayerAPI
}

func (me *SceneMultiplayer) BaseClass() string {
  return "SceneMultiplayer"
}

func NewSceneMultiplayer() *SceneMultiplayer {
  str := StringNameFromStr("SceneMultiplayer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SceneMultiplayer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SceneMultiplayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SceneMultiplayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneMultiplayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SceneMultiplayer) SetRootPath(path NodePath, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) GetRootPath() NodePath {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneMultiplayer) Clear()  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) DisconnectPeer(id int64, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) GetAuthenticatingPeers() PackedInt32Array {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_authenticating_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneMultiplayer) SendAuth(id int64, data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_auth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 506032537) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneMultiplayer) CompleteAuth(id int64, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("complete_auth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneMultiplayer) SetAuthCallback(callback Callable, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auth_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) GetAuthCallback() Callable {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auth_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307783378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCallable()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SceneMultiplayer) SetAuthTimeout(timeout float64, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auth_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) GetAuthTimeout() float64 {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auth_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SetRefuseNewConnections(refuse bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_refuse_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) IsRefusingNewConnections() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_refusing_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SetAllowObjectDecoding(enable bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_object_decoding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) IsObjectDecodingAllowed() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_object_decoding_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SetServerRelayEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_server_relay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) IsServerRelayEnabled() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_server_relay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SendBytes(bytes PackedByteArray, id int64, mode MultiplayerPeerTransferMode, channel int64, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307428718) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bytes.AsCTypePtr(), gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&id)
  pinner.Pin(&mode)
  pinner.Pin(&channel)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneMultiplayer) GetMaxSyncPacketSize() int64 {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_sync_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SetMaxSyncPacketSize(size int64, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_sync_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneMultiplayer) GetMaxDeltaPacketSize() int64 {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_delta_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneMultiplayer) SetMaxDeltaPacketSize(size int64, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_delta_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneMultiplayerPeerAuthenticatingSignalFn func(id int, )

func (me *SceneMultiplayer) ConnectPeerAuthenticating(subs SignalSubscribers, fn SceneMultiplayerPeerAuthenticatingSignalFn) {
  sig := StringNameFromStr("peer_authenticating")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneMultiplayer) DisconnectPeerAuthenticating(subs SignalSubscribers, fn SceneMultiplayerPeerAuthenticatingSignalFn) {
  sig := StringNameFromStr("peer_authenticating")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneMultiplayerPeerAuthenticationFailedSignalFn func(id int, )

func (me *SceneMultiplayer) ConnectPeerAuthenticationFailed(subs SignalSubscribers, fn SceneMultiplayerPeerAuthenticationFailedSignalFn) {
  sig := StringNameFromStr("peer_authentication_failed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneMultiplayer) DisconnectPeerAuthenticationFailed(subs SignalSubscribers, fn SceneMultiplayerPeerAuthenticationFailedSignalFn) {
  sig := StringNameFromStr("peer_authentication_failed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SceneMultiplayerPeerPacketSignalFn func(id int, packet PackedByteArray, )

func (me *SceneMultiplayer) ConnectPeerPacket(subs SignalSubscribers, fn SceneMultiplayerPeerPacketSignalFn) {
  sig := StringNameFromStr("peer_packet")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneMultiplayer) DisconnectPeerPacket(subs SignalSubscribers, fn SceneMultiplayerPeerPacketSignalFn) {
  sig := StringNameFromStr("peer_packet")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
