// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OggPacketSequencePlayback struct {
  obj gdc.ObjectPtr
}

func createOggPacketSequencePlayback(obj gdc.ObjectPtr) *OggPacketSequencePlayback {
  return &OggPacketSequencePlayback{
    obj: obj,
  }
}

func (me *OggPacketSequencePlayback) BaseClass() string {
  return "OggPacketSequencePlayback"
}
