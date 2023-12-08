// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorSphere3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesAttractorSphere3D(obj gdc.ObjectPtr) *GPUParticlesAttractorSphere3D {
  return &GPUParticlesAttractorSphere3D{
    obj: obj,
  }
}

func (me *GPUParticlesAttractorSphere3D) BaseClass() string {
  return "GPUParticlesAttractorSphere3D"
}
