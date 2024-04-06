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

type ptrsForSceneMultiplayerList struct {
	fnSetRootPath              gdc.MethodBindPtr
	fnGetRootPath              gdc.MethodBindPtr
	fnClear                    gdc.MethodBindPtr
	fnDisconnectPeer           gdc.MethodBindPtr
	fnGetAuthenticatingPeers   gdc.MethodBindPtr
	fnSendAuth                 gdc.MethodBindPtr
	fnCompleteAuth             gdc.MethodBindPtr
	fnSetAuthCallback          gdc.MethodBindPtr
	fnGetAuthCallback          gdc.MethodBindPtr
	fnSetAuthTimeout           gdc.MethodBindPtr
	fnGetAuthTimeout           gdc.MethodBindPtr
	fnSetRefuseNewConnections  gdc.MethodBindPtr
	fnIsRefusingNewConnections gdc.MethodBindPtr
	fnSetAllowObjectDecoding   gdc.MethodBindPtr
	fnIsObjectDecodingAllowed  gdc.MethodBindPtr
	fnSetServerRelayEnabled    gdc.MethodBindPtr
	fnIsServerRelayEnabled     gdc.MethodBindPtr
	fnSendBytes                gdc.MethodBindPtr
	fnGetMaxSyncPacketSize     gdc.MethodBindPtr
	fnSetMaxSyncPacketSize     gdc.MethodBindPtr
	fnGetMaxDeltaPacketSize    gdc.MethodBindPtr
	fnSetMaxDeltaPacketSize    gdc.MethodBindPtr
}

var ptrsForSceneMultiplayer ptrsForSceneMultiplayerList

func initSceneMultiplayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SceneMultiplayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_root_path")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetRootPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_root_path")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetRootPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("disconnect_peer")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnDisconnectPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_authenticating_peers")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetAuthenticatingPeers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("send_auth")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSendAuth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 506032537))
	}
	{
		methodName := StringNameFromStr("complete_auth")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnCompleteAuth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844576869))
	}
	{
		methodName := StringNameFromStr("set_auth_callback")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetAuthCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
	}
	{
		methodName := StringNameFromStr("get_auth_callback")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetAuthCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1307783378))
	}
	{
		methodName := StringNameFromStr("set_auth_timeout")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetAuthTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_auth_timeout")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetAuthTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_refuse_new_connections")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetRefuseNewConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_refusing_new_connections")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnIsRefusingNewConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_object_decoding")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetAllowObjectDecoding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_object_decoding_allowed")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnIsObjectDecodingAllowed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_server_relay_enabled")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetServerRelayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_server_relay_enabled")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnIsServerRelayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("send_bytes")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSendBytes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1307428718))
	}
	{
		methodName := StringNameFromStr("get_max_sync_packet_size")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetMaxSyncPacketSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_sync_packet_size")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetMaxSyncPacketSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_delta_packet_size")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnGetMaxDeltaPacketSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_delta_packet_size")
		defer methodName.Destroy()
		ptrsForSceneMultiplayer.fnSetMaxDeltaPacketSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
}

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

func (me *SceneMultiplayer) SetRootPath(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetRootPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) GetRootPath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetRootPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneMultiplayer) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) DisconnectPeer(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnDisconnectPeer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) GetAuthenticatingPeers() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetAuthenticatingPeers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneMultiplayer) SendAuth(id int64, data PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSendAuth), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SceneMultiplayer) CompleteAuth(id int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnCompleteAuth), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SceneMultiplayer) SetAuthCallback(callback Callable) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetAuthCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) GetAuthCallback() Callable {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetAuthCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneMultiplayer) SetAuthTimeout(timeout float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetAuthTimeout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) GetAuthTimeout() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetAuthTimeout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SetRefuseNewConnections(refuse bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetRefuseNewConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) IsRefusingNewConnections() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnIsRefusingNewConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SetAllowObjectDecoding(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetAllowObjectDecoding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) IsObjectDecodingAllowed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnIsObjectDecodingAllowed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SetServerRelayEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetServerRelayEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) IsServerRelayEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnIsServerRelayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SendBytes(bytes PackedByteArray, id int64, mode MultiplayerPeerTransferMode, channel int64) Error {
	cargs := []gdc.ConstTypePtr{bytes.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&channel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&id)
	pinner.Pin(&mode)
	pinner.Pin(&channel)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSendBytes), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SceneMultiplayer) GetMaxSyncPacketSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetMaxSyncPacketSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SetMaxSyncPacketSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetMaxSyncPacketSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneMultiplayer) GetMaxDeltaPacketSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnGetMaxDeltaPacketSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneMultiplayer) SetMaxDeltaPacketSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneMultiplayer.fnSetMaxDeltaPacketSize), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneMultiplayerPeerAuthenticatingSignalFn func(id int)

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

type SceneMultiplayerPeerAuthenticationFailedSignalFn func(id int)

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

type SceneMultiplayerPeerPacketSignalFn func(id int, packet PackedByteArray)

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
