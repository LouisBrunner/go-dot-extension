// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioBusLayout struct {
  Resource
}

func (me *AudioBusLayout) BaseClass() string {
  return "AudioBusLayout"
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
