// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractor3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractor3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractor3D) BaseClass() string {
  return "GPUParticlesAttractor3D"
}
