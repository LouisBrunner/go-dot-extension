// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HeightMapShape3D struct {
  obj gdc.ObjectPtr
}

func createHeightMapShape3D(obj gdc.ObjectPtr) *HeightMapShape3D {
  return &HeightMapShape3D{
    obj: obj,
  }
}

func (me *HeightMapShape3D) BaseClass() string {
  return "HeightMapShape3D"
}
