// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GradientTexture1D struct {
  obj gdc.ObjectPtr
}

func createGradientTexture1D(obj gdc.ObjectPtr) *GradientTexture1D {
  return &GradientTexture1D{
    obj: obj,
  }
}

func (me *GradientTexture1D) BaseClass() string {
  return "GradientTexture1D"
}
