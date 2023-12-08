// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionSphere3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionSphere3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionSphere3D) BaseClass() string {
  return "GPUParticlesCollisionSphere3D"
}
