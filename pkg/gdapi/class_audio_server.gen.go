// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioServer struct {
  obj gdc.ObjectPtr
}

func (me *AudioServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioServer) BaseClass() string {
  return "AudioServer"
}

type AudioServerSpeakerMode int
const (
  AudioServerSpeakerModeStereo AudioServerSpeakerMode = 0
  AudioServerSpeakerSurround31 AudioServerSpeakerMode = 1
  AudioServerSpeakerSurround51 AudioServerSpeakerMode = 2
  AudioServerSpeakerSurround71 AudioServerSpeakerMode = 3
)
