// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticles3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticles3D) BaseClass() string {
  return "GPUParticles3D"
}

const (
  GPUParticles3DMAX_DRAW_PASSES = 4
)

type GPUParticles3DDrawOrder int
const (
  GPUParticles3DDrawOrderIndex GPUParticles3DDrawOrder = 0
  GPUParticles3DDrawOrderLifetime GPUParticles3DDrawOrder = 1
  GPUParticles3DDrawOrderReverseLifetime GPUParticles3DDrawOrder = 2
  GPUParticles3DDrawOrderViewDepth GPUParticles3DDrawOrder = 3
)

type GPUParticles3DEmitFlags int
const (
  GPUParticles3DEmitFlagPosition GPUParticles3DEmitFlags = 1
  GPUParticles3DEmitFlagRotationScale GPUParticles3DEmitFlags = 2
  GPUParticles3DEmitFlagVelocity GPUParticles3DEmitFlags = 4
  GPUParticles3DEmitFlagColor GPUParticles3DEmitFlags = 8
  GPUParticles3DEmitFlagCustom GPUParticles3DEmitFlags = 16
)

type GPUParticles3DTransformAlign int
const (
  GPUParticles3DTransformAlignDisabled GPUParticles3DTransformAlign = 0
  GPUParticles3DTransformAlignZBillboard GPUParticles3DTransformAlign = 1
  GPUParticles3DTransformAlignYToVelocity GPUParticles3DTransformAlign = 2
  GPUParticles3DTransformAlignZBillboardYToVelocity GPUParticles3DTransformAlign = 3
)
