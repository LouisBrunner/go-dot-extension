// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerDTLS struct {
  obj gdc.ObjectPtr
}

func createPacketPeerDTLS(obj gdc.ObjectPtr) *PacketPeerDTLS {
  return &PacketPeerDTLS{
    obj: obj,
  }
}

func (me *PacketPeerDTLS) BaseClass() string {
  return "PacketPeerDTLS"
}
