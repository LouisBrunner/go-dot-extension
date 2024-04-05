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

type ptrsForEnvironmentList struct {
  fnSetBackground gdc.MethodBindPtr
  fnGetBackground gdc.MethodBindPtr
  fnSetSky gdc.MethodBindPtr
  fnGetSky gdc.MethodBindPtr
  fnSetSkyCustomFov gdc.MethodBindPtr
  fnGetSkyCustomFov gdc.MethodBindPtr
  fnSetSkyRotation gdc.MethodBindPtr
  fnGetSkyRotation gdc.MethodBindPtr
  fnSetBgColor gdc.MethodBindPtr
  fnGetBgColor gdc.MethodBindPtr
  fnSetBgEnergyMultiplier gdc.MethodBindPtr
  fnGetBgEnergyMultiplier gdc.MethodBindPtr
  fnSetBgIntensity gdc.MethodBindPtr
  fnGetBgIntensity gdc.MethodBindPtr
  fnSetCanvasMaxLayer gdc.MethodBindPtr
  fnGetCanvasMaxLayer gdc.MethodBindPtr
  fnSetCameraFeedId gdc.MethodBindPtr
  fnGetCameraFeedId gdc.MethodBindPtr
  fnSetAmbientLightColor gdc.MethodBindPtr
  fnGetAmbientLightColor gdc.MethodBindPtr
  fnSetAmbientSource gdc.MethodBindPtr
  fnGetAmbientSource gdc.MethodBindPtr
  fnSetAmbientLightEnergy gdc.MethodBindPtr
  fnGetAmbientLightEnergy gdc.MethodBindPtr
  fnSetAmbientLightSkyContribution gdc.MethodBindPtr
  fnGetAmbientLightSkyContribution gdc.MethodBindPtr
  fnSetReflectionSource gdc.MethodBindPtr
  fnGetReflectionSource gdc.MethodBindPtr
  fnSetTonemapper gdc.MethodBindPtr
  fnGetTonemapper gdc.MethodBindPtr
  fnSetTonemapExposure gdc.MethodBindPtr
  fnGetTonemapExposure gdc.MethodBindPtr
  fnSetTonemapWhite gdc.MethodBindPtr
  fnGetTonemapWhite gdc.MethodBindPtr
  fnSetSsrEnabled gdc.MethodBindPtr
  fnIsSsrEnabled gdc.MethodBindPtr
  fnSetSsrMaxSteps gdc.MethodBindPtr
  fnGetSsrMaxSteps gdc.MethodBindPtr
  fnSetSsrFadeIn gdc.MethodBindPtr
  fnGetSsrFadeIn gdc.MethodBindPtr
  fnSetSsrFadeOut gdc.MethodBindPtr
  fnGetSsrFadeOut gdc.MethodBindPtr
  fnSetSsrDepthTolerance gdc.MethodBindPtr
  fnGetSsrDepthTolerance gdc.MethodBindPtr
  fnSetSsaoEnabled gdc.MethodBindPtr
  fnIsSsaoEnabled gdc.MethodBindPtr
  fnSetSsaoRadius gdc.MethodBindPtr
  fnGetSsaoRadius gdc.MethodBindPtr
  fnSetSsaoIntensity gdc.MethodBindPtr
  fnGetSsaoIntensity gdc.MethodBindPtr
  fnSetSsaoPower gdc.MethodBindPtr
  fnGetSsaoPower gdc.MethodBindPtr
  fnSetSsaoDetail gdc.MethodBindPtr
  fnGetSsaoDetail gdc.MethodBindPtr
  fnSetSsaoHorizon gdc.MethodBindPtr
  fnGetSsaoHorizon gdc.MethodBindPtr
  fnSetSsaoSharpness gdc.MethodBindPtr
  fnGetSsaoSharpness gdc.MethodBindPtr
  fnSetSsaoDirectLightAffect gdc.MethodBindPtr
  fnGetSsaoDirectLightAffect gdc.MethodBindPtr
  fnSetSsaoAoChannelAffect gdc.MethodBindPtr
  fnGetSsaoAoChannelAffect gdc.MethodBindPtr
  fnSetSsilEnabled gdc.MethodBindPtr
  fnIsSsilEnabled gdc.MethodBindPtr
  fnSetSsilRadius gdc.MethodBindPtr
  fnGetSsilRadius gdc.MethodBindPtr
  fnSetSsilIntensity gdc.MethodBindPtr
  fnGetSsilIntensity gdc.MethodBindPtr
  fnSetSsilSharpness gdc.MethodBindPtr
  fnGetSsilSharpness gdc.MethodBindPtr
  fnSetSsilNormalRejection gdc.MethodBindPtr
  fnGetSsilNormalRejection gdc.MethodBindPtr
  fnSetSdfgiEnabled gdc.MethodBindPtr
  fnIsSdfgiEnabled gdc.MethodBindPtr
  fnSetSdfgiCascades gdc.MethodBindPtr
  fnGetSdfgiCascades gdc.MethodBindPtr
  fnSetSdfgiMinCellSize gdc.MethodBindPtr
  fnGetSdfgiMinCellSize gdc.MethodBindPtr
  fnSetSdfgiMaxDistance gdc.MethodBindPtr
  fnGetSdfgiMaxDistance gdc.MethodBindPtr
  fnSetSdfgiCascade0Distance gdc.MethodBindPtr
  fnGetSdfgiCascade0Distance gdc.MethodBindPtr
  fnSetSdfgiYScale gdc.MethodBindPtr
  fnGetSdfgiYScale gdc.MethodBindPtr
  fnSetSdfgiUseOcclusion gdc.MethodBindPtr
  fnIsSdfgiUsingOcclusion gdc.MethodBindPtr
  fnSetSdfgiBounceFeedback gdc.MethodBindPtr
  fnGetSdfgiBounceFeedback gdc.MethodBindPtr
  fnSetSdfgiReadSkyLight gdc.MethodBindPtr
  fnIsSdfgiReadingSkyLight gdc.MethodBindPtr
  fnSetSdfgiEnergy gdc.MethodBindPtr
  fnGetSdfgiEnergy gdc.MethodBindPtr
  fnSetSdfgiNormalBias gdc.MethodBindPtr
  fnGetSdfgiNormalBias gdc.MethodBindPtr
  fnSetSdfgiProbeBias gdc.MethodBindPtr
  fnGetSdfgiProbeBias gdc.MethodBindPtr
  fnSetGlowEnabled gdc.MethodBindPtr
  fnIsGlowEnabled gdc.MethodBindPtr
  fnSetGlowLevel gdc.MethodBindPtr
  fnGetGlowLevel gdc.MethodBindPtr
  fnSetGlowNormalized gdc.MethodBindPtr
  fnIsGlowNormalized gdc.MethodBindPtr
  fnSetGlowIntensity gdc.MethodBindPtr
  fnGetGlowIntensity gdc.MethodBindPtr
  fnSetGlowStrength gdc.MethodBindPtr
  fnGetGlowStrength gdc.MethodBindPtr
  fnSetGlowMix gdc.MethodBindPtr
  fnGetGlowMix gdc.MethodBindPtr
  fnSetGlowBloom gdc.MethodBindPtr
  fnGetGlowBloom gdc.MethodBindPtr
  fnSetGlowBlendMode gdc.MethodBindPtr
  fnGetGlowBlendMode gdc.MethodBindPtr
  fnSetGlowHdrBleedThreshold gdc.MethodBindPtr
  fnGetGlowHdrBleedThreshold gdc.MethodBindPtr
  fnSetGlowHdrBleedScale gdc.MethodBindPtr
  fnGetGlowHdrBleedScale gdc.MethodBindPtr
  fnSetGlowHdrLuminanceCap gdc.MethodBindPtr
  fnGetGlowHdrLuminanceCap gdc.MethodBindPtr
  fnSetGlowMapStrength gdc.MethodBindPtr
  fnGetGlowMapStrength gdc.MethodBindPtr
  fnSetGlowMap gdc.MethodBindPtr
  fnGetGlowMap gdc.MethodBindPtr
  fnSetFogEnabled gdc.MethodBindPtr
  fnIsFogEnabled gdc.MethodBindPtr
  fnSetFogLightColor gdc.MethodBindPtr
  fnGetFogLightColor gdc.MethodBindPtr
  fnSetFogLightEnergy gdc.MethodBindPtr
  fnGetFogLightEnergy gdc.MethodBindPtr
  fnSetFogSunScatter gdc.MethodBindPtr
  fnGetFogSunScatter gdc.MethodBindPtr
  fnSetFogDensity gdc.MethodBindPtr
  fnGetFogDensity gdc.MethodBindPtr
  fnSetFogHeight gdc.MethodBindPtr
  fnGetFogHeight gdc.MethodBindPtr
  fnSetFogHeightDensity gdc.MethodBindPtr
  fnGetFogHeightDensity gdc.MethodBindPtr
  fnSetFogAerialPerspective gdc.MethodBindPtr
  fnGetFogAerialPerspective gdc.MethodBindPtr
  fnSetFogSkyAffect gdc.MethodBindPtr
  fnGetFogSkyAffect gdc.MethodBindPtr
  fnSetVolumetricFogEnabled gdc.MethodBindPtr
  fnIsVolumetricFogEnabled gdc.MethodBindPtr
  fnSetVolumetricFogEmission gdc.MethodBindPtr
  fnGetVolumetricFogEmission gdc.MethodBindPtr
  fnSetVolumetricFogAlbedo gdc.MethodBindPtr
  fnGetVolumetricFogAlbedo gdc.MethodBindPtr
  fnSetVolumetricFogDensity gdc.MethodBindPtr
  fnGetVolumetricFogDensity gdc.MethodBindPtr
  fnSetVolumetricFogEmissionEnergy gdc.MethodBindPtr
  fnGetVolumetricFogEmissionEnergy gdc.MethodBindPtr
  fnSetVolumetricFogAnisotropy gdc.MethodBindPtr
  fnGetVolumetricFogAnisotropy gdc.MethodBindPtr
  fnSetVolumetricFogLength gdc.MethodBindPtr
  fnGetVolumetricFogLength gdc.MethodBindPtr
  fnSetVolumetricFogDetailSpread gdc.MethodBindPtr
  fnGetVolumetricFogDetailSpread gdc.MethodBindPtr
  fnSetVolumetricFogGiInject gdc.MethodBindPtr
  fnGetVolumetricFogGiInject gdc.MethodBindPtr
  fnSetVolumetricFogAmbientInject gdc.MethodBindPtr
  fnGetVolumetricFogAmbientInject gdc.MethodBindPtr
  fnSetVolumetricFogSkyAffect gdc.MethodBindPtr
  fnGetVolumetricFogSkyAffect gdc.MethodBindPtr
  fnSetVolumetricFogTemporalReprojectionEnabled gdc.MethodBindPtr
  fnIsVolumetricFogTemporalReprojectionEnabled gdc.MethodBindPtr
  fnSetVolumetricFogTemporalReprojectionAmount gdc.MethodBindPtr
  fnGetVolumetricFogTemporalReprojectionAmount gdc.MethodBindPtr
  fnSetAdjustmentEnabled gdc.MethodBindPtr
  fnIsAdjustmentEnabled gdc.MethodBindPtr
  fnSetAdjustmentBrightness gdc.MethodBindPtr
  fnGetAdjustmentBrightness gdc.MethodBindPtr
  fnSetAdjustmentContrast gdc.MethodBindPtr
  fnGetAdjustmentContrast gdc.MethodBindPtr
  fnSetAdjustmentSaturation gdc.MethodBindPtr
  fnGetAdjustmentSaturation gdc.MethodBindPtr
  fnSetAdjustmentColorCorrection gdc.MethodBindPtr
  fnGetAdjustmentColorCorrection gdc.MethodBindPtr
}

