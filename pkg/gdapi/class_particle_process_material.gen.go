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

func  (me *ParticleProcessMaterial) SetDirection(degrees Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetDirection() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSpread(degrees float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSpread() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetFlatness(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetFlatness() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetParamMin(param ParticleProcessMaterialParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetParamMin(param ParticleProcessMaterialParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetParamMax(param ParticleProcessMaterialParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetParamMax(param ParticleProcessMaterialParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetParamTexture(param ParticleProcessMaterialParameter, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetParamTexture(param ParticleProcessMaterialParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetColor() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetColorRamp(ramp Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetColorRamp() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetColorInitialRamp(ramp Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetColorInitialRamp() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionShape(shape ParticleProcessMaterialEmissionShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionShape() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionSphereRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionSphereRadius() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionBoxExtents(extents Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionBoxExtents() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionPointTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionPointTexture() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionNormalTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionNormalTexture() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionColorTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionColorTexture() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionPointCount(point_count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionRingAxis(axis Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionRingAxis() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionRingHeight(height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionRingHeight() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionRingRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionRingRadius() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetEmissionRingInnerRadius(inner_radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetEmissionRingInnerRadius() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetTurbulenceEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetTurbulenceEnabled(turbulence_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseStrength() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseStrength(turbulence_noise_strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseScale() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseScale(turbulence_noise_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeedRandom() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeed(turbulence_noise_speed Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetGravity() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetGravity(accel_vec Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetLifetimeRandomness(randomness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetLifetimeRandomness() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSubEmitterMode() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSubEmitterMode(mode ParticleProcessMaterialSubEmitterMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSubEmitterFrequency() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSubEmitterFrequency(hz float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtEnd() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtEnd(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtCollision() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtCollision(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetSubEmitterKeepVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetSubEmitterKeepVelocity(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetAttractorInteractionEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) IsAttractorInteractionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetCollisionMode(mode ParticleProcessMaterialCollisionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetCollisionMode() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetCollisionUseScale(radius bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) IsCollisionUsingScale() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetCollisionFriction(friction float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetCollisionFriction() { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) SetCollisionBounce(bounce float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ParticleProcessMaterial) GetCollisionBounce() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
