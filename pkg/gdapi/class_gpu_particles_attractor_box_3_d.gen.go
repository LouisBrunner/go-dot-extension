// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorBox3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractorBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractorBox3D) BaseClass() string {
  return "GPUParticlesAttractorBox3D"
}
