// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerStream struct {
  obj gdc.ObjectPtr
}

func createPacketPeerStream(obj gdc.ObjectPtr) *PacketPeerStream {
  return &PacketPeerStream{
    obj: obj,
  }
}

func (me *PacketPeerStream) BaseClass() string {
  return "PacketPeerStream"
}
