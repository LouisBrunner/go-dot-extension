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

type ptrsForCPUParticles3DList struct {
	fnSetEmitting                gdc.MethodBindPtr
	fnSetAmount                  gdc.MethodBindPtr
	fnSetLifetime                gdc.MethodBindPtr
	fnSetOneShot                 gdc.MethodBindPtr
	fnSetPreProcessTime          gdc.MethodBindPtr
	fnSetExplosivenessRatio      gdc.MethodBindPtr
	fnSetRandomnessRatio         gdc.MethodBindPtr
	fnSetLifetimeRandomness      gdc.MethodBindPtr
	fnSetUseLocalCoordinates     gdc.MethodBindPtr
	fnSetFixedFps                gdc.MethodBindPtr
	fnSetFractionalDelta         gdc.MethodBindPtr
	fnSetSpeedScale              gdc.MethodBindPtr
	fnIsEmitting                 gdc.MethodBindPtr
	fnGetAmount                  gdc.MethodBindPtr
	fnGetLifetime                gdc.MethodBindPtr
	fnGetOneShot                 gdc.MethodBindPtr
	fnGetPreProcessTime          gdc.MethodBindPtr
	fnGetExplosivenessRatio      gdc.MethodBindPtr
	fnGetRandomnessRatio         gdc.MethodBindPtr
	fnGetLifetimeRandomness      gdc.MethodBindPtr
	fnGetUseLocalCoordinates     gdc.MethodBindPtr
	fnGetFixedFps                gdc.MethodBindPtr
	fnGetFractionalDelta         gdc.MethodBindPtr
	fnGetSpeedScale              gdc.MethodBindPtr
	fnSetDrawOrder               gdc.MethodBindPtr
	fnGetDrawOrder               gdc.MethodBindPtr
	fnSetMesh                    gdc.MethodBindPtr
	fnGetMesh                    gdc.MethodBindPtr
	fnRestart                    gdc.MethodBindPtr
	fnSetDirection               gdc.MethodBindPtr
	fnGetDirection               gdc.MethodBindPtr
	fnSetSpread                  gdc.MethodBindPtr
	fnGetSpread                  gdc.MethodBindPtr
	fnSetFlatness                gdc.MethodBindPtr
	fnGetFlatness                gdc.MethodBindPtr
	fnSetParamMin                gdc.MethodBindPtr
	fnGetParamMin                gdc.MethodBindPtr
	fnSetParamMax                gdc.MethodBindPtr
	fnGetParamMax                gdc.MethodBindPtr
	fnSetParamCurve              gdc.MethodBindPtr
	fnGetParamCurve              gdc.MethodBindPtr
	fnSetColor                   gdc.MethodBindPtr
	fnGetColor                   gdc.MethodBindPtr
	fnSetColorRamp               gdc.MethodBindPtr
	fnGetColorRamp               gdc.MethodBindPtr
	fnSetColorInitialRamp        gdc.MethodBindPtr
	fnGetColorInitialRamp        gdc.MethodBindPtr
	fnSetParticleFlag            gdc.MethodBindPtr
	fnGetParticleFlag            gdc.MethodBindPtr
	fnSetEmissionShape           gdc.MethodBindPtr
	fnGetEmissionShape           gdc.MethodBindPtr
	fnSetEmissionSphereRadius    gdc.MethodBindPtr
	fnGetEmissionSphereRadius    gdc.MethodBindPtr
	fnSetEmissionBoxExtents      gdc.MethodBindPtr
	fnGetEmissionBoxExtents      gdc.MethodBindPtr
	fnSetEmissionPoints          gdc.MethodBindPtr
	fnGetEmissionPoints          gdc.MethodBindPtr
	fnSetEmissionNormals         gdc.MethodBindPtr
	fnGetEmissionNormals         gdc.MethodBindPtr
	fnSetEmissionColors          gdc.MethodBindPtr
	fnGetEmissionColors          gdc.MethodBindPtr
	fnSetEmissionRingAxis        gdc.MethodBindPtr
	fnGetEmissionRingAxis        gdc.MethodBindPtr
	fnSetEmissionRingHeight      gdc.MethodBindPtr
	fnGetEmissionRingHeight      gdc.MethodBindPtr
	fnSetEmissionRingRadius      gdc.MethodBindPtr
	fnGetEmissionRingRadius      gdc.MethodBindPtr
	fnSetEmissionRingInnerRadius gdc.MethodBindPtr
	fnGetEmissionRingInnerRadius gdc.MethodBindPtr
	fnGetGravity                 gdc.MethodBindPtr
	fnSetGravity                 gdc.MethodBindPtr
	fnGetSplitScale              gdc.MethodBindPtr
	fnSetSplitScale              gdc.MethodBindPtr
	fnGetScaleCurveX             gdc.MethodBindPtr
	fnSetScaleCurveX             gdc.MethodBindPtr
	fnGetScaleCurveY             gdc.MethodBindPtr
	fnSetScaleCurveY             gdc.MethodBindPtr
	fnGetScaleCurveZ             gdc.MethodBindPtr
	fnSetScaleCurveZ             gdc.MethodBindPtr
	fnConvertFromParticles       gdc.MethodBindPtr
}

