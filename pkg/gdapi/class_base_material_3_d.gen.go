// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BaseMaterial3D struct {
  obj gdc.ObjectPtr
}

func (me *BaseMaterial3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BaseMaterial3D) BaseClass() string {
  return "BaseMaterial3D"
}

type BaseMaterial3DTextureParam int
const (
  BaseMaterial3DTextureParamTextureAlbedo BaseMaterial3DTextureParam = 0
  BaseMaterial3DTextureParamTextureMetallic BaseMaterial3DTextureParam = 1
  BaseMaterial3DTextureParamTextureRoughness BaseMaterial3DTextureParam = 2
  BaseMaterial3DTextureParamTextureEmission BaseMaterial3DTextureParam = 3
  BaseMaterial3DTextureParamTextureNormal BaseMaterial3DTextureParam = 4
  BaseMaterial3DTextureParamTextureRim BaseMaterial3DTextureParam = 5
  BaseMaterial3DTextureParamTextureClearcoat BaseMaterial3DTextureParam = 6
  BaseMaterial3DTextureParamTextureFlowmap BaseMaterial3DTextureParam = 7
  BaseMaterial3DTextureParamTextureAmbientOcclusion BaseMaterial3DTextureParam = 8
  BaseMaterial3DTextureParamTextureHeightmap BaseMaterial3DTextureParam = 9
  BaseMaterial3DTextureParamTextureSubsurfaceScattering BaseMaterial3DTextureParam = 10
  BaseMaterial3DTextureParamTextureSubsurfaceTransmittance BaseMaterial3DTextureParam = 11
  BaseMaterial3DTextureParamTextureBacklight BaseMaterial3DTextureParam = 12
  BaseMaterial3DTextureParamTextureRefraction BaseMaterial3DTextureParam = 13
  BaseMaterial3DTextureParamTextureDetailMask BaseMaterial3DTextureParam = 14
  BaseMaterial3DTextureParamTextureDetailAlbedo BaseMaterial3DTextureParam = 15
  BaseMaterial3DTextureParamTextureDetailNormal BaseMaterial3DTextureParam = 16
  BaseMaterial3DTextureParamTextureOrm BaseMaterial3DTextureParam = 17
  BaseMaterial3DTextureParamTextureMax BaseMaterial3DTextureParam = 18
)

type BaseMaterial3DTextureFilter int
const (
  BaseMaterial3DTextureFilterTextureFilterNearest BaseMaterial3DTextureFilter = 0
  BaseMaterial3DTextureFilterTextureFilterLinear BaseMaterial3DTextureFilter = 1
  BaseMaterial3DTextureFilterTextureFilterNearestWithMipmaps BaseMaterial3DTextureFilter = 2
  BaseMaterial3DTextureFilterTextureFilterLinearWithMipmaps BaseMaterial3DTextureFilter = 3
  BaseMaterial3DTextureFilterTextureFilterNearestWithMipmapsAnisotropic BaseMaterial3DTextureFilter = 4
  BaseMaterial3DTextureFilterTextureFilterLinearWithMipmapsAnisotropic BaseMaterial3DTextureFilter = 5
  BaseMaterial3DTextureFilterTextureFilterMax BaseMaterial3DTextureFilter = 6
)

type BaseMaterial3DDetailUV int
const (
  BaseMaterial3DDetailUVDetailUv1 BaseMaterial3DDetailUV = 0
  BaseMaterial3DDetailUVDetailUv2 BaseMaterial3DDetailUV = 1
)

type BaseMaterial3DTransparency int
const (
  BaseMaterial3DTransparencyTransparencyDisabled BaseMaterial3DTransparency = 0
  BaseMaterial3DTransparencyTransparencyAlpha BaseMaterial3DTransparency = 1
  BaseMaterial3DTransparencyTransparencyAlphaScissor BaseMaterial3DTransparency = 2
  BaseMaterial3DTransparencyTransparencyAlphaHash BaseMaterial3DTransparency = 3
  BaseMaterial3DTransparencyTransparencyAlphaDepthPrePass BaseMaterial3DTransparency = 4
  BaseMaterial3DTransparencyTransparencyMax BaseMaterial3DTransparency = 5
)

