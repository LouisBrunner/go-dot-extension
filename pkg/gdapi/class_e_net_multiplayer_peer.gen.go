// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func createENetMultiplayerPeer(obj gdc.ObjectPtr) *ENetMultiplayerPeer {
  return &ENetMultiplayerPeer{
    obj: obj,
  }
}

func (me *ENetMultiplayerPeer) BaseClass() string {
  return "ENetMultiplayerPeer"
}
