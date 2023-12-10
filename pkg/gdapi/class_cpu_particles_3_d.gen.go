// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *CPUParticles3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CPUParticles3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CPUParticles3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CPUParticles3D) SetEmitting(emitting bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetAmount(amount int, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetLifetime(secs float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetPreProcessTime(secs float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetExplosivenessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetRandomnessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetLifetimeRandomness(random float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&random), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetUseLocalCoordinates(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetFixedFps(fps int, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetFractionalDelta(enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetSpeedScale(scale float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) IsEmitting() bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetAmount() int {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetLifetime() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetOneShot() bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetPreProcessTime() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetExplosivenessRatio() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetRandomnessRatio() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetLifetimeRandomness() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetUseLocalCoordinates() bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetFixedFps() int {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetFractionalDelta() bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetDrawOrder(order CPUParticles3DDrawOrder, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1427401774) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetDrawOrder() CPUParticles3DDrawOrder {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321900776) // FIXME: should cache?
  var ret CPUParticles3DDrawOrder
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetMesh() Mesh {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1808005922) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) Restart()  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("restart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) SetDirection(direction Vector3, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetDirection() Vector3 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetSpread(degrees float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetSpread() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetFlatness(amount float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetFlatness() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetParamMin(param CPUParticles3DParameter, value float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 557936109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetParamMin(param CPUParticles3DParameter, ) float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 597646162) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetParamMax(param CPUParticles3DParameter, value float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 557936109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetParamMax(param CPUParticles3DParameter, ) float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 597646162) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetParamCurve(param CPUParticles3DParameter, curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4044142537) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetParamCurve(param CPUParticles3DParameter, ) Curve {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4132790277) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetColor() Color {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetColorRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetColorRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetColorInitialRamp(ramp Gradient, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetColorInitialRamp() Gradient {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetParticleFlag(particle_flag CPUParticles3DParticleFlags, enable bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3515406498) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetParticleFlag(particle_flag CPUParticles3DParticleFlags, ) bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2845201987) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionShape(shape CPUParticles3DEmissionShape, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 491823814) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionShape() CPUParticles3DEmissionShape {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961454842) // FIXME: should cache?
  var ret CPUParticles3DEmissionShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionSphereRadius(radius float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionSphereRadius() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionBoxExtents(extents Vector3, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_box_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extents.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionBoxExtents() Vector3 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_box_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionPoints(array PackedVector3Array, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionPoints() PackedVector3Array {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionNormals(array PackedVector3Array, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionNormals() PackedVector3Array {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionColors(array PackedColorArray, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionColors() PackedColorArray {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  var ret PackedColorArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionRingAxis(axis Vector3, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionRingAxis() Vector3 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionRingHeight(height float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionRingHeight() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionRingRadius(radius float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionRingRadius() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetEmissionRingInnerRadius(inner_radius float32, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetEmissionRingInnerRadius() float32 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) GetGravity() Vector3 {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetGravity(accel_vec Vector3, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(accel_vec.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetSplitScale() bool {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetSplitScale(split_scale bool, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_split_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&split_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetScaleCurveX() Curve {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetScaleCurveX(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale_curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetScaleCurveY() Curve {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetScaleCurveY(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale_curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) GetScaleCurveZ() Curve {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_curve_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CPUParticles3D) SetScaleCurveZ(scale_curve Curve, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_curve_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale_curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CPUParticles3D) ConvertFromParticles(particles Node, )  {
  classNameV := StringNameFromStr("CPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_from_particles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(particles.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *CPUParticles3D) GetPropEmitting() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmitting(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAmount() int {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAmount(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLifetime() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLifetime(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropOneShot() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropOneShot(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropPreprocess() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropPreprocess(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropSpeedScale() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropSpeedScale(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropExplosiveness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropExplosiveness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropRandomness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropRandomness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLifetimeRandomness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLifetimeRandomness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropFixedFps() int {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropFixedFps(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropFractDelta() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropFractDelta(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLocalCoords() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLocalCoords(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropDrawOrder() int {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropDrawOrder(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropMesh() Mesh {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropMesh(value Mesh) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionShape() int {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionShape(value int) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionSphereRadius() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionSphereRadius(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionBoxExtents() Vector3 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionBoxExtents(value Vector3) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionPoints() PackedVector3Array {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionPoints(value PackedVector3Array) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionNormals() PackedVector3Array {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionNormals(value PackedVector3Array) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionColors() PackedColorArray {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionColors(value PackedColorArray) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionRingAxis() Vector3 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionRingAxis(value Vector3) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionRingHeight() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionRingHeight(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionRingRadius() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionRingRadius(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropEmissionRingInnerRadius() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropEmissionRingInnerRadius(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropParticleFlagAlignY() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropParticleFlagAlignY(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropParticleFlagRotateY() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropParticleFlagRotateY(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropParticleFlagDisableZ() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropParticleFlagDisableZ(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropDirection() Vector3 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropDirection(value Vector3) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropSpread() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropSpread(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropFlatness() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropFlatness(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropGravity() Vector3 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropGravity(value Vector3) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropInitialVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropInitialVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropInitialVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropInitialVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngularVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngularVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngularVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngularVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngularVelocityCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngularVelocityCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropOrbitVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropOrbitVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropOrbitVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropOrbitVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropOrbitVelocityCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropOrbitVelocityCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLinearAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLinearAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLinearAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLinearAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropLinearAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropLinearAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropRadialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropRadialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropRadialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropRadialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropRadialAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropRadialAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropTangentialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropTangentialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropTangentialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropTangentialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropTangentialAccelCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropTangentialAccelCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropDampingMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropDampingMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropDampingMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropDampingMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropDampingCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropDampingCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngleMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngleMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngleMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngleMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAngleCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAngleCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleAmountMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleAmountMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleAmountMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleAmountMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleAmountCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleAmountCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropSplitScale() bool {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropSplitScale(value bool) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleCurveX() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleCurveX(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleCurveY() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleCurveY(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropScaleCurveZ() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropScaleCurveZ(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropColorRamp() Gradient {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropColorRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropColorInitialRamp() Gradient {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropColorInitialRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropHueVariationMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropHueVariationMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropHueVariationMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropHueVariationMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropHueVariationCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropHueVariationCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimSpeedMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimSpeedMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimSpeedMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimSpeedMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimSpeedCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimSpeedCurve(value Curve) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimOffsetMin() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimOffsetMin(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimOffsetMax() float32 {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimOffsetMax(value float32) {
  panic("TODO: implement")
}

func (me *CPUParticles3D) GetPropAnimOffsetCurve() Curve {
  panic("TODO: implement")
}

func (me *CPUParticles3D) SetPropAnimOffsetCurve(value Curve) {
  panic("TODO: implement")
}