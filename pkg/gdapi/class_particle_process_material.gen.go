// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ParticleProcessMaterial struct {
  obj gdc.ObjectPtr
}

func (me *ParticleProcessMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ParticleProcessMaterial) BaseClass() string {
  return "ParticleProcessMaterial"
}

type ParticleProcessMaterialParameter int
const (
  ParticleProcessMaterialParamInitialLinearVelocity ParticleProcessMaterialParameter = 0
  ParticleProcessMaterialParamAngularVelocity ParticleProcessMaterialParameter = 1
  ParticleProcessMaterialParamOrbitVelocity ParticleProcessMaterialParameter = 2
  ParticleProcessMaterialParamLinearAccel ParticleProcessMaterialParameter = 3
  ParticleProcessMaterialParamRadialAccel ParticleProcessMaterialParameter = 4
  ParticleProcessMaterialParamTangentialAccel ParticleProcessMaterialParameter = 5
  ParticleProcessMaterialParamDamping ParticleProcessMaterialParameter = 6
  ParticleProcessMaterialParamAngle ParticleProcessMaterialParameter = 7
  ParticleProcessMaterialParamScale ParticleProcessMaterialParameter = 8
  ParticleProcessMaterialParamHueVariation ParticleProcessMaterialParameter = 9
  ParticleProcessMaterialParamAnimSpeed ParticleProcessMaterialParameter = 10
  ParticleProcessMaterialParamAnimOffset ParticleProcessMaterialParameter = 11
  ParticleProcessMaterialParamMax ParticleProcessMaterialParameter = 15
  ParticleProcessMaterialParamTurbVelInfluence ParticleProcessMaterialParameter = 13
  ParticleProcessMaterialParamTurbInitDisplacement ParticleProcessMaterialParameter = 14
  ParticleProcessMaterialParamTurbInfluenceOverLife ParticleProcessMaterialParameter = 12
)

type ParticleProcessMaterialParticleFlags int
const (
  ParticleProcessMaterialParticleFlagAlignYToVelocity ParticleProcessMaterialParticleFlags = 0
  ParticleProcessMaterialParticleFlagRotateY ParticleProcessMaterialParticleFlags = 1
  ParticleProcessMaterialParticleFlagDisableZ ParticleProcessMaterialParticleFlags = 2
  ParticleProcessMaterialParticleFlagMax ParticleProcessMaterialParticleFlags = 3
)

type ParticleProcessMaterialEmissionShape int
const (
  ParticleProcessMaterialEmissionShapePoint ParticleProcessMaterialEmissionShape = 0
  ParticleProcessMaterialEmissionShapeSphere ParticleProcessMaterialEmissionShape = 1
  ParticleProcessMaterialEmissionShapeSphereSurface ParticleProcessMaterialEmissionShape = 2
  ParticleProcessMaterialEmissionShapeBox ParticleProcessMaterialEmissionShape = 3
  ParticleProcessMaterialEmissionShapePoints ParticleProcessMaterialEmissionShape = 4
  ParticleProcessMaterialEmissionShapeDirectedPoints ParticleProcessMaterialEmissionShape = 5
  ParticleProcessMaterialEmissionShapeRing ParticleProcessMaterialEmissionShape = 6
  ParticleProcessMaterialEmissionShapeMax ParticleProcessMaterialEmissionShape = 7
)

type ParticleProcessMaterialSubEmitterMode int
const (
  ParticleProcessMaterialSubEmitterDisabled ParticleProcessMaterialSubEmitterMode = 0
  ParticleProcessMaterialSubEmitterConstant ParticleProcessMaterialSubEmitterMode = 1
  ParticleProcessMaterialSubEmitterAtEnd ParticleProcessMaterialSubEmitterMode = 2
  ParticleProcessMaterialSubEmitterAtCollision ParticleProcessMaterialSubEmitterMode = 3
  ParticleProcessMaterialSubEmitterMax ParticleProcessMaterialSubEmitterMode = 4
)

type ParticleProcessMaterialCollisionMode int
const (
  ParticleProcessMaterialCollisionDisabled ParticleProcessMaterialCollisionMode = 0
  ParticleProcessMaterialCollisionRigid ParticleProcessMaterialCollisionMode = 1
  ParticleProcessMaterialCollisionHideOnContact ParticleProcessMaterialCollisionMode = 2
  ParticleProcessMaterialCollisionMax ParticleProcessMaterialCollisionMode = 3
)
