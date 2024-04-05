// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForCPUParticles2DList struct {
  fnSetEmitting gdc.MethodBindPtr
  fnSetAmount gdc.MethodBindPtr
  fnSetLifetime gdc.MethodBindPtr
  fnSetOneShot gdc.MethodBindPtr
  fnSetPreProcessTime gdc.MethodBindPtr
  fnSetExplosivenessRatio gdc.MethodBindPtr
  fnSetRandomnessRatio gdc.MethodBindPtr
  fnSetLifetimeRandomness gdc.MethodBindPtr
  fnSetUseLocalCoordinates gdc.MethodBindPtr
  fnSetFixedFps gdc.MethodBindPtr
  fnSetFractionalDelta gdc.MethodBindPtr
  fnSetSpeedScale gdc.MethodBindPtr
  fnIsEmitting gdc.MethodBindPtr
  fnGetAmount gdc.MethodBindPtr
  fnGetLifetime gdc.MethodBindPtr
  fnGetOneShot gdc.MethodBindPtr
  fnGetPreProcessTime gdc.MethodBindPtr
  fnGetExplosivenessRatio gdc.MethodBindPtr
  fnGetRandomnessRatio gdc.MethodBindPtr
  fnGetLifetimeRandomness gdc.MethodBindPtr
  fnGetUseLocalCoordinates gdc.MethodBindPtr
  fnGetFixedFps gdc.MethodBindPtr
  fnGetFractionalDelta gdc.MethodBindPtr
  fnGetSpeedScale gdc.MethodBindPtr
  fnSetDrawOrder gdc.MethodBindPtr
  fnGetDrawOrder gdc.MethodBindPtr
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnRestart gdc.MethodBindPtr
  fnSetDirection gdc.MethodBindPtr
  fnGetDirection gdc.MethodBindPtr
  fnSetSpread gdc.MethodBindPtr
  fnGetSpread gdc.MethodBindPtr
  fnSetParamMin gdc.MethodBindPtr
  fnGetParamMin gdc.MethodBindPtr
  fnSetParamMax gdc.MethodBindPtr
  fnGetParamMax gdc.MethodBindPtr
  fnSetParamCurve gdc.MethodBindPtr
  fnGetParamCurve gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSetColorRamp gdc.MethodBindPtr
  fnGetColorRamp gdc.MethodBindPtr
  fnSetColorInitialRamp gdc.MethodBindPtr
  fnGetColorInitialRamp gdc.MethodBindPtr
  fnSetParticleFlag gdc.MethodBindPtr
  fnGetParticleFlag gdc.MethodBindPtr
  fnSetEmissionShape gdc.MethodBindPtr
  fnGetEmissionShape gdc.MethodBindPtr
  fnSetEmissionSphereRadius gdc.MethodBindPtr
  fnGetEmissionSphereRadius gdc.MethodBindPtr
  fnSetEmissionRectExtents gdc.MethodBindPtr
  fnGetEmissionRectExtents gdc.MethodBindPtr
  fnSetEmissionPoints gdc.MethodBindPtr
  fnGetEmissionPoints gdc.MethodBindPtr
  fnSetEmissionNormals gdc.MethodBindPtr
  fnGetEmissionNormals gdc.MethodBindPtr
  fnSetEmissionColors gdc.MethodBindPtr
  fnGetEmissionColors gdc.MethodBindPtr
  fnGetGravity gdc.MethodBindPtr
  fnSetGravity gdc.MethodBindPtr
  fnGetSplitScale gdc.MethodBindPtr
  fnSetSplitScale gdc.MethodBindPtr
  fnGetScaleCurveX gdc.MethodBindPtr
  fnSetScaleCurveX gdc.MethodBindPtr
  fnGetScaleCurveY gdc.MethodBindPtr
  fnSetScaleCurveY gdc.MethodBindPtr
  fnConvertFromParticles gdc.MethodBindPtr
}

var ptrsForCPUParticles2D ptrsForCPUParticles2DList

