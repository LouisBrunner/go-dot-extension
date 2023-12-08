// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerExtension struct {
  obj gdc.ObjectPtr
}

func createPacketPeerExtension(obj gdc.ObjectPtr) *PacketPeerExtension {
  return &PacketPeerExtension{
    obj: obj,
  }
}

func (me *PacketPeerExtension) BaseClass() string {
  return "PacketPeerExtension"
}
