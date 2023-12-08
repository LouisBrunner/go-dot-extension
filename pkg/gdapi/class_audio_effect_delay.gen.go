// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectDelay struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectDelay) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectDelay) BaseClass() string {
  return "AudioEffectDelay"
}
