// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CPUParticles3D struct {
  obj gdc.ObjectPtr
}

func createCPUParticles3D(obj gdc.ObjectPtr) *CPUParticles3D {
  return &CPUParticles3D{
    obj: obj,
  }
}

func (me *CPUParticles3D) BaseClass() string {
  return "CPUParticles3D"
}
