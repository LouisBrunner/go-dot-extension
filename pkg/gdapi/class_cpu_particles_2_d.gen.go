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



// Enums

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

func (me *CPUParticles2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CPUParticles2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CPUParticles2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CPUParticles2D) SetEmitting(emitting bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetAmount(amount int, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetOneShot(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetPreProcessTime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetExplosivenessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetRandomnessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetLifetimeRandomness(random float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetUseLocalCoordinates(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetFixedFps(fps int, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetFractionalDelta(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) IsEmitting()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetAmount()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetLifetime()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetOneShot()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetPreProcessTime()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetExplosivenessRatio()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetRandomnessRatio()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetLifetimeRandomness()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetUseLocalCoordinates()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetFixedFps()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetFractionalDelta()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetDrawOrder(order CPUParticles2DDrawOrder, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetDrawOrder()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetTexture()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) Restart()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetDirection(direction Vector2, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetDirection()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetSpread(spread float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetSpread()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetParamMin(param CPUParticles2DParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetParamMin(param CPUParticles2DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetParamMax(param CPUParticles2DParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetParamMax(param CPUParticles2DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetParamCurve(param CPUParticles2DParameter, curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetParamCurve(param CPUParticles2DParameter, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetColor()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetColorRamp(ramp Gradient, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetColorRamp()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetColorInitialRamp(ramp Gradient, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetColorInitialRamp()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetParticleFlag(particle_flag CPUParticles2DParticleFlags, enable bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetParticleFlag(particle_flag CPUParticles2DParticleFlags, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionShape(shape CPUParticles2DEmissionShape, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionShape()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionSphereRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionSphereRadius()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionRectExtents(extents Vector2, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionRectExtents()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionPoints(array PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionPoints()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionNormals(array PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionNormals()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetEmissionColors(array PackedColorArray, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetEmissionColors()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetGravity()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetGravity(accel_vec Vector2, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetSplitScale()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetSplitScale(split_scale bool, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetScaleCurveX()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetScaleCurveX(scale_curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) GetScaleCurveY()  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) SetScaleCurveY(scale_curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CPUParticles2D) ConvertFromParticles(particles Node, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
