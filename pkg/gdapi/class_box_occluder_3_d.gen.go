// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxOccluder3D struct {
  obj gdc.ObjectPtr
}

func createBoxOccluder3D(obj gdc.ObjectPtr) *BoxOccluder3D {
  return &BoxOccluder3D{
    obj: obj,
  }
}

func (me *BoxOccluder3D) BaseClass() string {
  return "BoxOccluder3D"
}
