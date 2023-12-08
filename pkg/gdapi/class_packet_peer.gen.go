// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeer struct {
  obj gdc.ObjectPtr
}

func createPacketPeer(obj gdc.ObjectPtr) *PacketPeer {
  return &PacketPeer{
    obj: obj,
  }
}

func (me *PacketPeer) BaseClass() string {
  return "PacketPeer"
}
