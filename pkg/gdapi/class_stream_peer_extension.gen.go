// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StreamPeerExtension struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerExtension) BaseClass() string {
  return "StreamPeerExtension"
}



// Enums

func (me *StreamPeerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
