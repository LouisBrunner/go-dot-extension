// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionSphere3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesCollisionSphere3D(obj gdc.ObjectPtr) *GPUParticlesCollisionSphere3D {
  return &GPUParticlesCollisionSphere3D{
    obj: obj,
  }
}

func (me *GPUParticlesCollisionSphere3D) BaseClass() string {
  return "GPUParticlesCollisionSphere3D"
}
