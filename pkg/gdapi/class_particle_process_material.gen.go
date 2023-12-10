// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ParticleProcessMaterial struct {
  obj gdc.ObjectPtr
}

func (me *ParticleProcessMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ParticleProcessMaterial) BaseClass() string {
  return "ParticleProcessMaterial"
}



// Enums

type ParticleProcessMaterialParameter int
const (
  ParticleProcessMaterialParameterParamInitialLinearVelocity ParticleProcessMaterialParameter = 0
  ParticleProcessMaterialParameterParamAngularVelocity ParticleProcessMaterialParameter = 1
  ParticleProcessMaterialParameterParamOrbitVelocity ParticleProcessMaterialParameter = 2
  ParticleProcessMaterialParameterParamLinearAccel ParticleProcessMaterialParameter = 3
  ParticleProcessMaterialParameterParamRadialAccel ParticleProcessMaterialParameter = 4
  ParticleProcessMaterialParameterParamTangentialAccel ParticleProcessMaterialParameter = 5
  ParticleProcessMaterialParameterParamDamping ParticleProcessMaterialParameter = 6
  ParticleProcessMaterialParameterParamAngle ParticleProcessMaterialParameter = 7
  ParticleProcessMaterialParameterParamScale ParticleProcessMaterialParameter = 8
  ParticleProcessMaterialParameterParamHueVariation ParticleProcessMaterialParameter = 9
  ParticleProcessMaterialParameterParamAnimSpeed ParticleProcessMaterialParameter = 10
  ParticleProcessMaterialParameterParamAnimOffset ParticleProcessMaterialParameter = 11
  ParticleProcessMaterialParameterParamMax ParticleProcessMaterialParameter = 15
  ParticleProcessMaterialParameterParamTurbVelInfluence ParticleProcessMaterialParameter = 13
  ParticleProcessMaterialParameterParamTurbInitDisplacement ParticleProcessMaterialParameter = 14
  ParticleProcessMaterialParameterParamTurbInfluenceOverLife ParticleProcessMaterialParameter = 12
)

type ParticleProcessMaterialParticleFlags int
const (
  ParticleProcessMaterialParticleFlagsParticleFlagAlignYToVelocity ParticleProcessMaterialParticleFlags = 0
  ParticleProcessMaterialParticleFlagsParticleFlagRotateY ParticleProcessMaterialParticleFlags = 1
  ParticleProcessMaterialParticleFlagsParticleFlagDisableZ ParticleProcessMaterialParticleFlags = 2
  ParticleProcessMaterialParticleFlagsParticleFlagMax ParticleProcessMaterialParticleFlags = 3
)

type ParticleProcessMaterialEmissionShape int
const (
  ParticleProcessMaterialEmissionShapeEmissionShapePoint ParticleProcessMaterialEmissionShape = 0
  ParticleProcessMaterialEmissionShapeEmissionShapeSphere ParticleProcessMaterialEmissionShape = 1
  ParticleProcessMaterialEmissionShapeEmissionShapeSphereSurface ParticleProcessMaterialEmissionShape = 2
  ParticleProcessMaterialEmissionShapeEmissionShapeBox ParticleProcessMaterialEmissionShape = 3
  ParticleProcessMaterialEmissionShapeEmissionShapePoints ParticleProcessMaterialEmissionShape = 4
  ParticleProcessMaterialEmissionShapeEmissionShapeDirectedPoints ParticleProcessMaterialEmissionShape = 5
  ParticleProcessMaterialEmissionShapeEmissionShapeRing ParticleProcessMaterialEmissionShape = 6
  ParticleProcessMaterialEmissionShapeEmissionShapeMax ParticleProcessMaterialEmissionShape = 7
)

type ParticleProcessMaterialSubEmitterMode int
const (
  ParticleProcessMaterialSubEmitterModeSubEmitterDisabled ParticleProcessMaterialSubEmitterMode = 0
  ParticleProcessMaterialSubEmitterModeSubEmitterConstant ParticleProcessMaterialSubEmitterMode = 1
  ParticleProcessMaterialSubEmitterModeSubEmitterAtEnd ParticleProcessMaterialSubEmitterMode = 2
  ParticleProcessMaterialSubEmitterModeSubEmitterAtCollision ParticleProcessMaterialSubEmitterMode = 3
  ParticleProcessMaterialSubEmitterModeSubEmitterMax ParticleProcessMaterialSubEmitterMode = 4
)

