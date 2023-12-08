// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  BaseMaterial3DTextureAlbedo BaseMaterial3DTextureParam = 0
  BaseMaterial3DTextureMetallic BaseMaterial3DTextureParam = 1
  BaseMaterial3DTextureRoughness BaseMaterial3DTextureParam = 2
  BaseMaterial3DTextureEmission BaseMaterial3DTextureParam = 3
  BaseMaterial3DTextureNormal BaseMaterial3DTextureParam = 4
  BaseMaterial3DTextureRim BaseMaterial3DTextureParam = 5
  BaseMaterial3DTextureClearcoat BaseMaterial3DTextureParam = 6
  BaseMaterial3DTextureFlowmap BaseMaterial3DTextureParam = 7
  BaseMaterial3DTextureAmbientOcclusion BaseMaterial3DTextureParam = 8
  BaseMaterial3DTextureHeightmap BaseMaterial3DTextureParam = 9
  BaseMaterial3DTextureSubsurfaceScattering BaseMaterial3DTextureParam = 10
  BaseMaterial3DTextureSubsurfaceTransmittance BaseMaterial3DTextureParam = 11
  BaseMaterial3DTextureBacklight BaseMaterial3DTextureParam = 12
  BaseMaterial3DTextureRefraction BaseMaterial3DTextureParam = 13
  BaseMaterial3DTextureDetailMask BaseMaterial3DTextureParam = 14
  BaseMaterial3DTextureDetailAlbedo BaseMaterial3DTextureParam = 15
  BaseMaterial3DTextureDetailNormal BaseMaterial3DTextureParam = 16
  BaseMaterial3DTextureOrm BaseMaterial3DTextureParam = 17
  BaseMaterial3DTextureMax BaseMaterial3DTextureParam = 18
)

type BaseMaterial3DTextureFilter int
const (
  BaseMaterial3DTextureFilterNearest BaseMaterial3DTextureFilter = 0
  BaseMaterial3DTextureFilterLinear BaseMaterial3DTextureFilter = 1
  BaseMaterial3DTextureFilterNearestWithMipmaps BaseMaterial3DTextureFilter = 2
  BaseMaterial3DTextureFilterLinearWithMipmaps BaseMaterial3DTextureFilter = 3
  BaseMaterial3DTextureFilterNearestWithMipmapsAnisotropic BaseMaterial3DTextureFilter = 4
  BaseMaterial3DTextureFilterLinearWithMipmapsAnisotropic BaseMaterial3DTextureFilter = 5
  BaseMaterial3DTextureFilterMax BaseMaterial3DTextureFilter = 6
)

type BaseMaterial3DDetailUV int
const (
  BaseMaterial3DDetailUv1 BaseMaterial3DDetailUV = 0
  BaseMaterial3DDetailUv2 BaseMaterial3DDetailUV = 1
)

type BaseMaterial3DTransparency int
const (
  BaseMaterial3DTransparencyDisabled BaseMaterial3DTransparency = 0
  BaseMaterial3DTransparencyAlpha BaseMaterial3DTransparency = 1
  BaseMaterial3DTransparencyAlphaScissor BaseMaterial3DTransparency = 2
  BaseMaterial3DTransparencyAlphaHash BaseMaterial3DTransparency = 3
  BaseMaterial3DTransparencyAlphaDepthPrePass BaseMaterial3DTransparency = 4
  BaseMaterial3DTransparencyMax BaseMaterial3DTransparency = 5
)

type BaseMaterial3DShadingMode int
const (
  BaseMaterial3DShadingModeUnshaded BaseMaterial3DShadingMode = 0
  BaseMaterial3DShadingModePerPixel BaseMaterial3DShadingMode = 1
  BaseMaterial3DShadingModePerVertex BaseMaterial3DShadingMode = 2
  BaseMaterial3DShadingModeMax BaseMaterial3DShadingMode = 3
)

type BaseMaterial3DFeature int
const (
  BaseMaterial3DFeatureEmission BaseMaterial3DFeature = 0
  BaseMaterial3DFeatureNormalMapping BaseMaterial3DFeature = 1
  BaseMaterial3DFeatureRim BaseMaterial3DFeature = 2
  BaseMaterial3DFeatureClearcoat BaseMaterial3DFeature = 3
  BaseMaterial3DFeatureAnisotropy BaseMaterial3DFeature = 4
  BaseMaterial3DFeatureAmbientOcclusion BaseMaterial3DFeature = 5
  BaseMaterial3DFeatureHeightMapping BaseMaterial3DFeature = 6
  BaseMaterial3DFeatureSubsurfaceScattering BaseMaterial3DFeature = 7
  BaseMaterial3DFeatureSubsurfaceTransmittance BaseMaterial3DFeature = 8
  BaseMaterial3DFeatureBacklight BaseMaterial3DFeature = 9
  BaseMaterial3DFeatureRefraction BaseMaterial3DFeature = 10
  BaseMaterial3DFeatureDetail BaseMaterial3DFeature = 11
  BaseMaterial3DFeatureMax BaseMaterial3DFeature = 12
)

type BaseMaterial3DBlendMode int
const (
  BaseMaterial3DBlendModeMix BaseMaterial3DBlendMode = 0
  BaseMaterial3DBlendModeAdd BaseMaterial3DBlendMode = 1
  BaseMaterial3DBlendModeSub BaseMaterial3DBlendMode = 2
  BaseMaterial3DBlendModeMul BaseMaterial3DBlendMode = 3
)

