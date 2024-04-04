// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  BaseMaterial3DFlagsFlagDisableFog BaseMaterial3DFlags = 21
  BaseMaterial3DFlagsFlagMax BaseMaterial3DFlags = 22
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

func  (me *BaseMaterial3D) SetAlbedo(albedo Color, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{albedo.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAlbedo() Color {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetTransparency(transparency BaseMaterial3DTransparency, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transparency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3435651667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transparency) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTransparency() BaseMaterial3DTransparency {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transparency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990903061) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTransparency

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3212649852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2889939400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DAlphaAntiAliasing

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetAlphaAntialiasingEdge(edge float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAlphaAntialiasingEdge() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetShadingMode(shading_mode BaseMaterial3DShadingMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shading_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3368750322) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shading_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetShadingMode() BaseMaterial3DShadingMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shading_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2132070559) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DShadingMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetSpecular(specular float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&specular) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetSpecular() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetMetallic(metallic float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metallic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&metallic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetMetallic() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metallic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetRoughness(roughness float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_roughness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&roughness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRoughness() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_roughness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetEmission(emission Color, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{emission.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetEmission() Color {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetEmissionEnergyMultiplier(emission_energy_multiplier float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emission_energy_multiplier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetEmissionEnergyMultiplier() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetEmissionIntensity(emission_energy_multiplier float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emission_energy_multiplier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetEmissionIntensity() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetNormalScale(normal_scale float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetNormalScale() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetRim(rim float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rim")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rim) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRim() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rim")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetRimTint(rim_tint float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rim_tint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rim_tint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRimTint() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rim_tint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetClearcoat(clearcoat float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clearcoat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clearcoat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetClearcoat() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clearcoat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetClearcoatRoughness(clearcoat_roughness float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clearcoat_roughness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clearcoat_roughness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetClearcoatRoughness() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clearcoat_roughness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetAnisotropy(anisotropy float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anisotropy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAnisotropy() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapScale(heightmap_scale float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heightmap_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetHeightmapScale() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_heightmap_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetSubsurfaceScatteringStrength(strength float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subsurface_scattering_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetSubsurfaceScatteringStrength() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subsurface_scattering_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetTransmittanceColor(color Color, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transmittance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTransmittanceColor() Color {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transmittance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetTransmittanceDepth(depth float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transmittance_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTransmittanceDepth() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transmittance_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetTransmittanceBoost(boost float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transmittance_boost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&boost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTransmittanceBoost() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transmittance_boost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetBacklight(backlight Color, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_backlight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{backlight.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetBacklight() Color {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_backlight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetRefraction(refraction float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_refraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refraction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRefraction() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_refraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetPointSize(point_size float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetPointSize() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetDetailUv(detail_uv BaseMaterial3DDetailUV, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_detail_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 456801921) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_uv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDetailUv() BaseMaterial3DDetailUV {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_detail_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2306920512) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DDetailUV

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetBlendMode(blend_mode BaseMaterial3DBlendMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2830186259) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetBlendMode() BaseMaterial3DBlendMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4022690962) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetDepthDrawMode(depth_draw_mode BaseMaterial3DDepthDrawMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_draw_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1456584748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth_draw_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDepthDrawMode() BaseMaterial3DDepthDrawMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_draw_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2578197639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DDepthDrawMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetCullMode(cull_mode BaseMaterial3DCullMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2338909218) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetCullMode() BaseMaterial3DCullMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1941499586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DCullMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetDiffuseMode(diffuse_mode BaseMaterial3DDiffuseMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1045299638) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&diffuse_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDiffuseMode() BaseMaterial3DDiffuseMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3973617136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DDiffuseMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetSpecularMode(specular_mode BaseMaterial3DSpecularMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 584737147) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&specular_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetSpecularMode() BaseMaterial3DSpecularMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2569953298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DSpecularMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetFlag(flag BaseMaterial3DFlags, enable bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3070159527) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetFlag(flag BaseMaterial3DFlags, ) bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410065) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&flag)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 22904437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTextureFilter() BaseMaterial3DTextureFilter {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289213076) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetFeature(feature BaseMaterial3DFeature, enable bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2819288693) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetFeature(feature BaseMaterial3DFeature, ) bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965241794) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&feature)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetTexture(param BaseMaterial3DTextureParam, texture Texture2D, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 464208135) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetTexture(param BaseMaterial3DTextureParam, ) Texture2D {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 329605813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetDetailBlendMode(detail_blend_mode BaseMaterial3DBlendMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_detail_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2830186259) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_blend_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDetailBlendMode() BaseMaterial3DBlendMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_detail_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4022690962) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetUv1Scale(scale Vector3, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv1_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv1Scale() Vector3 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv1_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetUv1Offset(offset Vector3, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv1_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv1Offset() Vector3 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv1_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetUv1TriplanarBlendSharpness(sharpness float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv1_triplanar_blend_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv1TriplanarBlendSharpness() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv1_triplanar_blend_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetUv2Scale(scale Vector3, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv2_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv2Scale() Vector3 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv2_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetUv2Offset(offset Vector3, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv2_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv2Offset() Vector3 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv2_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BaseMaterial3D) SetUv2TriplanarBlendSharpness(sharpness float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv2_triplanar_blend_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetUv2TriplanarBlendSharpness() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv2_triplanar_blend_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4202036497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetBillboardMode() BaseMaterial3DBillboardMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1283840139) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DBillboardMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetParticlesAnimHFrames(frames int64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_h_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetParticlesAnimHFrames() int64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_h_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetParticlesAnimVFrames(frames int64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_v_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetParticlesAnimVFrames() int64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_v_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetParticlesAnimLoop(loop bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetParticlesAnimLoop() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallax(enable bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_deep_parallax")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) IsHeightmapDeepParallaxEnabled() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_heightmap_deep_parallax_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMinLayers(layer int64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_deep_parallax_min_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMinLayers() int64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_heightmap_deep_parallax_min_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxMaxLayers(layer int64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_deep_parallax_max_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxMaxLayers() int64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_heightmap_deep_parallax_max_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipTangent(flip bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_deep_parallax_flip_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipTangent() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_heightmap_deep_parallax_flip_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetHeightmapDeepParallaxFlipBinormal(flip bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_heightmap_deep_parallax_flip_binormal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetHeightmapDeepParallaxFlipBinormal() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_heightmap_deep_parallax_flip_binormal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetGrow(amount float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_grow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetGrow() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_grow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetEmissionOperator(operator BaseMaterial3DEmissionOperator, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3825128922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operator) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetEmissionOperator() BaseMaterial3DEmissionOperator {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 974205018) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DEmissionOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetAoLightAffect(amount float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ao_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAoLightAffect() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ao_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetAlphaScissorThreshold(threshold float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAlphaScissorThreshold() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetAlphaHashScale(threshold float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAlphaHashScale() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetGrowEnabled(enable bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_grow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) IsGrowEnabled() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_grow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetMetallicTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metallic_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 744167988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetMetallicTextureChannel() BaseMaterial3DTextureChannel {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metallic_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 568133867) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureChannel

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetRoughnessTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_roughness_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 744167988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRoughnessTextureChannel() BaseMaterial3DTextureChannel {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_roughness_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 568133867) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureChannel

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetAoTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ao_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 744167988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetAoTextureChannel() BaseMaterial3DTextureChannel {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ao_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 568133867) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureChannel

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetRefractionTextureChannel(channel BaseMaterial3DTextureChannel, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_refraction_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 744167988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetRefractionTextureChannel() BaseMaterial3DTextureChannel {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_refraction_texture_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 568133867) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DTextureChannel

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetProximityFadeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_proximity_fade_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) IsProximityFadeEnabled() bool {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_proximity_fade_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetProximityFadeDistance(distance float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_proximity_fade_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetProximityFadeDistance() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_proximity_fade_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetMsdfPixelRange(range_ float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&range_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetMsdfPixelRange() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetMsdfOutlineSize(size float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msdf_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetMsdfOutlineSize() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msdf_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetDistanceFade(mode BaseMaterial3DDistanceFadeMode, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1379478617) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDistanceFade() BaseMaterial3DDistanceFadeMode {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2694575734) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret BaseMaterial3DDistanceFadeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *BaseMaterial3D) SetDistanceFadeMaxDistance(distance float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDistanceFadeMaxDistance() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BaseMaterial3D) SetDistanceFadeMinDistance(distance float64, )  {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_min_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BaseMaterial3D) GetDistanceFadeMinDistance() float64 {
  classNameV := StringNameFromStr("BaseMaterial3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_min_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
