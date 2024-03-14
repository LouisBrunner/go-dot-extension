// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WebRTCDataChannelExtension struct {
  WebRTCDataChannel
}

func (me *WebRTCDataChannelExtension) BaseClass() string {
  return "WebRTCDataChannelExtension"
}



// Enums

func (me *WebRTCDataChannelExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebRTCDataChannelExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCDataChannelExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
