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

type ptrsForMultiplayerAPIList struct {
	fnHasMultiplayerPeer        gdc.MethodBindPtr
	fnGetMultiplayerPeer        gdc.MethodBindPtr
	fnSetMultiplayerPeer        gdc.MethodBindPtr
	fnGetUniqueId               gdc.MethodBindPtr
	fnIsServer                  gdc.MethodBindPtr
	fnGetRemoteSenderId         gdc.MethodBindPtr
	fnPoll                      gdc.MethodBindPtr
	fnRpc                       gdc.MethodBindPtr
	fnObjectConfigurationAdd    gdc.MethodBindPtr
	fnObjectConfigurationRemove gdc.MethodBindPtr
	fnGetPeers                  gdc.MethodBindPtr
	fnSetDefaultInterface       gdc.MethodBindPtr
	fnGetDefaultInterface       gdc.MethodBindPtr
	fnCreateDefaultInterface    gdc.MethodBindPtr
}

var ptrsForMultiplayerAPI ptrsForMultiplayerAPIList

func initMultiplayerAPIPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiplayerAPI")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_multiplayer_peer")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnHasMultiplayerPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_multiplayer_peer")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnGetMultiplayerPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3223692825))
	}
	{
		methodName := StringNameFromStr("set_multiplayer_peer")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnSetMultiplayerPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3694835298))
	}
	{
		methodName := StringNameFromStr("get_unique_id")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnGetUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("is_server")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnIsServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_remote_sender_id")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnGetRemoteSenderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("poll")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("rpc")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnRpc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2077486355))
	}
	{
		methodName := StringNameFromStr("object_configuration_add")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnObjectConfigurationAdd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1171879464))
	}
	{
		methodName := StringNameFromStr("object_configuration_remove")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnObjectConfigurationRemove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1171879464))
	}
	{
		methodName := StringNameFromStr("get_peers")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnGetPeers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_default_interface")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnSetDefaultInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_default_interface")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnGetDefaultInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2737447660))
	}
	{
		methodName := StringNameFromStr("create_default_interface")
		defer methodName.Destroy()
		ptrsForMultiplayerAPI.fnCreateDefaultInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3294156723))
	}
}

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
	MultiplayerAPIRPCModeRpcModeDisabled  MultiplayerAPIRPCMode = 0
	MultiplayerAPIRPCModeRpcModeAnyPeer   MultiplayerAPIRPCMode = 1
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

func (me *MultiplayerAPI) HasMultiplayerPeer() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnHasMultiplayerPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerAPI) GetMultiplayerPeer() MultiplayerPeer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMultiplayerPeer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnGetMultiplayerPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiplayerAPI) SetMultiplayerPeer(peer MultiplayerPeer) {
	cargs := []gdc.ConstTypePtr{peer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnSetMultiplayerPeer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiplayerAPI) GetUniqueId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnGetUniqueId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerAPI) IsServer() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnIsServer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerAPI) GetRemoteSenderId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnGetRemoteSenderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiplayerAPI) Poll() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MultiplayerAPI) Rpc(peer int64, object Object, method StringName, arguments Array) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer), object.AsCTypePtr(), method.AsCTypePtr(), arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&peer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnRpc), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MultiplayerAPI) ObjectConfigurationAdd(object Object, configuration Variant) Error {
	cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), configuration.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnObjectConfigurationAdd), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MultiplayerAPI) ObjectConfigurationRemove(object Object, configuration Variant) Error {
	cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), configuration.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnObjectConfigurationRemove), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MultiplayerAPI) GetPeers() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnGetPeers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func MultiplayerAPISetDefaultInterface(interface_name StringName) {
	cargs := []gdc.ConstTypePtr{interface_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnSetDefaultInterface), nil, unsafe.SliceData(cargs), nil)

}

func MultiplayerAPIGetDefaultInterface() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnGetDefaultInterface), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func MultiplayerAPICreateDefaultInterface() MultiplayerAPI {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMultiplayerAPI()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiplayerAPI.fnCreateDefaultInterface), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiplayerAPIPeerConnectedSignalFn func(id int)

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

type MultiplayerAPIPeerDisconnectedSignalFn func(id int)

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
