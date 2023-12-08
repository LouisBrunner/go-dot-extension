// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionHeightField3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionHeightField3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
  return "GPUParticlesCollisionHeightField3D"
}

type GPUParticlesCollisionHeightField3DResolution int
const (
  GPUParticlesCollisionHeightField3DResolution256 GPUParticlesCollisionHeightField3DResolution = 0
  GPUParticlesCollisionHeightField3DResolution512 GPUParticlesCollisionHeightField3DResolution = 1
  GPUParticlesCollisionHeightField3DResolution1024 GPUParticlesCollisionHeightField3DResolution = 2
  GPUParticlesCollisionHeightField3DResolution2048 GPUParticlesCollisionHeightField3DResolution = 3
  GPUParticlesCollisionHeightField3DResolution4096 GPUParticlesCollisionHeightField3DResolution = 4
  GPUParticlesCollisionHeightField3DResolution8192 GPUParticlesCollisionHeightField3DResolution = 5
  GPUParticlesCollisionHeightField3DResolutionMax GPUParticlesCollisionHeightField3DResolution = 6
)

type GPUParticlesCollisionHeightField3DUpdateMode int
const (
  GPUParticlesCollisionHeightField3DUpdateModeWhenMoved GPUParticlesCollisionHeightField3DUpdateMode = 0
  GPUParticlesCollisionHeightField3DUpdateModeAlways GPUParticlesCollisionHeightField3DUpdateMode = 1
)
