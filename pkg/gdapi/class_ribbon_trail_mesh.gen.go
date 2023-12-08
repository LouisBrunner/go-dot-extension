// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RibbonTrailMesh struct {
  obj gdc.ObjectPtr
}

func createRibbonTrailMesh(obj gdc.ObjectPtr) *RibbonTrailMesh {
  return &RibbonTrailMesh{
    obj: obj,
  }
}

func (me *RibbonTrailMesh) BaseClass() string {
  return "RibbonTrailMesh"
}