type ParticleProcessMaterialCollisionMode int
const (
  ParticleProcessMaterialCollisionModeCollisionDisabled ParticleProcessMaterialCollisionMode = 0
  ParticleProcessMaterialCollisionModeCollisionRigid ParticleProcessMaterialCollisionMode = 1
  ParticleProcessMaterialCollisionModeCollisionHideOnContact ParticleProcessMaterialCollisionMode = 2
  ParticleProcessMaterialCollisionModeCollisionMax ParticleProcessMaterialCollisionMode = 3
)

func (me *ParticleProcessMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ParticleProcessMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ParticleProcessMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ParticleProcessMaterial) SetDirection(degrees Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(degrees.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetDirection() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSpread(degrees float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetSpread() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetFlatness(amount float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetFlatness() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetParamMin(param ParticleProcessMaterialParameter, value float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295964248) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetParamMin(param ParticleProcessMaterialParameter, ) float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3903786503) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetParamMax(param ParticleProcessMaterialParameter, value float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295964248) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetParamMax(param ParticleProcessMaterialParameter, ) float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3903786503) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetParamTexture(param ParticleProcessMaterialParameter, texture Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 526976089) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetParamTexture(param ParticleProcessMaterialParameter, ) Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3489372978) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetColor(color Color, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetColor() Color {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetColorRamp(ramp Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetColorRamp() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetColorInitialRamp(ramp Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ramp.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetColorInitialRamp() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_initial_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, enable bool, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1711815571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, ) bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particle_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3895316907) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionShape(shape ParticleProcessMaterialEmissionShape, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 461501442) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionShape() ParticleProcessMaterialEmissionShape {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3719733018) // FIXME: should cache?
  var ret ParticleProcessMaterialEmissionShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionSphereRadius(radius float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionSphereRadius() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionBoxExtents(extents Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_box_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extents.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionBoxExtents() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_box_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionPointTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_point_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionPointTexture() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_point_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionNormalTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionNormalTexture() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionColorTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_color_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionColorTexture() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_color_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionPointCount(point_count int, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionPointCount() int {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionRingAxis(axis Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionRingAxis() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_axis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionRingHeight(height float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionRingHeight() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionRingRadius(radius float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionRingRadius() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionRingInnerRadius(inner_radius float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetEmissionRingInnerRadius() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) GetTurbulenceEnabled() bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetTurbulenceEnabled(turbulence_enabled bool, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseStrength() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseStrength(turbulence_noise_strength float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseScale() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseScale(turbulence_noise_scale float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeedRandom() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_speed_random")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_speed_random")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_speed_random), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeed() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeed(turbulence_noise_speed Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(turbulence_noise_speed.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetGravity() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetGravity(accel_vec Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(accel_vec.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) SetLifetimeRandomness(randomness float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&randomness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetLifetimeRandomness() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) GetSubEmitterMode() ParticleProcessMaterialSubEmitterMode {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2399052877) // FIXME: should cache?
  var ret ParticleProcessMaterialSubEmitterMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSubEmitterMode(mode ParticleProcessMaterialSubEmitterMode, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2161806672) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetSubEmitterFrequency() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSubEmitterFrequency(hz float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtEnd() int {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_amount_at_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtEnd(amount int, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_amount_at_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtCollision() int {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_amount_at_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtCollision(amount int, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_amount_at_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetSubEmitterKeepVelocity() bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_keep_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetSubEmitterKeepVelocity(enable bool, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_keep_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) SetAttractorInteractionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attractor_interaction_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) IsAttractorInteractionEnabled() bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_attractor_interaction_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetCollisionMode(mode ParticleProcessMaterialCollisionMode, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 653804659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetCollisionMode() ParticleProcessMaterialCollisionMode {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 139371864) // FIXME: should cache?
  var ret ParticleProcessMaterialCollisionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetCollisionUseScale(radius bool, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_use_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) IsCollisionUsingScale() bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collision_using_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetCollisionFriction(friction float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetCollisionFriction() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParticleProcessMaterial) SetCollisionBounce(bounce float32, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParticleProcessMaterial) GetCollisionBounce() float32 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ParticleProcessMaterial) GetPropLifetimeRandomness() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropLifetimeRandomness(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionShape() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionShape(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionSphereRadius() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionSphereRadius(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionBoxExtents() Vector3 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionBoxExtents(value Vector3) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionPointTexture() Texture2D {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionPointTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionNormalTexture() Texture2D {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionNormalTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionColorTexture() Texture2D {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionColorTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionPointCount() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionPointCount(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionRingAxis() Vector3 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionRingAxis(value Vector3) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionRingHeight() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionRingHeight(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionRingRadius() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionRingRadius(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropEmissionRingInnerRadius() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropEmissionRingInnerRadius(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropParticleFlagAlignY() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropParticleFlagAlignY(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropParticleFlagRotateY() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropParticleFlagRotateY(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropParticleFlagDisableZ() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropParticleFlagDisableZ(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropDirection() Vector3 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropDirection(value Vector3) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSpread() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSpread(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropFlatness() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropFlatness(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropGravity() Vector3 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropGravity(value Vector3) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropInitialVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropInitialVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropInitialVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropInitialVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngularVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngularVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngularVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngularVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngularVelocityCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngularVelocityCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropOrbitVelocityMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropOrbitVelocityMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropOrbitVelocityMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropOrbitVelocityMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropOrbitVelocityCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropOrbitVelocityCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropLinearAccelMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropLinearAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropLinearAccelMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropLinearAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropLinearAccelCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropLinearAccelCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropRadialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropRadialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropRadialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropRadialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropRadialAccelCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropRadialAccelCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTangentialAccelMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTangentialAccelMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTangentialAccelMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTangentialAccelMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTangentialAccelCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTangentialAccelCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropDampingMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropDampingMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropDampingMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropDampingMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropDampingCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropDampingCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngleMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngleMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngleMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngleMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAngleCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAngleCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropScaleMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropScaleMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropScaleMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropScaleMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropScaleCurve() any {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropScaleCurve(value any) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropColorRamp() GradientTexture1D {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropColorRamp(value GradientTexture1D) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropColorInitialRamp() GradientTexture1D {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropColorInitialRamp(value GradientTexture1D) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropHueVariationMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropHueVariationMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropHueVariationMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropHueVariationMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropHueVariationCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropHueVariationCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceEnabled() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceEnabled(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceNoiseStrength() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceNoiseStrength(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceNoiseScale() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceNoiseScale(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceNoiseSpeed() Vector3 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceNoiseSpeed(value Vector3) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceNoiseSpeedRandom() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceNoiseSpeedRandom(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceInfluenceMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceInfluenceMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceInfluenceMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceInfluenceMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceInitialDisplacementMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceInitialDisplacementMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceInitialDisplacementMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceInitialDisplacementMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropTurbulenceInfluenceOverLife() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropTurbulenceInfluenceOverLife(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimSpeedMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimSpeedMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimSpeedMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimSpeedMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimSpeedCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimSpeedCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimOffsetMin() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimOffsetMin(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimOffsetMax() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimOffsetMax(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAnimOffsetCurve() CurveTexture {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAnimOffsetCurve(value CurveTexture) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSubEmitterMode() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSubEmitterMode(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSubEmitterFrequency() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSubEmitterFrequency(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSubEmitterAmountAtEnd() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSubEmitterAmountAtEnd(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSubEmitterAmountAtCollision() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSubEmitterAmountAtCollision(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropSubEmitterKeepVelocity() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropSubEmitterKeepVelocity(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropAttractorInteractionEnabled() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropAttractorInteractionEnabled(value bool) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropCollisionMode() int {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropCollisionMode(value int) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropCollisionFriction() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropCollisionFriction(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropCollisionBounce() float32 {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropCollisionBounce(value float32) {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) GetPropCollisionUseScale() bool {
  panic("TODO: implement")
}

func (me *ParticleProcessMaterial) SetPropCollisionUseScale(value bool) {
  panic("TODO: implement")
}