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



// Enums

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

func (me *BaseMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BaseMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BaseMaterial3D) SetAlbedo(albedo Color, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAlbedo()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTransparency(transparency BaseMaterial3DTransparency, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTransparency()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAlphaAntialiasing()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAlphaAntialiasingEdge(edge float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAlphaAntialiasingEdge()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetShadingMode(shading_mode BaseMaterial3DShadingMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetShadingMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetSpecular(specular float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetSpecular()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetMetallic(metallic float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetMetallic()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRoughness(roughness float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRoughness()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetEmission(emission Color, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetEmission()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetEmissionEnergyMultiplier(emission_energy_multiplier float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetEmissionEnergyMultiplier()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetEmissionIntensity(emission_energy_multiplier float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetEmissionIntensity()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetNormalScale(normal_scale float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetNormalScale()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRim(rim float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRim()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRimTint(rim_tint float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRimTint()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetClearcoat(clearcoat float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetClearcoat()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetClearcoatRoughness(clearcoat_roughness float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetClearcoatRoughness()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAnisotropy(anisotropy float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAnisotropy()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapScale(heightmap_scale float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetHeightmapScale()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetSubsurfaceScatteringStrength(strength float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetSubsurfaceScatteringStrength()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTransmittanceColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTransmittanceColor()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTransmittanceDepth(depth float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTransmittanceDepth()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTransmittanceBoost(boost float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTransmittanceBoost()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetBacklight(backlight Color, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetBacklight()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRefraction(refraction float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRefraction()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetPointSize(point_size float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetPointSize()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDetailUv(detail_uv BaseMaterial3DDetailUV, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDetailUv()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetBlendMode(blend_mode BaseMaterial3DBlendMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetBlendMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDepthDrawMode(depth_draw_mode BaseMaterial3DDepthDrawMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDepthDrawMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetCullMode(cull_mode BaseMaterial3DCullMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetCullMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDiffuseMode(diffuse_mode BaseMaterial3DDiffuseMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDiffuseMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetSpecularMode(specular_mode BaseMaterial3DSpecularMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetSpecularMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetFlag(flag BaseMaterial3DFlags, enable bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetFlag(flag BaseMaterial3DFlags, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTextureFilter()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetFeature(feature BaseMaterial3DFeature, enable bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetFeature(feature BaseMaterial3DFeature, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetTexture(param BaseMaterial3DTextureParam, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetTexture(param BaseMaterial3DTextureParam, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDetailBlendMode(detail_blend_mode BaseMaterial3DBlendMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDetailBlendMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv1Scale(scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv1Scale()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv1Offset(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv1Offset()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv1TriplanarBlendSharpness(sharpness float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv1TriplanarBlendSharpness()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv2Scale(scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv2Scale()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv2Offset(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv2Offset()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetUv2TriplanarBlendSharpness(sharpness float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetUv2TriplanarBlendSharpness()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetBillboardMode()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetParticlesAnimHFrames(frames int, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetParticlesAnimHFrames()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetParticlesAnimVFrames(frames int, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetParticlesAnimVFrames()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetParticlesAnimLoop(loop bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetParticlesAnimLoop()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallax(enable bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) IsHeightmapDeepParallaxEnabled()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMinLayers(layer int, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMinLayers()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMaxLayers(layer int, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMaxLayers()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipTangent(flip bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipTangent()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipBinormal(flip bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipBinormal()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetGrow(amount float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetGrow()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetEmissionOperator(operator BaseMaterial3DEmissionOperator, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetEmissionOperator()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAoLightAffect(amount float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAoLightAffect()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAlphaScissorThreshold(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAlphaScissorThreshold()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAlphaHashScale(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAlphaHashScale()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetGrowEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) IsGrowEnabled()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetMetallicTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetMetallicTextureChannel()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRoughnessTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRoughnessTextureChannel()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetAoTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetAoTextureChannel()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetRefractionTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetRefractionTextureChannel()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetProximityFadeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) IsProximityFadeEnabled()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetProximityFadeDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetProximityFadeDistance()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetMsdfPixelRange(range_ float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetMsdfPixelRange()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetMsdfOutlineSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetMsdfOutlineSize()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDistanceFade(mode BaseMaterial3DDistanceFadeMode, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDistanceFade()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDistanceFadeMaxDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDistanceFadeMaxDistance()  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) SetDistanceFadeMinDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *BaseMaterial3D) GetDistanceFadeMinDistance()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
