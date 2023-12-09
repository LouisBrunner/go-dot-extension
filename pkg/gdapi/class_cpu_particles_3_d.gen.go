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



// Enums

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

func (me *CPUParticles3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CPUParticles3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CPUParticles3D) SetEmitting(emitting bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetAmount(amount int, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetOneShot(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetPreProcessTime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetExplosivenessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetRandomnessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetLifetimeRandomness(random float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetUseLocalCoordinates(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetFixedFps(fps int, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetFractionalDelta(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) IsEmitting()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetAmount()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetLifetime()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetOneShot()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetPreProcessTime()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetExplosivenessRatio()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetRandomnessRatio()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetLifetimeRandomness()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetUseLocalCoordinates()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetFixedFps()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetFractionalDelta()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetDrawOrder(order CPUParticles3DDrawOrder, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetDrawOrder()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetMesh()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) Restart()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetDirection(direction Vector3, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetDirection()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetSpread(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetSpread()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetFlatness(amount float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetFlatness()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetParamMin(param CPUParticles3DParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetParamMin(param CPUParticles3DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetParamMax(param CPUParticles3DParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetParamMax(param CPUParticles3DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetParamCurve(param CPUParticles3DParameter, curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetParamCurve(param CPUParticles3DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetColor()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetColorRamp(ramp Gradient, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetColorRamp()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetColorInitialRamp(ramp Gradient, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetColorInitialRamp()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetParticleFlag(particle_flag CPUParticles3DParticleFlags, enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetParticleFlag(particle_flag CPUParticles3DParticleFlags, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionShape(shape CPUParticles3DEmissionShape, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionShape()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionSphereRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionSphereRadius()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionBoxExtents(extents Vector3, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionBoxExtents()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionPoints(array PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionPoints()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionNormals(array PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionNormals()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionColors(array PackedColorArray, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionColors()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionRingAxis(axis Vector3, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionRingAxis()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionRingHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionRingHeight()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionRingRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionRingRadius()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetEmissionRingInnerRadius(inner_radius float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetEmissionRingInnerRadius()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetGravity()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetGravity(accel_vec Vector3, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetSplitScale()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetSplitScale(split_scale bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetScaleCurveX()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetScaleCurveX(scale_curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetScaleCurveY()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetScaleCurveY(scale_curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) GetScaleCurveZ()  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) SetScaleCurveZ(scale_curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles3D) ConvertFromParticles(particles Node, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
