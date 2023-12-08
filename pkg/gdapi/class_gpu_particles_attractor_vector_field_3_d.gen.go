// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorVectorField3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesAttractorVectorField3D(obj gdc.ObjectPtr) *GPUParticlesAttractorVectorField3D {
  return &GPUParticlesAttractorVectorField3D{
    obj: obj,
  }
}

func (me *GPUParticlesAttractorVectorField3D) BaseClass() string {
  return "GPUParticlesAttractorVectorField3D"
}
