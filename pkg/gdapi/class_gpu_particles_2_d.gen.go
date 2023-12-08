// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles2D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticles2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticles2D) BaseClass() string {
  return "GPUParticles2D"
}

type GPUParticles2DDrawOrder int
const (
  GPUParticles2DDrawOrderIndex GPUParticles2DDrawOrder = 0
  GPUParticles2DDrawOrderLifetime GPUParticles2DDrawOrder = 1
  GPUParticles2DDrawOrderReverseLifetime GPUParticles2DDrawOrder = 2
)

type GPUParticles2DEmitFlags int
const (
  GPUParticles2DEmitFlagPosition GPUParticles2DEmitFlags = 1
  GPUParticles2DEmitFlagRotationScale GPUParticles2DEmitFlags = 2
  GPUParticles2DEmitFlagVelocity GPUParticles2DEmitFlags = 4
  GPUParticles2DEmitFlagColor GPUParticles2DEmitFlags = 8
  GPUParticles2DEmitFlagCustom GPUParticles2DEmitFlags = 16
)
