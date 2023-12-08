// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpotLight3D struct {
  obj gdc.ObjectPtr
}

func createSpotLight3D(obj gdc.ObjectPtr) *SpotLight3D {
  return &SpotLight3D{
    obj: obj,
  }
}

func (me *SpotLight3D) BaseClass() string {
  return "SpotLight3D"
}
