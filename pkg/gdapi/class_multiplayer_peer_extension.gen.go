// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerPeerExtension struct {
  obj gdc.ObjectPtr
}

func createMultiplayerPeerExtension(obj gdc.ObjectPtr) *MultiplayerPeerExtension {
  return &MultiplayerPeerExtension{
    obj: obj,
  }
}

func (me *MultiplayerPeerExtension) BaseClass() string {
  return "MultiplayerPeerExtension"
}
