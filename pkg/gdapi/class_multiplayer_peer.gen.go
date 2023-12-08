// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func createMultiplayerPeer(obj gdc.ObjectPtr) *MultiplayerPeer {
  return &MultiplayerPeer{
    obj: obj,
  }
}

func (me *MultiplayerPeer) BaseClass() string {
  return "MultiplayerPeer"
}
