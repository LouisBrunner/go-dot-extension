// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SeparationRayShape3D struct {
  obj gdc.ObjectPtr
}

func createSeparationRayShape3D(obj gdc.ObjectPtr) *SeparationRayShape3D {
  return &SeparationRayShape3D{
    obj: obj,
  }
}

func (me *SeparationRayShape3D) BaseClass() string {
  return "SeparationRayShape3D"
}
