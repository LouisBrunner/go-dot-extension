// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  WebRTCPeerConnectionStateNew WebRTCPeerConnectionConnectionState = 0
  WebRTCPeerConnectionStateConnecting WebRTCPeerConnectionConnectionState = 1
  WebRTCPeerConnectionStateConnected WebRTCPeerConnectionConnectionState = 2
  WebRTCPeerConnectionStateDisconnected WebRTCPeerConnectionConnectionState = 3
  WebRTCPeerConnectionStateFailed WebRTCPeerConnectionConnectionState = 4
  WebRTCPeerConnectionStateClosed WebRTCPeerConnectionConnectionState = 5
)

type WebRTCPeerConnectionGatheringState int
const (
  WebRTCPeerConnectionGatheringStateNew WebRTCPeerConnectionGatheringState = 0
  WebRTCPeerConnectionGatheringStateGathering WebRTCPeerConnectionGatheringState = 1
  WebRTCPeerConnectionGatheringStateComplete WebRTCPeerConnectionGatheringState = 2
)

type WebRTCPeerConnectionSignalingState int
const (
  WebRTCPeerConnectionSignalingStateStable WebRTCPeerConnectionSignalingState = 0
  WebRTCPeerConnectionSignalingStateHaveLocalOffer WebRTCPeerConnectionSignalingState = 1
  WebRTCPeerConnectionSignalingStateHaveRemoteOffer WebRTCPeerConnectionSignalingState = 2
  WebRTCPeerConnectionSignalingStateHaveLocalPranswer WebRTCPeerConnectionSignalingState = 3
  WebRTCPeerConnectionSignalingStateHaveRemotePranswer WebRTCPeerConnectionSignalingState = 4
  WebRTCPeerConnectionSignalingStateClosed WebRTCPeerConnectionSignalingState = 5
)
