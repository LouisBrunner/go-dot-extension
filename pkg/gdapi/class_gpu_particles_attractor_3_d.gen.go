// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractor3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesAttractor3D(obj gdc.ObjectPtr) *GPUParticlesAttractor3D {
  return &GPUParticlesAttractor3D{
    obj: obj,
  }
}

func (me *GPUParticlesAttractor3D) BaseClass() string {
  return "GPUParticlesAttractor3D"
}
