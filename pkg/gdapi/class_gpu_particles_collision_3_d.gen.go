// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollision3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesCollision3D(obj gdc.ObjectPtr) *GPUParticlesCollision3D {
  return &GPUParticlesCollision3D{
    obj: obj,
  }
}

func (me *GPUParticlesCollision3D) BaseClass() string {
  return "GPUParticlesCollision3D"
}
