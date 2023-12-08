// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports
  "unsafe"






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

func  (me *AudioEffectInstance) XProcess(src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectInstance) XProcessSilence() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
