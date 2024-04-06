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

type ptrsForParticleProcessMaterialList struct {
	fnSetDirection                   gdc.MethodBindPtr
	fnGetDirection                   gdc.MethodBindPtr
	fnSetInheritVelocityRatio        gdc.MethodBindPtr
	fnGetInheritVelocityRatio        gdc.MethodBindPtr
	fnSetSpread                      gdc.MethodBindPtr
	fnGetSpread                      gdc.MethodBindPtr
	fnSetFlatness                    gdc.MethodBindPtr
	fnGetFlatness                    gdc.MethodBindPtr
	fnSetParamMin                    gdc.MethodBindPtr
	fnGetParamMin                    gdc.MethodBindPtr
	fnSetParamMax                    gdc.MethodBindPtr
	fnGetParamMax                    gdc.MethodBindPtr
	fnSetParamTexture                gdc.MethodBindPtr
	fnGetParamTexture                gdc.MethodBindPtr
	fnSetColor                       gdc.MethodBindPtr
	fnGetColor                       gdc.MethodBindPtr
	fnSetColorRamp                   gdc.MethodBindPtr
	fnGetColorRamp                   gdc.MethodBindPtr
	fnSetAlphaCurve                  gdc.MethodBindPtr
	fnGetAlphaCurve                  gdc.MethodBindPtr
	fnSetEmissionCurve               gdc.MethodBindPtr
	fnGetEmissionCurve               gdc.MethodBindPtr
	fnSetColorInitialRamp            gdc.MethodBindPtr
	fnGetColorInitialRamp            gdc.MethodBindPtr
	fnSetVelocityLimitCurve          gdc.MethodBindPtr
	fnGetVelocityLimitCurve          gdc.MethodBindPtr
	fnSetParticleFlag                gdc.MethodBindPtr
	fnGetParticleFlag                gdc.MethodBindPtr
	fnSetVelocityPivot               gdc.MethodBindPtr
	fnGetVelocityPivot               gdc.MethodBindPtr
	fnSetEmissionShape               gdc.MethodBindPtr
	fnGetEmissionShape               gdc.MethodBindPtr
	fnSetEmissionSphereRadius        gdc.MethodBindPtr
	fnGetEmissionSphereRadius        gdc.MethodBindPtr
	fnSetEmissionBoxExtents          gdc.MethodBindPtr
	fnGetEmissionBoxExtents          gdc.MethodBindPtr
	fnSetEmissionPointTexture        gdc.MethodBindPtr
	fnGetEmissionPointTexture        gdc.MethodBindPtr
	fnSetEmissionNormalTexture       gdc.MethodBindPtr
	fnGetEmissionNormalTexture       gdc.MethodBindPtr
	fnSetEmissionColorTexture        gdc.MethodBindPtr
	fnGetEmissionColorTexture        gdc.MethodBindPtr
	fnSetEmissionPointCount          gdc.MethodBindPtr
	fnGetEmissionPointCount          gdc.MethodBindPtr
	fnSetEmissionRingAxis            gdc.MethodBindPtr
	fnGetEmissionRingAxis            gdc.MethodBindPtr
	fnSetEmissionRingHeight          gdc.MethodBindPtr
	fnGetEmissionRingHeight          gdc.MethodBindPtr
	fnSetEmissionRingRadius          gdc.MethodBindPtr
	fnGetEmissionRingRadius          gdc.MethodBindPtr
	fnSetEmissionRingInnerRadius     gdc.MethodBindPtr
	fnGetEmissionRingInnerRadius     gdc.MethodBindPtr
	fnSetEmissionShapeOffset         gdc.MethodBindPtr
	fnGetEmissionShapeOffset         gdc.MethodBindPtr
	fnSetEmissionShapeScale          gdc.MethodBindPtr
	fnGetEmissionShapeScale          gdc.MethodBindPtr
	fnGetTurbulenceEnabled           gdc.MethodBindPtr
	fnSetTurbulenceEnabled           gdc.MethodBindPtr
	fnGetTurbulenceNoiseStrength     gdc.MethodBindPtr
	fnSetTurbulenceNoiseStrength     gdc.MethodBindPtr
	fnGetTurbulenceNoiseScale        gdc.MethodBindPtr
	fnSetTurbulenceNoiseScale        gdc.MethodBindPtr
	fnGetTurbulenceNoiseSpeedRandom  gdc.MethodBindPtr
	fnSetTurbulenceNoiseSpeedRandom  gdc.MethodBindPtr
	fnGetTurbulenceNoiseSpeed        gdc.MethodBindPtr
	fnSetTurbulenceNoiseSpeed        gdc.MethodBindPtr
	fnGetGravity                     gdc.MethodBindPtr
	fnSetGravity                     gdc.MethodBindPtr
	fnSetLifetimeRandomness          gdc.MethodBindPtr
	fnGetLifetimeRandomness          gdc.MethodBindPtr
	fnGetSubEmitterMode              gdc.MethodBindPtr
	fnSetSubEmitterMode              gdc.MethodBindPtr
	fnGetSubEmitterFrequency         gdc.MethodBindPtr
	fnSetSubEmitterFrequency         gdc.MethodBindPtr
	fnGetSubEmitterAmountAtEnd       gdc.MethodBindPtr
	fnSetSubEmitterAmountAtEnd       gdc.MethodBindPtr
	fnGetSubEmitterAmountAtCollision gdc.MethodBindPtr
	fnSetSubEmitterAmountAtCollision gdc.MethodBindPtr
	fnGetSubEmitterKeepVelocity      gdc.MethodBindPtr
	fnSetSubEmitterKeepVelocity      gdc.MethodBindPtr
	fnSetAttractorInteractionEnabled gdc.MethodBindPtr
	fnIsAttractorInteractionEnabled  gdc.MethodBindPtr
	fnSetCollisionMode               gdc.MethodBindPtr
	fnGetCollisionMode               gdc.MethodBindPtr
	fnSetCollisionUseScale           gdc.MethodBindPtr
	fnIsCollisionUsingScale          gdc.MethodBindPtr
	fnSetCollisionFriction           gdc.MethodBindPtr
	fnGetCollisionFriction           gdc.MethodBindPtr
	fnSetCollisionBounce             gdc.MethodBindPtr
	fnGetCollisionBounce             gdc.MethodBindPtr
}