type BaseMaterial3DShadingMode int
const (
  BaseMaterial3DShadingModeShadingModeUnshaded BaseMaterial3DShadingMode = 0
  BaseMaterial3DShadingModeShadingModePerPixel BaseMaterial3DShadingMode = 1
  BaseMaterial3DShadingModeShadingModePerVertex BaseMaterial3DShadingMode = 2
  BaseMaterial3DShadingModeShadingModeMax BaseMaterial3DShadingMode = 3
)

type BaseMaterial3DFeature int
const (
  BaseMaterial3DFeatureFeatureEmission BaseMaterial3DFeature = 0
  BaseMaterial3DFeatureFeatureNormalMapping BaseMaterial3DFeature = 1
  BaseMaterial3DFeatureFeatureRim BaseMaterial3DFeature = 2
  BaseMaterial3DFeatureFeatureClearcoat BaseMaterial3DFeature = 3
  BaseMaterial3DFeatureFeatureAnisotropy BaseMaterial3DFeature = 4
  BaseMaterial3DFeatureFeatureAmbientOcclusion BaseMaterial3DFeature = 5
  BaseMaterial3DFeatureFeatureHeightMapping BaseMaterial3DFeature = 6
  BaseMaterial3DFeatureFeatureSubsurfaceScattering BaseMaterial3DFeature = 7
  BaseMaterial3DFeatureFeatureSubsurfaceTransmittance BaseMaterial3DFeature = 8
  BaseMaterial3DFeatureFeatureBacklight BaseMaterial3DFeature = 9
  BaseMaterial3DFeatureFeatureRefraction BaseMaterial3DFeature = 10
  BaseMaterial3DFeatureFeatureDetail BaseMaterial3DFeature = 11
  BaseMaterial3DFeatureFeatureMax BaseMaterial3DFeature = 12
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
  BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingOff BaseMaterial3DAlphaAntiAliasing = 0
  BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingAlphaToCoverage BaseMaterial3DAlphaAntiAliasing = 1
  BaseMaterial3DAlphaAntiAliasingAlphaAntialiasingAlphaToCoverageAndToOne BaseMaterial3DAlphaAntiAliasing = 2
)

type BaseMaterial3DDepthDrawMode int
const (
  BaseMaterial3DDepthDrawModeDepthDrawOpaqueOnly BaseMaterial3DDepthDrawMode = 0
  BaseMaterial3DDepthDrawModeDepthDrawAlways BaseMaterial3DDepthDrawMode = 1
  BaseMaterial3DDepthDrawModeDepthDrawDisabled BaseMaterial3DDepthDrawMode = 2
)

type BaseMaterial3DCullMode int
const (
  BaseMaterial3DCullModeCullBack BaseMaterial3DCullMode = 0
  BaseMaterial3DCullModeCullFront BaseMaterial3DCullMode = 1
  BaseMaterial3DCullModeCullDisabled BaseMaterial3DCullMode = 2
)

type BaseMaterial3DFlags int
const (
  BaseMaterial3DFlagsFlagDisableDepthTest BaseMaterial3DFlags = 0
  BaseMaterial3DFlagsFlagAlbedoFromVertexColor BaseMaterial3DFlags = 1
  BaseMaterial3DFlagsFlagSrgbVertexColor BaseMaterial3DFlags = 2
  BaseMaterial3DFlagsFlagUsePointSize BaseMaterial3DFlags = 3
  BaseMaterial3DFlagsFlagFixedSize BaseMaterial3DFlags = 4
  BaseMaterial3DFlagsFlagBillboardKeepScale BaseMaterial3DFlags = 5
  BaseMaterial3DFlagsFlagUv1UseTriplanar BaseMaterial3DFlags = 6
  BaseMaterial3DFlagsFlagUv2UseTriplanar BaseMaterial3DFlags = 7
  BaseMaterial3DFlagsFlagUv1UseWorldTriplanar BaseMaterial3DFlags = 8
  BaseMaterial3DFlagsFlagUv2UseWorldTriplanar BaseMaterial3DFlags = 9
  BaseMaterial3DFlagsFlagAoOnUv2 BaseMaterial3DFlags = 10
  BaseMaterial3DFlagsFlagEmissionOnUv2 BaseMaterial3DFlags = 11
  BaseMaterial3DFlagsFlagAlbedoTextureForceSrgb BaseMaterial3DFlags = 12
  BaseMaterial3DFlagsFlagDontReceiveShadows BaseMaterial3DFlags = 13
  BaseMaterial3DFlagsFlagDisableAmbientLight BaseMaterial3DFlags = 14
  BaseMaterial3DFlagsFlagUseShadowToOpacity BaseMaterial3DFlags = 15
  BaseMaterial3DFlagsFlagUseTextureRepeat BaseMaterial3DFlags = 16
  BaseMaterial3DFlagsFlagInvertHeightmap BaseMaterial3DFlags = 17
  BaseMaterial3DFlagsFlagSubsurfaceModeSkin BaseMaterial3DFlags = 18
  BaseMaterial3DFlagsFlagParticleTrailsMode BaseMaterial3DFlags = 19
  BaseMaterial3DFlagsFlagAlbedoTextureMsdf BaseMaterial3DFlags = 20
  BaseMaterial3DFlagsFlagMax BaseMaterial3DFlags = 21
)

