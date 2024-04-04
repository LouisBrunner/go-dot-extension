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

type WebRTCPeerConnectionExtension struct {
  WebRTCPeerConnection
}

func (me *WebRTCPeerConnectionExtension) BaseClass() string {
  return "WebRTCPeerConnectionExtension"
}

func NewWebRTCPeerConnectionExtension() *WebRTCPeerConnectionExtension {
  str := StringNameFromStr("WebRTCPeerConnectionExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WebRTCPeerConnectionExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
