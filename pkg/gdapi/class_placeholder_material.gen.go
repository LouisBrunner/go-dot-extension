// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderMaterial struct {
  obj gdc.ObjectPtr
}

func createPlaceholderMaterial(obj gdc.ObjectPtr) *PlaceholderMaterial {
  return &PlaceholderMaterial{
    obj: obj,
  }
}

func (me *PlaceholderMaterial) BaseClass() string {
  return "PlaceholderMaterial"
}
