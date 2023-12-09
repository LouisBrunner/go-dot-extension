// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Environment struct {
  obj gdc.ObjectPtr
}

func (me *Environment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Environment) BaseClass() string {
  return "Environment"
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
  panic("TODO: implement")
}

func  (me *Environment) GetBackground()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSky(sky Sky, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSky()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSkyCustomFov(scale float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSkyCustomFov()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSkyRotation(euler_radians Vector3, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSkyRotation()  {
  panic("TODO: implement")
}

func  (me *Environment) SetBgColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetBgColor()  {
  panic("TODO: implement")
}

func  (me *Environment) SetBgEnergyMultiplier(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetBgEnergyMultiplier()  {
  panic("TODO: implement")
}

func  (me *Environment) SetBgIntensity(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetBgIntensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetCanvasMaxLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetCanvasMaxLayer()  {
  panic("TODO: implement")
}

func  (me *Environment) SetCameraFeedId(id int, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetCameraFeedId()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAmbientLightColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAmbientLightColor()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAmbientSource(source EnvironmentAmbientSource, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAmbientSource()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAmbientLightEnergy(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAmbientLightEnergy()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAmbientLightSkyContribution(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAmbientLightSkyContribution()  {
  panic("TODO: implement")
}

func  (me *Environment) SetReflectionSource(source EnvironmentReflectionSource, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetReflectionSource()  {
  panic("TODO: implement")
}

func  (me *Environment) SetTonemapper(mode EnvironmentToneMapper, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetTonemapper()  {
  panic("TODO: implement")
}

func  (me *Environment) SetTonemapExposure(exposure float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetTonemapExposure()  {
  panic("TODO: implement")
}

func  (me *Environment) SetTonemapWhite(white float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetTonemapWhite()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsrEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSsrEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsrMaxSteps(max_steps int, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsrMaxSteps()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsrFadeIn(fade_in float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsrFadeIn()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsrFadeOut(fade_out float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsrFadeOut()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsrDepthTolerance(depth_tolerance float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsrDepthTolerance()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSsaoEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoRadius()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoIntensity(intensity float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoIntensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoPower(power float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoPower()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoDetail(detail float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoDetail()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoHorizon(horizon float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoHorizon()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoSharpness(sharpness float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoSharpness()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoDirectLightAffect(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoDirectLightAffect()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsaoAoChannelAffect(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsaoAoChannelAffect()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsilEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSsilEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsilRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsilRadius()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsilIntensity(intensity float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsilIntensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsilSharpness(sharpness float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsilSharpness()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSsilNormalRejection(normal_rejection float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSsilNormalRejection()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSdfgiEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiCascades(amount int, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiCascades()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiMinCellSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiMinCellSize()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiMaxDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiMaxDistance()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiCascade0Distance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiCascade0Distance()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiYScale(scale EnvironmentSDFGIYScale, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiYScale()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiUseOcclusion(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSdfgiUsingOcclusion()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiBounceFeedback(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiBounceFeedback()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiReadSkyLight(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsSdfgiReadingSkyLight()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiEnergy(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiEnergy()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiNormalBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiNormalBias()  {
  panic("TODO: implement")
}

func  (me *Environment) SetSdfgiProbeBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetSdfgiProbeBias()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsGlowEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowLevel(idx int, intensity float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowLevel(idx int, )  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowNormalized(normalize bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsGlowNormalized()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowIntensity(intensity float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowIntensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowStrength(strength float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowStrength()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowMix(mix float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowMix()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowBloom(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowBloom()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowBlendMode(mode EnvironmentGlowBlendMode, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowBlendMode()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowHdrBleedThreshold(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowHdrBleedThreshold()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowHdrBleedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowHdrBleedScale()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowHdrLuminanceCap(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowHdrLuminanceCap()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowMapStrength(strength float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowMapStrength()  {
  panic("TODO: implement")
}

func  (me *Environment) SetGlowMap(mode Texture, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetGlowMap()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsFogEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogLightColor(light_color Color, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogLightColor()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogLightEnergy(light_energy float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogLightEnergy()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogSunScatter(sun_scatter float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogSunScatter()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogDensity(density float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogDensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogHeight()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogHeightDensity(height_density float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogHeightDensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogAerialPerspective(aerial_perspective float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogAerialPerspective()  {
  panic("TODO: implement")
}

func  (me *Environment) SetFogSkyAffect(sky_affect float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetFogSkyAffect()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsVolumetricFogEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogEmission(color Color, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogEmission()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogAlbedo(color Color, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogAlbedo()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogDensity(density float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogDensity()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogEmissionEnergy(begin float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogEmissionEnergy()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogAnisotropy(anisotropy float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogAnisotropy()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogLength()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogDetailSpread(detail_spread float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogDetailSpread()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogGiInject(gi_inject float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogGiInject()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogAmbientInject(enabled float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogAmbientInject()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogSkyAffect(sky_affect float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogSkyAffect()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsVolumetricFogTemporalReprojectionEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetVolumetricFogTemporalReprojectionAmount()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAdjustmentEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Environment) IsAdjustmentEnabled()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAdjustmentBrightness(brightness float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAdjustmentBrightness()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAdjustmentContrast(contrast float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAdjustmentContrast()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAdjustmentSaturation(saturation float32, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAdjustmentSaturation()  {
  panic("TODO: implement")
}

func  (me *Environment) SetAdjustmentColorCorrection(color_correction Texture, )  {
  panic("TODO: implement")
}

func  (me *Environment) GetAdjustmentColorCorrection()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
