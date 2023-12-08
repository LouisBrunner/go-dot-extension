// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ParticleProcessMaterial struct {
  obj gdc.ObjectPtr
}

func createParticleProcessMaterial(obj gdc.ObjectPtr) *ParticleProcessMaterial {
  return &ParticleProcessMaterial{
    obj: obj,
  }
}

func (me *ParticleProcessMaterial) BaseClass() string {
  return "ParticleProcessMaterial"
}
