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



// Enums

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

func (me *GPUParticles2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticles2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticles2D) SetEmitting(emitting bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetAmount(amount int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetOneShot(secs bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetPreProcessTime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetExplosivenessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetRandomnessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetVisibilityRect(visibility_rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetUseLocalCoordinates(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetFixedFps(fps int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetFractionalDelta(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetInterpolate(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetProcessMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetCollisionBaseSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) IsEmitting()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetAmount()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetLifetime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetOneShot()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetPreProcessTime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetExplosivenessRatio()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetRandomnessRatio()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetVisibilityRect()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetUseLocalCoordinates()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetFixedFps()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetFractionalDelta()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetInterpolate()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetProcessMaterial()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetCollisionBaseSize()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetDrawOrder(order GPUParticles2DDrawOrder, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetDrawOrder()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetTexture()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) CaptureRect()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) Restart()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetSubEmitter(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetSubEmitter()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) EmitParticle(xform Transform2D, velocity Vector2, color Color, custom Color, flags int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetTrailEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetTrailLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) IsTrailEnabled()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetTrailLifetime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetTrailSections(sections int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetTrailSections()  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) SetTrailSectionSubdivisions(subdivisions int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles2D) GetTrailSectionSubdivisions()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
