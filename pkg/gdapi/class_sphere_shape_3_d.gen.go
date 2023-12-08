// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereShape3D struct {
  obj gdc.ObjectPtr
}

func createSphereShape3D(obj gdc.ObjectPtr) *SphereShape3D {
  return &SphereShape3D{
    obj: obj,
  }
}

func (me *SphereShape3D) BaseClass() string {
  return "SphereShape3D"
}
