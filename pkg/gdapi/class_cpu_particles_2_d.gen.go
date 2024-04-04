// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type CPUParticles2D struct {
  Node2D
}

func (me *CPUParticles2D) BaseClass() string {
  return "CPUParticles2D"
}

func NewCPUParticles2D() *CPUParticles2D {
  str := StringNameFromStr("CPUParticles2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CPUParticles2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetAmount(amount int64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetLifetime(secs float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetPreProcessTime(secs float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetExplosivenessRatio(ratio float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetRandomnessRatio(ratio float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetLifetimeRandomness(random float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&random) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetUseLocalCoordinates(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetFixedFps(fps int64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetFractionalDelta(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetSpeedScale(scale float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) IsEmitting() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetAmount() int64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetLifetime() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetOneShot() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetPreProcessTime() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetExplosivenessRatio() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetRandomnessRatio() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetLifetimeRandomness() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetUseLocalCoordinates() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetFixedFps() int64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetFractionalDelta() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetSpeedScale() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetDrawOrder(order CPUParticles2DDrawOrder, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4183193490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetDrawOrder() CPUParticles2DDrawOrder {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1668655735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CPUParticles2DDrawOrder

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CPUParticles2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) Restart()  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("restart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetDirection(direction Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{direction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetDirection() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetSpread(spread float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetSpread() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamMin(param CPUParticles2DParameter, value float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3320615296) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamMin(param CPUParticles2DParameter, ) float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2038050600) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamMax(param CPUParticles2DParameter, value float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3320615296) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamMax(param CPUParticles2DParameter, ) float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2038050600) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamCurve(param CPUParticles2DParameter, curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2959350143) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamCurve(param CPUParticles2DParameter, ) Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2603158474) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColor() Color {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColorRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColorRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColorInitialRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColorInitialRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetParticleFlag(particle_flag CPUParticles2DParticleFlags, enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4178137949) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParticleFlag(particle_flag CPUParticles2DParticleFlags, ) bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2829976507) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&particle_flag)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetEmissionShape(shape CPUParticles2DEmissionShape, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 393763892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionShape() CPUParticles2DEmissionShape {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740246024) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CPUParticles2DEmissionShape

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CPUParticles2D) SetEmissionSphereRadius(radius float64, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionSphereRadius() float64 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetEmissionRectExtents(extents Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_rect_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extents.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionRectExtents() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_rect_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionPoints(array PackedVector2Array, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionPoints() PackedVector2Array {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionNormals(array PackedVector2Array, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionNormals() PackedVector2Array {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionColors(array PackedColorArray, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionColors() PackedColorArray {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) GetGravity() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetGravity(accel_vec Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{accel_vec.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetSplitScale() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetSplitScale(split_scale bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&split_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetScaleCurveX() Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetScaleCurveX(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetScaleCurveY() Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetScaleCurveY(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) ConvertFromParticles(particles Node, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_from_particles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CPUParticles2DFinishedSignalFn func()

func (me *CPUParticles2D) ConnectFinished(subs SignalSubscribers, fn CPUParticles2DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CPUParticles2D) DisconnectFinished(subs SignalSubscribers, fn CPUParticles2DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
