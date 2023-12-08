// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles2D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticles2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticles2D) BaseClass() string {
  return "GPUParticles2D"
}

type GPUParticles2DDrawOrder int
const (
  GPUParticles2DDrawOrderDrawOrderIndex GPUParticles2DDrawOrder = 0
  GPUParticles2DDrawOrderDrawOrderLifetime GPUParticles2DDrawOrder = 1
  GPUParticles2DDrawOrderDrawOrderReverseLifetime GPUParticles2DDrawOrder = 2
)

type GPUParticles2DEmitFlags int
const (
  GPUParticles2DEmitFlagsEmitFlagPosition GPUParticles2DEmitFlags = 1
  GPUParticles2DEmitFlagsEmitFlagRotationScale GPUParticles2DEmitFlags = 2
  GPUParticles2DEmitFlagsEmitFlagVelocity GPUParticles2DEmitFlags = 4
  GPUParticles2DEmitFlagsEmitFlagColor GPUParticles2DEmitFlags = 8
  GPUParticles2DEmitFlagsEmitFlagCustom GPUParticles2DEmitFlags = 16
)

func  (me *GPUParticles2D) SetEmitting(emitting bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetAmount(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetOneShot(secs bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetPreProcessTime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetExplosivenessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetRandomnessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetVisibilityRect(visibility_rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetUseLocalCoordinates(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetFixedFps(fps int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetFractionalDelta(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetInterpolate(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetProcessMaterial(material Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetSpeedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetCollisionBaseSize(size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) IsEmitting() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetAmount() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetOneShot() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetPreProcessTime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetExplosivenessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetRandomnessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetVisibilityRect() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetUseLocalCoordinates() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetFixedFps() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetFractionalDelta() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetInterpolate() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetProcessMaterial() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetCollisionBaseSize() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetDrawOrder(order GPUParticles2DDrawOrder, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetDrawOrder() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetTexture() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) CaptureRect() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) Restart() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetSubEmitter(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetSubEmitter() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) EmitParticle(xform Transform2D, velocity Vector2, color Color, custom Color, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetTrailEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetTrailLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) IsTrailEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetTrailLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetTrailSections(sections int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetTrailSections() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) SetTrailSectionSubdivisions(subdivisions int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles2D) GetTrailSectionSubdivisions() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
