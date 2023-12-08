// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionSDF3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesCollisionSDF3D(obj gdc.ObjectPtr) *GPUParticlesCollisionSDF3D {
  return &GPUParticlesCollisionSDF3D{
    obj: obj,
  }
}

func (me *GPUParticlesCollisionSDF3D) BaseClass() string {
  return "GPUParticlesCollisionSDF3D"
}
