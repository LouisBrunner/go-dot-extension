// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles2D struct {
  obj gdc.ObjectPtr
}

func createGPUParticles2D(obj gdc.ObjectPtr) *GPUParticles2D {
  return &GPUParticles2D{
    obj: obj,
  }
}

func (me *GPUParticles2D) BaseClass() string {
  return "GPUParticles2D"
}
