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

type OggPacketSequencePlayback struct {
  RefCounted
}

func (me *OggPacketSequencePlayback) BaseClass() string {
  return "OggPacketSequencePlayback"
}

func NewOggPacketSequencePlayback() *OggPacketSequencePlayback {
  str := StringNameFromStr("OggPacketSequencePlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OggPacketSequencePlayback{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OggPacketSequencePlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OggPacketSequencePlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OggPacketSequencePlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