var ptrsForCPUParticles3D ptrsForCPUParticles3DList

func initCPUParticles3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CPUParticles3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_emitting")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_amount")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_lifetime")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_one_shot")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_pre_process_time")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_explosiveness_ratio")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_randomness_ratio")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_lifetime_randomness")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_use_local_coordinates")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_fixed_fps")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_fractional_delta")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_speed_scale")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("is_emitting")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnIsEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_amount")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_lifetime")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_one_shot")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_pre_process_time")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_explosiveness_ratio")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_randomness_ratio")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_lifetime_randomness")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetLifetimeRandomness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_use_local_coordinates")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_fixed_fps")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_fractional_delta")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_speed_scale")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_draw_order")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1427401774))
	}
	{
		methodName := StringNameFromStr("get_draw_order")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321900776))
	}
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
	}
	{
		methodName := StringNameFromStr("restart")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnRestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_direction")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_direction")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_spread")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_spread")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_flatness")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetFlatness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_flatness")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetFlatness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_param_min")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 557936109))
	}
	{
		methodName := StringNameFromStr("get_param_min")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetParamMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 597646162))
	}
	{
		methodName := StringNameFromStr("set_param_max")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 557936109))
	}
	{
		methodName := StringNameFromStr("get_param_max")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetParamMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 597646162))
	}
	{
		methodName := StringNameFromStr("set_param_curve")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetParamCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4044142537))
	}
	{
		methodName := StringNameFromStr("get_param_curve")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetParamCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4132790277))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_color_ramp")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
	}
	{
		methodName := StringNameFromStr("get_color_ramp")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
	}
	{
		methodName := StringNameFromStr("set_color_initial_ramp")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
	}
	{
		methodName := StringNameFromStr("get_color_initial_ramp")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetColorInitialRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
	}
	{
		methodName := StringNameFromStr("set_particle_flag")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3515406498))
	}
	{
		methodName := StringNameFromStr("get_particle_flag")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetParticleFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2845201987))
	}
	{
		methodName := StringNameFromStr("set_emission_shape")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 491823814))
	}
	{
		methodName := StringNameFromStr("get_emission_shape")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961454842))
	}
	{
		methodName := StringNameFromStr("set_emission_sphere_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_sphere_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_box_extents")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionBoxExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_box_extents")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionBoxExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_emission_points")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
	}
	{
		methodName := StringNameFromStr("get_emission_points")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("set_emission_normals")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionNormals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
	}
	{
		methodName := StringNameFromStr("get_emission_normals")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionNormals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("set_emission_colors")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3546319833))
	}
	{
		methodName := StringNameFromStr("get_emission_colors")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_axis")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionRingAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_axis")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionRingAxis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_height")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionRingHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_height")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionRingHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionRingRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionRingRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_ring_inner_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetEmissionRingInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_ring_inner_radius")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetEmissionRingInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_gravity")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_gravity")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_split_scale")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetSplitScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_split_scale")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetSplitScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_scale_curve_x")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetScaleCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_scale_curve_x")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetScaleCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_scale_curve_y")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetScaleCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_scale_curve_y")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetScaleCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_scale_curve_z")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnGetScaleCurveZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_scale_curve_z")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnSetScaleCurveZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("convert_from_particles")
		defer methodName.Destroy()
		ptrsForCPUParticles3D.fnConvertFromParticles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
}

