// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetPacketPeer struct {
  obj gdc.ObjectPtr
}

func createENetPacketPeer(obj gdc.ObjectPtr) *ENetPacketPeer {
  return &ENetPacketPeer{
    obj: obj,
  }
}

func (me *ENetPacketPeer) BaseClass() string {
  return "ENetPacketPeer"
}
