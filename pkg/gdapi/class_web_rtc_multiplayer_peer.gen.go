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

type ptrsForWebRTCMultiplayerPeerList struct {
	fnCreateServer gdc.MethodBindPtr
	fnCreateClient gdc.MethodBindPtr
	fnCreateMesh   gdc.MethodBindPtr
	fnAddPeer      gdc.MethodBindPtr
	fnRemovePeer   gdc.MethodBindPtr
	fnHasPeer      gdc.MethodBindPtr
	fnGetPeer      gdc.MethodBindPtr
	fnGetPeers     gdc.MethodBindPtr
}

var ptrsForWebRTCMultiplayerPeer ptrsForWebRTCMultiplayerPeerList

func initWebRTCMultiplayerPeerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WebRTCMultiplayerPeer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_server")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnCreateServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2865356025))
	}
	{
		methodName := StringNameFromStr("create_client")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnCreateClient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2641732907))
	}
	{
		methodName := StringNameFromStr("create_mesh")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnCreateMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2641732907))
	}
	{
		methodName := StringNameFromStr("add_peer")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnAddPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4078953270))
	}
	{
		methodName := StringNameFromStr("remove_peer")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnRemovePeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("has_peer")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnHasPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("get_peer")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnGetPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3554694381))
	}
	{
		methodName := StringNameFromStr("get_peers")
		defer methodName.Destroy()
		ptrsForWebRTCMultiplayerPeer.fnGetPeers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
}

type WebRTCMultiplayerPeer struct {
	MultiplayerPeer
}

func (me *WebRTCMultiplayerPeer) BaseClass() string {
	return "WebRTCMultiplayerPeer"
}

func NewWebRTCMultiplayerPeer() *WebRTCMultiplayerPeer {
	str := StringNameFromStr("WebRTCMultiplayerPeer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WebRTCMultiplayerPeer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *WebRTCMultiplayerPeer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WebRTCMultiplayerPeer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WebRTCMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *WebRTCMultiplayerPeer) CreateServer(channels_config Array) Error {
	cargs := []gdc.ConstTypePtr{channels_config.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnCreateServer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCMultiplayerPeer) CreateClient(peer_id int64, channels_config Array) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), channels_config.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&peer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnCreateClient), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCMultiplayerPeer) CreateMesh(peer_id int64, channels_config Array) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), channels_config.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&peer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnCreateMesh), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCMultiplayerPeer) AddPeer(peer WebRTCPeerConnection, peer_id int64, unreliable_lifetime int64) Error {
	cargs := []gdc.ConstTypePtr{peer.AsCTypePtr(), gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(&unreliable_lifetime)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&peer_id)
	pinner.Pin(&unreliable_lifetime)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnAddPeer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCMultiplayerPeer) RemovePeer(peer_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnRemovePeer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebRTCMultiplayerPeer) HasPeer(peer_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&peer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnHasPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCMultiplayerPeer) GetPeer(peer_id int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&peer_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnGetPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebRTCMultiplayerPeer) GetPeers() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCMultiplayerPeer.fnGetPeers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
