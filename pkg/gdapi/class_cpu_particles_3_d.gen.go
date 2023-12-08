// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CPUParticles3D struct {
  obj gdc.ObjectPtr
}

func (me *CPUParticles3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CPUParticles3D) BaseClass() string {
  return "CPUParticles3D"
}

type CPUParticles3DDrawOrder int
const (
  CPUParticles3DDrawOrderIndex CPUParticles3DDrawOrder = 0
  CPUParticles3DDrawOrderLifetime CPUParticles3DDrawOrder = 1
  CPUParticles3DDrawOrderViewDepth CPUParticles3DDrawOrder = 2
)

type CPUParticles3DParameter int
const (
  CPUParticles3DParamInitialLinearVelocity CPUParticles3DParameter = 0
  CPUParticles3DParamAngularVelocity CPUParticles3DParameter = 1
  CPUParticles3DParamOrbitVelocity CPUParticles3DParameter = 2
  CPUParticles3DParamLinearAccel CPUParticles3DParameter = 3
  CPUParticles3DParamRadialAccel CPUParticles3DParameter = 4
  CPUParticles3DParamTangentialAccel CPUParticles3DParameter = 5
  CPUParticles3DParamDamping CPUParticles3DParameter = 6
  CPUParticles3DParamAngle CPUParticles3DParameter = 7
  CPUParticles3DParamScale CPUParticles3DParameter = 8
  CPUParticles3DParamHueVariation CPUParticles3DParameter = 9
  CPUParticles3DParamAnimSpeed CPUParticles3DParameter = 10
  CPUParticles3DParamAnimOffset CPUParticles3DParameter = 11
  CPUParticles3DParamMax CPUParticles3DParameter = 12
)

type CPUParticles3DParticleFlags int
const (
  CPUParticles3DParticleFlagAlignYToVelocity CPUParticles3DParticleFlags = 0
  CPUParticles3DParticleFlagRotateY CPUParticles3DParticleFlags = 1
  CPUParticles3DParticleFlagDisableZ CPUParticles3DParticleFlags = 2
  CPUParticles3DParticleFlagMax CPUParticles3DParticleFlags = 3
)

type CPUParticles3DEmissionShape int
const (
  CPUParticles3DEmissionShapePoint CPUParticles3DEmissionShape = 0
  CPUParticles3DEmissionShapeSphere CPUParticles3DEmissionShape = 1
  CPUParticles3DEmissionShapeSphereSurface CPUParticles3DEmissionShape = 2
  CPUParticles3DEmissionShapeBox CPUParticles3DEmissionShape = 3
  CPUParticles3DEmissionShapePoints CPUParticles3DEmissionShape = 4
  CPUParticles3DEmissionShapeDirectedPoints CPUParticles3DEmissionShape = 5
  CPUParticles3DEmissionShapeRing CPUParticles3DEmissionShape = 6
  CPUParticles3DEmissionShapeMax CPUParticles3DEmissionShape = 7
)
