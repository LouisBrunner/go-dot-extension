// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamPlayback struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlayback) BaseClass() string {
  return "AudioStreamPlayback"
}



// Enums

func (me *AudioStreamPlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