type CPUParticles3D struct {
	GeometryInstance3D
}

func (me *CPUParticles3D) BaseClass() string {
	return "CPUParticles3D"
}

func NewCPUParticles3D() *CPUParticles3D {
	str := StringNameFromStr("CPUParticles3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CPUParticles3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CPUParticles3DDrawOrder int

const (
	CPUParticles3DDrawOrderDrawOrderIndex     CPUParticles3DDrawOrder = 0
	CPUParticles3DDrawOrderDrawOrderLifetime  CPUParticles3DDrawOrder = 1
	CPUParticles3DDrawOrderDrawOrderViewDepth CPUParticles3DDrawOrder = 2
)

type CPUParticles3DParameter int

const (
	CPUParticles3DParameterParamInitialLinearVelocity CPUParticles3DParameter = 0
	CPUParticles3DParameterParamAngularVelocity       CPUParticles3DParameter = 1
	CPUParticles3DParameterParamOrbitVelocity         CPUParticles3DParameter = 2
	CPUParticles3DParameterParamLinearAccel           CPUParticles3DParameter = 3
	CPUParticles3DParameterParamRadialAccel           CPUParticles3DParameter = 4
	CPUParticles3DParameterParamTangentialAccel       CPUParticles3DParameter = 5
	CPUParticles3DParameterParamDamping               CPUParticles3DParameter = 6
	CPUParticles3DParameterParamAngle                 CPUParticles3DParameter = 7
	CPUParticles3DParameterParamScale                 CPUParticles3DParameter = 8
	CPUParticles3DParameterParamHueVariation          CPUParticles3DParameter = 9
	CPUParticles3DParameterParamAnimSpeed             CPUParticles3DParameter = 10
	CPUParticles3DParameterParamAnimOffset            CPUParticles3DParameter = 11
	CPUParticles3DParameterParamMax                   CPUParticles3DParameter = 12
)

type CPUParticles3DParticleFlags int

const (
	CPUParticles3DParticleFlagsParticleFlagAlignYToVelocity CPUParticles3DParticleFlags = 0
	CPUParticles3DParticleFlagsParticleFlagRotateY          CPUParticles3DParticleFlags = 1
	CPUParticles3DParticleFlagsParticleFlagDisableZ         CPUParticles3DParticleFlags = 2
	CPUParticles3DParticleFlagsParticleFlagMax              CPUParticles3DParticleFlags = 3
)

type CPUParticles3DEmissionShape int

const (
	CPUParticles3DEmissionShapeEmissionShapePoint          CPUParticles3DEmissionShape = 0
	CPUParticles3DEmissionShapeEmissionShapeSphere         CPUParticles3DEmissionShape = 1
	CPUParticles3DEmissionShapeEmissionShapeSphereSurface  CPUParticles3DEmissionShape = 2
	CPUParticles3DEmissionShapeEmissionShapeBox            CPUParticles3DEmissionShape = 3
	CPUParticles3DEmissionShapeEmissionShapePoints         CPUParticles3DEmissionShape = 4
	CPUParticles3DEmissionShapeEmissionShapeDirectedPoints CPUParticles3DEmissionShape = 5
	CPUParticles3DEmissionShapeEmissionShapeRing           CPUParticles3DEmissionShape = 6
	CPUParticles3DEmissionShapeEmissionShapeMax            CPUParticles3DEmissionShape = 7
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

func (me *CPUParticles3D) SetEmitting(emitting bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmitting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetAmount(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetLifetime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetOneShot(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetPreProcessTime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetPreProcessTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetExplosivenessRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetRandomnessRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetRandomnessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetLifetimeRandomness(random float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&random)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetUseLocalCoordinates(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetFixedFps(fps int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetFixedFps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetFractionalDelta(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetFractionalDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetSpeedScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) IsEmitting() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnIsEmitting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetAmount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetLifetime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetOneShot() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetPreProcessTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetPreProcessTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetExplosivenessRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetRandomnessRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetRandomnessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetLifetimeRandomness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetLifetimeRandomness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetUseLocalCoordinates() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetFixedFps() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetFixedFps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetFractionalDelta() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetFractionalDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetSpeedScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetDrawOrder(order CPUParticles3DDrawOrder) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetDrawOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetDrawOrder() CPUParticles3DDrawOrder {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CPUParticles3DDrawOrder

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetDrawOrder), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CPUParticles3D) SetMesh(mesh Mesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetMesh() Mesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) Restart() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnRestart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) SetDirection(direction Vector3) {
	cargs := []gdc.ConstTypePtr{direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetDirection() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetSpread(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetSpread), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetSpread() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetSpread), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetFlatness(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetFlatness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetFlatness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetFlatness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetParamMin(param CPUParticles3DParameter, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetParamMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetParamMin(param CPUParticles3DParameter) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetParamMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetParamMax(param CPUParticles3DParameter, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetParamMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetParamMax(param CPUParticles3DParameter) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetParamMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetParamCurve(param CPUParticles3DParameter, curve Curve) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetParamCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetParamCurve(param CPUParticles3DParameter) Curve {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetParamCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetColorRamp(ramp Gradient) {
	cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetColorRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetColorRamp() Gradient {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGradient()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetColorRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetColorInitialRamp(ramp Gradient) {
	cargs := []gdc.ConstTypePtr{ramp.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetColorInitialRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetColorInitialRamp() Gradient {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGradient()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetColorInitialRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetParticleFlag(particle_flag CPUParticles3DParticleFlags, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetParticleFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetParticleFlag(particle_flag CPUParticles3DParticleFlags) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particle_flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&particle_flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetParticleFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetEmissionShape(shape CPUParticles3DEmissionShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionShape() CPUParticles3DEmissionShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CPUParticles3DEmissionShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CPUParticles3D) SetEmissionSphereRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionSphereRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionSphereRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetEmissionBoxExtents(extents Vector3) {
	cargs := []gdc.ConstTypePtr{extents.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionBoxExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionBoxExtents() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionBoxExtents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetEmissionPoints(array PackedVector3Array) {
	cargs := []gdc.ConstTypePtr{array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionPoints() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetEmissionNormals(array PackedVector3Array) {
	cargs := []gdc.ConstTypePtr{array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionNormals), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionNormals() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionNormals), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetEmissionColors(array PackedColorArray) {
	cargs := []gdc.ConstTypePtr{array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionColors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionColors() PackedColorArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedColorArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetEmissionRingAxis(axis Vector3) {
	cargs := []gdc.ConstTypePtr{axis.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionRingAxis), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionRingAxis() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionRingAxis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetEmissionRingHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionRingHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionRingHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionRingHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetEmissionRingRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionRingRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionRingRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionRingRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetEmissionRingInnerRadius(inner_radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetEmissionRingInnerRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetEmissionRingInnerRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetEmissionRingInnerRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) GetGravity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetGravity(accel_vec Vector3) {
	cargs := []gdc.ConstTypePtr{accel_vec.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetSplitScale() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetSplitScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CPUParticles3D) SetSplitScale(split_scale bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&split_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetSplitScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetScaleCurveX() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetScaleCurveX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetScaleCurveX(scale_curve Curve) {
	cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetScaleCurveX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetScaleCurveY() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetScaleCurveY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetScaleCurveY(scale_curve Curve) {
	cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetScaleCurveY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) GetScaleCurveZ() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnGetScaleCurveZ), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CPUParticles3D) SetScaleCurveZ(scale_curve Curve) {
	cargs := []gdc.ConstTypePtr{scale_curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnSetScaleCurveZ), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CPUParticles3D) ConvertFromParticles(particles Node) {
	cargs := []gdc.ConstTypePtr{particles.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCPUParticles3D.fnConvertFromParticles), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CPUParticles3DFinishedSignalFn func()

func (me *CPUParticles3D) ConnectFinished(subs SignalSubscribers, fn CPUParticles3DFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CPUParticles3D) DisconnectFinished(subs SignalSubscribers, fn CPUParticles3DFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