type BaseMaterial3DDiffuseMode int
const (
  BaseMaterial3DDiffuseModeDiffuseBurley BaseMaterial3DDiffuseMode = 0
  BaseMaterial3DDiffuseModeDiffuseLambert BaseMaterial3DDiffuseMode = 1
  BaseMaterial3DDiffuseModeDiffuseLambertWrap BaseMaterial3DDiffuseMode = 2
  BaseMaterial3DDiffuseModeDiffuseToon BaseMaterial3DDiffuseMode = 3
)

type BaseMaterial3DSpecularMode int
const (
  BaseMaterial3DSpecularModeSpecularSchlickGgx BaseMaterial3DSpecularMode = 0
  BaseMaterial3DSpecularModeSpecularToon BaseMaterial3DSpecularMode = 1
  BaseMaterial3DSpecularModeSpecularDisabled BaseMaterial3DSpecularMode = 2
)

type BaseMaterial3DBillboardMode int
const (
  BaseMaterial3DBillboardModeBillboardDisabled BaseMaterial3DBillboardMode = 0
  BaseMaterial3DBillboardModeBillboardEnabled BaseMaterial3DBillboardMode = 1
  BaseMaterial3DBillboardModeBillboardFixedY BaseMaterial3DBillboardMode = 2
  BaseMaterial3DBillboardModeBillboardParticles BaseMaterial3DBillboardMode = 3
)

type BaseMaterial3DTextureChannel int
const (
  BaseMaterial3DTextureChannelTextureChannelRed BaseMaterial3DTextureChannel = 0
  BaseMaterial3DTextureChannelTextureChannelGreen BaseMaterial3DTextureChannel = 1
  BaseMaterial3DTextureChannelTextureChannelBlue BaseMaterial3DTextureChannel = 2
  BaseMaterial3DTextureChannelTextureChannelAlpha BaseMaterial3DTextureChannel = 3
  BaseMaterial3DTextureChannelTextureChannelGrayscale BaseMaterial3DTextureChannel = 4
)

type BaseMaterial3DEmissionOperator int
const (
  BaseMaterial3DEmissionOperatorEmissionOpAdd BaseMaterial3DEmissionOperator = 0
  BaseMaterial3DEmissionOperatorEmissionOpMultiply BaseMaterial3DEmissionOperator = 1
)

type BaseMaterial3DDistanceFadeMode int
const (
  BaseMaterial3DDistanceFadeModeDistanceFadeDisabled BaseMaterial3DDistanceFadeMode = 0
  BaseMaterial3DDistanceFadeModeDistanceFadePixelAlpha BaseMaterial3DDistanceFadeMode = 1
  BaseMaterial3DDistanceFadeModeDistanceFadePixelDither BaseMaterial3DDistanceFadeMode = 2
  BaseMaterial3DDistanceFadeModeDistanceFadeObjectDither BaseMaterial3DDistanceFadeMode = 3
)