func initCPUParticles2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CPUParticles2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_emitting")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_amount")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_lifetime")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_one_shot")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_pre_process_time")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_explosiveness_ratio")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_randomness_ratio")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_lifetime_randomness")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_use_local_coordinates")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_fixed_fps")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_fractional_delta")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_speed_scale")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("is_emitting")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnIsEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_amount")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_lifetime")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_one_shot")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_pre_process_time")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_explosiveness_ratio")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_randomness_ratio")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_lifetime_randomness")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_use_local_coordinates")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_fixed_fps")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_fractional_delta")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_speed_scale")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_draw_order")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4183193490))
  }
  {
    methodName := StringNameFromStr("get_draw_order")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1668655735))
  }
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("restart")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnRestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_direction")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_direction")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_spread")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_spread")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_param_min")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3320615296))
  }
  {
    methodName := StringNameFromStr("get_param_min")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2038050600))
  }
  {
    methodName := StringNameFromStr("set_param_max")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3320615296))
  }
  {
    methodName := StringNameFromStr("get_param_max")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2038050600))
  }
  {
    methodName := StringNameFromStr("set_param_curve")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetParamCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2959350143))
  }
  {
    methodName := StringNameFromStr("get_param_curve")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetParamCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2603158474))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_color_ramp")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
  }
  {
    methodName := StringNameFromStr("get_color_ramp")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
  }
  {
    methodName := StringNameFromStr("set_color_initial_ramp")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
  }
  {
    methodName := StringNameFromStr("get_color_initial_ramp")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
  }
  {
    methodName := StringNameFromStr("set_particle_flag")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4178137949))
  }
  {
    methodName := StringNameFromStr("get_particle_flag")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2829976507))
  }
  {
    methodName := StringNameFromStr("set_emission_shape")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 393763892))
  }
  {
    methodName := StringNameFromStr("get_emission_shape")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740246024))
  }
  {
    methodName := StringNameFromStr("set_emission_sphere_radius")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_emission_sphere_radius")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_emission_rect_extents")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionRectExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_emission_rect_extents")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionRectExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_emission_points")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_emission_points")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
  {
    methodName := StringNameFromStr("set_emission_normals")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionNormals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_emission_normals")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionNormals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
  {
    methodName := StringNameFromStr("set_emission_colors")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetEmissionColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3546319833))
  }
  {
    methodName := StringNameFromStr("get_emission_colors")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetEmissionColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
  }
  {
    methodName := StringNameFromStr("get_gravity")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_gravity")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_split_scale")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetSplitScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_split_scale")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetSplitScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_scale_curve_x")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetScaleCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
  }
  {
    methodName := StringNameFromStr("set_scale_curve_x")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetScaleCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
  }
  {
    methodName := StringNameFromStr("get_scale_curve_y")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnGetScaleCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
  }
  {
    methodName := StringNameFromStr("set_scale_curve_y")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnSetScaleCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
  }
  {
    methodName := StringNameFromStr("convert_from_particles")
    defer methodName.Destroy()
    ptrsForCPUParticles2D.fnConvertFromParticles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmitting), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetAmount(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetLifetime(secs float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetOneShot(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetPreProcessTime(secs float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetPreProcessTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetExplosivenessRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetRandomnessRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetRandomnessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetLifetimeRandomness(random float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&random) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetUseLocalCoordinates(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetFixedFps(fps int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetFixedFps), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetFractionalDelta(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetFractionalDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetSpeedScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) IsEmitting() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnIsEmitting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetAmount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetLifetime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetOneShot() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetPreProcessTime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetPreProcessTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetExplosivenessRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetRandomnessRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetRandomnessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetLifetimeRandomness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetUseLocalCoordinates() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetFixedFps() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetFixedFps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetFractionalDelta() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetFractionalDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) GetSpeedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetDrawOrder(order CPUParticles2DDrawOrder, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetDrawOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetDrawOrder() CPUParticles2DDrawOrder {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CPUParticles2DDrawOrder

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetDrawOrder), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CPUParticles2D) SetTexture(texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) Restart()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnRestart), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) SetDirection(direction Vector2, )  {
  cargs := []gdc.ConstTypePtr{direction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetDirection() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetSpread(spread float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetSpread), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetSpread() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetSpread), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamMin(param CPUParticles2DParameter, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetParamMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamMin(param CPUParticles2DParameter, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetParamMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamMax(param CPUParticles2DParameter, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetParamMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamMax(param CPUParticles2DParameter, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetParamMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetParamCurve(param CPUParticles2DParameter, curve Curve, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetParamCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParamCurve(param CPUParticles2DParameter, ) Curve {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetParamCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColorRamp(ramp Gradient, )  {
  cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetColorRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColorRamp() Gradient {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetColorRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetColorInitialRamp(ramp Gradient, )  {
  cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetColorInitialRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetColorInitialRamp() Gradient {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetColorInitialRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetParticleFlag(particle_flag CPUParticles2DParticleFlags, enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetParticleFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetParticleFlag(particle_flag CPUParticles2DParticleFlags, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&particle_flag)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetParticleFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetEmissionShape(shape CPUParticles2DEmissionShape, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionShape() CPUParticles2DEmissionShape {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CPUParticles2DEmissionShape

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CPUParticles2D) SetEmissionSphereRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionSphereRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetEmissionRectExtents(extents Vector2, )  {
  cargs := []gdc.ConstTypePtr{extents.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionRectExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionRectExtents() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionRectExtents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionPoints(array PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionPoints() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionNormals(array PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionNormals), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionNormals() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionNormals), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetEmissionColors(array PackedColorArray, )  {
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetEmissionColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetEmissionColors() PackedColorArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetEmissionColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) GetGravity() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetGravity(accel_vec Vector2, )  {
  cargs := []gdc.ConstTypePtr{accel_vec.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetSplitScale() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetSplitScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CPUParticles2D) SetSplitScale(split_scale bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&split_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetSplitScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetScaleCurveX() Curve {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetScaleCurveX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetScaleCurveX(scale_curve Curve, )  {
  cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetScaleCurveX), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) GetScaleCurveY() Curve {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnGetScaleCurveY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CPUParticles2D) SetScaleCurveY(scale_curve Curve, )  {
  cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnSetScaleCurveY), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CPUParticles2D) ConvertFromParticles(particles Node, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles2D.fnConvertFromParticles), me.obj, unsafe.SliceData(cargs), nil)

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
