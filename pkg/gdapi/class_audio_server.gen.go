// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioServer struct {
  obj gdc.ObjectPtr
}

func createAudioServer(obj gdc.ObjectPtr) *AudioServer {
  return &AudioServer{
    obj: obj,
  }
}

func (me *AudioServer) BaseClass() string {
  return "AudioServer"
}
