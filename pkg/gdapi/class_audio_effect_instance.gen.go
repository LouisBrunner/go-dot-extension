// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectInstance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectInstance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectInstance) BaseClass() string {
  return "AudioEffectInstance"
}
