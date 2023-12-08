// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTexture3D struct {
  obj gdc.ObjectPtr
}

func createPlaceholderTexture3D(obj gdc.ObjectPtr) *PlaceholderTexture3D {
  return &PlaceholderTexture3D{
    obj: obj,
  }
}

func (me *PlaceholderTexture3D) BaseClass() string {
  return "PlaceholderTexture3D"
}
