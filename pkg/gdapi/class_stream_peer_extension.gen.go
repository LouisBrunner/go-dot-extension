// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type StreamPeerExtension struct {
  StreamPeer
}

func (me *StreamPeerExtension) BaseClass() string {
  return "StreamPeerExtension"
}

func NewStreamPeerExtension() *StreamPeerExtension {
  str := StringNameFromStr("StreamPeerExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StreamPeerExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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
