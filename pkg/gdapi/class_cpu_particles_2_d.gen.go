// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CPUParticles2D struct {
  obj gdc.ObjectPtr
}

func (me *CPUParticles2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CPUParticles2D) BaseClass() string {
  return "CPUParticles2D"
}

type CPUParticles2DDrawOrder int
const (
  CPUParticles2DDrawOrderIndex CPUParticles2DDrawOrder = 0
  CPUParticles2DDrawOrderLifetime CPUParticles2DDrawOrder = 1
)

type CPUParticles2DParameter int
const (
  CPUParticles2DParamInitialLinearVelocity CPUParticles2DParameter = 0
  CPUParticles2DParamAngularVelocity CPUParticles2DParameter = 1
  CPUParticles2DParamOrbitVelocity CPUParticles2DParameter = 2
  CPUParticles2DParamLinearAccel CPUParticles2DParameter = 3
  CPUParticles2DParamRadialAccel CPUParticles2DParameter = 4
  CPUParticles2DParamTangentialAccel CPUParticles2DParameter = 5
  CPUParticles2DParamDamping CPUParticles2DParameter = 6
  CPUParticles2DParamAngle CPUParticles2DParameter = 7
  CPUParticles2DParamScale CPUParticles2DParameter = 8
  CPUParticles2DParamHueVariation CPUParticles2DParameter = 9
  CPUParticles2DParamAnimSpeed CPUParticles2DParameter = 10
  CPUParticles2DParamAnimOffset CPUParticles2DParameter = 11
  CPUParticles2DParamMax CPUParticles2DParameter = 12
)

type CPUParticles2DParticleFlags int
const (
  CPUParticles2DParticleFlagAlignYToVelocity CPUParticles2DParticleFlags = 0
  CPUParticles2DParticleFlagRotateY CPUParticles2DParticleFlags = 1
  CPUParticles2DParticleFlagDisableZ CPUParticles2DParticleFlags = 2
  CPUParticles2DParticleFlagMax CPUParticles2DParticleFlags = 3
)

type CPUParticles2DEmissionShape int
const (
  CPUParticles2DEmissionShapePoint CPUParticles2DEmissionShape = 0
  CPUParticles2DEmissionShapeSphere CPUParticles2DEmissionShape = 1
  CPUParticles2DEmissionShapeSphereSurface CPUParticles2DEmissionShape = 2
  CPUParticles2DEmissionShapeRectangle CPUParticles2DEmissionShape = 3
  CPUParticles2DEmissionShapePoints CPUParticles2DEmissionShape = 4
  CPUParticles2DEmissionShapeDirectedPoints CPUParticles2DEmissionShape = 5
  CPUParticles2DEmissionShapeMax CPUParticles2DEmissionShape = 6
)
