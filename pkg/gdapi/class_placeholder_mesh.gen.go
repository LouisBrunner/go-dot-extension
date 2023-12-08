// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderMesh struct {
  obj gdc.ObjectPtr
}

func createPlaceholderMesh(obj gdc.ObjectPtr) *PlaceholderMesh {
  return &PlaceholderMesh{
    obj: obj,
  }
}

func (me *PlaceholderMesh) BaseClass() string {
  return "PlaceholderMesh"
}