func  (me *BaseMaterial3D) SetAlbedo(albedo Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAlbedo() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTransparency(transparency BaseMaterial3DTransparency, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTransparency() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAlphaAntialiasing() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAlphaAntialiasingEdge(edge float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAlphaAntialiasingEdge() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetShadingMode(shading_mode BaseMaterial3DShadingMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetShadingMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetSpecular(specular float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetSpecular() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetMetallic(metallic float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetMetallic() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRoughness(roughness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRoughness() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetEmission(emission Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetEmission() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetEmissionEnergyMultiplier(emission_energy_multiplier float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetEmissionEnergyMultiplier() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetEmissionIntensity(emission_energy_multiplier float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetEmissionIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetNormalScale(normal_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetNormalScale() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRim(rim float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRim() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRimTint(rim_tint float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRimTint() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetClearcoat(clearcoat float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetClearcoat() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetClearcoatRoughness(clearcoat_roughness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetClearcoatRoughness() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAnisotropy(anisotropy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAnisotropy() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapScale(heightmap_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetHeightmapScale() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetSubsurfaceScatteringStrength(strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetSubsurfaceScatteringStrength() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTransmittanceColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTransmittanceColor() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTransmittanceDepth(depth float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTransmittanceDepth() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTransmittanceBoost(boost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTransmittanceBoost() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetBacklight(backlight Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetBacklight() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRefraction(refraction float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRefraction() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetPointSize(point_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetPointSize() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDetailUv(detail_uv BaseMaterial3DDetailUV, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDetailUv() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetBlendMode(blend_mode BaseMaterial3DBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDepthDrawMode(depth_draw_mode BaseMaterial3DDepthDrawMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDepthDrawMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetCullMode(cull_mode BaseMaterial3DCullMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetCullMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDiffuseMode(diffuse_mode BaseMaterial3DDiffuseMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDiffuseMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetSpecularMode(specular_mode BaseMaterial3DSpecularMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetSpecularMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetFlag(flag BaseMaterial3DFlags, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetFlag(flag BaseMaterial3DFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTextureFilter() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetFeature(feature BaseMaterial3DFeature, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetFeature(feature BaseMaterial3DFeature, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetTexture(param BaseMaterial3DTextureParam, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetTexture(param BaseMaterial3DTextureParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDetailBlendMode(detail_blend_mode BaseMaterial3DBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDetailBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv1Scale(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv1Scale() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv1Offset(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv1Offset() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv1TriplanarBlendSharpness(sharpness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv1TriplanarBlendSharpness() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv2Scale(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv2Scale() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv2Offset(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv2Offset() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetUv2TriplanarBlendSharpness(sharpness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetUv2TriplanarBlendSharpness() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetBillboardMode() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetParticlesAnimHFrames(frames int, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetParticlesAnimHFrames() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetParticlesAnimVFrames(frames int, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetParticlesAnimVFrames() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetParticlesAnimLoop(loop bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetParticlesAnimLoop() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallax(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) IsHeightmapDeepParallaxEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMinLayers(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMinLayers() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMaxLayers(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMaxLayers() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipTangent(flip bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipTangent() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipBinormal(flip bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipBinormal() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetGrow(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetGrow() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetEmissionOperator(operator BaseMaterial3DEmissionOperator, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetEmissionOperator() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAoLightAffect(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAoLightAffect() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAlphaScissorThreshold(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAlphaScissorThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAlphaHashScale(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAlphaHashScale() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetGrowEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) IsGrowEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetMetallicTextureChannel(channel BaseMaterial3DTextureChannel, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetMetallicTextureChannel() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRoughnessTextureChannel(channel BaseMaterial3DTextureChannel, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRoughnessTextureChannel() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetAoTextureChannel(channel BaseMaterial3DTextureChannel, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetAoTextureChannel() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetRefractionTextureChannel(channel BaseMaterial3DTextureChannel, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetRefractionTextureChannel() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetProximityFadeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) IsProximityFadeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetProximityFadeDistance(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetProximityFadeDistance() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetMsdfPixelRange(range_ float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetMsdfPixelRange() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetMsdfOutlineSize(size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetMsdfOutlineSize() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDistanceFade(mode BaseMaterial3DDistanceFadeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDistanceFade() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDistanceFadeMaxDistance(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDistanceFadeMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) SetDistanceFadeMinDistance(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *BaseMaterial3D) GetDistanceFadeMinDistance() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
