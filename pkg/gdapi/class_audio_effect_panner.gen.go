// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPanner struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPanner) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPanner) BaseClass() string {
  return "AudioEffectPanner"
}