var ptrsForParticleProcessMaterial ptrsForParticleProcessMaterialList

func initParticleProcessMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ParticleProcessMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_direction")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_direction")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_inherit_velocity_ratio")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetInheritVelocityRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_inherit_velocity_ratio")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetInheritVelocityRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_spread")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_spread")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_flatness")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetFlatness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_flatness")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetFlatness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_param_min")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2295964248))
	}
	{
		methodName := StringNameFromStr("get_param_min")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3903786503))
	}
	{
		methodName := StringNameFromStr("set_param_max")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2295964248))
	}
	{
		methodName := StringNameFromStr("get_param_max")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3903786503))
	}
	{
		methodName := StringNameFromStr("set_param_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetParamTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 526976089))
	}
	{
		methodName := StringNameFromStr("get_param_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetParamTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3489372978))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_color_ramp")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_color_ramp")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_alpha_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetAlphaCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_alpha_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetAlphaCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_emission_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_emission_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_color_initial_ramp")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_color_initial_ramp")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_velocity_limit_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetVelocityLimitCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_velocity_limit_curve")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetVelocityLimitCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_particle_flag")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1711815571))
	}
	{
		methodName := StringNameFromStr("get_particle_flag")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3895316907))
	}
	{
		methodName := StringNameFromStr("set_velocity_pivot")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetVelocityPivot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_velocity_pivot")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetVelocityPivot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
	}
	{
		methodName := StringNameFromStr("set_emission_shape")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 461501442))
	}
	{
		methodName := StringNameFromStr("get_emission_shape")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3719733018))
	}
	{
		methodName := StringNameFromStr("set_emission_sphere_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_sphere_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_box_extents")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionBoxExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_box_extents")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionBoxExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_emission_point_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionPointTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_emission_point_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionPointTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_emission_normal_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionNormalTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_emission_normal_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionNormalTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_emission_color_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionColorTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_emission_color_texture")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionColorTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_emission_point_count")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_emission_point_count")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_axis")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionRingAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_axis")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionRingAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_height")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionRingHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_height")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionRingHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionRingRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionRingRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_inner_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionRingInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_inner_radius")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionRingInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_shape_offset")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionShapeOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_shape_offset")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionShapeOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_emission_shape_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetEmissionShapeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_shape_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetEmissionShapeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_turbulence_enabled")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetTurbulenceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_turbulence_enabled")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetTurbulenceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_turbulence_noise_strength")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_turbulence_noise_strength")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_turbulence_noise_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_turbulence_noise_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_turbulence_noise_speed_random")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseSpeedRandom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_turbulence_noise_speed_random")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseSpeedRandom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_turbulence_noise_speed")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_turbulence_noise_speed")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_gravity")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_gravity")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_lifetime_randomness")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_lifetime_randomness")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter_mode")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSubEmitterMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2399052877))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter_mode")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSubEmitterMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2161806672))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter_frequency")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSubEmitterFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter_frequency")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSubEmitterFrequency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter_amount_at_end")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSubEmitterAmountAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter_amount_at_end")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSubEmitterAmountAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter_amount_at_collision")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSubEmitterAmountAtCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter_amount_at_collision")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSubEmitterAmountAtCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter_keep_velocity")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetSubEmitterKeepVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter_keep_velocity")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetSubEmitterKeepVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_attractor_interaction_enabled")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetAttractorInteractionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_attractor_interaction_enabled")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnIsAttractorInteractionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collision_mode")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetCollisionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 653804659))
	}
	{
		methodName := StringNameFromStr("get_collision_mode")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetCollisionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 139371864))
	}
	{
		methodName := StringNameFromStr("set_collision_use_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetCollisionUseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collision_using_scale")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnIsCollisionUsingScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collision_friction")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetCollisionFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_collision_friction")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetCollisionFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_collision_bounce")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnSetCollisionBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_collision_bounce")
		defer methodName.Destroy()
		ptrsForParticleProcessMaterial.fnGetCollisionBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

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
	ParticleProcessMaterialParameterParamAngularVelocity       ParticleProcessMaterialParameter = 1
	ParticleProcessMaterialParameterParamOrbitVelocity         ParticleProcessMaterialParameter = 2
	ParticleProcessMaterialParameterParamLinearAccel           ParticleProcessMaterialParameter = 3
	ParticleProcessMaterialParameterParamRadialAccel           ParticleProcessMaterialParameter = 4
	ParticleProcessMaterialParameterParamTangentialAccel       ParticleProcessMaterialParameter = 5
	ParticleProcessMaterialParameterParamDamping               ParticleProcessMaterialParameter = 6
	ParticleProcessMaterialParameterParamAngle                 ParticleProcessMaterialParameter = 7
	ParticleProcessMaterialParameterParamScale                 ParticleProcessMaterialParameter = 8
	ParticleProcessMaterialParameterParamHueVariation          ParticleProcessMaterialParameter = 9
	ParticleProcessMaterialParameterParamAnimSpeed             ParticleProcessMaterialParameter = 10
	ParticleProcessMaterialParameterParamAnimOffset            ParticleProcessMaterialParameter = 11
	ParticleProcessMaterialParameterParamRadialVelocity        ParticleProcessMaterialParameter = 15
	ParticleProcessMaterialParameterParamDirectionalVelocity   ParticleProcessMaterialParameter = 16
	ParticleProcessMaterialParameterParamScaleOverVelocity     ParticleProcessMaterialParameter = 17
	ParticleProcessMaterialParameterParamMax                   ParticleProcessMaterialParameter = 18
	ParticleProcessMaterialParameterParamTurbVelInfluence      ParticleProcessMaterialParameter = 13
	ParticleProcessMaterialParameterParamTurbInitDisplacement  ParticleProcessMaterialParameter = 14
	ParticleProcessMaterialParameterParamTurbInfluenceOverLife ParticleProcessMaterialParameter = 12
)

