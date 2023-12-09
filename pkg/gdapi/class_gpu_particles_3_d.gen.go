// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticles3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticles3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticles3D) BaseClass() string {
  return "GPUParticles3D"
}



// Constants

var (
  GPUParticles3DMaxDrawPasses = "4" // TODO: construct correctly
)

// Enums

type GPUParticles3DDrawOrder int
const (
  GPUParticles3DDrawOrderDrawOrderIndex GPUParticles3DDrawOrder = 0
  GPUParticles3DDrawOrderDrawOrderLifetime GPUParticles3DDrawOrder = 1
  GPUParticles3DDrawOrderDrawOrderReverseLifetime GPUParticles3DDrawOrder = 2
  GPUParticles3DDrawOrderDrawOrderViewDepth GPUParticles3DDrawOrder = 3
)

type GPUParticles3DEmitFlags int
const (
  GPUParticles3DEmitFlagsEmitFlagPosition GPUParticles3DEmitFlags = 1
  GPUParticles3DEmitFlagsEmitFlagRotationScale GPUParticles3DEmitFlags = 2
  GPUParticles3DEmitFlagsEmitFlagVelocity GPUParticles3DEmitFlags = 4
  GPUParticles3DEmitFlagsEmitFlagColor GPUParticles3DEmitFlags = 8
  GPUParticles3DEmitFlagsEmitFlagCustom GPUParticles3DEmitFlags = 16
)

type GPUParticles3DTransformAlign int
const (
  GPUParticles3DTransformAlignTransformAlignDisabled GPUParticles3DTransformAlign = 0
  GPUParticles3DTransformAlignTransformAlignZBillboard GPUParticles3DTransformAlign = 1
  GPUParticles3DTransformAlignTransformAlignYToVelocity GPUParticles3DTransformAlign = 2
  GPUParticles3DTransformAlignTransformAlignZBillboardYToVelocity GPUParticles3DTransformAlign = 3
)

func (me *GPUParticles3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticles3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticles3D) SetEmitting(emitting bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetAmount(amount int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetOneShot(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetPreProcessTime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetExplosivenessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetRandomnessRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetVisibilityAabb(aabb AABB, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetUseLocalCoordinates(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetFixedFps(fps int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetFractionalDelta(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetInterpolate(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetProcessMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetCollisionBaseSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) IsEmitting()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetAmount()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetLifetime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetOneShot()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetPreProcessTime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetExplosivenessRatio()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetRandomnessRatio()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetVisibilityAabb()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetUseLocalCoordinates()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetFixedFps()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetFractionalDelta()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetInterpolate()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetProcessMaterial()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetSpeedScale()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetCollisionBaseSize()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetDrawOrder(order GPUParticles3DDrawOrder, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetDrawOrder()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetDrawPasses(passes int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetDrawPassMesh(pass int, mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetDrawPasses()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetDrawPassMesh(pass int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetSkin(skin Skin, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetSkin()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) Restart()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) CaptureAabb()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetSubEmitter(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetSubEmitter()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) EmitParticle(xform Transform3D, velocity Vector3, color Color, custom Color, flags int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetTrailEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetTrailLifetime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) IsTrailEnabled()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetTrailLifetime()  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) SetTransformAlign(align GPUParticles3DTransformAlign, )  {
  panic("TODO: implement")
}

func  (me *GPUParticles3D) GetTransformAlign()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
