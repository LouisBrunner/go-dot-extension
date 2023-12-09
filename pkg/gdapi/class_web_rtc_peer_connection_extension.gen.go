// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebRTCPeerConnectionExtension struct {
  obj gdc.ObjectPtr
}

func (me *WebRTCPeerConnectionExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebRTCPeerConnectionExtension) BaseClass() string {
  return "WebRTCPeerConnectionExtension"
}



// Enums

func (me *WebRTCPeerConnectionExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebRTCPeerConnectionExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCPeerConnectionExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WebRTCPeerConnectionExtension) XGetConnectionState()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XGetGatheringState()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XGetSignalingState()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XInitialize(p_config Dictionary, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XCreateDataChannel(p_label String, p_config Dictionary, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XCreateOffer()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XSetRemoteDescription(p_type String, p_sdp String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XSetLocalDescription(p_type String, p_sdp String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XAddIceCandidate(p_sdp_mid_name String, p_sdp_mline_index int, p_sdp_name String, )  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XPoll()  {
  panic("TODO: implement")
}

func  (me *WebRTCPeerConnectionExtension) XClose()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
