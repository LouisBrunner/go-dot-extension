// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OggPacketSequencePlayback struct {
  obj gdc.ObjectPtr
}

func (me *OggPacketSequencePlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OggPacketSequencePlayback) BaseClass() string {
  return "OggPacketSequencePlayback"
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

// Properties