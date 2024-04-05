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

type ptrsForStreamPeerExtensionList struct {
  fnXGetData gdc.MethodBindPtr
  fnXGetPartialData gdc.MethodBindPtr
  fnXPutData gdc.MethodBindPtr
  fnXPutPartialData gdc.MethodBindPtr
  fnXGetAvailableBytes gdc.MethodBindPtr
}

var ptrsForStreamPeerExtension ptrsForStreamPeerExtensionList

func initStreamPeerExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("StreamPeerExtension")
  defer className.Destroy()
}

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
