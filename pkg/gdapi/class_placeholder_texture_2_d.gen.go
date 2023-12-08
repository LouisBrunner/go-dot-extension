// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTexture2D struct {
  obj gdc.ObjectPtr
}

func createPlaceholderTexture2D(obj gdc.ObjectPtr) *PlaceholderTexture2D {
  return &PlaceholderTexture2D{
    obj: obj,
  }
}

func (me *PlaceholderTexture2D) BaseClass() string {
  return "PlaceholderTexture2D"
}
