// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerAPIExtension struct {
  obj gdc.ObjectPtr
}

func createMultiplayerAPIExtension(obj gdc.ObjectPtr) *MultiplayerAPIExtension {
  return &MultiplayerAPIExtension{
    obj: obj,
  }
}

func (me *MultiplayerAPIExtension) BaseClass() string {
  return "MultiplayerAPIExtension"
}
