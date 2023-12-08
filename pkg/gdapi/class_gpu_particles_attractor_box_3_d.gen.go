// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorBox3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesAttractorBox3D(obj gdc.ObjectPtr) *GPUParticlesAttractorBox3D {
  return &GPUParticlesAttractorBox3D{
    obj: obj,
  }
}

func (me *GPUParticlesAttractorBox3D) BaseClass() string {
  return "GPUParticlesAttractorBox3D"
}
