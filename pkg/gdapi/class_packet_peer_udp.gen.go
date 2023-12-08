// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerUDP struct {
  obj gdc.ObjectPtr
}

func createPacketPeerUDP(obj gdc.ObjectPtr) *PacketPeerUDP {
  return &PacketPeerUDP{
    obj: obj,
  }
}

func (me *PacketPeerUDP) BaseClass() string {
  return "PacketPeerUDP"
}
