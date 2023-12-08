// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerDTLS struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeerDTLS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PacketPeerDTLS) BaseClass() string {
  return "PacketPeerDTLS"
}

type PacketPeerDTLSStatus int
const (
  PacketPeerDTLSStatusDisconnected PacketPeerDTLSStatus = 0
  PacketPeerDTLSStatusHandshaking PacketPeerDTLSStatus = 1
  PacketPeerDTLSStatusConnected PacketPeerDTLSStatus = 2
  PacketPeerDTLSStatusError PacketPeerDTLSStatus = 3
  PacketPeerDTLSStatusErrorHostnameMismatch PacketPeerDTLSStatus = 4
)
