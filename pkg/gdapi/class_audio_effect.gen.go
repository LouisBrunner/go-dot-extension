// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffect struct {
  Resource
}

func (me *AudioEffect) BaseClass() string {
  return "AudioEffect"
}

func NewAudioEffect() *AudioEffect {
  str := StringNameFromStr("AudioEffect") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffect{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
