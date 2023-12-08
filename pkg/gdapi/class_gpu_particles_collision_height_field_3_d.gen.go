// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionHeightField3D struct {
  obj gdc.ObjectPtr
}

func createGPUParticlesCollisionHeightField3D(obj gdc.ObjectPtr) *GPUParticlesCollisionHeightField3D {
  return &GPUParticlesCollisionHeightField3D{
    obj: obj,
  }
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
  return "GPUParticlesCollisionHeightField3D"
}
