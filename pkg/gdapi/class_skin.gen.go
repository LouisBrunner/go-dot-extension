// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skin struct {
  obj gdc.ObjectPtr
}

func createSkin(obj gdc.ObjectPtr) *Skin {
  return &Skin{
    obj: obj,
  }
}

func (me *Skin) BaseClass() string {
  return "Skin"
}
