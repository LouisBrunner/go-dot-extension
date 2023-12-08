// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  CPUParticles2DDrawOrderDrawOrderIndex CPUParticles2DDrawOrder = 0
  CPUParticles2DDrawOrderDrawOrderLifetime CPUParticles2DDrawOrder = 1
)

type CPUParticles2DParameter int
const (
  CPUParticles2DParameterParamInitialLinearVelocity CPUParticles2DParameter = 0
  CPUParticles2DParameterParamAngularVelocity CPUParticles2DParameter = 1
  CPUParticles2DParameterParamOrbitVelocity CPUParticles2DParameter = 2
  CPUParticles2DParameterParamLinearAccel CPUParticles2DParameter = 3
  CPUParticles2DParameterParamRadialAccel CPUParticles2DParameter = 4
  CPUParticles2DParameterParamTangentialAccel CPUParticles2DParameter = 5
  CPUParticles2DParameterParamDamping CPUParticles2DParameter = 6
  CPUParticles2DParameterParamAngle CPUParticles2DParameter = 7
  CPUParticles2DParameterParamScale CPUParticles2DParameter = 8
  CPUParticles2DParameterParamHueVariation CPUParticles2DParameter = 9
  CPUParticles2DParameterParamAnimSpeed CPUParticles2DParameter = 10
  CPUParticles2DParameterParamAnimOffset CPUParticles2DParameter = 11
  CPUParticles2DParameterParamMax CPUParticles2DParameter = 12
)

type CPUParticles2DParticleFlags int
const (
  CPUParticles2DParticleFlagsParticleFlagAlignYToVelocity CPUParticles2DParticleFlags = 0
  CPUParticles2DParticleFlagsParticleFlagRotateY CPUParticles2DParticleFlags = 1
  CPUParticles2DParticleFlagsParticleFlagDisableZ CPUParticles2DParticleFlags = 2
  CPUParticles2DParticleFlagsParticleFlagMax CPUParticles2DParticleFlags = 3
)

type CPUParticles2DEmissionShape int
const (
  CPUParticles2DEmissionShapeEmissionShapePoint CPUParticles2DEmissionShape = 0
  CPUParticles2DEmissionShapeEmissionShapeSphere CPUParticles2DEmissionShape = 1
  CPUParticles2DEmissionShapeEmissionShapeSphereSurface CPUParticles2DEmissionShape = 2
  CPUParticles2DEmissionShapeEmissionShapeRectangle CPUParticles2DEmissionShape = 3
  CPUParticles2DEmissionShapeEmissionShapePoints CPUParticles2DEmissionShape = 4
  CPUParticles2DEmissionShapeEmissionShapeDirectedPoints CPUParticles2DEmissionShape = 5
  CPUParticles2DEmissionShapeEmissionShapeMax CPUParticles2DEmissionShape = 6
)

func  (me *CPUParticles2D) SetEmitting(emitting bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetAmount(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetOneShot(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetPreProcessTime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetExplosivenessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetRandomnessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetLifetimeRandomness(random float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetUseLocalCoordinates(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetFixedFps(fps int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetFractionalDelta(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetSpeedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) IsEmitting() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetAmount() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetOneShot() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetPreProcessTime() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetExplosivenessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetRandomnessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetLifetimeRandomness() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetUseLocalCoordinates() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetFixedFps() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetFractionalDelta() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetDrawOrder(order CPUParticles2DDrawOrder, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetDrawOrder() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetTexture() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) Restart() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetDirection(direction Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetDirection() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetSpread(spread float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetSpread() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetParamMin(param CPUParticles2DParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetParamMin(param CPUParticles2DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetParamMax(param CPUParticles2DParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetParamMax(param CPUParticles2DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetParamCurve(param CPUParticles2DParameter, curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetParamCurve(param CPUParticles2DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetColor() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetColorRamp(ramp Gradient, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetColorRamp() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetColorInitialRamp(ramp Gradient, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetColorInitialRamp() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetParticleFlag(particle_flag CPUParticles2DParticleFlags, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetParticleFlag(particle_flag CPUParticles2DParticleFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionShape(shape CPUParticles2DEmissionShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionShape() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionSphereRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionSphereRadius() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionRectExtents(extents Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionRectExtents() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionPoints(array PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionPoints() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionNormals(array PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionNormals() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetEmissionColors(array PackedColorArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetEmissionColors() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetGravity() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetGravity(accel_vec Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetSplitScale() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetSplitScale(split_scale bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetScaleCurveX() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetScaleCurveX(scale_curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) GetScaleCurveY() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) SetScaleCurveY(scale_curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles2D) ConvertFromParticles(particles Node, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