type ParticleProcessMaterialParticleFlags int

const (
	ParticleProcessMaterialParticleFlagsParticleFlagAlignYToVelocity  ParticleProcessMaterialParticleFlags = 0
	ParticleProcessMaterialParticleFlagsParticleFlagRotateY           ParticleProcessMaterialParticleFlags = 1
	ParticleProcessMaterialParticleFlagsParticleFlagDisableZ          ParticleProcessMaterialParticleFlags = 2
	ParticleProcessMaterialParticleFlagsParticleFlagDampingAsFriction ParticleProcessMaterialParticleFlags = 3
	ParticleProcessMaterialParticleFlagsParticleFlagMax               ParticleProcessMaterialParticleFlags = 4
)

type ParticleProcessMaterialEmissionShape int

const (
	ParticleProcessMaterialEmissionShapeEmissionShapePoint          ParticleProcessMaterialEmissionShape = 0
	ParticleProcessMaterialEmissionShapeEmissionShapeSphere         ParticleProcessMaterialEmissionShape = 1
	ParticleProcessMaterialEmissionShapeEmissionShapeSphereSurface  ParticleProcessMaterialEmissionShape = 2
	ParticleProcessMaterialEmissionShapeEmissionShapeBox            ParticleProcessMaterialEmissionShape = 3
	ParticleProcessMaterialEmissionShapeEmissionShapePoints         ParticleProcessMaterialEmissionShape = 4
	ParticleProcessMaterialEmissionShapeEmissionShapeDirectedPoints ParticleProcessMaterialEmissionShape = 5
	ParticleProcessMaterialEmissionShapeEmissionShapeRing           ParticleProcessMaterialEmissionShape = 6
	ParticleProcessMaterialEmissionShapeEmissionShapeMax            ParticleProcessMaterialEmissionShape = 7
)

