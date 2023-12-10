// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetAmount(amount int, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetLifetime(secs float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetPreProcessTime(secs float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetExplosivenessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetRandomnessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetLifetimeRandomness(random float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&random), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetUseLocalCoordinates(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetFixedFps(fps int, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetFractionalDelta(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetSpeedScale(scale float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) IsEmitting() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetAmount() int {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetLifetime() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetOneShot() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetPreProcessTime() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetExplosivenessRatio() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetRandomnessRatio() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetLifetimeRandomness() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetUseLocalCoordinates() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetFixedFps() int {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetFractionalDelta() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetDrawOrder(order CPUParticles2DDrawOrder, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4183193490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetDrawOrder() CPUParticles2DDrawOrder {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1668655735) // FIXME: should cache?
  var ret CPUParticles2DDrawOrder
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) Restart()  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("restart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) SetDirection(direction Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetDirection() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetSpread(spread float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spread), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetSpread() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetParamMin(param CPUParticles2DParameter, value float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3320615296) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetParamMin(param CPUParticles2DParameter, ) float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2038050600) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetParamMax(param CPUParticles2DParameter, value float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3320615296) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetParamMax(param CPUParticles2DParameter, ) float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2038050600) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetParamCurve(param CPUParticles2DParameter, curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2959350143) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetParamCurve(param CPUParticles2DParameter, ) Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2603158474) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetColor() Color {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetColorRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetColorRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetColorInitialRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetColorInitialRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetParticleFlag(particle_flag CPUParticles2DParticleFlags, enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4178137949) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetParticleFlag(particle_flag CPUParticles2DParticleFlags, ) bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2829976507) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionShape(shape CPUParticles2DEmissionShape, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 393763892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionShape() CPUParticles2DEmissionShape {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740246024) // FIXME: should cache?
  var ret CPUParticles2DEmissionShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionSphereRadius(radius float32, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionSphereRadius() float32 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionRectExtents(extents Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_rect_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extents.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionRectExtents() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_rect_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionPoints(array PackedVector2Array, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionPoints() PackedVector2Array {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionNormals(array PackedVector2Array, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionNormals() PackedVector2Array {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetEmissionColors(array PackedColorArray, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetEmissionColors() PackedColorArray {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  var ret PackedColorArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) GetGravity() Vector2 {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetGravity(accel_vec Vector2, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(accel_vec.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetSplitScale() bool {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetSplitScale(split_scale bool, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&split_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetScaleCurveX() Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetScaleCurveX(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale_curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) GetScaleCurveY() Curve {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles2D) SetScaleCurveY(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale_curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles2D) ConvertFromParticles(particles Node, )  {
  classNameV := StringNameFromStr("CPUParticles2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_from_particles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(particles.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *CPUParticles2D) GetPropEmitting() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmitting(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAmount() int {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAmount(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLifetime() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLifetime(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropOneShot() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropOneShot(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropPreprocess() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropPreprocess(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropSpeedScale() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropSpeedScale(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropExplosiveness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropExplosiveness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropRandomness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropRandomness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLifetimeRandomness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLifetimeRandomness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropFixedFps() int {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropFixedFps(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropFractDelta() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropFractDelta(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLocalCoords() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLocalCoords(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropDrawOrder() int {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropDrawOrder(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionShape() int {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionShape(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionSphereRadius() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionSphereRadius(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionRectExtents() Vector2 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionRectExtents(value Vector2) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionPoints() PackedVector2Array {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionPoints(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionNormals() PackedVector2Array {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionNormals(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropEmissionColors() PackedColorArray {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropEmissionColors(value PackedColorArray) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropParticleFlagAlignY() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropParticleFlagAlignY(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropDirection() Vector2 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropDirection(value Vector2) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropSpread() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropSpread(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropGravity() Vector2 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropGravity(value Vector2) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropInitialVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropInitialVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropInitialVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropInitialVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngularVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngularVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngularVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngularVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngularVelocityCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngularVelocityCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropOrbitVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropOrbitVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropOrbitVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropOrbitVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropOrbitVelocityCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropOrbitVelocityCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLinearAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLinearAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLinearAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLinearAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropLinearAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropLinearAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropRadialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropRadialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropRadialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropRadialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropRadialAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropRadialAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropTangentialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropTangentialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropTangentialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropTangentialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropTangentialAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropTangentialAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropDampingMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropDampingMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropDampingMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropDampingMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropDampingCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropDampingCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngleMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngleMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngleMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngleMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAngleCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAngleCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropScaleAmountMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropScaleAmountMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropScaleAmountMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropScaleAmountMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropScaleAmountCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropScaleAmountCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropSplitScale() bool {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropSplitScale(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropScaleCurveX() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropScaleCurveX(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropScaleCurveY() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropScaleCurveY(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropColorRamp() Gradient {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropColorRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropColorInitialRamp() Gradient {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropColorInitialRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropHueVariationMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropHueVariationMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropHueVariationMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropHueVariationMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropHueVariationCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropHueVariationCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimSpeedMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimSpeedMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimSpeedMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimSpeedMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimSpeedCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimSpeedCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimOffsetMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimOffsetMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimOffsetMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimOffsetMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles2D) GetPropAnimOffsetCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles2D) SetPropAnimOffsetCurve(value Curve) {
  panic("TODO: implement")
}