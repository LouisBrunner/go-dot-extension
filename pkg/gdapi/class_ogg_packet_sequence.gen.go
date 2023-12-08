// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OggPacketSequence struct {
  obj gdc.ObjectPtr
}

func createOggPacketSequence(obj gdc.ObjectPtr) *OggPacketSequence {
  return &OggPacketSequence{
    obj: obj,
  }
}

func (me *OggPacketSequence) BaseClass() string {
  return "OggPacketSequence"
}
