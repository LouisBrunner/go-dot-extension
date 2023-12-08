// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionSDF3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionSDF3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionSDF3D) BaseClass() string {
  return "GPUParticlesCollisionSDF3D"
}

type GPUParticlesCollisionSDF3DResolution int
const (
  GPUParticlesCollisionSDF3DResolution16 GPUParticlesCollisionSDF3DResolution = 0
  GPUParticlesCollisionSDF3DResolution32 GPUParticlesCollisionSDF3DResolution = 1
  GPUParticlesCollisionSDF3DResolution64 GPUParticlesCollisionSDF3DResolution = 2
  GPUParticlesCollisionSDF3DResolution128 GPUParticlesCollisionSDF3DResolution = 3
  GPUParticlesCollisionSDF3DResolution256 GPUParticlesCollisionSDF3DResolution = 4
  GPUParticlesCollisionSDF3DResolution512 GPUParticlesCollisionSDF3DResolution = 5
  GPUParticlesCollisionSDF3DResolutionMax GPUParticlesCollisionSDF3DResolution = 6
)
