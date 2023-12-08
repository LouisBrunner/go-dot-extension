// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SurfaceTool struct {
  obj gdc.ObjectPtr
}

func createSurfaceTool(obj gdc.ObjectPtr) *SurfaceTool {
  return &SurfaceTool{
    obj: obj,
  }
}

func (me *SurfaceTool) BaseClass() string {
  return "SurfaceTool"
}
