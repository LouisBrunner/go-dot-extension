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

type ptrsForBaseMaterial3DList struct {
	fnSetAlbedo                            gdc.MethodBindPtr
	fnGetAlbedo                            gdc.MethodBindPtr
	fnSetTransparency                      gdc.MethodBindPtr
	fnGetTransparency                      gdc.MethodBindPtr
	fnSetAlphaAntialiasing                 gdc.MethodBindPtr
	fnGetAlphaAntialiasing                 gdc.MethodBindPtr
	fnSetAlphaAntialiasingEdge             gdc.MethodBindPtr
	fnGetAlphaAntialiasingEdge             gdc.MethodBindPtr
	fnSetShadingMode                       gdc.MethodBindPtr
	fnGetShadingMode                       gdc.MethodBindPtr
	fnSetSpecular                          gdc.MethodBindPtr
	fnGetSpecular                          gdc.MethodBindPtr
	fnSetMetallic                          gdc.MethodBindPtr
	fnGetMetallic                          gdc.MethodBindPtr
	fnSetRoughness                         gdc.MethodBindPtr
	fnGetRoughness                         gdc.MethodBindPtr
	fnSetEmission                          gdc.MethodBindPtr
	fnGetEmission                          gdc.MethodBindPtr
	fnSetEmissionEnergyMultiplier          gdc.MethodBindPtr
	fnGetEmissionEnergyMultiplier          gdc.MethodBindPtr
	fnSetEmissionIntensity                 gdc.MethodBindPtr
	fnGetEmissionIntensity                 gdc.MethodBindPtr
	fnSetNormalScale                       gdc.MethodBindPtr
	fnGetNormalScale                       gdc.MethodBindPtr
	fnSetRim                               gdc.MethodBindPtr
	fnGetRim                               gdc.MethodBindPtr
	fnSetRimTint                           gdc.MethodBindPtr
	fnGetRimTint                           gdc.MethodBindPtr
	fnSetClearcoat                         gdc.MethodBindPtr
	fnGetClearcoat                         gdc.MethodBindPtr
	fnSetClearcoatRoughness                gdc.MethodBindPtr
	fnGetClearcoatRoughness                gdc.MethodBindPtr
	fnSetAnisotropy                        gdc.MethodBindPtr
	fnGetAnisotropy                        gdc.MethodBindPtr
	fnSetHeightmapScale                    gdc.MethodBindPtr
	fnGetHeightmapScale                    gdc.MethodBindPtr
	fnSetSubsurfaceScatteringStrength      gdc.MethodBindPtr
	fnGetSubsurfaceScatteringStrength      gdc.MethodBindPtr
	fnSetTransmittanceColor                gdc.MethodBindPtr
	fnGetTransmittanceColor                gdc.MethodBindPtr
	fnSetTransmittanceDepth                gdc.MethodBindPtr
	fnGetTransmittanceDepth                gdc.MethodBindPtr
	fnSetTransmittanceBoost                gdc.MethodBindPtr
	fnGetTransmittanceBoost                gdc.MethodBindPtr
	fnSetBacklight                         gdc.MethodBindPtr
	fnGetBacklight                         gdc.MethodBindPtr
	fnSetRefraction                        gdc.MethodBindPtr
	fnGetRefraction                        gdc.MethodBindPtr
	fnSetPointSize                         gdc.MethodBindPtr
	fnGetPointSize                         gdc.MethodBindPtr
	fnSetDetailUv                          gdc.MethodBindPtr
	fnGetDetailUv                          gdc.MethodBindPtr
	fnSetBlendMode                         gdc.MethodBindPtr
	fnGetBlendMode                         gdc.MethodBindPtr
	fnSetDepthDrawMode                     gdc.MethodBindPtr
	fnGetDepthDrawMode                     gdc.MethodBindPtr
	fnSetCullMode                          gdc.MethodBindPtr
	fnGetCullMode                          gdc.MethodBindPtr
	fnSetDiffuseMode                       gdc.MethodBindPtr
	fnGetDiffuseMode                       gdc.MethodBindPtr
	fnSetSpecularMode                      gdc.MethodBindPtr
	fnGetSpecularMode                      gdc.MethodBindPtr
	fnSetFlag                              gdc.MethodBindPtr
	fnGetFlag                              gdc.MethodBindPtr
	fnSetTextureFilter                     gdc.MethodBindPtr
	fnGetTextureFilter                     gdc.MethodBindPtr
	fnSetFeature                           gdc.MethodBindPtr
	fnGetFeature                           gdc.MethodBindPtr
	fnSetTexture                           gdc.MethodBindPtr
	fnGetTexture                           gdc.MethodBindPtr
	fnSetDetailBlendMode                   gdc.MethodBindPtr
	fnGetDetailBlendMode                   gdc.MethodBindPtr
	fnSetUv1Scale                          gdc.MethodBindPtr
	fnGetUv1Scale                          gdc.MethodBindPtr
	fnSetUv1Offset                         gdc.MethodBindPtr
	fnGetUv1Offset                         gdc.MethodBindPtr
	fnSetUv1TriplanarBlendSharpness        gdc.MethodBindPtr
	fnGetUv1TriplanarBlendSharpness        gdc.MethodBindPtr
	fnSetUv2Scale                          gdc.MethodBindPtr
	fnGetUv2Scale                          gdc.MethodBindPtr
	fnSetUv2Offset                         gdc.MethodBindPtr
	fnGetUv2Offset                         gdc.MethodBindPtr
	fnSetUv2TriplanarBlendSharpness        gdc.MethodBindPtr
	fnGetUv2TriplanarBlendSharpness        gdc.MethodBindPtr
	fnSetBillboardMode                     gdc.MethodBindPtr
	fnGetBillboardMode                     gdc.MethodBindPtr
	fnSetParticlesAnimHFrames              gdc.MethodBindPtr
	fnGetParticlesAnimHFrames              gdc.MethodBindPtr
	fnSetParticlesAnimVFrames              gdc.MethodBindPtr
	fnGetParticlesAnimVFrames              gdc.MethodBindPtr
	fnSetParticlesAnimLoop                 gdc.MethodBindPtr
	fnGetParticlesAnimLoop                 gdc.MethodBindPtr
	fnSetHeightmapDeepParallax             gdc.MethodBindPtr
	fnIsHeightmapDeepParallaxEnabled       gdc.MethodBindPtr
	fnSetHeightmapDeepParallaxMinLayers    gdc.MethodBindPtr
	fnGetHeightmapDeepParallaxMinLayers    gdc.MethodBindPtr
	fnSetHeightmapDeepParallaxMaxLayers    gdc.MethodBindPtr
	fnGetHeightmapDeepParallaxMaxLayers    gdc.MethodBindPtr
	fnSetHeightmapDeepParallaxFlipTangent  gdc.MethodBindPtr
	fnGetHeightmapDeepParallaxFlipTangent  gdc.MethodBindPtr
	fnSetHeightmapDeepParallaxFlipBinormal gdc.MethodBindPtr
	fnGetHeightmapDeepParallaxFlipBinormal gdc.MethodBindPtr
	fnSetGrow                              gdc.MethodBindPtr
	fnGetGrow                              gdc.MethodBindPtr
	fnSetEmissionOperator                  gdc.MethodBindPtr
	fnGetEmissionOperator                  gdc.MethodBindPtr
	fnSetAoLightAffect                     gdc.MethodBindPtr
	fnGetAoLightAffect                     gdc.MethodBindPtr
	fnSetAlphaScissorThreshold             gdc.MethodBindPtr
	fnGetAlphaScissorThreshold             gdc.MethodBindPtr
	fnSetAlphaHashScale                    gdc.MethodBindPtr
	fnGetAlphaHashScale                    gdc.MethodBindPtr
	fnSetGrowEnabled                       gdc.MethodBindPtr
	fnIsGrowEnabled                        gdc.MethodBindPtr
	fnSetMetallicTextureChannel            gdc.MethodBindPtr
	fnGetMetallicTextureChannel            gdc.MethodBindPtr
	fnSetRoughnessTextureChannel           gdc.MethodBindPtr
	fnGetRoughnessTextureChannel           gdc.MethodBindPtr
	fnSetAoTextureChannel                  gdc.MethodBindPtr
	fnGetAoTextureChannel                  gdc.MethodBindPtr
	fnSetRefractionTextureChannel          gdc.MethodBindPtr
	fnGetRefractionTextureChannel          gdc.MethodBindPtr
	fnSetProximityFadeEnabled              gdc.MethodBindPtr
	fnIsProximityFadeEnabled               gdc.MethodBindPtr
	fnSetProximityFadeDistance             gdc.MethodBindPtr
	fnGetProximityFadeDistance             gdc.MethodBindPtr
	fnSetMsdfPixelRange                    gdc.MethodBindPtr
	fnGetMsdfPixelRange                    gdc.MethodBindPtr
	fnSetMsdfOutlineSize                   gdc.MethodBindPtr
	fnGetMsdfOutlineSize                   gdc.MethodBindPtr
	fnSetDistanceFade                      gdc.MethodBindPtr
	fnGetDistanceFade                      gdc.MethodBindPtr
	fnSetDistanceFadeMaxDistance           gdc.MethodBindPtr
	fnGetDistanceFadeMaxDistance           gdc.MethodBindPtr
	fnSetDistanceFadeMinDistance           gdc.MethodBindPtr
	fnGetDistanceFadeMinDistance           gdc.MethodBindPtr
}

