// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectFilter) BaseClass() string {
  return "AudioEffectFilter"
}

type AudioEffectFilterFilterDB int
const (
  AudioEffectFilterFilter6Db AudioEffectFilterFilterDB = 0
  AudioEffectFilterFilter12Db AudioEffectFilterFilterDB = 1
  AudioEffectFilterFilter18Db AudioEffectFilterFilterDB = 2
  AudioEffectFilterFilter24Db AudioEffectFilterFilterDB = 3
)
