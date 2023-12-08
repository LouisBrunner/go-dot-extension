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

func  WebRTCPeerConnectionSetDefaultExtension(extension_class StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) Initialize(configuration Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) CreateDataChannel(label String, options Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) CreateOffer() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) SetLocalDescription(type_ String, sdp String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) SetRemoteDescription(type_ String, sdp String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) AddIceCandidate(media String, index int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) Close() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) GetConnectionState() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) GetGatheringState() { // TODO: return value
  // TODO: implement
}

func  (me *WebRTCPeerConnection) GetSignalingState() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
