// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamMicrophone struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamMicrophone) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamMicrophone) BaseClass() string {
  return "AudioStreamMicrophone"
}



// Enums

func (me *AudioStreamMicrophone) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamMicrophone) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamMicrophone) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
