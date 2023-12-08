// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionBox3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesCollisionBox3D(obj gdc.ObjectPtr) *GPUParticlesCollisionBox3D {
  return &GPUParticlesCollisionBox3D{
    obj: obj,
  }
}

func (me *GPUParticlesCollisionBox3D) BaseClass() string {
  return "GPUParticlesCollisionBox3D"
}
