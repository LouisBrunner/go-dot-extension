// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLimiter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectLimiter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectLimiter) BaseClass() string {
  return "AudioEffectLimiter"
}