var ptrsForEnvironment ptrsForEnvironmentList

func initEnvironmentPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Environment")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_background")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4071623990))
  }
  {
    methodName := StringNameFromStr("get_background")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1843210413))
  }
  {
    methodName := StringNameFromStr("set_sky")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3336722921))
  }
  {
    methodName := StringNameFromStr("get_sky")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1177136966))
  }
  {
    methodName := StringNameFromStr("set_sky_custom_fov")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSkyCustomFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sky_custom_fov")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSkyCustomFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sky_rotation")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSkyRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_sky_rotation")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSkyRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_bg_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_bg_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_bg_energy_multiplier")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetBgEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bg_energy_multiplier")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetBgEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_bg_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetBgIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bg_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetBgIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_canvas_max_layer")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetCanvasMaxLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_canvas_max_layer")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetCanvasMaxLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_camera_feed_id")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetCameraFeedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_camera_feed_id")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetCameraFeedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_ambient_light_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAmbientLightColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_ambient_light_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAmbientLightColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_ambient_source")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAmbientSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2607780160))
  }
  {
    methodName := StringNameFromStr("get_ambient_source")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAmbientSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 67453933))
  }
  {
    methodName := StringNameFromStr("set_ambient_light_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAmbientLightEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ambient_light_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAmbientLightEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ambient_light_sky_contribution")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAmbientLightSkyContribution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ambient_light_sky_contribution")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAmbientLightSkyContribution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_reflection_source")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetReflectionSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 299673197))
  }
  {
    methodName := StringNameFromStr("get_reflection_source")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetReflectionSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 777700713))
  }
  {
    methodName := StringNameFromStr("set_tonemapper")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetTonemapper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509116664))
  }
  {
    methodName := StringNameFromStr("get_tonemapper")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetTonemapper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2908408137))
  }
  {
    methodName := StringNameFromStr("set_tonemap_exposure")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetTonemapExposure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_tonemap_exposure")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetTonemapExposure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_tonemap_white")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetTonemapWhite = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_tonemap_white")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetTonemapWhite = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssr_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsrEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_ssr_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSsrEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_ssr_max_steps")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsrMaxSteps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_ssr_max_steps")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsrMaxSteps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_ssr_fade_in")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsrFadeIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssr_fade_in")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsrFadeIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssr_fade_out")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsrFadeOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssr_fade_out")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsrFadeOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssr_depth_tolerance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsrDepthTolerance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssr_depth_tolerance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsrDepthTolerance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_ssao_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSsaoEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_ssao_radius")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_radius")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_power")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoPower = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_power")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoPower = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_detail")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoDetail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_detail")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoDetail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_horizon")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoHorizon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_horizon")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoHorizon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_sharpness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_sharpness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_direct_light_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoDirectLightAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_direct_light_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoDirectLightAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssao_ao_channel_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsaoAoChannelAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssao_ao_channel_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsaoAoChannelAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssil_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsilEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_ssil_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSsilEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_ssil_radius")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsilRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssil_radius")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsilRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssil_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsilIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssil_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsilIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssil_sharpness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsilSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssil_sharpness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsilSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ssil_normal_rejection")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSsilNormalRejection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ssil_normal_rejection")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSsilNormalRejection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_sdfgi_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSdfgiEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_cascades")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiCascades = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_cascades")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiCascades = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_min_cell_size")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiMinCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_min_cell_size")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiMinCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_max_distance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_max_distance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_cascade0_distance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiCascade0Distance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_cascade0_distance")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiCascade0Distance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_y_scale")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiYScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3608608372))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_y_scale")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiYScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2568002245))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_use_occlusion")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiUseOcclusion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_sdfgi_using_occlusion")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSdfgiUsingOcclusion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_bounce_feedback")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiBounceFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_bounce_feedback")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiBounceFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_read_sky_light")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiReadSkyLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_sdfgi_reading_sky_light")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsSdfgiReadingSkyLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_normal_bias")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiNormalBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_normal_bias")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiNormalBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sdfgi_probe_bias")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetSdfgiProbeBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sdfgi_probe_bias")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetSdfgiProbeBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_glow_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsGlowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_glow_level")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_glow_level")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_glow_normalized")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_glow_normalized")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsGlowNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_glow_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_intensity")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_strength")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_strength")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_mix")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_mix")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_bloom")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowBloom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_bloom")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowBloom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_blend_mode")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2561587761))
  }
  {
    methodName := StringNameFromStr("get_glow_blend_mode")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1529667332))
  }
  {
    methodName := StringNameFromStr("set_glow_hdr_bleed_threshold")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowHdrBleedThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_hdr_bleed_threshold")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowHdrBleedThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_hdr_bleed_scale")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowHdrBleedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_hdr_bleed_scale")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowHdrBleedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_hdr_luminance_cap")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowHdrLuminanceCap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_hdr_luminance_cap")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowHdrLuminanceCap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_map_strength")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowMapStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_glow_map_strength")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowMapStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_glow_map")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetGlowMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1790811099))
  }
  {
    methodName := StringNameFromStr("get_glow_map")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetGlowMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4037048985))
  }
  {
    methodName := StringNameFromStr("set_fog_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_fog_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsFogEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_fog_light_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogLightColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_fog_light_color")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogLightColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_fog_light_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogLightEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_light_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogLightEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_sun_scatter")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogSunScatter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_sun_scatter")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogSunScatter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_height")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_height")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_height_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogHeightDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_height_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogHeightDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_aerial_perspective")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogAerialPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_aerial_perspective")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogAerialPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_fog_sky_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetFogSkyAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fog_sky_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetFogSkyAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_volumetric_fog_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsVolumetricFogEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_emission")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_emission")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogEmission = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_albedo")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_albedo")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogAlbedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_density")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogDensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_emission_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogEmissionEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_emission_energy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogEmissionEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_anisotropy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_anisotropy")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogAnisotropy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_length")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_length")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_detail_spread")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogDetailSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_detail_spread")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogDetailSpread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_gi_inject")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogGiInject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_gi_inject")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogGiInject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_ambient_inject")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogAmbientInject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_ambient_inject")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogAmbientInject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_sky_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogSkyAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_sky_affect")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogSkyAffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_temporal_reprojection_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogTemporalReprojectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_volumetric_fog_temporal_reprojection_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsVolumetricFogTemporalReprojectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_volumetric_fog_temporal_reprojection_amount")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetVolumetricFogTemporalReprojectionAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volumetric_fog_temporal_reprojection_amount")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetVolumetricFogTemporalReprojectionAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_adjustment_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAdjustmentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_adjustment_enabled")
    defer methodName.Destroy()
    ptrsForEnvironment.fnIsAdjustmentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_adjustment_brightness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAdjustmentBrightness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_adjustment_brightness")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAdjustmentBrightness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_adjustment_contrast")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAdjustmentContrast = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_adjustment_contrast")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAdjustmentContrast = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_adjustment_saturation")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAdjustmentSaturation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_adjustment_saturation")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAdjustmentSaturation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_adjustment_color_correction")
    defer methodName.Destroy()
    ptrsForEnvironment.fnSetAdjustmentColorCorrection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1790811099))
  }
  {
    methodName := StringNameFromStr("get_adjustment_color_correction")
    defer methodName.Destroy()
    ptrsForEnvironment.fnGetAdjustmentColorCorrection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4037048985))
  }
}