var ptrsForBaseMaterial3D ptrsForBaseMaterial3DList

func initBaseMaterial3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BaseMaterial3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_albedo")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_albedo")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_transparency")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTransparency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3435651667))
	}
	{
		methodName := StringNameFromStr("get_transparency")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTransparency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990903061))
	}
	{
		methodName := StringNameFromStr("set_alpha_antialiasing")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3212649852))
	}
	{
		methodName := StringNameFromStr("get_alpha_antialiasing")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2889939400))
	}
	{
		methodName := StringNameFromStr("set_alpha_antialiasing_edge")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_antialiasing_edge")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_shading_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetShadingMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3368750322))
	}
	{
		methodName := StringNameFromStr("get_shading_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetShadingMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2132070559))
	}
	{
		methodName := StringNameFromStr("set_specular")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetSpecular = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_specular")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetSpecular = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_metallic")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetMetallic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_metallic")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetMetallic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_roughness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_roughness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_emission")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_emission_energy_multiplier")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetEmissionEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_energy_multiplier")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetEmissionEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_intensity")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetEmissionIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_emission_intensity")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetEmissionIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_normal_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetNormalScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_normal_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetNormalScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rim")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRim = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rim")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRim = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rim_tint")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRimTint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rim_tint")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRimTint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_clearcoat")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetClearcoat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_clearcoat")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetClearcoat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_clearcoat_roughness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetClearcoatRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_clearcoat_roughness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetClearcoatRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_anisotropy")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_anisotropy")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_heightmap_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_heightmap_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetHeightmapScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_subsurface_scattering_strength")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetSubsurfaceScatteringStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_subsurface_scattering_strength")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetSubsurfaceScatteringStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_transmittance_color")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTransmittanceColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_transmittance_color")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTransmittanceColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_transmittance_depth")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTransmittanceDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_transmittance_depth")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTransmittanceDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_transmittance_boost")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTransmittanceBoost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_transmittance_boost")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTransmittanceBoost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_backlight")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetBacklight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_backlight")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetBacklight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_refraction")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRefraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_refraction")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRefraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_point_size")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_point_size")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_detail_uv")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDetailUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 456801921))
	}
	{
		methodName := StringNameFromStr("get_detail_uv")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDetailUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2306920512))
	}
	{
		methodName := StringNameFromStr("set_blend_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2830186259))
	}
	{
		methodName := StringNameFromStr("get_blend_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4022690962))
	}
	{
		methodName := StringNameFromStr("set_depth_draw_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDepthDrawMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1456584748))
	}
	{
		methodName := StringNameFromStr("get_depth_draw_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDepthDrawMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2578197639))
	}
	{
		methodName := StringNameFromStr("set_cull_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2338909218))
	}
	{
		methodName := StringNameFromStr("get_cull_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1941499586))
	}
	{
		methodName := StringNameFromStr("set_diffuse_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDiffuseMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1045299638))
	}
	{
		methodName := StringNameFromStr("get_diffuse_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDiffuseMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3973617136))
	}
	{
		methodName := StringNameFromStr("set_specular_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetSpecularMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 584737147))
	}
	{
		methodName := StringNameFromStr("get_specular_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetSpecularMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2569953298))
	}
	{
		methodName := StringNameFromStr("set_flag")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3070159527))
	}
	{
		methodName := StringNameFromStr("get_flag")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410065))
	}
	{
		methodName := StringNameFromStr("set_texture_filter")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 22904437))
	}
	{
		methodName := StringNameFromStr("get_texture_filter")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289213076))
	}
	{
		methodName := StringNameFromStr("set_feature")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2819288693))
	}
	{
		methodName := StringNameFromStr("get_feature")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965241794))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 464208135))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 329605813))
	}
	{
		methodName := StringNameFromStr("set_detail_blend_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDetailBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2830186259))
	}
	{
		methodName := StringNameFromStr("get_detail_blend_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDetailBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4022690962))
	}
	{
		methodName := StringNameFromStr("set_uv1_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv1Scale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_uv1_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv1Scale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_uv1_offset")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv1Offset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_uv1_offset")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv1Offset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_uv1_triplanar_blend_sharpness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv1TriplanarBlendSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_uv1_triplanar_blend_sharpness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv1TriplanarBlendSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_uv2_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv2Scale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_uv2_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv2Scale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_uv2_offset")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv2Offset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_uv2_offset")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv2Offset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_uv2_triplanar_blend_sharpness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetUv2TriplanarBlendSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_uv2_triplanar_blend_sharpness")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetUv2TriplanarBlendSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_billboard_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4202036497))
	}
	{
		methodName := StringNameFromStr("get_billboard_mode")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1283840139))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_h_frames")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetParticlesAnimHFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_h_frames")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetParticlesAnimHFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_v_frames")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetParticlesAnimVFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_v_frames")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetParticlesAnimVFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_loop")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetParticlesAnimLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_loop")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetParticlesAnimLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_heightmap_deep_parallax")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapDeepParallax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_heightmap_deep_parallax_enabled")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnIsHeightmapDeepParallaxEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_heightmap_deep_parallax_min_layers")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxMinLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_heightmap_deep_parallax_min_layers")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxMinLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_heightmap_deep_parallax_max_layers")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxMaxLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_heightmap_deep_parallax_max_layers")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxMaxLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_heightmap_deep_parallax_flip_tangent")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxFlipTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_heightmap_deep_parallax_flip_tangent")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxFlipTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_heightmap_deep_parallax_flip_binormal")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxFlipBinormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_heightmap_deep_parallax_flip_binormal")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxFlipBinormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_grow")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetGrow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_grow")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetGrow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_emission_operator")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetEmissionOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3825128922))
	}
	{
		methodName := StringNameFromStr("get_emission_operator")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetEmissionOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974205018))
	}
	{
		methodName := StringNameFromStr("set_ao_light_affect")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAoLightAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_ao_light_affect")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAoLightAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_alpha_scissor_threshold")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_scissor_threshold")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_alpha_hash_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_hash_scale")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_grow_enabled")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetGrowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_grow_enabled")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnIsGrowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_metallic_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetMetallicTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 744167988))
	}
	{
		methodName := StringNameFromStr("get_metallic_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetMetallicTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 568133867))
	}
	{
		methodName := StringNameFromStr("set_roughness_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRoughnessTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 744167988))
	}
	{
		methodName := StringNameFromStr("get_roughness_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRoughnessTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 568133867))
	}
	{
		methodName := StringNameFromStr("set_ao_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetAoTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 744167988))
	}
	{
		methodName := StringNameFromStr("get_ao_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetAoTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 568133867))
	}
	{
		methodName := StringNameFromStr("set_refraction_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetRefractionTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 744167988))
	}
	{
		methodName := StringNameFromStr("get_refraction_texture_channel")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetRefractionTextureChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 568133867))
	}
	{
		methodName := StringNameFromStr("set_proximity_fade_enabled")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetProximityFadeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_proximity_fade_enabled")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnIsProximityFadeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_proximity_fade_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetProximityFadeDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_proximity_fade_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetProximityFadeDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_msdf_outline_size")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetMsdfOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_msdf_outline_size")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetMsdfOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_distance_fade")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1379478617))
	}
	{
		methodName := StringNameFromStr("get_distance_fade")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2694575734))
	}
	{
		methodName := StringNameFromStr("set_distance_fade_max_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDistanceFadeMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_distance_fade_max_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDistanceFadeMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_distance_fade_min_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnSetDistanceFadeMinDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_distance_fade_min_distance")
		defer methodName.Destroy()
		ptrsForBaseMaterial3D.fnGetDistanceFadeMinDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type BaseMaterial3D struct {
	Material
}

func (me *BaseMaterial3D) BaseClass() string {
	return "BaseMaterial3D"
}

func NewBaseMaterial3D() *BaseMaterial3D {
	str := StringNameFromStr("BaseMaterial3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BaseMaterial3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type BaseMaterial3DTextureParam int

const (
	BaseMaterial3DTextureParamTextureAlbedo                  BaseMaterial3DTextureParam = 0
	BaseMaterial3DTextureParamTextureMetallic                BaseMaterial3DTextureParam = 1
	BaseMaterial3DTextureParamTextureRoughness               BaseMaterial3DTextureParam = 2
	BaseMaterial3DTextureParamTextureEmission                BaseMaterial3DTextureParam = 3
	BaseMaterial3DTextureParamTextureNormal                  BaseMaterial3DTextureParam = 4
	BaseMaterial3DTextureParamTextureRim                     BaseMaterial3DTextureParam = 5
	BaseMaterial3DTextureParamTextureClearcoat               BaseMaterial3DTextureParam = 6
	BaseMaterial3DTextureParamTextureFlowmap                 BaseMaterial3DTextureParam = 7
	BaseMaterial3DTextureParamTextureAmbientOcclusion        BaseMaterial3DTextureParam = 8
	BaseMaterial3DTextureParamTextureHeightmap               BaseMaterial3DTextureParam = 9
	BaseMaterial3DTextureParamTextureSubsurfaceScattering    BaseMaterial3DTextureParam = 10
	BaseMaterial3DTextureParamTextureSubsurfaceTransmittance BaseMaterial3DTextureParam = 11
	BaseMaterial3DTextureParamTextureBacklight               BaseMaterial3DTextureParam = 12
	BaseMaterial3DTextureParamTextureRefraction              BaseMaterial3DTextureParam = 13
	BaseMaterial3DTextureParamTextureDetailMask              BaseMaterial3DTextureParam = 14
	BaseMaterial3DTextureParamTextureDetailAlbedo            BaseMaterial3DTextureParam = 15
	BaseMaterial3DTextureParamTextureDetailNormal            BaseMaterial3DTextureParam = 16
	BaseMaterial3DTextureParamTextureOrm                     BaseMaterial3DTextureParam = 17
	BaseMaterial3DTextureParamTextureMax                     BaseMaterial3DTextureParam = 18
)

type BaseMaterial3DTextureFilter int

const (
	BaseMaterial3DTextureFilterTextureFilterNearest                       BaseMaterial3DTextureFilter = 0
	BaseMaterial3DTextureFilterTextureFilterLinear                        BaseMaterial3DTextureFilter = 1
	BaseMaterial3DTextureFilterTextureFilterNearestWithMipmaps            BaseMaterial3DTextureFilter = 2
	BaseMaterial3DTextureFilterTextureFilterLinearWithMipmaps             BaseMaterial3DTextureFilter = 3
	BaseMaterial3DTextureFilterTextureFilterNearestWithMipmapsAnisotropic BaseMaterial3DTextureFilter = 4
	BaseMaterial3DTextureFilterTextureFilterLinearWithMipmapsAnisotropic  BaseMaterial3DTextureFilter = 5
	BaseMaterial3DTextureFilterTextureFilterMax                           BaseMaterial3DTextureFilter = 6
)

type BaseMaterial3DDetailUV int

const (
	BaseMaterial3DDetailUVDetailUv1 BaseMaterial3DDetailUV = 0
	BaseMaterial3DDetailUVDetailUv2 BaseMaterial3DDetailUV = 1
)

type BaseMaterial3DTransparency int

const (
	BaseMaterial3DTransparencyTransparencyDisabled          BaseMaterial3DTransparency = 0
	BaseMaterial3DTransparencyTransparencyAlpha             BaseMaterial3DTransparency = 1
	BaseMaterial3DTransparencyTransparencyAlphaScissor      BaseMaterial3DTransparency = 2
	BaseMaterial3DTransparencyTransparencyAlphaHash         BaseMaterial3DTransparency = 3
	BaseMaterial3DTransparencyTransparencyAlphaDepthPrePass BaseMaterial3DTransparency = 4
	BaseMaterial3DTransparencyTransparencyMax               BaseMaterial3DTransparency = 5
)

type BaseMaterial3DShadingMode int

const (
	BaseMaterial3DShadingModeShadingModeUnshaded  BaseMaterial3DShadingMode = 0
	BaseMaterial3DShadingModeShadingModePerPixel  BaseMaterial3DShadingMode = 1
	BaseMaterial3DShadingModeShadingModePerVertex BaseMaterial3DShadingMode = 2
	BaseMaterial3DShadingModeShadingModeMax       BaseMaterial3DShadingMode = 3
)

type BaseMaterial3DFeature int

const (
	BaseMaterial3DFeatureFeatureEmission                BaseMaterial3DFeature = 0
	BaseMaterial3DFeatureFeatureNormalMapping           BaseMaterial3DFeature = 1
	BaseMaterial3DFeatureFeatureRim                     BaseMaterial3DFeature = 2
	BaseMaterial3DFeatureFeatureClearcoat               BaseMaterial3DFeature = 3
	BaseMaterial3DFeatureFeatureAnisotropy              BaseMaterial3DFeature = 4
	BaseMaterial3DFeatureFeatureAmbientOcclusion        BaseMaterial3DFeature = 5
	BaseMaterial3DFeatureFeatureHeightMapping           BaseMaterial3DFeature = 6
	BaseMaterial3DFeatureFeatureSubsurfaceScattering    BaseMaterial3DFeature = 7
	BaseMaterial3DFeatureFeatureSubsurfaceTransmittance BaseMaterial3DFeature = 8
	BaseMaterial3DFeatureFeatureBacklight               BaseMaterial3DFeature = 9
	BaseMaterial3DFeatureFeatureRefraction              BaseMaterial3DFeature = 10
	BaseMaterial3DFeatureFeatureDetail                  BaseMaterial3DFeature = 11
	BaseMaterial3DFeatureFeatureMax                     BaseMaterial3DFeature = 12
)

type BaseMaterial3DBlendMode int

const (
	BaseMaterial3DBlendModeBlendModeMix BaseMaterial3DBlendMode = 0
	BaseMaterial3DBlendModeBlendModeAdd BaseMaterial3DBlendMode = 1
	BaseMaterial3DBlendModeBlendModeSub BaseMaterial3DBlendMode = 2
	BaseMaterial3DBlendModeBlendModeMul BaseMaterial3DBlendMode = 3
)

type BaseMaterial3DAlphaAntiAliasing int

const (
	BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingOff                     BaseMaterial3DAlphaAntiAliasing = 0
	BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingAlphaToCoverage         BaseMaterial3DAlphaAntiAliasing = 1
	BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingAlphaToCoverageAndToOne BaseMaterial3DAlphaAntiAliasing = 2
)

type BaseMaterial3DDepthDrawMode int

const (
	BaseMaterial3DDepthDrawModeDepthDrawOpaqueOnly BaseMaterial3DDepthDrawMode = 0
	BaseMaterial3DDepthDrawModeDepthDrawAlways     BaseMaterial3DDepthDrawMode = 1
	BaseMaterial3DDepthDrawModeDepthDrawDisabled   BaseMaterial3DDepthDrawMode = 2
)

type BaseMaterial3DCullMode int

const (
	BaseMaterial3DCullModeCullBack     BaseMaterial3DCullMode = 0
	BaseMaterial3DCullModeCullFront    BaseMaterial3DCullMode = 1
	BaseMaterial3DCullModeCullDisabled BaseMaterial3DCullMode = 2
)

type BaseMaterial3DFlags int

const (
	BaseMaterial3DFlagsFlagDisableDepthTest       BaseMaterial3DFlags = 0
	BaseMaterial3DFlagsFlagAlbedoFromVertexColor  BaseMaterial3DFlags = 1
	BaseMaterial3DFlagsFlagSrgbVertexColor        BaseMaterial3DFlags = 2
	BaseMaterial3DFlagsFlagUsePointSize           BaseMaterial3DFlags = 3
	BaseMaterial3DFlagsFlagFixedSize              BaseMaterial3DFlags = 4
	BaseMaterial3DFlagsFlagBillboardKeepScale     BaseMaterial3DFlags = 5
	BaseMaterial3DFlagsFlagUv1UseTriplanar        BaseMaterial3DFlags = 6
	BaseMaterial3DFlagsFlagUv2UseTriplanar        BaseMaterial3DFlags = 7
	BaseMaterial3DFlagsFlagUv1UseWorldTriplanar   BaseMaterial3DFlags = 8
	BaseMaterial3DFlagsFlagUv2UseWorldTriplanar   BaseMaterial3DFlags = 9
	BaseMaterial3DFlagsFlagAoOnUv2                BaseMaterial3DFlags = 10
	BaseMaterial3DFlagsFlagEmissionOnUv2          BaseMaterial3DFlags = 11
	BaseMaterial3DFlagsFlagAlbedoTextureForceSrgb BaseMaterial3DFlags = 12
	BaseMaterial3DFlagsFlagDontReceiveShadows     BaseMaterial3DFlags = 13
	BaseMaterial3DFlagsFlagDisableAmbientLight    BaseMaterial3DFlags = 14
	BaseMaterial3DFlagsFlagUseShadowToOpacity     BaseMaterial3DFlags = 15
	BaseMaterial3DFlagsFlagUseTextureRepeat       BaseMaterial3DFlags = 16
	BaseMaterial3DFlagsFlagInvertHeightmap        BaseMaterial3DFlags = 17
	BaseMaterial3DFlagsFlagSubsurfaceModeSkin     BaseMaterial3DFlags = 18
	BaseMaterial3DFlagsFlagParticleTrailsMode     BaseMaterial3DFlags = 19
	BaseMaterial3DFlagsFlagAlbedoTextureMsdf      BaseMaterial3DFlags = 20
	BaseMaterial3DFlagsFlagDisableFog             BaseMaterial3DFlags = 21
	BaseMaterial3DFlagsFlagMax                    BaseMaterial3DFlags = 22
)

type BaseMaterial3DDiffuseMode int

const (
	BaseMaterial3DDiffuseModeDiffuseBurley      BaseMaterial3DDiffuseMode = 0
	BaseMaterial3DDiffuseModeDiffuseLambert     BaseMaterial3DDiffuseMode = 1
	BaseMaterial3DDiffuseModeDiffuseLambertWrap BaseMaterial3DDiffuseMode = 2
	BaseMaterial3DDiffuseModeDiffuseToon        BaseMaterial3DDiffuseMode = 3
)

type BaseMaterial3DSpecularMode int

const (
	BaseMaterial3DSpecularModeSpecularSchlickGgx BaseMaterial3DSpecularMode = 0
	BaseMaterial3DSpecularModeSpecularToon       BaseMaterial3DSpecularMode = 1
	BaseMaterial3DSpecularModeSpecularDisabled   BaseMaterial3DSpecularMode = 2
)

type BaseMaterial3DBillboardMode int

const (
	BaseMaterial3DBillboardModeBillboardDisabled  BaseMaterial3DBillboardMode = 0
	BaseMaterial3DBillboardModeBillboardEnabled   BaseMaterial3DBillboardMode = 1
	BaseMaterial3DBillboardModeBillboardFixedY    BaseMaterial3DBillboardMode = 2
	BaseMaterial3DBillboardModeBillboardParticles BaseMaterial3DBillboardMode = 3
)

type BaseMaterial3DTextureChannel int

const (
	BaseMaterial3DTextureChannelTextureChannelRed       BaseMaterial3DTextureChannel = 0
	BaseMaterial3DTextureChannelTextureChannelGreen     BaseMaterial3DTextureChannel = 1
	BaseMaterial3DTextureChannelTextureChannelBlue      BaseMaterial3DTextureChannel = 2
	BaseMaterial3DTextureChannelTextureChannelAlpha     BaseMaterial3DTextureChannel = 3
	BaseMaterial3DTextureChannelTextureChannelGrayscale BaseMaterial3DTextureChannel = 4
)

type BaseMaterial3DEmissionOperator int

const (
	BaseMaterial3DEmissionOperatorEmissionOpAdd      BaseMaterial3DEmissionOperator = 0
	BaseMaterial3DEmissionOperatorEmissionOpMultiply BaseMaterial3DEmissionOperator = 1
)

type BaseMaterial3DDistanceFadeMode int

const (
	BaseMaterial3DDistanceFadeModeDistanceFadeDisabled     BaseMaterial3DDistanceFadeMode = 0
	BaseMaterial3DDistanceFadeModeDistanceFadePixelAlpha   BaseMaterial3DDistanceFadeMode = 1
	BaseMaterial3DDistanceFadeModeDistanceFadePixelDither  BaseMaterial3DDistanceFadeMode = 2
	BaseMaterial3DDistanceFadeModeDistanceFadeObjectDither BaseMaterial3DDistanceFadeMode = 3
)

func (me *BaseMaterial3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BaseMaterial3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BaseMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BaseMaterial3D) SetAlbedo(albedo Color) {
	cargs := []gdc.ConstTypePtr{albedo.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAlbedo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAlbedo() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAlbedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetTransparency(transparency BaseMaterial3DTransparency) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transparency)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTransparency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTransparency() BaseMaterial3DTransparency {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTransparency

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTransparency), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DAlphaAntiAliasing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetAlphaAntialiasingEdge(edge float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAlphaAntialiasingEdge() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetShadingMode(shading_mode BaseMaterial3DShadingMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shading_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetShadingMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetShadingMode() BaseMaterial3DShadingMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DShadingMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetShadingMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetSpecular(specular float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&specular)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetSpecular), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetSpecular() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetSpecular), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetMetallic(metallic float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&metallic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetMetallic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetMetallic() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetMetallic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetRoughness(roughness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&roughness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRoughness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRoughness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRoughness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetEmission(emission Color) {
	cargs := []gdc.ConstTypePtr{emission.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetEmission), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetEmission() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetEmission), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetEmissionEnergyMultiplier(emission_energy_multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emission_energy_multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetEmissionEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetEmissionEnergyMultiplier() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetEmissionEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetEmissionIntensity(emission_energy_multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emission_energy_multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetEmissionIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetEmissionIntensity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetEmissionIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetNormalScale(normal_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetNormalScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetNormalScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetNormalScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetRim(rim float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rim)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRim), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRim() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRim), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetRimTint(rim_tint float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rim_tint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRimTint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRimTint() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRimTint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetClearcoat(clearcoat float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clearcoat)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetClearcoat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetClearcoat() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetClearcoat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetClearcoatRoughness(clearcoat_roughness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clearcoat_roughness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetClearcoatRoughness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetClearcoatRoughness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetClearcoatRoughness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetAnisotropy(anisotropy float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anisotropy)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAnisotropy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAnisotropy() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAnisotropy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapScale(heightmap_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heightmap_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetHeightmapScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetHeightmapScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetSubsurfaceScatteringStrength(strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetSubsurfaceScatteringStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetSubsurfaceScatteringStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetSubsurfaceScatteringStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetTransmittanceColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTransmittanceColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTransmittanceColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTransmittanceColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetTransmittanceDepth(depth float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTransmittanceDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTransmittanceDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTransmittanceDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetTransmittanceBoost(boost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&boost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTransmittanceBoost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTransmittanceBoost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTransmittanceBoost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetBacklight(backlight Color) {
	cargs := []gdc.ConstTypePtr{backlight.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetBacklight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetBacklight() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetBacklight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetRefraction(refraction float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refraction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRefraction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRefraction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRefraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetPointSize(point_size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetPointSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetPointSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetPointSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetDetailUv(detail_uv BaseMaterial3DDetailUV) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_uv)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDetailUv), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDetailUv() BaseMaterial3DDetailUV {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DDetailUV

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDetailUv), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetBlendMode(blend_mode BaseMaterial3DBlendMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetBlendMode() BaseMaterial3DBlendMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DBlendMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetDepthDrawMode(depth_draw_mode BaseMaterial3DDepthDrawMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth_draw_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDepthDrawMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDepthDrawMode() BaseMaterial3DDepthDrawMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DDepthDrawMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDepthDrawMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetCullMode(cull_mode BaseMaterial3DCullMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetCullMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetCullMode() BaseMaterial3DCullMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DCullMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetCullMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetDiffuseMode(diffuse_mode BaseMaterial3DDiffuseMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&diffuse_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDiffuseMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDiffuseMode() BaseMaterial3DDiffuseMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DDiffuseMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDiffuseMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetSpecularMode(specular_mode BaseMaterial3DSpecularMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&specular_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetSpecularMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetSpecularMode() BaseMaterial3DSpecularMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DSpecularMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetSpecularMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetFlag(flag BaseMaterial3DFlags, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetFlag(flag BaseMaterial3DFlags) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetTextureFilter(mode BaseMaterial3DTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTextureFilter() BaseMaterial3DTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetFeature(feature BaseMaterial3DFeature, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetFeature), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetFeature(feature BaseMaterial3DFeature) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetTexture(param BaseMaterial3DTextureParam, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetTexture(param BaseMaterial3DTextureParam) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetDetailBlendMode(detail_blend_mode BaseMaterial3DBlendMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_blend_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDetailBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDetailBlendMode() BaseMaterial3DBlendMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DBlendMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDetailBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetUv1Scale(scale Vector3) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv1Scale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv1Scale() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv1Scale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetUv1Offset(offset Vector3) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv1Offset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv1Offset() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv1Offset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetUv1TriplanarBlendSharpness(sharpness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv1TriplanarBlendSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv1TriplanarBlendSharpness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv1TriplanarBlendSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetUv2Scale(scale Vector3) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv2Scale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv2Scale() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv2Scale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetUv2Offset(offset Vector3) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv2Offset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv2Offset() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv2Offset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BaseMaterial3D) SetUv2TriplanarBlendSharpness(sharpness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetUv2TriplanarBlendSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetUv2TriplanarBlendSharpness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetUv2TriplanarBlendSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetBillboardMode(mode BaseMaterial3DBillboardMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetBillboardMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetBillboardMode() BaseMaterial3DBillboardMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DBillboardMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetBillboardMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetParticlesAnimHFrames(frames int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetParticlesAnimHFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetParticlesAnimHFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetParticlesAnimHFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetParticlesAnimVFrames(frames int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetParticlesAnimVFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetParticlesAnimVFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetParticlesAnimVFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetParticlesAnimLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetParticlesAnimLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetParticlesAnimLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetParticlesAnimLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapDeepParallax(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapDeepParallax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) IsHeightmapDeepParallaxEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnIsHeightmapDeepParallaxEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapDeepParallaxMinLayers(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxMinLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetHeightmapDeepParallaxMinLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxMinLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapDeepParallaxMaxLayers(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxMaxLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetHeightmapDeepParallaxMaxLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxMaxLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipTangent(flip bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxFlipTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipTangent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxFlipTangent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipBinormal(flip bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetHeightmapDeepParallaxFlipBinormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipBinormal() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetHeightmapDeepParallaxFlipBinormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetGrow(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetGrow), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetGrow() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetGrow), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetEmissionOperator(operator BaseMaterial3DEmissionOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operator)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetEmissionOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetEmissionOperator() BaseMaterial3DEmissionOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DEmissionOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetEmissionOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetAoLightAffect(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAoLightAffect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAoLightAffect() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAoLightAffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetAlphaScissorThreshold(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAlphaScissorThreshold() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetAlphaHashScale(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAlphaHashScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAlphaHashScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAlphaHashScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetGrowEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetGrowEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) IsGrowEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnIsGrowEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetMetallicTextureChannel(channel BaseMaterial3DTextureChannel) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetMetallicTextureChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetMetallicTextureChannel() BaseMaterial3DTextureChannel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureChannel

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetMetallicTextureChannel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetRoughnessTextureChannel(channel BaseMaterial3DTextureChannel) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRoughnessTextureChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRoughnessTextureChannel() BaseMaterial3DTextureChannel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureChannel

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRoughnessTextureChannel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetAoTextureChannel(channel BaseMaterial3DTextureChannel) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetAoTextureChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetAoTextureChannel() BaseMaterial3DTextureChannel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureChannel

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetAoTextureChannel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetRefractionTextureChannel(channel BaseMaterial3DTextureChannel) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetRefractionTextureChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetRefractionTextureChannel() BaseMaterial3DTextureChannel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureChannel

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetRefractionTextureChannel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetProximityFadeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetProximityFadeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) IsProximityFadeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnIsProximityFadeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetProximityFadeDistance(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetProximityFadeDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetProximityFadeDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetProximityFadeDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetMsdfPixelRange(range_ float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&range_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetMsdfPixelRange() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetMsdfOutlineSize(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetMsdfOutlineSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetMsdfOutlineSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetMsdfOutlineSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetDistanceFade(mode BaseMaterial3DDistanceFadeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDistanceFade), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDistanceFade() BaseMaterial3DDistanceFadeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DDistanceFadeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDistanceFade), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BaseMaterial3D) SetDistanceFadeMaxDistance(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDistanceFadeMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDistanceFadeMaxDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDistanceFadeMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BaseMaterial3D) SetDistanceFadeMinDistance(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnSetDistanceFadeMinDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BaseMaterial3D) GetDistanceFadeMinDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBaseMaterial3D.fnGetDistanceFadeMinDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
