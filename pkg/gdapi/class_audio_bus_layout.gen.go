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

type AudioBusLayout struct {
  Resource
}

func (me *AudioBusLayout) BaseClass() string {
  return "AudioBusLayout"
}

func NewAudioBusLayout() *AudioBusLayout {
  str := StringNameFromStr("AudioBusLayout") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioBusLayout{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioBusLayout) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioBusLayout) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioBusLayout) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
