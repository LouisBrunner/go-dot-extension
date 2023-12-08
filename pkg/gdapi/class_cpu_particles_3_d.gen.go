// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  CPUParticles3DDrawOrderDrawOrderIndex CPUParticles3DDrawOrder = 0
  CPUParticles3DDrawOrderDrawOrderLifetime CPUParticles3DDrawOrder = 1
  CPUParticles3DDrawOrderDrawOrderViewDepth CPUParticles3DDrawOrder = 2
)

type CPUParticles3DParameter int
const (
  CPUParticles3DParameterParamInitialLinearVelocity CPUParticles3DParameter = 0
  CPUParticles3DParameterParamAngularVelocity CPUParticles3DParameter = 1
  CPUParticles3DParameterParamOrbitVelocity CPUParticles3DParameter = 2
  CPUParticles3DParameterParamLinearAccel CPUParticles3DParameter = 3
  CPUParticles3DParameterParamRadialAccel CPUParticles3DParameter = 4
  CPUParticles3DParameterParamTangentialAccel CPUParticles3DParameter = 5
  CPUParticles3DParameterParamDamping CPUParticles3DParameter = 6
  CPUParticles3DParameterParamAngle CPUParticles3DParameter = 7
  CPUParticles3DParameterParamScale CPUParticles3DParameter = 8
  CPUParticles3DParameterParamHueVariation CPUParticles3DParameter = 9
  CPUParticles3DParameterParamAnimSpeed CPUParticles3DParameter = 10
  CPUParticles3DParameterParamAnimOffset CPUParticles3DParameter = 11
  CPUParticles3DParameterParamMax CPUParticles3DParameter = 12
)

type CPUParticles3DParticleFlags int
const (
  CPUParticles3DParticleFlagsParticleFlagAlignYToVelocity CPUParticles3DParticleFlags = 0
  CPUParticles3DParticleFlagsParticleFlagRotateY CPUParticles3DParticleFlags = 1
  CPUParticles3DParticleFlagsParticleFlagDisableZ CPUParticles3DParticleFlags = 2
  CPUParticles3DParticleFlagsParticleFlagMax CPUParticles3DParticleFlags = 3
)

type CPUParticles3DEmissionShape int
const (
  CPUParticles3DEmissionShapeEmissionShapePoint CPUParticles3DEmissionShape = 0
  CPUParticles3DEmissionShapeEmissionShapeSphere CPUParticles3DEmissionShape = 1
  CPUParticles3DEmissionShapeEmissionShapeSphereSurface CPUParticles3DEmissionShape = 2
  CPUParticles3DEmissionShapeEmissionShapeBox CPUParticles3DEmissionShape = 3
  CPUParticles3DEmissionShapeEmissionShapePoints CPUParticles3DEmissionShape = 4
  CPUParticles3DEmissionShapeEmissionShapeDirectedPoints CPUParticles3DEmissionShape = 5
  CPUParticles3DEmissionShapeEmissionShapeRing CPUParticles3DEmissionShape = 6
  CPUParticles3DEmissionShapeEmissionShapeMax CPUParticles3DEmissionShape = 7
)

func  (me *CPUParticles3D) SetEmitting(emitting bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetAmount(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetOneShot(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetPreProcessTime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetExplosivenessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetRandomnessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetLifetimeRandomness(random float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetUseLocalCoordinates(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetFixedFps(fps int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetFractionalDelta(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetSpeedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) IsEmitting() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetAmount() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetOneShot() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetPreProcessTime() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetExplosivenessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetRandomnessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetLifetimeRandomness() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetUseLocalCoordinates() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetFixedFps() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetFractionalDelta() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetDrawOrder(order CPUParticles3DDrawOrder, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetDrawOrder() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetMesh(mesh Mesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetMesh() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) Restart() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetDirection(direction Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetDirection() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetSpread(degrees float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetSpread() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetFlatness(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetFlatness() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetParamMin(param CPUParticles3DParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetParamMin(param CPUParticles3DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetParamMax(param CPUParticles3DParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetParamMax(param CPUParticles3DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetParamCurve(param CPUParticles3DParameter, curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetParamCurve(param CPUParticles3DParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetColor() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetColorRamp(ramp Gradient, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetColorRamp() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetColorInitialRamp(ramp Gradient, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetColorInitialRamp() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetParticleFlag(particle_flag CPUParticles3DParticleFlags, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetParticleFlag(particle_flag CPUParticles3DParticleFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionShape(shape CPUParticles3DEmissionShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionShape() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionSphereRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionSphereRadius() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionBoxExtents(extents Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionBoxExtents() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionPoints(array PackedVector3Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionPoints() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionNormals(array PackedVector3Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionNormals() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionColors(array PackedColorArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionColors() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionRingAxis(axis Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionRingAxis() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionRingHeight(height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionRingHeight() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionRingRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionRingRadius() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetEmissionRingInnerRadius(inner_radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetEmissionRingInnerRadius() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetGravity() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetGravity(accel_vec Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetSplitScale() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetSplitScale(split_scale bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetScaleCurveX() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetScaleCurveX(scale_curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetScaleCurveY() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetScaleCurveY(scale_curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) GetScaleCurveZ() { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) SetScaleCurveZ(scale_curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *CPUParticles3D) ConvertFromParticles(particles Node, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
