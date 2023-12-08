// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerAPI struct {
  obj gdc.ObjectPtr
}

func createMultiplayerAPI(obj gdc.ObjectPtr) *MultiplayerAPI {
  return &MultiplayerAPI{
    obj: obj,
  }
}

func (me *MultiplayerAPI) BaseClass() string {
  return "MultiplayerAPI"
}
