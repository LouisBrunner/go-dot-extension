// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Occluder3D struct {
  obj gdc.ObjectPtr
}

func createOccluder3D(obj gdc.ObjectPtr) *Occluder3D {
  return &Occluder3D{
    obj: obj,
  }
}

func (me *Occluder3D) BaseClass() string {
  return "Occluder3D"
}
