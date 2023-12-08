// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SeparationRayShape2D struct {
  obj gdc.ObjectPtr
}

func createSeparationRayShape2D(obj gdc.ObjectPtr) *SeparationRayShape2D {
  return &SeparationRayShape2D{
    obj: obj,
  }
}

func (me *SeparationRayShape2D) BaseClass() string {
  return "SeparationRayShape2D"
}
