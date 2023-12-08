// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerSpawner struct {
  obj gdc.ObjectPtr
}

func createMultiplayerSpawner(obj gdc.ObjectPtr) *MultiplayerSpawner {
  return &MultiplayerSpawner{
    obj: obj,
  }
}

func (me *MultiplayerSpawner) BaseClass() string {
  return "MultiplayerSpawner"
}
