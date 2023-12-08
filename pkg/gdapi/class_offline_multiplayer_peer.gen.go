// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OfflineMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func createOfflineMultiplayerPeer(obj gdc.ObjectPtr) *OfflineMultiplayerPeer {
  return &OfflineMultiplayerPeer{
    obj: obj,
  }
}

func (me *OfflineMultiplayerPeer) BaseClass() string {
  return "OfflineMultiplayerPeer"
}
