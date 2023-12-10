// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SceneMultiplayer struct {
  obj gdc.ObjectPtr
}

func (me *SceneMultiplayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneMultiplayer) BaseClass() string {
  return "SceneMultiplayer"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) GetRootPath() NodePath {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) Clear()  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) DisconnectPeer(id int, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) GetAuthenticatingPeers() PackedInt32Array {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_authenticating_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SendAuth(id int, data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_auth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 506032537) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) CompleteAuth(id int, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("complete_auth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetAuthCallback(callback Callable, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auth_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) GetAuthCallback() Callable {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auth_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307783378) // FIXME: should cache?
  var ret Callable
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetAuthTimeout(timeout float32, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auth_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) GetAuthTimeout() float32 {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auth_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetRefuseNewConnections(refuse bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_refuse_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) IsRefusingNewConnections() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_refusing_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetAllowObjectDecoding(enable bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_object_decoding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) IsObjectDecodingAllowed() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_object_decoding_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetServerRelayEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_server_relay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) IsServerRelayEnabled() bool {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_server_relay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SendBytes(bytes PackedByteArray, id int, mode MultiplayerPeerTransferMode, channel int, ) Error {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2742700601) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bytes.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) GetMaxSyncPacketSize() int {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_sync_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetMaxSyncPacketSize(size int, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_sync_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneMultiplayer) GetMaxDeltaPacketSize() int {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_delta_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneMultiplayer) SetMaxDeltaPacketSize(size int, )  {
  classNameV := StringNameFromStr("SceneMultiplayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_delta_packet_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *SceneMultiplayer) GetPropRootPath() NodePath {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropRootPath(value NodePath) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropAuthCallback() Callable {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropAuthCallback(value Callable) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropAuthTimeout() float32 {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropAuthTimeout(value float32) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropAllowObjectDecoding() bool {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropAllowObjectDecoding(value bool) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropRefuseNewConnections() bool {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropRefuseNewConnections(value bool) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropServerRelay() bool {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropServerRelay(value bool) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropMaxSyncPacketSize() int {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropMaxSyncPacketSize(value int) {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) GetPropMaxDeltaPacketSize() int {
  panic("TODO: implement")
}

func (me *SceneMultiplayer) SetPropMaxDeltaPacketSize(value int) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API