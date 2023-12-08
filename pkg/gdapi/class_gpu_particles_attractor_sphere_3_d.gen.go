// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorSphere3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractorSphere3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractorSphere3D) BaseClass() string {
  return "GPUParticlesAttractorSphere3D"
}
