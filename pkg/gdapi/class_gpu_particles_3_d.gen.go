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

// TODO: needed?
// const (
// // )

var (
  GPUParticles3DMaxDrawPasses = "4" // TODO: construct correctly
)

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

func  (me *GPUParticles3D) SetEmitting(emitting bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetAmount(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetOneShot(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetPreProcessTime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetExplosivenessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetRandomnessRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetVisibilityAabb(aabb AABB, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetUseLocalCoordinates(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetFixedFps(fps int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetFractionalDelta(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetInterpolate(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetProcessMaterial(material Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetSpeedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetCollisionBaseSize(size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) IsEmitting() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetAmount() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetOneShot() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetPreProcessTime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetExplosivenessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetRandomnessRatio() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetVisibilityAabb() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetUseLocalCoordinates() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetFixedFps() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetFractionalDelta() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetInterpolate() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetProcessMaterial() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetCollisionBaseSize() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetDrawOrder(order GPUParticles3DDrawOrder, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetDrawOrder() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetDrawPasses(passes int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetDrawPassMesh(pass int, mesh Mesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetDrawPasses() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetDrawPassMesh(pass int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetSkin(skin Skin, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetSkin() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) Restart() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) CaptureAabb() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetSubEmitter(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetSubEmitter() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) EmitParticle(xform Transform3D, velocity Vector3, color Color, custom Color, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetTrailEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetTrailLifetime(secs float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) IsTrailEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetTrailLifetime() { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) SetTransformAlign(align GPUParticles3DTransformAlign, ) { // TODO: return value
  // TODO: implement
}

func  (me *GPUParticles3D) GetTransformAlign() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
