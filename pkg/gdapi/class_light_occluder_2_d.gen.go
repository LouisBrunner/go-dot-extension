// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightOccluder2D struct {
  obj gdc.ObjectPtr
}

func createLightOccluder2D(obj gdc.ObjectPtr) *LightOccluder2D {
  return &LightOccluder2D{
    obj: obj,
  }
}

func (me *LightOccluder2D) BaseClass() string {
  return "LightOccluder2D"
}
