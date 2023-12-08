// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NoiseTexture3D struct {
  obj gdc.ObjectPtr
}

func createNoiseTexture3D(obj gdc.ObjectPtr) *NoiseTexture3D {
  return &NoiseTexture3D{
    obj: obj,
  }
}

func (me *NoiseTexture3D) BaseClass() string {
  return "NoiseTexture3D"
}