type ParticleProcessMaterialSubEmitterMode int

const (
	ParticleProcessMaterialSubEmitterModeSubEmitterDisabled    ParticleProcessMaterialSubEmitterMode = 0
	ParticleProcessMaterialSubEmitterModeSubEmitterConstant    ParticleProcessMaterialSubEmitterMode = 1
	ParticleProcessMaterialSubEmitterModeSubEmitterAtEnd       ParticleProcessMaterialSubEmitterMode = 2
	ParticleProcessMaterialSubEmitterModeSubEmitterAtCollision ParticleProcessMaterialSubEmitterMode = 3
	ParticleProcessMaterialSubEmitterModeSubEmitterMax         ParticleProcessMaterialSubEmitterMode = 4
)

type ParticleProcessMaterialCollisionMode int

const (
	ParticleProcessMaterialCollisionModeCollisionDisabled      ParticleProcessMaterialCollisionMode = 0
	ParticleProcessMaterialCollisionModeCollisionRigid         ParticleProcessMaterialCollisionMode = 1
	ParticleProcessMaterialCollisionModeCollisionHideOnContact ParticleProcessMaterialCollisionMode = 2
	ParticleProcessMaterialCollisionModeCollisionMax           ParticleProcessMaterialCollisionMode = 3
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

func (me *ParticleProcessMaterial) SetDirection(degrees Vector3) {
	cargs := []gdc.ConstTypePtr{degrees.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetDirection() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetInheritVelocityRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetInheritVelocityRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetInheritVelocityRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetInheritVelocityRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetSpread(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSpread), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetSpread() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSpread), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetFlatness(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetFlatness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetFlatness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetFlatness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetParamMin(param ParticleProcessMaterialParameter, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetParamMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetParamMin(param ParticleProcessMaterialParameter) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetParamMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetParamMax(param ParticleProcessMaterialParameter, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetParamMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetParamMax(param ParticleProcessMaterialParameter) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetParamMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetParamTexture(param ParticleProcessMaterialParameter, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetParamTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetParamTexture(param ParticleProcessMaterialParameter) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetParamTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetColorRamp(ramp Texture2D) {
	cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetColorRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetColorRamp() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetColorRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetAlphaCurve(curve Texture2D) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetAlphaCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetAlphaCurve() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetAlphaCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionCurve(curve Texture2D) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionCurve() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetColorInitialRamp(ramp Texture2D) {
	cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetColorInitialRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetColorInitialRamp() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetColorInitialRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetVelocityLimitCurve(curve Texture2D) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetVelocityLimitCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetVelocityLimitCurve() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetVelocityLimitCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetParticleFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetParticleFlag(particle_flag ParticleProcessMaterialParticleFlags) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&particle_flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetParticleFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetVelocityPivot(pivot Vector3) {
	cargs := []gdc.ConstTypePtr{pivot.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetVelocityPivot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetVelocityPivot() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetVelocityPivot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionShape(shape ParticleProcessMaterialEmissionShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionShape() ParticleProcessMaterialEmissionShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ParticleProcessMaterialEmissionShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ParticleProcessMaterial) SetEmissionSphereRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionSphereRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetEmissionBoxExtents(extents Vector3) {
	cargs := []gdc.ConstTypePtr{extents.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionBoxExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionBoxExtents() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionBoxExtents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionPointTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionPointTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionPointTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionPointTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionNormalTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionNormalTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionNormalTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionNormalTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionColorTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionColorTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionColorTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionColorTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionPointCount(point_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionPointCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionPointCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetEmissionRingAxis(axis Vector3) {
	cargs := []gdc.ConstTypePtr{axis.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionRingAxis), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionRingAxis() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionRingAxis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionRingHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionRingHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionRingHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionRingHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetEmissionRingRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionRingRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionRingRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionRingRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetEmissionRingInnerRadius(inner_radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionRingInnerRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionRingInnerRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionRingInnerRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetEmissionShapeOffset(emission_shape_offset Vector3) {
	cargs := []gdc.ConstTypePtr{emission_shape_offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionShapeOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionShapeOffset() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionShapeOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetEmissionShapeScale(emission_shape_scale Vector3) {
	cargs := []gdc.ConstTypePtr{emission_shape_scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetEmissionShapeScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetEmissionShapeScale() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetEmissionShapeScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) GetTurbulenceEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetTurbulenceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetTurbulenceEnabled(turbulence_enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetTurbulenceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetTurbulenceNoiseStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetTurbulenceNoiseStrength(turbulence_noise_strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetTurbulenceNoiseScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetTurbulenceNoiseScale(turbulence_noise_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeedRandom() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseSpeedRandom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbulence_noise_speed_random)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseSpeedRandom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetTurbulenceNoiseSpeed() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetTurbulenceNoiseSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetTurbulenceNoiseSpeed(turbulence_noise_speed Vector3) {
	cargs := []gdc.ConstTypePtr{turbulence_noise_speed.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetTurbulenceNoiseSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetGravity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParticleProcessMaterial) SetGravity(accel_vec Vector3) {
	cargs := []gdc.ConstTypePtr{accel_vec.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) SetLifetimeRandomness(randomness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&randomness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetLifetimeRandomness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) GetSubEmitterMode() ParticleProcessMaterialSubEmitterMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ParticleProcessMaterialSubEmitterMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSubEmitterMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ParticleProcessMaterial) SetSubEmitterMode(mode ParticleProcessMaterialSubEmitterMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSubEmitterMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetSubEmitterFrequency() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSubEmitterFrequency), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetSubEmitterFrequency(hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSubEmitterFrequency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetSubEmitterAmountAtEnd() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSubEmitterAmountAtEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetSubEmitterAmountAtEnd(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSubEmitterAmountAtEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetSubEmitterAmountAtCollision() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSubEmitterAmountAtCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetSubEmitterAmountAtCollision(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSubEmitterAmountAtCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetSubEmitterKeepVelocity() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetSubEmitterKeepVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetSubEmitterKeepVelocity(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetSubEmitterKeepVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) SetAttractorInteractionEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetAttractorInteractionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) IsAttractorInteractionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnIsAttractorInteractionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetCollisionMode(mode ParticleProcessMaterialCollisionMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetCollisionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetCollisionMode() ParticleProcessMaterialCollisionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ParticleProcessMaterialCollisionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetCollisionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ParticleProcessMaterial) SetCollisionUseScale(radius bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetCollisionUseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) IsCollisionUsingScale() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnIsCollisionUsingScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetCollisionFriction(friction float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetCollisionFriction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetCollisionFriction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetCollisionFriction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ParticleProcessMaterial) SetCollisionBounce(bounce float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnSetCollisionBounce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParticleProcessMaterial) GetCollisionBounce() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParticleProcessMaterial.fnGetCollisionBounce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
