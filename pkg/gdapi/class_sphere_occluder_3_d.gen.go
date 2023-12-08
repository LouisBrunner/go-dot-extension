// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereOccluder3D struct {
  obj gdc.ObjectPtr
}

func createSphereOccluder3D(obj gdc.ObjectPtr) *SphereOccluder3D {
  return &SphereOccluder3D{
    obj: obj,
  }
}

func (me *SphereOccluder3D) BaseClass() string {
  return "SphereOccluder3D"
}
