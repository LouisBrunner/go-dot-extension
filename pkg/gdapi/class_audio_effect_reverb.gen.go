// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectReverb struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectReverb) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectReverb) BaseClass() string {
  return "AudioEffectReverb"
}