type Environment struct {
  Resource
}

func (me *Environment) BaseClass() string {
  return "Environment"
}

func NewEnvironment() *Environment {
  str := StringNameFromStr("Environment") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Environment{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type EnvironmentBGMode int
const (
  EnvironmentBGModeBgClearColor EnvironmentBGMode = 0
  EnvironmentBGModeBgColor EnvironmentBGMode = 1
  EnvironmentBGModeBgSky EnvironmentBGMode = 2
  EnvironmentBGModeBgCanvas EnvironmentBGMode = 3
  EnvironmentBGModeBgKeep EnvironmentBGMode = 4
  EnvironmentBGModeBgCameraFeed EnvironmentBGMode = 5
  EnvironmentBGModeBgMax EnvironmentBGMode = 6
)

type EnvironmentAmbientSource int
const (
  EnvironmentAmbientSourceAmbientSourceBg EnvironmentAmbientSource = 0
  EnvironmentAmbientSourceAmbientSourceDisabled EnvironmentAmbientSource = 1
  EnvironmentAmbientSourceAmbientSourceColor EnvironmentAmbientSource = 2
  EnvironmentAmbientSourceAmbientSourceSky EnvironmentAmbientSource = 3
)

type EnvironmentReflectionSource int
const (
  EnvironmentReflectionSourceReflectionSourceBg EnvironmentReflectionSource = 0
  EnvironmentReflectionSourceReflectionSourceDisabled EnvironmentReflectionSource = 1
  EnvironmentReflectionSourceReflectionSourceSky EnvironmentReflectionSource = 2
)

type EnvironmentToneMapper int
const (
  EnvironmentToneMapperToneMapperLinear EnvironmentToneMapper = 0
  EnvironmentToneMapperToneMapperReinhardt EnvironmentToneMapper = 1
  EnvironmentToneMapperToneMapperFilmic EnvironmentToneMapper = 2
  EnvironmentToneMapperToneMapperAces EnvironmentToneMapper = 3
)

type EnvironmentGlowBlendMode int
const (
  EnvironmentGlowBlendModeGlowBlendModeAdditive EnvironmentGlowBlendMode = 0
  EnvironmentGlowBlendModeGlowBlendModeScreen EnvironmentGlowBlendMode = 1
  EnvironmentGlowBlendModeGlowBlendModeSoftlight EnvironmentGlowBlendMode = 2
  EnvironmentGlowBlendModeGlowBlendModeReplace EnvironmentGlowBlendMode = 3
  EnvironmentGlowBlendModeGlowBlendModeMix EnvironmentGlowBlendMode = 4
)

type EnvironmentSDFGIYScale int
const (
  EnvironmentSDFGIYScaleSdfgiYScale50Percent EnvironmentSDFGIYScale = 0
  EnvironmentSDFGIYScaleSdfgiYScale75Percent EnvironmentSDFGIYScale = 1
  EnvironmentSDFGIYScaleSdfgiYScale100Percent EnvironmentSDFGIYScale = 2
)

func (me *Environment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Environment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Environment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Environment) SetBackground(mode EnvironmentBGMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetBackground), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBackground() EnvironmentBGMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentBGMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetBackground), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetSky(sky Sky, )  {
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSky), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSky() Sky {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSky()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSky), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetSkyCustomFov(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSkyCustomFov), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSkyCustomFov() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSkyCustomFov), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSkyRotation(euler_radians Vector3, )  {
  cargs := []gdc.ConstTypePtr{euler_radians.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSkyRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSkyRotation() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSkyRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetBgColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetBgColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetBgEnergyMultiplier(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetBgEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgEnergyMultiplier() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetBgEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetBgIntensity(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetBgIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgIntensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetBgIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetCanvasMaxLayer(layer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetCanvasMaxLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetCanvasMaxLayer() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetCanvasMaxLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetCameraFeedId(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetCameraFeedId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetCameraFeedId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetCameraFeedId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAmbientLightColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAmbientLightColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAmbientLightColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetAmbientSource(source EnvironmentAmbientSource, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAmbientSource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientSource() EnvironmentAmbientSource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentAmbientSource

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAmbientSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetAmbientLightEnergy(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAmbientLightEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAmbientLightEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAmbientLightSkyContribution(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAmbientLightSkyContribution), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightSkyContribution() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAmbientLightSkyContribution), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetReflectionSource(source EnvironmentReflectionSource, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetReflectionSource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetReflectionSource() EnvironmentReflectionSource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentReflectionSource

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetReflectionSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetTonemapper(mode EnvironmentToneMapper, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetTonemapper), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapper() EnvironmentToneMapper {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentToneMapper

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetTonemapper), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetTonemapExposure(exposure float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetTonemapExposure), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapExposure() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetTonemapExposure), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetTonemapWhite(white float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&white) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetTonemapWhite), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapWhite() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetTonemapWhite), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsrEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsrEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSsrEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrMaxSteps(max_steps int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_steps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsrMaxSteps), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrMaxSteps() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsrMaxSteps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrFadeIn(fade_in float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_in) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsrFadeIn), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrFadeIn() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsrFadeIn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrFadeOut(fade_out float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_out) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsrFadeOut), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrFadeOut() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsrFadeOut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrDepthTolerance(depth_tolerance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth_tolerance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsrDepthTolerance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrDepthTolerance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsrDepthTolerance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsaoEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSsaoEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoIntensity(intensity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoIntensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoPower(power float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&power) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoPower), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoPower() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoPower), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoDetail(detail float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoDetail), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoDetail() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoDetail), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoHorizon(horizon float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoHorizon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoHorizon() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoHorizon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoSharpness(sharpness float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoSharpness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoDirectLightAffect(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoDirectLightAffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoDirectLightAffect() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoDirectLightAffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoAoChannelAffect(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsaoAoChannelAffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoAoChannelAffect() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsaoAoChannelAffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsilEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsilEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSsilEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsilRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsilRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilIntensity(intensity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsilIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilIntensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsilIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilSharpness(sharpness float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsilSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilSharpness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsilSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilNormalRejection(normal_rejection float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_rejection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSsilNormalRejection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilNormalRejection() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSsilNormalRejection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSdfgiEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiCascades(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiCascades), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiCascades() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiCascades), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiMinCellSize(size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiMinCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiMinCellSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiMinCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiMaxDistance(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiMaxDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiCascade0Distance(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiCascade0Distance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiCascade0Distance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiCascade0Distance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiYScale(scale EnvironmentSDFGIYScale, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiYScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiYScale() EnvironmentSDFGIYScale {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentSDFGIYScale

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiYScale), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetSdfgiUseOcclusion(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiUseOcclusion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiUsingOcclusion() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSdfgiUsingOcclusion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiBounceFeedback(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiBounceFeedback), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiBounceFeedback() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiBounceFeedback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiReadSkyLight(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiReadSkyLight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiReadingSkyLight() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsSdfgiReadingSkyLight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiEnergy(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiNormalBias(bias float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiNormalBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiNormalBias() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiNormalBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiProbeBias(bias float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetSdfgiProbeBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiProbeBias() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetSdfgiProbeBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsGlowEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsGlowEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowLevel(idx int64, intensity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowLevel(idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowNormalized(normalize bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowNormalized), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsGlowNormalized() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsGlowNormalized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowIntensity(intensity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowIntensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowStrength(strength float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowStrength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMix(mix float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowMix), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMix() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowMix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowBloom(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowBloom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowBloom() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowBloom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowBlendMode(mode EnvironmentGlowBlendMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowBlendMode() EnvironmentGlowBlendMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentGlowBlendMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetGlowHdrBleedThreshold(threshold float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowHdrBleedThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrBleedThreshold() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowHdrBleedThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowHdrBleedScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowHdrBleedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrBleedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowHdrBleedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowHdrLuminanceCap(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowHdrLuminanceCap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrLuminanceCap() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowHdrLuminanceCap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMapStrength(strength float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowMapStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMapStrength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowMapStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMap(mode Texture, )  {
  cargs := []gdc.ConstTypePtr{mode.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetGlowMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMap() Texture {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetGlowMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetFogEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsFogEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsFogEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogLightColor(light_color Color, )  {
  cargs := []gdc.ConstTypePtr{light_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogLightColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogLightColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogLightColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetFogLightEnergy(light_energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogLightEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogLightEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogLightEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogSunScatter(sun_scatter float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sun_scatter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogSunScatter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogSunScatter() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogSunScatter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogDensity(density float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogDensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogDensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogDensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogHeightDensity(height_density float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogHeightDensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogHeightDensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogHeightDensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogAerialPerspective(aerial_perspective float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aerial_perspective) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogAerialPerspective), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogAerialPerspective() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogAerialPerspective), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogSkyAffect(sky_affect float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetFogSkyAffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogSkyAffect() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetFogSkyAffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsVolumetricFogEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsVolumetricFogEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEmission(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogEmission), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogEmission() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogEmission), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetVolumetricFogAlbedo(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogAlbedo), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAlbedo() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogAlbedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetVolumetricFogDensity(density float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogDensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogDensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogDensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEmissionEnergy(begin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogEmissionEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogEmissionEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogEmissionEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogAnisotropy(anisotropy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anisotropy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogAnisotropy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAnisotropy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogAnisotropy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogLength(length float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogDetailSpread(detail_spread float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_spread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogDetailSpread), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogDetailSpread() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogDetailSpread), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogGiInject(gi_inject float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gi_inject) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogGiInject), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogGiInject() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogGiInject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogAmbientInject(enabled float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogAmbientInject), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAmbientInject() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogAmbientInject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogSkyAffect(sky_affect float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogSkyAffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogSkyAffect() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogSkyAffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogTemporalReprojectionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsVolumetricFogTemporalReprojectionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsVolumetricFogTemporalReprojectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temporal_reprojection_amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetVolumetricFogTemporalReprojectionAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogTemporalReprojectionAmount() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetVolumetricFogTemporalReprojectionAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAdjustmentEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsAdjustmentEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnIsAdjustmentEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentBrightness(brightness float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brightness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAdjustmentBrightness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentBrightness() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAdjustmentBrightness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentContrast(contrast float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contrast) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAdjustmentContrast), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentContrast() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAdjustmentContrast), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentSaturation(saturation float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&saturation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAdjustmentSaturation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentSaturation() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAdjustmentSaturation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentColorCorrection(color_correction Texture, )  {
  cargs := []gdc.ConstTypePtr{color_correction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnSetAdjustmentColorCorrection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentColorCorrection() Texture {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEnvironment.fnGetAdjustmentColorCorrection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
