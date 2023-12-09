// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type ParticleProcessMaterialParameter int
const (
  ParticleProcessMaterialParameterParamInitialLinearVelocity ParticleProcessMaterialParameter = 0
  ParticleProcessMaterialParameterParamAngularVelocity ParticleProcessMaterialParameter = 1
  ParticleProcessMaterialParameterParamOrbitVelocity ParticleProcessMaterialParameter = 2
  ParticleProcessMaterialParameterParamLinearAccel ParticleProcessMaterialParameter = 3
  ParticleProcessMaterialParameterParamRadialAccel ParticleProcessMaterialParameter = 4
  ParticleProcessMaterialParameterParamTangentialAccel ParticleProcessMaterialParameter = 5
  ParticleProcessMaterialParameterParamDamping ParticleProcessMaterialParameter = 6
  ParticleProcessMaterialParameterParamAngle ParticleProcessMaterialParameter = 7
  ParticleProcessMaterialParameterParamScale ParticleProcessMaterialParameter = 8
  ParticleProcessMaterialParameterParamHueVariation ParticleProcessMaterialParameter = 9
  ParticleProcessMaterialParameterParamAnimSpeed ParticleProcessMaterialParameter = 10
  ParticleProcessMaterialParameterParamAnimOffset ParticleProcessMaterialParameter = 11
  ParticleProcessMaterialParameterParamMax ParticleProcessMaterialParameter = 15
  ParticleProcessMaterialParameterParamTurbVelInfluence ParticleProcessMaterialParameter = 13
  ParticleProcessMaterialParameterParamTurbInitDisplacement ParticleProcessMaterialParameter = 14
  ParticleProcessMaterialParameterParamTurbInfluenceOverLife ParticleProcessMaterialParameter = 12
)

type ParticleProcessMaterialParticleFlags int
const (
  ParticleProcessMaterialParticleFlagsParticleFlagAlignYToVelocity ParticleProcessMaterialParticleFlags = 0
  ParticleProcessMaterialParticleFlagsParticleFlagRotateY ParticleProcessMaterialParticleFlags = 1
  ParticleProcessMaterialParticleFlagsParticleFlagDisableZ ParticleProcessMaterialParticleFlags = 2
  ParticleProcessMaterialParticleFlagsParticleFlagMax ParticleProcessMaterialParticleFlags = 3
)

type ParticleProcessMaterialEmissionShape int
const (
  ParticleProcessMaterialEmissionShapeEmissionShapePoint ParticleProcessMaterialEmissionShape = 0
  ParticleProcessMaterialEmissionShapeEmissionShapeSphere ParticleProcessMaterialEmissionShape = 1
  ParticleProcessMaterialEmissionShapeEmissionShapeSphereSurface ParticleProcessMaterialEmissionShape = 2
  ParticleProcessMaterialEmissionShapeEmissionShapeBox ParticleProcessMaterialEmissionShape = 3
  ParticleProcessMaterialEmissionShapeEmissionShapePoints ParticleProcessMaterialEmissionShape = 4
  ParticleProcessMaterialEmissionShapeEmissionShapeDirectedPoints ParticleProcessMaterialEmissionShape = 5
  ParticleProcessMaterialEmissionShapeEmissionShapeRing ParticleProcessMaterialEmissionShape = 6
  ParticleProcessMaterialEmissionShapeEmissionShapeMax ParticleProcessMaterialEmissionShape = 7
)

type ParticleProcessMaterialSubEmitterMode int
const (
  ParticleProcessMaterialSubEmitterModeSubEmitterDisabled ParticleProcessMaterialSubEmitterMode = 0
  ParticleProcessMaterialSubEmitterModeSubEmitterConstant ParticleProcessMaterialSubEmitterMode = 1
  ParticleProcessMaterialSubEmitterModeSubEmitterAtEnd ParticleProcessMaterialSubEmitterMode = 2
  ParticleProcessMaterialSubEmitterModeSubEmitterAtCollision ParticleProcessMaterialSubEmitterMode = 3
  ParticleProcessMaterialSubEmitterModeSubEmitterMax ParticleProcessMaterialSubEmitterMode = 4
)

type ParticleProcessMaterialCollisionMode int
const (
  ParticleProcessMaterialCollisionModeCollisionDisabled ParticleProcessMaterialCollisionMode = 0
  ParticleProcessMaterialCollisionModeCollisionRigid ParticleProcessMaterialCollisionMode = 1
  ParticleProcessMaterialCollisionModeCollisionHideOnContact ParticleProcessMaterialCollisionMode = 2
  ParticleProcessMaterialCollisionModeCollisionMax ParticleProcessMaterialCollisionMode = 3
)

func (me *ParticleProcessMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ParticleProcessMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ParticleProcessMaterial) SetDirection(degrees Vector3, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetDirection()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSpread(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSpread()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetFlatness(amount float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetFlatness()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetParamMin(param ParticleProcessMaterialParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetParamMin(param ParticleProcessMaterialParameter, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetParamMax(param ParticleProcessMaterialParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetParamMax(param ParticleProcessMaterialParameter, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetParamTexture(param ParticleProcessMaterialParameter, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetParamTexture(param ParticleProcessMaterialParameter, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetColor()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetColorRamp(ramp Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetColorRamp()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetColorInitialRamp(ramp Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetColorInitialRamp()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, enable bool, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionShape(shape ParticleProcessMaterialEmissionShape, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionShape()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionSphereRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionSphereRadius()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionBoxExtents(extents Vector3, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionBoxExtents()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionPointTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionPointTexture()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionNormalTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionNormalTexture()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionColorTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionColorTexture()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionPointCount(point_count int, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionPointCount()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionRingAxis(axis Vector3, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionRingAxis()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionRingHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionRingHeight()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionRingRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionRingRadius()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetEmissionRingInnerRadius(inner_radius float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetEmissionRingInnerRadius()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetTurbulenceEnabled()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetTurbulenceEnabled(turbulence_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseStrength()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseStrength(turbulence_noise_strength float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseScale()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseScale(turbulence_noise_scale float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeedRandom()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeed()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeed(turbulence_noise_speed Vector3, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetGravity()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetGravity(accel_vec Vector3, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetLifetimeRandomness(randomness float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetLifetimeRandomness()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSubEmitterMode()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSubEmitterMode(mode ParticleProcessMaterialSubEmitterMode, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSubEmitterFrequency()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSubEmitterFrequency(hz float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtEnd()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtEnd(amount int, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtCollision()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtCollision(amount int, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetSubEmitterKeepVelocity()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetSubEmitterKeepVelocity(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetAttractorInteractionEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) IsAttractorInteractionEnabled()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetCollisionMode(mode ParticleProcessMaterialCollisionMode, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetCollisionMode()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetCollisionUseScale(radius bool, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) IsCollisionUsingScale()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetCollisionFriction(friction float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetCollisionFriction()  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) SetCollisionBounce(bounce float32, )  {
  panic("TODO: implement")
}

func  (me *ParticleProcessMaterial) GetCollisionBounce()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
