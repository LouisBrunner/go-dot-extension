// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NoiseTexture2D struct {
  obj gdc.ObjectPtr
}

func createNoiseTexture2D(obj gdc.ObjectPtr) *NoiseTexture2D {
  return &NoiseTexture2D{
    obj: obj,
  }
}

func (me *NoiseTexture2D) BaseClass() string {
  return "NoiseTexture2D"
}
