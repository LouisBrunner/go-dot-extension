// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AtlasTexture struct {
  obj gdc.ObjectPtr
}

func createAtlasTexture(obj gdc.ObjectPtr) *AtlasTexture {
  return &AtlasTexture{
    obj: obj,
  }
}

func (me *AtlasTexture) BaseClass() string {
  return "AtlasTexture"
}
