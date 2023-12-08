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

func  (me *Environment) SetBackground(mode EnvironmentBGMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetBackground() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSky(sky Sky, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSky() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSkyCustomFov(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSkyCustomFov() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSkyRotation(euler_radians Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSkyRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetBgColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetBgColor() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetBgEnergyMultiplier(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetBgEnergyMultiplier() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetBgIntensity(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetBgIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetCanvasMaxLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetCanvasMaxLayer() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetCameraFeedId(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetCameraFeedId() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAmbientLightColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAmbientLightColor() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAmbientSource(source EnvironmentAmbientSource, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAmbientSource() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAmbientLightEnergy(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAmbientLightEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAmbientLightSkyContribution(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAmbientLightSkyContribution() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetReflectionSource(source EnvironmentReflectionSource, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetReflectionSource() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetTonemapper(mode EnvironmentToneMapper, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetTonemapper() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetTonemapExposure(exposure float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetTonemapExposure() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetTonemapWhite(white float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetTonemapWhite() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsrEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSsrEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsrMaxSteps(max_steps int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsrMaxSteps() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsrFadeIn(fade_in float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsrFadeIn() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsrFadeOut(fade_out float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsrFadeOut() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsrDepthTolerance(depth_tolerance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsrDepthTolerance() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSsaoEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoRadius() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoIntensity(intensity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoPower(power float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoPower() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoDetail(detail float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoDetail() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoHorizon(horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoHorizon() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoSharpness(sharpness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoSharpness() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoDirectLightAffect(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoDirectLightAffect() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsaoAoChannelAffect(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsaoAoChannelAffect() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsilEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSsilEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsilRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsilRadius() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsilIntensity(intensity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsilIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsilSharpness(sharpness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsilSharpness() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSsilNormalRejection(normal_rejection float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSsilNormalRejection() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSdfgiEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiCascades(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiCascades() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiMinCellSize(size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiMinCellSize() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiMaxDistance(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiCascade0Distance(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiCascade0Distance() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiYScale(scale EnvironmentSDFGIYScale, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiYScale() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiUseOcclusion(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSdfgiUsingOcclusion() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiBounceFeedback(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiBounceFeedback() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiReadSkyLight(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsSdfgiReadingSkyLight() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiEnergy(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiNormalBias(bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiNormalBias() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetSdfgiProbeBias(bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetSdfgiProbeBias() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsGlowEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowLevel(idx int, intensity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowLevel(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowNormalized(normalize bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsGlowNormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowIntensity(intensity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowStrength(strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowStrength() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowMix(mix float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowMix() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowBloom(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowBloom() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowBlendMode(mode EnvironmentGlowBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowHdrBleedThreshold(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowHdrBleedThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowHdrBleedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowHdrBleedScale() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowHdrLuminanceCap(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowHdrLuminanceCap() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowMapStrength(strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowMapStrength() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetGlowMap(mode Texture, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetGlowMap() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsFogEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogLightColor(light_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogLightColor() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogLightEnergy(light_energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogLightEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogSunScatter(sun_scatter float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogSunScatter() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogDensity(density float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogDensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogHeight(height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogHeight() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogHeightDensity(height_density float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogHeightDensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogAerialPerspective(aerial_perspective float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogAerialPerspective() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetFogSkyAffect(sky_affect float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetFogSkyAffect() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsVolumetricFogEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogEmission(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogEmission() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogAlbedo(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogAlbedo() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogDensity(density float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogDensity() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogEmissionEnergy(begin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogEmissionEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogAnisotropy(anisotropy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogAnisotropy() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogLength(length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogLength() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogDetailSpread(detail_spread float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogDetailSpread() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogGiInject(gi_inject float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogGiInject() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogAmbientInject(enabled float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogAmbientInject() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogSkyAffect(sky_affect float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogSkyAffect() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsVolumetricFogTemporalReprojectionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetVolumetricFogTemporalReprojectionAmount() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAdjustmentEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) IsAdjustmentEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAdjustmentBrightness(brightness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAdjustmentBrightness() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAdjustmentContrast(contrast float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAdjustmentContrast() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAdjustmentSaturation(saturation float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAdjustmentSaturation() { // TODO: return value
  // TODO: implement
}

func  (me *Environment) SetAdjustmentColorCorrection(color_correction Texture, ) { // TODO: return value
  // TODO: implement
}

func  (me *Environment) GetAdjustmentColorCorrection() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