type BaseMaterial3DAlphaAntiAliasing int
const (
  BaseMaterial3DAlphaAntialiasingOff BaseMaterial3DAlphaAntiAliasing = 0
  BaseMaterial3DAlphaAntialiasingAlphaToCoverage BaseMaterial3DAlphaAntiAliasing = 1
  BaseMaterial3DAlphaAntialiasingAlphaToCoverageAndToOne BaseMaterial3DAlphaAntiAliasing = 2
)

type BaseMaterial3DDepthDrawMode int
const (
  BaseMaterial3DDepthDrawOpaqueOnly BaseMaterial3DDepthDrawMode = 0
  BaseMaterial3DDepthDrawAlways BaseMaterial3DDepthDrawMode = 1
  BaseMaterial3DDepthDrawDisabled BaseMaterial3DDepthDrawMode = 2
)

type BaseMaterial3DCullMode int
const (
  BaseMaterial3DCullBack BaseMaterial3DCullMode = 0
  BaseMaterial3DCullFront BaseMaterial3DCullMode = 1
  BaseMaterial3DCullDisabled BaseMaterial3DCullMode = 2
)

type BaseMaterial3DFlags int
const (
  BaseMaterial3DFlagDisableDepthTest BaseMaterial3DFlags = 0
  BaseMaterial3DFlagAlbedoFromVertexColor BaseMaterial3DFlags = 1
  BaseMaterial3DFlagSrgbVertexColor BaseMaterial3DFlags = 2
  BaseMaterial3DFlagUsePointSize BaseMaterial3DFlags = 3
  BaseMaterial3DFlagFixedSize BaseMaterial3DFlags = 4
  BaseMaterial3DFlagBillboardKeepScale BaseMaterial3DFlags = 5
  BaseMaterial3DFlagUv1UseTriplanar BaseMaterial3DFlags = 6
  BaseMaterial3DFlagUv2UseTriplanar BaseMaterial3DFlags = 7
  BaseMaterial3DFlagUv1UseWorldTriplanar BaseMaterial3DFlags = 8
  BaseMaterial3DFlagUv2UseWorldTriplanar BaseMaterial3DFlags = 9
  BaseMaterial3DFlagAoOnUv2 BaseMaterial3DFlags = 10
  BaseMaterial3DFlagEmissionOnUv2 BaseMaterial3DFlags = 11
  BaseMaterial3DFlagAlbedoTextureForceSrgb BaseMaterial3DFlags = 12
  BaseMaterial3DFlagDontReceiveShadows BaseMaterial3DFlags = 13
  BaseMaterial3DFlagDisableAmbientLight BaseMaterial3DFlags = 14
  BaseMaterial3DFlagUseShadowToOpacity BaseMaterial3DFlags = 15
  BaseMaterial3DFlagUseTextureRepeat BaseMaterial3DFlags = 16
  BaseMaterial3DFlagInvertHeightmap BaseMaterial3DFlags = 17
  BaseMaterial3DFlagSubsurfaceModeSkin BaseMaterial3DFlags = 18
  BaseMaterial3DFlagParticleTrailsMode BaseMaterial3DFlags = 19
  BaseMaterial3DFlagAlbedoTextureMsdf BaseMaterial3DFlags = 20
  BaseMaterial3DFlagMax BaseMaterial3DFlags = 21
)

type BaseMaterial3DDiffuseMode int
const (
  BaseMaterial3DDiffuseBurley BaseMaterial3DDiffuseMode = 0
  BaseMaterial3DDiffuseLambert BaseMaterial3DDiffuseMode = 1
  BaseMaterial3DDiffuseLambertWrap BaseMaterial3DDiffuseMode = 2
  BaseMaterial3DDiffuseToon BaseMaterial3DDiffuseMode = 3
)

type BaseMaterial3DSpecularMode int
const (
  BaseMaterial3DSpecularSchlickGgx BaseMaterial3DSpecularMode = 0
  BaseMaterial3DSpecularToon BaseMaterial3DSpecularMode = 1
  BaseMaterial3DSpecularDisabled BaseMaterial3DSpecularMode = 2
)

type BaseMaterial3DBillboardMode int
const (
  BaseMaterial3DBillboardDisabled BaseMaterial3DBillboardMode = 0
  BaseMaterial3DBillboardEnabled BaseMaterial3DBillboardMode = 1
  BaseMaterial3DBillboardFixedY BaseMaterial3DBillboardMode = 2
  BaseMaterial3DBillboardParticles BaseMaterial3DBillboardMode = 3
)

type BaseMaterial3DTextureChannel int
const (
  BaseMaterial3DTextureChannelRed BaseMaterial3DTextureChannel = 0
  BaseMaterial3DTextureChannelGreen BaseMaterial3DTextureChannel = 1
  BaseMaterial3DTextureChannelBlue BaseMaterial3DTextureChannel = 2
  BaseMaterial3DTextureChannelAlpha BaseMaterial3DTextureChannel = 3
  BaseMaterial3DTextureChannelGrayscale BaseMaterial3DTextureChannel = 4
)

type BaseMaterial3DEmissionOperator int
const (
  BaseMaterial3DEmissionOpAdd BaseMaterial3DEmissionOperator = 0
  BaseMaterial3DEmissionOpMultiply BaseMaterial3DEmissionOperator = 1
)

type BaseMaterial3DDistanceFadeMode int
const (
  BaseMaterial3DDistanceFadeDisabled BaseMaterial3DDistanceFadeMode = 0
  BaseMaterial3DDistanceFadePixelAlpha BaseMaterial3DDistanceFadeMode = 1
  BaseMaterial3DDistanceFadePixelDither BaseMaterial3DDistanceFadeMode = 2
  BaseMaterial3DDistanceFadeObjectDither BaseMaterial3DDistanceFadeMode = 3
)
