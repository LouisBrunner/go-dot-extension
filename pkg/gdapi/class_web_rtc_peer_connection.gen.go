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

type ptrsForWebRTCPeerConnectionList struct {
	fnSetDefaultExtension  gdc.MethodBindPtr
	fnInitialize           gdc.MethodBindPtr
	fnCreateDataChannel    gdc.MethodBindPtr
	fnCreateOffer          gdc.MethodBindPtr
	fnSetLocalDescription  gdc.MethodBindPtr
	fnSetRemoteDescription gdc.MethodBindPtr
	fnAddIceCandidate      gdc.MethodBindPtr
	fnPoll                 gdc.MethodBindPtr
	fnClose                gdc.MethodBindPtr
	fnGetConnectionState   gdc.MethodBindPtr
	fnGetGatheringState    gdc.MethodBindPtr
	fnGetSignalingState    gdc.MethodBindPtr
}

var ptrsForWebRTCPeerConnection ptrsForWebRTCPeerConnectionList

func initWebRTCPeerConnectionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WebRTCPeerConnection")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_extension")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnSetDefaultExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("initialize")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnInitialize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2625064318))
	}
	{
		methodName := StringNameFromStr("create_data_channel")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnCreateDataChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1288557393))
	}
	{
		methodName := StringNameFromStr("create_offer")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnCreateOffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("set_local_description")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnSetLocalDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("set_remote_description")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnSetRemoteDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("add_ice_candidate")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnAddIceCandidate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3958950400))
	}
	{
		methodName := StringNameFromStr("poll")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("close")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_connection_state")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnGetConnectionState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2275710506))
	}
	{
		methodName := StringNameFromStr("get_gathering_state")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnGetGatheringState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4262591401))
	}
	{
		methodName := StringNameFromStr("get_signaling_state")
		defer methodName.Destroy()
		ptrsForWebRTCPeerConnection.fnGetSignalingState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3342956226))
	}

}

type WebRTCPeerConnection struct {
	RefCounted
}

func (me *WebRTCPeerConnection) BaseClass() string {
	return "WebRTCPeerConnection"
}

func NewWebRTCPeerConnection() *WebRTCPeerConnection {
	str := StringNameFromStr("WebRTCPeerConnection") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WebRTCPeerConnection{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type WebRTCPeerConnectionConnectionState int

const (
	WebRTCPeerConnectionConnectionStateStateNew          WebRTCPeerConnectionConnectionState = 0
	WebRTCPeerConnectionConnectionStateStateConnecting   WebRTCPeerConnectionConnectionState = 1
	WebRTCPeerConnectionConnectionStateStateConnected    WebRTCPeerConnectionConnectionState = 2
	WebRTCPeerConnectionConnectionStateStateDisconnected WebRTCPeerConnectionConnectionState = 3
	WebRTCPeerConnectionConnectionStateStateFailed       WebRTCPeerConnectionConnectionState = 4
	WebRTCPeerConnectionConnectionStateStateClosed       WebRTCPeerConnectionConnectionState = 5
)

type WebRTCPeerConnectionGatheringState int

const (
	WebRTCPeerConnectionGatheringStateGatheringStateNew       WebRTCPeerConnectionGatheringState = 0
	WebRTCPeerConnectionGatheringStateGatheringStateGathering WebRTCPeerConnectionGatheringState = 1
	WebRTCPeerConnectionGatheringStateGatheringStateComplete  WebRTCPeerConnectionGatheringState = 2
)

type WebRTCPeerConnectionSignalingState int

const (
	WebRTCPeerConnectionSignalingStateSignalingStateStable             WebRTCPeerConnectionSignalingState = 0
	WebRTCPeerConnectionSignalingStateSignalingStateHaveLocalOffer     WebRTCPeerConnectionSignalingState = 1
	WebRTCPeerConnectionSignalingStateSignalingStateHaveRemoteOffer    WebRTCPeerConnectionSignalingState = 2
	WebRTCPeerConnectionSignalingStateSignalingStateHaveLocalPranswer  WebRTCPeerConnectionSignalingState = 3
	WebRTCPeerConnectionSignalingStateSignalingStateHaveRemotePranswer WebRTCPeerConnectionSignalingState = 4
	WebRTCPeerConnectionSignalingStateSignalingStateClosed             WebRTCPeerConnectionSignalingState = 5
)

func (me *WebRTCPeerConnection) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WebRTCPeerConnection) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WebRTCPeerConnection) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func WebRTCPeerConnectionSetDefaultExtension(extension_class StringName) {
	cargs := []gdc.ConstTypePtr{extension_class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnSetDefaultExtension), nil, unsafe.SliceData(cargs), nil)

}

func (me *WebRTCPeerConnection) Initialize(configuration Dictionary) Error {
	cargs := []gdc.ConstTypePtr{configuration.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnInitialize), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) CreateDataChannel(label String, options Dictionary) WebRTCDataChannel {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), options.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWebRTCDataChannel()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnCreateDataChannel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebRTCPeerConnection) CreateOffer() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnCreateOffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) SetLocalDescription(type_ String, sdp String) Error {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), sdp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnSetLocalDescription), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) SetRemoteDescription(type_ String, sdp String) Error {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), sdp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnSetRemoteDescription), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) AddIceCandidate(media String, index int64, name String) Error {
	cargs := []gdc.ConstTypePtr{media.AsCTypePtr(), gdc.ConstTypePtr(&index), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnAddIceCandidate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) Poll() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) Close() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebRTCPeerConnection) GetConnectionState() WebRTCPeerConnectionConnectionState {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebRTCPeerConnectionConnectionState

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnGetConnectionState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) GetGatheringState() WebRTCPeerConnectionGatheringState {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebRTCPeerConnectionGatheringState

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnGetGatheringState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCPeerConnection) GetSignalingState() WebRTCPeerConnectionSignalingState {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebRTCPeerConnectionSignalingState

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCPeerConnection.fnGetSignalingState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals

type WebRTCPeerConnectionSessionDescriptionCreatedSignalFn func(type_ String, sdp String)

func (me *WebRTCPeerConnection) ConnectSessionDescriptionCreated(subs SignalSubscribers, fn WebRTCPeerConnectionSessionDescriptionCreatedSignalFn) {
	sig := StringNameFromStr("session_description_created")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebRTCPeerConnection) DisconnectSessionDescriptionCreated(subs SignalSubscribers, fn WebRTCPeerConnectionSessionDescriptionCreatedSignalFn) {
	sig := StringNameFromStr("session_description_created")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebRTCPeerConnectionIceCandidateCreatedSignalFn func(media String, index int, name String)

func (me *WebRTCPeerConnection) ConnectIceCandidateCreated(subs SignalSubscribers, fn WebRTCPeerConnectionIceCandidateCreatedSignalFn) {
	sig := StringNameFromStr("ice_candidate_created")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebRTCPeerConnection) DisconnectIceCandidateCreated(subs SignalSubscribers, fn WebRTCPeerConnectionIceCandidateCreatedSignalFn) {
	sig := StringNameFromStr("ice_candidate_created")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WebRTCPeerConnectionDataChannelReceivedSignalFn func(channel WebRTCDataChannel)

func (me *WebRTCPeerConnection) ConnectDataChannelReceived(subs SignalSubscribers, fn WebRTCPeerConnectionDataChannelReceivedSignalFn) {
	sig := StringNameFromStr("data_channel_received")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *WebRTCPeerConnection) DisconnectDataChannelReceived(subs SignalSubscribers, fn WebRTCPeerConnectionDataChannelReceivedSignalFn) {
	sig := StringNameFromStr("data_channel_received")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
