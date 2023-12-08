// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerSynchronizer struct {
  obj gdc.ObjectPtr
}

func createMultiplayerSynchronizer(obj gdc.ObjectPtr) *MultiplayerSynchronizer {
  return &MultiplayerSynchronizer{
    obj: obj,
  }
}

func (me *MultiplayerSynchronizer) BaseClass() string {
  return "MultiplayerSynchronizer"
}
