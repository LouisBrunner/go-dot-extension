// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TubeTrailMesh struct {
  obj gdc.ObjectPtr
}

func createTubeTrailMesh(obj gdc.ObjectPtr) *TubeTrailMesh {
  return &TubeTrailMesh{
    obj: obj,
  }
}

func (me *TubeTrailMesh) BaseClass() string {
  return "TubeTrailMesh"
}
