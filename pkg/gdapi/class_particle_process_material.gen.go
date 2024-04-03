// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ParticleProcessMaterial struct {
  Material
}

func (me *ParticleProcessMaterial) BaseClass() string {
  return "ParticleProcessMaterial"
}

func NewParticleProcessMaterial() *ParticleProcessMaterial {
  str := StringNameFromStr("ParticleProcessMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ParticleProcessMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
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
  ParticleProcessMaterialParameterParamRadialVelocity ParticleProcessMaterialParameter = 15
  ParticleProcessMaterialParameterParamDirectionalVelocity ParticleProcessMaterialParameter = 16
  ParticleProcessMaterialParameterParamScaleOverVelocity ParticleProcessMaterialParameter = 17
  ParticleProcessMaterialParameterParamMax ParticleProcessMaterialParameter = 18
  ParticleProcessMaterialParameterParamTurbVelInfluence ParticleProcessMaterialParameter = 13
  ParticleProcessMaterialParameterParamTurbInitDisplacement ParticleProcessMaterialParameter = 14
  ParticleProcessMaterialParameterParamTurbInfluenceOverLife ParticleProcessMaterialParameter = 12
)

type ParticleProcessMaterialParticleFlags int
const (
  ParticleProcessMaterialParticleFlagsParticleFlagAlignYToVelocity ParticleProcessMaterialParticleFlags = 0
  ParticleProcessMaterialParticleFlagsParticleFlagRotateY ParticleProcessMaterialParticleFlags = 1
  ParticleProcessMaterialParticleFlagsParticleFlagDisableZ ParticleProcessMaterialParticleFlags = 2
  ParticleProcessMaterialParticleFlagsParticleFlagDampingAsFriction ParticleProcessMaterialParticleFlags = 3
  ParticleProcessMaterialParticleFlagsParticleFlagMax ParticleProcessMaterialParticleFlags = 4
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetInheritVelocityRatio(ratio float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inherit_velocity_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetInheritVelocityRatio() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inherit_velocity_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetSpread(degrees float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetSpread() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetFlatness(amount float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetFlatness() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flatness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetParamMin(param ParticleProcessMaterialParameter, value float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295964248) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetParamMin(param ParticleProcessMaterialParameter, ) float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3903786503) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetParamMax(param ParticleProcessMaterialParameter, value float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295964248) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetParamMax(param ParticleProcessMaterialParameter, ) float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3903786503) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetAlphaCurve(curve Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetAlphaCurve() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetEmissionCurve(curve Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionCurve() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetVelocityLimitCurve(curve Texture2D, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity_limit_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetVelocityLimitCurve() Texture2D {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_limit_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetVelocityPivot(pivot Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity_pivot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pivot.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetVelocityPivot() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_pivot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3783033775) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  var ret ParticleProcessMaterialEmissionShape

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ParticleProcessMaterial) SetEmissionSphereRadius(radius float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionSphereRadius() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_sphere_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetEmissionPointCount(point_count int64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionPointCount() int64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetEmissionRingHeight(height float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionRingHeight() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetEmissionRingRadius(radius float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionRingRadius() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetEmissionRingInnerRadius(inner_radius float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_radius), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionRingInnerRadius() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_ring_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetEmissionShapeOffset(emission_shape_offset Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(emission_shape_offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionShapeOffset() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) SetEmissionShapeScale(emission_shape_scale Vector3, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_shape_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(emission_shape_scale.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetEmissionShapeScale() Vector3 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_shape_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParticleProcessMaterial) GetTurbulenceEnabled() bool {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseStrength() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseStrength(turbulence_noise_strength float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_strength), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseScale() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseScale(turbulence_noise_scale float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbulence_noise_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_scale), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeedRandom() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbulence_noise_speed_random")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float64, )  {
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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

func  (me *ParticleProcessMaterial) SetLifetimeRandomness(randomness float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&randomness), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetLifetimeRandomness() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime_randomness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) GetSubEmitterMode() ParticleProcessMaterialSubEmitterMode {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2399052877) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret ParticleProcessMaterialSubEmitterMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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

func  (me *ParticleProcessMaterial) GetSubEmitterFrequency() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetSubEmitterFrequency(hz float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtEnd() int64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_amount_at_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtEnd(amount int64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter_amount_at_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetSubEmitterAmountAtCollision() int64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter_amount_at_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetSubEmitterAmountAtCollision(amount int64, )  {
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  var ret ParticleProcessMaterialCollisionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetCollisionFriction(friction float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetCollisionFriction() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ParticleProcessMaterial) SetCollisionBounce(bounce float64, )  {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParticleProcessMaterial) GetCollisionBounce() float64 {
  classNameV := StringNameFromStr("ParticleProcessMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
