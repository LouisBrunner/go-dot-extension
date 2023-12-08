// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticles3D(obj gdc.ObjectPtr) *GPUParticles3D {
  return &GPUParticles3D{
    obj: obj,
  }
}

func (me *GPUParticles3D) BaseClass() string {
  return "GPUParticles3D"
}
