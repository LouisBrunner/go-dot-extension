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
  WebRTCPeerConnectionConnectionStateStateNew WebRTCPeerConnectionConnectionState = 0
  WebRTCPeerConnectionConnectionStateStateConnecting WebRTCPeerConnectionConnectionState = 1
  WebRTCPeerConnectionConnectionStateStateConnected WebRTCPeerConnectionConnectionState = 2
  WebRTCPeerConnectionConnectionStateStateDisconnected WebRTCPeerConnectionConnectionState = 3
  WebRTCPeerConnectionConnectionStateStateFailed WebRTCPeerConnectionConnectionState = 4
  WebRTCPeerConnectionConnectionStateStateClosed WebRTCPeerConnectionConnectionState = 5
)

type WebRTCPeerConnectionGatheringState int
const (
  WebRTCPeerConnectionGatheringStateGatheringStateNew WebRTCPeerConnectionGatheringState = 0
  WebRTCPeerConnectionGatheringStateGatheringStateGathering WebRTCPeerConnectionGatheringState = 1
  WebRTCPeerConnectionGatheringStateGatheringStateComplete WebRTCPeerConnectionGatheringState = 2
)

type WebRTCPeerConnectionSignalingState int
const (
  WebRTCPeerConnectionSignalingStateSignalingStateStable WebRTCPeerConnectionSignalingState = 0
  WebRTCPeerConnectionSignalingStateSignalingStateHaveLocalOffer WebRTCPeerConnectionSignalingState = 1
  WebRTCPeerConnectionSignalingStateSignalingStateHaveRemoteOffer WebRTCPeerConnectionSignalingState = 2
  WebRTCPeerConnectionSignalingStateSignalingStateHaveLocalPranswer WebRTCPeerConnectionSignalingState = 3
  WebRTCPeerConnectionSignalingStateSignalingStateHaveRemotePranswer WebRTCPeerConnectionSignalingState = 4
  WebRTCPeerConnectionSignalingStateSignalingStateClosed WebRTCPeerConnectionSignalingState = 5
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

func  WebRTCPeerConnectionSetDefaultExtension(extension_class StringName, )  {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension_class.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}

func  (me *WebRTCPeerConnection) Initialize(configuration Dictionary, ) Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("initialize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2625064318) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{configuration.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) CreateDataChannel(label String, options Dictionary, ) WebRTCDataChannel {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_data_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1288557393) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWebRTCDataChannel()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebRTCPeerConnection) CreateOffer() Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_offer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) SetLocalDescription(type_ String, sdp String, ) Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_local_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), sdp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) SetRemoteDescription(type_ String, sdp String, ) Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_remote_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), sdp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) AddIceCandidate(media String, index int64, name String, ) Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ice_candidate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3958950400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{media.AsCTypePtr(), gdc.ConstTypePtr(&index) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) Poll() Error {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
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

func  (me *WebRTCPeerConnection) Close()  {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebRTCPeerConnection) GetConnectionState() WebRTCPeerConnectionConnectionState {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2275710506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebRTCPeerConnectionConnectionState

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) GetGatheringState() WebRTCPeerConnectionGatheringState {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gathering_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4262591401) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebRTCPeerConnectionGatheringState

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCPeerConnection) GetSignalingState() WebRTCPeerConnectionSignalingState {
  classNameV := StringNameFromStr("WebRTCPeerConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_signaling_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3342956226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebRTCPeerConnectionSignalingState

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals

type WebRTCPeerConnectionSessionDescriptionCreatedSignalFn func(type_ String, sdp String, )

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

type WebRTCPeerConnectionIceCandidateCreatedSignalFn func(media String, index int, name String, )

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

type WebRTCPeerConnectionDataChannelReceivedSignalFn func(channel WebRTCDataChannel, )

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
