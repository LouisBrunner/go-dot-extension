// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTexture2DArray struct {
  obj gdc.ObjectPtr
}

func createPlaceholderTexture2DArray(obj gdc.ObjectPtr) *PlaceholderTexture2DArray {
  return &PlaceholderTexture2DArray{
    obj: obj,
  }
}

func (me *PlaceholderTexture2DArray) BaseClass() string {
  return "PlaceholderTexture2DArray"
}
