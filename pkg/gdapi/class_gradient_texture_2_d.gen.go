// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GradientTexture2D struct {
  obj gdc.ObjectPtr
}

func createGradientTexture2D(obj gdc.ObjectPtr) *GradientTexture2D {
  return &GradientTexture2D{
    obj: obj,
  }
}

func (me *GradientTexture2D) BaseClass() string {
  return "GradientTexture2D"
}
