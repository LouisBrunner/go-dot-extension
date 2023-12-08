// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CPUParticles2D struct {
  obj gdc.ObjectPtr
}

func createCPUParticles2D(obj gdc.ObjectPtr) *CPUParticles2D {
  return &CPUParticles2D{
    obj: obj,
  }
}

func (me *CPUParticles2D) BaseClass() string {
  return "CPUParticles2D"
}
