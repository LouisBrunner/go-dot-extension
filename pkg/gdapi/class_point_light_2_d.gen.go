// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointLight2D struct {
  obj gdc.ObjectPtr
}

func createPointLight2D(obj gdc.ObjectPtr) *PointLight2D {
  return &PointLight2D{
    obj: obj,
  }
}

func (me *PointLight2D) BaseClass() string {
  return "PointLight2D"
}
