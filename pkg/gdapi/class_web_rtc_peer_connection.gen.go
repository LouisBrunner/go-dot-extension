// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCPeerConnection struct {
  obj gdc.ObjectPtr
}

func (me *WebRTCPeerConnection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebRTCPeerConnection) BaseClass() string {
  return "WebRTCPeerConnection"
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
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) Initialize(configuration Dictionary, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) CreateDataChannel(label String, options Dictionary, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) CreateOffer()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) SetLocalDescription(type_ String, sdp String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) SetRemoteDescription(type_ String, sdp String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) AddIceCandidate(media String, index int, name String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) Poll()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) Close()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) GetConnectionState()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) GetGatheringState()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnection) GetSignalingState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
