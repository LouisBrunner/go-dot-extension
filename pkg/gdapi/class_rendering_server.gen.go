// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RenderingServer struct {
  obj gdc.ObjectPtr
}

func (me *RenderingServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RenderingServer) BaseClass() string {
  return "RenderingServer"
}

const (
  RenderingServerNO_INDEX_ARRAY = -1
  RenderingServerARRAY_WEIGHTS_SIZE = 4
  RenderingServerCANVAS_ITEM_Z_MIN = -4096
  RenderingServerCANVAS_ITEM_Z_MAX = 4096
  RenderingServerMAX_GLOW_LEVELS = 7
  RenderingServerMAX_CURSORS = 8
  RenderingServerMAX_2D_DIRECTIONAL_LIGHTS = 8
  RenderingServerMATERIAL_RENDER_PRIORITY_MIN = -128
  RenderingServerMATERIAL_RENDER_PRIORITY_MAX = 127
  RenderingServerARRAY_CUSTOM_COUNT = 4
  RenderingServerPARTICLES_EMIT_FLAG_POSITION = 1
  RenderingServerPARTICLES_EMIT_FLAG_ROTATION_SCALE = 2
  RenderingServerPARTICLES_EMIT_FLAG_VELOCITY = 4
  RenderingServerPARTICLES_EMIT_FLAG_COLOR = 8
  RenderingServerPARTICLES_EMIT_FLAG_CUSTOM = 16
)

type RenderingServerTextureLayeredType int
const (
  RenderingServerTextureLayered2DArray RenderingServerTextureLayeredType = 0
  RenderingServerTextureLayeredCubemap RenderingServerTextureLayeredType = 1
  RenderingServerTextureLayeredCubemapArray RenderingServerTextureLayeredType = 2
)

type RenderingServerCubeMapLayer int
const (
  RenderingServerCubemapLayerLeft RenderingServerCubeMapLayer = 0
  RenderingServerCubemapLayerRight RenderingServerCubeMapLayer = 1
  RenderingServerCubemapLayerBottom RenderingServerCubeMapLayer = 2
  RenderingServerCubemapLayerTop RenderingServerCubeMapLayer = 3
  RenderingServerCubemapLayerFront RenderingServerCubeMapLayer = 4
  RenderingServerCubemapLayerBack RenderingServerCubeMapLayer = 5
)

type RenderingServerShaderMode int
const (
  RenderingServerShaderSpatial RenderingServerShaderMode = 0
  RenderingServerShaderCanvasItem RenderingServerShaderMode = 1
  RenderingServerShaderParticles RenderingServerShaderMode = 2
  RenderingServerShaderSky RenderingServerShaderMode = 3
  RenderingServerShaderFog RenderingServerShaderMode = 4
  RenderingServerShaderMax RenderingServerShaderMode = 5
)

type RenderingServerArrayType int
const (
  RenderingServerArrayVertex RenderingServerArrayType = 0
  RenderingServerArrayNormal RenderingServerArrayType = 1
  RenderingServerArrayTangent RenderingServerArrayType = 2
  RenderingServerArrayColor RenderingServerArrayType = 3
  RenderingServerArrayTexUv RenderingServerArrayType = 4
  RenderingServerArrayTexUv2 RenderingServerArrayType = 5
  RenderingServerArrayCustom0 RenderingServerArrayType = 6
  RenderingServerArrayCustom1 RenderingServerArrayType = 7
  RenderingServerArrayCustom2 RenderingServerArrayType = 8
  RenderingServerArrayCustom3 RenderingServerArrayType = 9
  RenderingServerArrayBones RenderingServerArrayType = 10
  RenderingServerArrayWeights RenderingServerArrayType = 11
  RenderingServerArrayIndex RenderingServerArrayType = 12
  RenderingServerArrayMax RenderingServerArrayType = 13
)

type RenderingServerArrayCustomFormat int
const (
  RenderingServerArrayCustomRgba8Unorm RenderingServerArrayCustomFormat = 0
  RenderingServerArrayCustomRgba8Snorm RenderingServerArrayCustomFormat = 1
  RenderingServerArrayCustomRgHalf RenderingServerArrayCustomFormat = 2
  RenderingServerArrayCustomRgbaHalf RenderingServerArrayCustomFormat = 3
  RenderingServerArrayCustomRFloat RenderingServerArrayCustomFormat = 4
  RenderingServerArrayCustomRgFloat RenderingServerArrayCustomFormat = 5
  RenderingServerArrayCustomRgbFloat RenderingServerArrayCustomFormat = 6
  RenderingServerArrayCustomRgbaFloat RenderingServerArrayCustomFormat = 7
  RenderingServerArrayCustomMax RenderingServerArrayCustomFormat = 8
)

type RenderingServerArrayFormat int
const (
  RenderingServerArrayFormatVertex RenderingServerArrayFormat = 1
  RenderingServerArrayFormatNormal RenderingServerArrayFormat = 2
  RenderingServerArrayFormatTangent RenderingServerArrayFormat = 4
  RenderingServerArrayFormatColor RenderingServerArrayFormat = 8
  RenderingServerArrayFormatTexUv RenderingServerArrayFormat = 16
  RenderingServerArrayFormatTexUv2 RenderingServerArrayFormat = 32
  RenderingServerArrayFormatCustom0 RenderingServerArrayFormat = 64
  RenderingServerArrayFormatCustom1 RenderingServerArrayFormat = 128
  RenderingServerArrayFormatCustom2 RenderingServerArrayFormat = 256
  RenderingServerArrayFormatCustom3 RenderingServerArrayFormat = 512
  RenderingServerArrayFormatBones RenderingServerArrayFormat = 1024
  RenderingServerArrayFormatWeights RenderingServerArrayFormat = 2048
  RenderingServerArrayFormatIndex RenderingServerArrayFormat = 4096
  RenderingServerArrayFormatBlendShapeMask RenderingServerArrayFormat = 7
  RenderingServerArrayFormatCustomBase RenderingServerArrayFormat = 13
  RenderingServerArrayFormatCustomBits RenderingServerArrayFormat = 3
  RenderingServerArrayFormatCustom0Shift RenderingServerArrayFormat = 13
  RenderingServerArrayFormatCustom1Shift RenderingServerArrayFormat = 16
  RenderingServerArrayFormatCustom2Shift RenderingServerArrayFormat = 19
  RenderingServerArrayFormatCustom3Shift RenderingServerArrayFormat = 22
  RenderingServerArrayFormatCustomMask RenderingServerArrayFormat = 7
  RenderingServerArrayCompressFlagsBase RenderingServerArrayFormat = 25
  RenderingServerArrayFlagUse2DVertices RenderingServerArrayFormat = 33554432
  RenderingServerArrayFlagUseDynamicUpdate RenderingServerArrayFormat = 67108864
  RenderingServerArrayFlagUse8BoneWeights RenderingServerArrayFormat = 134217728
  RenderingServerArrayFlagUsesEmptyVertexArray RenderingServerArrayFormat = 268435456
)

type RenderingServerPrimitiveType int
const (
  RenderingServerPrimitivePoints RenderingServerPrimitiveType = 0
  RenderingServerPrimitiveLines RenderingServerPrimitiveType = 1
  RenderingServerPrimitiveLineStrip RenderingServerPrimitiveType = 2
  RenderingServerPrimitiveTriangles RenderingServerPrimitiveType = 3
  RenderingServerPrimitiveTriangleStrip RenderingServerPrimitiveType = 4
  RenderingServerPrimitiveMax RenderingServerPrimitiveType = 5
)

type RenderingServerBlendShapeMode int
const (
  RenderingServerBlendShapeModeNormalized RenderingServerBlendShapeMode = 0
  RenderingServerBlendShapeModeRelative RenderingServerBlendShapeMode = 1
)

type RenderingServerMultimeshTransformFormat int
const (
  RenderingServerMultimeshTransform2D RenderingServerMultimeshTransformFormat = 0
  RenderingServerMultimeshTransform3D RenderingServerMultimeshTransformFormat = 1
)

type RenderingServerLightProjectorFilter int
const (
  RenderingServerLightProjectorFilterNearest RenderingServerLightProjectorFilter = 0
  RenderingServerLightProjectorFilterLinear RenderingServerLightProjectorFilter = 1
  RenderingServerLightProjectorFilterNearestMipmaps RenderingServerLightProjectorFilter = 2
  RenderingServerLightProjectorFilterLinearMipmaps RenderingServerLightProjectorFilter = 3
  RenderingServerLightProjectorFilterNearestMipmapsAnisotropic RenderingServerLightProjectorFilter = 4
  RenderingServerLightProjectorFilterLinearMipmapsAnisotropic RenderingServerLightProjectorFilter = 5
)

type RenderingServerLightType int
const (
  RenderingServerLightDirectional RenderingServerLightType = 0
  RenderingServerLightOmni RenderingServerLightType = 1
  RenderingServerLightSpot RenderingServerLightType = 2
)

type RenderingServerLightParam int
const (
  RenderingServerLightParamEnergy RenderingServerLightParam = 0
  RenderingServerLightParamIndirectEnergy RenderingServerLightParam = 1
  RenderingServerLightParamVolumetricFogEnergy RenderingServerLightParam = 2
  RenderingServerLightParamSpecular RenderingServerLightParam = 3
  RenderingServerLightParamRange RenderingServerLightParam = 4
  RenderingServerLightParamSize RenderingServerLightParam = 5
  RenderingServerLightParamAttenuation RenderingServerLightParam = 6
  RenderingServerLightParamSpotAngle RenderingServerLightParam = 7
  RenderingServerLightParamSpotAttenuation RenderingServerLightParam = 8
  RenderingServerLightParamShadowMaxDistance RenderingServerLightParam = 9
  RenderingServerLightParamShadowSplit1Offset RenderingServerLightParam = 10
  RenderingServerLightParamShadowSplit2Offset RenderingServerLightParam = 11
  RenderingServerLightParamShadowSplit3Offset RenderingServerLightParam = 12
  RenderingServerLightParamShadowFadeStart RenderingServerLightParam = 13
  RenderingServerLightParamShadowNormalBias RenderingServerLightParam = 14
  RenderingServerLightParamShadowBias RenderingServerLightParam = 15
  RenderingServerLightParamShadowPancakeSize RenderingServerLightParam = 16
  RenderingServerLightParamShadowOpacity RenderingServerLightParam = 17
  RenderingServerLightParamShadowBlur RenderingServerLightParam = 18
  RenderingServerLightParamTransmittanceBias RenderingServerLightParam = 19
  RenderingServerLightParamIntensity RenderingServerLightParam = 20
  RenderingServerLightParamMax RenderingServerLightParam = 21
)

type RenderingServerLightBakeMode int
const (
  RenderingServerLightBakeDisabled RenderingServerLightBakeMode = 0
  RenderingServerLightBakeStatic RenderingServerLightBakeMode = 1
  RenderingServerLightBakeDynamic RenderingServerLightBakeMode = 2
)

type RenderingServerLightOmniShadowMode int
const (
  RenderingServerLightOmniShadowDualParaboloid RenderingServerLightOmniShadowMode = 0
  RenderingServerLightOmniShadowCube RenderingServerLightOmniShadowMode = 1
)

type RenderingServerLightDirectionalShadowMode int
const (
  RenderingServerLightDirectionalShadowOrthogonal RenderingServerLightDirectionalShadowMode = 0
  RenderingServerLightDirectionalShadowParallel2Splits RenderingServerLightDirectionalShadowMode = 1
  RenderingServerLightDirectionalShadowParallel4Splits RenderingServerLightDirectionalShadowMode = 2
)

type RenderingServerLightDirectionalSkyMode int
const (
  RenderingServerLightDirectionalSkyModeLightAndSky RenderingServerLightDirectionalSkyMode = 0
  RenderingServerLightDirectionalSkyModeLightOnly RenderingServerLightDirectionalSkyMode = 1
  RenderingServerLightDirectionalSkyModeSkyOnly RenderingServerLightDirectionalSkyMode = 2
)

type RenderingServerShadowQuality int
const (
  RenderingServerShadowQualityHard RenderingServerShadowQuality = 0
  RenderingServerShadowQualitySoftVeryLow RenderingServerShadowQuality = 1
  RenderingServerShadowQualitySoftLow RenderingServerShadowQuality = 2
  RenderingServerShadowQualitySoftMedium RenderingServerShadowQuality = 3
  RenderingServerShadowQualitySoftHigh RenderingServerShadowQuality = 4
  RenderingServerShadowQualitySoftUltra RenderingServerShadowQuality = 5
  RenderingServerShadowQualityMax RenderingServerShadowQuality = 6
)

type RenderingServerReflectionProbeUpdateMode int
const (
  RenderingServerReflectionProbeUpdateOnce RenderingServerReflectionProbeUpdateMode = 0
  RenderingServerReflectionProbeUpdateAlways RenderingServerReflectionProbeUpdateMode = 1
)

type RenderingServerReflectionProbeAmbientMode int
const (
  RenderingServerReflectionProbeAmbientDisabled RenderingServerReflectionProbeAmbientMode = 0
  RenderingServerReflectionProbeAmbientEnvironment RenderingServerReflectionProbeAmbientMode = 1
  RenderingServerReflectionProbeAmbientColor RenderingServerReflectionProbeAmbientMode = 2
)

type RenderingServerDecalTexture int
const (
  RenderingServerDecalTextureAlbedo RenderingServerDecalTexture = 0
  RenderingServerDecalTextureNormal RenderingServerDecalTexture = 1
  RenderingServerDecalTextureOrm RenderingServerDecalTexture = 2
  RenderingServerDecalTextureEmission RenderingServerDecalTexture = 3
  RenderingServerDecalTextureMax RenderingServerDecalTexture = 4
)

type RenderingServerDecalFilter int
const (
  RenderingServerDecalFilterNearest RenderingServerDecalFilter = 0
  RenderingServerDecalFilterLinear RenderingServerDecalFilter = 1
  RenderingServerDecalFilterNearestMipmaps RenderingServerDecalFilter = 2
  RenderingServerDecalFilterLinearMipmaps RenderingServerDecalFilter = 3
  RenderingServerDecalFilterNearestMipmapsAnisotropic RenderingServerDecalFilter = 4
  RenderingServerDecalFilterLinearMipmapsAnisotropic RenderingServerDecalFilter = 5
)

type RenderingServerVoxelGIQuality int
const (
  RenderingServerVoxelGiQualityLow RenderingServerVoxelGIQuality = 0
  RenderingServerVoxelGiQualityHigh RenderingServerVoxelGIQuality = 1
)

type RenderingServerParticlesMode int
const (
  RenderingServerParticlesMode2D RenderingServerParticlesMode = 0
  RenderingServerParticlesMode3D RenderingServerParticlesMode = 1
)

type RenderingServerParticlesTransformAlign int
const (
  RenderingServerParticlesTransformAlignDisabled RenderingServerParticlesTransformAlign = 0
  RenderingServerParticlesTransformAlignZBillboard RenderingServerParticlesTransformAlign = 1
  RenderingServerParticlesTransformAlignYToVelocity RenderingServerParticlesTransformAlign = 2
  RenderingServerParticlesTransformAlignZBillboardYToVelocity RenderingServerParticlesTransformAlign = 3
)

type RenderingServerParticlesDrawOrder int
const (
  RenderingServerParticlesDrawOrderIndex RenderingServerParticlesDrawOrder = 0
  RenderingServerParticlesDrawOrderLifetime RenderingServerParticlesDrawOrder = 1
  RenderingServerParticlesDrawOrderReverseLifetime RenderingServerParticlesDrawOrder = 2
  RenderingServerParticlesDrawOrderViewDepth RenderingServerParticlesDrawOrder = 3
)

type RenderingServerParticlesCollisionType int
const (
  RenderingServerParticlesCollisionTypeSphereAttract RenderingServerParticlesCollisionType = 0
  RenderingServerParticlesCollisionTypeBoxAttract RenderingServerParticlesCollisionType = 1
  RenderingServerParticlesCollisionTypeVectorFieldAttract RenderingServerParticlesCollisionType = 2
  RenderingServerParticlesCollisionTypeSphereCollide RenderingServerParticlesCollisionType = 3
  RenderingServerParticlesCollisionTypeBoxCollide RenderingServerParticlesCollisionType = 4
  RenderingServerParticlesCollisionTypeSdfCollide RenderingServerParticlesCollisionType = 5
  RenderingServerParticlesCollisionTypeHeightfieldCollide RenderingServerParticlesCollisionType = 6
)

type RenderingServerParticlesCollisionHeightfieldResolution int
const (
  RenderingServerParticlesCollisionHeightfieldResolution256 RenderingServerParticlesCollisionHeightfieldResolution = 0
  RenderingServerParticlesCollisionHeightfieldResolution512 RenderingServerParticlesCollisionHeightfieldResolution = 1
  RenderingServerParticlesCollisionHeightfieldResolution1024 RenderingServerParticlesCollisionHeightfieldResolution = 2
  RenderingServerParticlesCollisionHeightfieldResolution2048 RenderingServerParticlesCollisionHeightfieldResolution = 3
  RenderingServerParticlesCollisionHeightfieldResolution4096 RenderingServerParticlesCollisionHeightfieldResolution = 4
  RenderingServerParticlesCollisionHeightfieldResolution8192 RenderingServerParticlesCollisionHeightfieldResolution = 5
  RenderingServerParticlesCollisionHeightfieldResolutionMax RenderingServerParticlesCollisionHeightfieldResolution = 6
)

type RenderingServerFogVolumeShape int
const (
  RenderingServerFogVolumeShapeEllipsoid RenderingServerFogVolumeShape = 0
  RenderingServerFogVolumeShapeCone RenderingServerFogVolumeShape = 1
  RenderingServerFogVolumeShapeCylinder RenderingServerFogVolumeShape = 2
  RenderingServerFogVolumeShapeBox RenderingServerFogVolumeShape = 3
  RenderingServerFogVolumeShapeWorld RenderingServerFogVolumeShape = 4
  RenderingServerFogVolumeShapeMax RenderingServerFogVolumeShape = 5
)

type RenderingServerViewportScaling3DMode int
const (
  RenderingServerViewportScaling3DModeBilinear RenderingServerViewportScaling3DMode = 0
  RenderingServerViewportScaling3DModeFsr RenderingServerViewportScaling3DMode = 1
  RenderingServerViewportScaling3DModeMax RenderingServerViewportScaling3DMode = 2
)

type RenderingServerViewportUpdateMode int
const (
  RenderingServerViewportUpdateDisabled RenderingServerViewportUpdateMode = 0
  RenderingServerViewportUpdateOnce RenderingServerViewportUpdateMode = 1
  RenderingServerViewportUpdateWhenVisible RenderingServerViewportUpdateMode = 2
  RenderingServerViewportUpdateWhenParentVisible RenderingServerViewportUpdateMode = 3
  RenderingServerViewportUpdateAlways RenderingServerViewportUpdateMode = 4
)

type RenderingServerViewportClearMode int
const (
  RenderingServerViewportClearAlways RenderingServerViewportClearMode = 0
  RenderingServerViewportClearNever RenderingServerViewportClearMode = 1
  RenderingServerViewportClearOnlyNextFrame RenderingServerViewportClearMode = 2
)

type RenderingServerViewportEnvironmentMode int
const (
  RenderingServerViewportEnvironmentDisabled RenderingServerViewportEnvironmentMode = 0
  RenderingServerViewportEnvironmentEnabled RenderingServerViewportEnvironmentMode = 1
  RenderingServerViewportEnvironmentInherit RenderingServerViewportEnvironmentMode = 2
  RenderingServerViewportEnvironmentMax RenderingServerViewportEnvironmentMode = 3
)

type RenderingServerViewportSDFOversize int
const (
  RenderingServerViewportSdfOversize100Percent RenderingServerViewportSDFOversize = 0
  RenderingServerViewportSdfOversize120Percent RenderingServerViewportSDFOversize = 1
  RenderingServerViewportSdfOversize150Percent RenderingServerViewportSDFOversize = 2
  RenderingServerViewportSdfOversize200Percent RenderingServerViewportSDFOversize = 3
  RenderingServerViewportSdfOversizeMax RenderingServerViewportSDFOversize = 4
)

type RenderingServerViewportSDFScale int
const (
  RenderingServerViewportSdfScale100Percent RenderingServerViewportSDFScale = 0
  RenderingServerViewportSdfScale50Percent RenderingServerViewportSDFScale = 1
  RenderingServerViewportSdfScale25Percent RenderingServerViewportSDFScale = 2
  RenderingServerViewportSdfScaleMax RenderingServerViewportSDFScale = 3
)

type RenderingServerViewportMSAA int
const (
  RenderingServerViewportMsaaDisabled RenderingServerViewportMSAA = 0
  RenderingServerViewportMsaa2X RenderingServerViewportMSAA = 1
  RenderingServerViewportMsaa4X RenderingServerViewportMSAA = 2
  RenderingServerViewportMsaa8X RenderingServerViewportMSAA = 3
  RenderingServerViewportMsaaMax RenderingServerViewportMSAA = 4
)

type RenderingServerViewportScreenSpaceAA int
const (
  RenderingServerViewportScreenSpaceAaDisabled RenderingServerViewportScreenSpaceAA = 0
  RenderingServerViewportScreenSpaceAaFxaa RenderingServerViewportScreenSpaceAA = 1
  RenderingServerViewportScreenSpaceAaMax RenderingServerViewportScreenSpaceAA = 2
)

type RenderingServerViewportOcclusionCullingBuildQuality int
const (
  RenderingServerViewportOcclusionBuildQualityLow RenderingServerViewportOcclusionCullingBuildQuality = 0
  RenderingServerViewportOcclusionBuildQualityMedium RenderingServerViewportOcclusionCullingBuildQuality = 1
  RenderingServerViewportOcclusionBuildQualityHigh RenderingServerViewportOcclusionCullingBuildQuality = 2
)

type RenderingServerViewportRenderInfo int
const (
  RenderingServerViewportRenderInfoObjectsInFrame RenderingServerViewportRenderInfo = 0
  RenderingServerViewportRenderInfoPrimitivesInFrame RenderingServerViewportRenderInfo = 1
  RenderingServerViewportRenderInfoDrawCallsInFrame RenderingServerViewportRenderInfo = 2
  RenderingServerViewportRenderInfoMax RenderingServerViewportRenderInfo = 3
)

type RenderingServerViewportRenderInfoType int
const (
  RenderingServerViewportRenderInfoTypeVisible RenderingServerViewportRenderInfoType = 0
  RenderingServerViewportRenderInfoTypeShadow RenderingServerViewportRenderInfoType = 1
  RenderingServerViewportRenderInfoTypeMax RenderingServerViewportRenderInfoType = 2
)

type RenderingServerViewportDebugDraw int
const (
  RenderingServerViewportDebugDrawDisabled RenderingServerViewportDebugDraw = 0
  RenderingServerViewportDebugDrawUnshaded RenderingServerViewportDebugDraw = 1
  RenderingServerViewportDebugDrawLighting RenderingServerViewportDebugDraw = 2
  RenderingServerViewportDebugDrawOverdraw RenderingServerViewportDebugDraw = 3
  RenderingServerViewportDebugDrawWireframe RenderingServerViewportDebugDraw = 4
  RenderingServerViewportDebugDrawNormalBuffer RenderingServerViewportDebugDraw = 5
  RenderingServerViewportDebugDrawVoxelGiAlbedo RenderingServerViewportDebugDraw = 6
  RenderingServerViewportDebugDrawVoxelGiLighting RenderingServerViewportDebugDraw = 7
  RenderingServerViewportDebugDrawVoxelGiEmission RenderingServerViewportDebugDraw = 8
  RenderingServerViewportDebugDrawShadowAtlas RenderingServerViewportDebugDraw = 9
  RenderingServerViewportDebugDrawDirectionalShadowAtlas RenderingServerViewportDebugDraw = 10
  RenderingServerViewportDebugDrawSceneLuminance RenderingServerViewportDebugDraw = 11
  RenderingServerViewportDebugDrawSsao RenderingServerViewportDebugDraw = 12
  RenderingServerViewportDebugDrawSsil RenderingServerViewportDebugDraw = 13
  RenderingServerViewportDebugDrawPssmSplits RenderingServerViewportDebugDraw = 14
  RenderingServerViewportDebugDrawDecalAtlas RenderingServerViewportDebugDraw = 15
  RenderingServerViewportDebugDrawSdfgi RenderingServerViewportDebugDraw = 16
  RenderingServerViewportDebugDrawSdfgiProbes RenderingServerViewportDebugDraw = 17
  RenderingServerViewportDebugDrawGiBuffer RenderingServerViewportDebugDraw = 18
  RenderingServerViewportDebugDrawDisableLod RenderingServerViewportDebugDraw = 19
  RenderingServerViewportDebugDrawClusterOmniLights RenderingServerViewportDebugDraw = 20
  RenderingServerViewportDebugDrawClusterSpotLights RenderingServerViewportDebugDraw = 21
  RenderingServerViewportDebugDrawClusterDecals RenderingServerViewportDebugDraw = 22
  RenderingServerViewportDebugDrawClusterReflectionProbes RenderingServerViewportDebugDraw = 23
  RenderingServerViewportDebugDrawOccluders RenderingServerViewportDebugDraw = 24
  RenderingServerViewportDebugDrawMotionVectors RenderingServerViewportDebugDraw = 25
)

type RenderingServerViewportVRSMode int
const (
  RenderingServerViewportVrsDisabled RenderingServerViewportVRSMode = 0
  RenderingServerViewportVrsTexture RenderingServerViewportVRSMode = 1
  RenderingServerViewportVrsXr RenderingServerViewportVRSMode = 2
  RenderingServerViewportVrsMax RenderingServerViewportVRSMode = 3
)

type RenderingServerSkyMode int
const (
  RenderingServerSkyModeAutomatic RenderingServerSkyMode = 0
  RenderingServerSkyModeQuality RenderingServerSkyMode = 1
  RenderingServerSkyModeIncremental RenderingServerSkyMode = 2
  RenderingServerSkyModeRealtime RenderingServerSkyMode = 3
)

type RenderingServerEnvironmentBG int
const (
  RenderingServerEnvBgClearColor RenderingServerEnvironmentBG = 0
  RenderingServerEnvBgColor RenderingServerEnvironmentBG = 1
  RenderingServerEnvBgSky RenderingServerEnvironmentBG = 2
  RenderingServerEnvBgCanvas RenderingServerEnvironmentBG = 3
  RenderingServerEnvBgKeep RenderingServerEnvironmentBG = 4
  RenderingServerEnvBgCameraFeed RenderingServerEnvironmentBG = 5
  RenderingServerEnvBgMax RenderingServerEnvironmentBG = 6
)

type RenderingServerEnvironmentAmbientSource int
const (
  RenderingServerEnvAmbientSourceBg RenderingServerEnvironmentAmbientSource = 0
  RenderingServerEnvAmbientSourceDisabled RenderingServerEnvironmentAmbientSource = 1
  RenderingServerEnvAmbientSourceColor RenderingServerEnvironmentAmbientSource = 2
  RenderingServerEnvAmbientSourceSky RenderingServerEnvironmentAmbientSource = 3
)

type RenderingServerEnvironmentReflectionSource int
const (
  RenderingServerEnvReflectionSourceBg RenderingServerEnvironmentReflectionSource = 0
  RenderingServerEnvReflectionSourceDisabled RenderingServerEnvironmentReflectionSource = 1
  RenderingServerEnvReflectionSourceSky RenderingServerEnvironmentReflectionSource = 2
)

type RenderingServerEnvironmentGlowBlendMode int
const (
  RenderingServerEnvGlowBlendModeAdditive RenderingServerEnvironmentGlowBlendMode = 0
  RenderingServerEnvGlowBlendModeScreen RenderingServerEnvironmentGlowBlendMode = 1
  RenderingServerEnvGlowBlendModeSoftlight RenderingServerEnvironmentGlowBlendMode = 2
  RenderingServerEnvGlowBlendModeReplace RenderingServerEnvironmentGlowBlendMode = 3
  RenderingServerEnvGlowBlendModeMix RenderingServerEnvironmentGlowBlendMode = 4
)

type RenderingServerEnvironmentToneMapper int
const (
  RenderingServerEnvToneMapperLinear RenderingServerEnvironmentToneMapper = 0
  RenderingServerEnvToneMapperReinhard RenderingServerEnvironmentToneMapper = 1
  RenderingServerEnvToneMapperFilmic RenderingServerEnvironmentToneMapper = 2
  RenderingServerEnvToneMapperAces RenderingServerEnvironmentToneMapper = 3
)

type RenderingServerEnvironmentSSRRoughnessQuality int
const (
  RenderingServerEnvSsrRoughnessQualityDisabled RenderingServerEnvironmentSSRRoughnessQuality = 0
  RenderingServerEnvSsrRoughnessQualityLow RenderingServerEnvironmentSSRRoughnessQuality = 1
  RenderingServerEnvSsrRoughnessQualityMedium RenderingServerEnvironmentSSRRoughnessQuality = 2
  RenderingServerEnvSsrRoughnessQualityHigh RenderingServerEnvironmentSSRRoughnessQuality = 3
)

type RenderingServerEnvironmentSSAOQuality int
const (
  RenderingServerEnvSsaoQualityVeryLow RenderingServerEnvironmentSSAOQuality = 0
  RenderingServerEnvSsaoQualityLow RenderingServerEnvironmentSSAOQuality = 1
  RenderingServerEnvSsaoQualityMedium RenderingServerEnvironmentSSAOQuality = 2
  RenderingServerEnvSsaoQualityHigh RenderingServerEnvironmentSSAOQuality = 3
  RenderingServerEnvSsaoQualityUltra RenderingServerEnvironmentSSAOQuality = 4
)

type RenderingServerEnvironmentSSILQuality int
const (
  RenderingServerEnvSsilQualityVeryLow RenderingServerEnvironmentSSILQuality = 0
  RenderingServerEnvSsilQualityLow RenderingServerEnvironmentSSILQuality = 1
  RenderingServerEnvSsilQualityMedium RenderingServerEnvironmentSSILQuality = 2
  RenderingServerEnvSsilQualityHigh RenderingServerEnvironmentSSILQuality = 3
  RenderingServerEnvSsilQualityUltra RenderingServerEnvironmentSSILQuality = 4
)

type RenderingServerEnvironmentSDFGIYScale int
const (
  RenderingServerEnvSdfgiYScale50Percent RenderingServerEnvironmentSDFGIYScale = 0
  RenderingServerEnvSdfgiYScale75Percent RenderingServerEnvironmentSDFGIYScale = 1
  RenderingServerEnvSdfgiYScale100Percent RenderingServerEnvironmentSDFGIYScale = 2
)

type RenderingServerEnvironmentSDFGIRayCount int
const (
  RenderingServerEnvSdfgiRayCount4 RenderingServerEnvironmentSDFGIRayCount = 0
  RenderingServerEnvSdfgiRayCount8 RenderingServerEnvironmentSDFGIRayCount = 1
  RenderingServerEnvSdfgiRayCount16 RenderingServerEnvironmentSDFGIRayCount = 2
  RenderingServerEnvSdfgiRayCount32 RenderingServerEnvironmentSDFGIRayCount = 3
  RenderingServerEnvSdfgiRayCount64 RenderingServerEnvironmentSDFGIRayCount = 4
  RenderingServerEnvSdfgiRayCount96 RenderingServerEnvironmentSDFGIRayCount = 5
  RenderingServerEnvSdfgiRayCount128 RenderingServerEnvironmentSDFGIRayCount = 6
  RenderingServerEnvSdfgiRayCountMax RenderingServerEnvironmentSDFGIRayCount = 7
)

type RenderingServerEnvironmentSDFGIFramesToConverge int
const (
  RenderingServerEnvSdfgiConvergeIn5Frames RenderingServerEnvironmentSDFGIFramesToConverge = 0
  RenderingServerEnvSdfgiConvergeIn10Frames RenderingServerEnvironmentSDFGIFramesToConverge = 1
  RenderingServerEnvSdfgiConvergeIn15Frames RenderingServerEnvironmentSDFGIFramesToConverge = 2
  RenderingServerEnvSdfgiConvergeIn20Frames RenderingServerEnvironmentSDFGIFramesToConverge = 3
  RenderingServerEnvSdfgiConvergeIn25Frames RenderingServerEnvironmentSDFGIFramesToConverge = 4
  RenderingServerEnvSdfgiConvergeIn30Frames RenderingServerEnvironmentSDFGIFramesToConverge = 5
  RenderingServerEnvSdfgiConvergeMax RenderingServerEnvironmentSDFGIFramesToConverge = 6
)

type RenderingServerEnvironmentSDFGIFramesToUpdateLight int
const (
  RenderingServerEnvSdfgiUpdateLightIn1Frame RenderingServerEnvironmentSDFGIFramesToUpdateLight = 0
  RenderingServerEnvSdfgiUpdateLightIn2Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 1
  RenderingServerEnvSdfgiUpdateLightIn4Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 2
  RenderingServerEnvSdfgiUpdateLightIn8Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 3
  RenderingServerEnvSdfgiUpdateLightIn16Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 4
  RenderingServerEnvSdfgiUpdateLightMax RenderingServerEnvironmentSDFGIFramesToUpdateLight = 5
)

type RenderingServerSubSurfaceScatteringQuality int
const (
  RenderingServerSubSurfaceScatteringQualityDisabled RenderingServerSubSurfaceScatteringQuality = 0
  RenderingServerSubSurfaceScatteringQualityLow RenderingServerSubSurfaceScatteringQuality = 1
  RenderingServerSubSurfaceScatteringQualityMedium RenderingServerSubSurfaceScatteringQuality = 2
  RenderingServerSubSurfaceScatteringQualityHigh RenderingServerSubSurfaceScatteringQuality = 3
)

type RenderingServerDOFBokehShape int
const (
  RenderingServerDofBokehBox RenderingServerDOFBokehShape = 0
  RenderingServerDofBokehHexagon RenderingServerDOFBokehShape = 1
  RenderingServerDofBokehCircle RenderingServerDOFBokehShape = 2
)

type RenderingServerDOFBlurQuality int
const (
  RenderingServerDofBlurQualityVeryLow RenderingServerDOFBlurQuality = 0
  RenderingServerDofBlurQualityLow RenderingServerDOFBlurQuality = 1
  RenderingServerDofBlurQualityMedium RenderingServerDOFBlurQuality = 2
  RenderingServerDofBlurQualityHigh RenderingServerDOFBlurQuality = 3
)

type RenderingServerInstanceType int
const (
  RenderingServerInstanceNone RenderingServerInstanceType = 0
  RenderingServerInstanceMesh RenderingServerInstanceType = 1
  RenderingServerInstanceMultimesh RenderingServerInstanceType = 2
  RenderingServerInstanceParticles RenderingServerInstanceType = 3
  RenderingServerInstanceParticlesCollision RenderingServerInstanceType = 4
  RenderingServerInstanceLight RenderingServerInstanceType = 5
  RenderingServerInstanceReflectionProbe RenderingServerInstanceType = 6
  RenderingServerInstanceDecal RenderingServerInstanceType = 7
  RenderingServerInstanceVoxelGi RenderingServerInstanceType = 8
  RenderingServerInstanceLightmap RenderingServerInstanceType = 9
  RenderingServerInstanceOccluder RenderingServerInstanceType = 10
  RenderingServerInstanceVisiblityNotifier RenderingServerInstanceType = 11
  RenderingServerInstanceFogVolume RenderingServerInstanceType = 12
  RenderingServerInstanceMax RenderingServerInstanceType = 13
  RenderingServerInstanceGeometryMask RenderingServerInstanceType = 14
)

type RenderingServerInstanceFlags int
const (
  RenderingServerInstanceFlagUseBakedLight RenderingServerInstanceFlags = 0
  RenderingServerInstanceFlagUseDynamicGi RenderingServerInstanceFlags = 1
  RenderingServerInstanceFlagDrawNextFrameIfVisible RenderingServerInstanceFlags = 2
  RenderingServerInstanceFlagIgnoreOcclusionCulling RenderingServerInstanceFlags = 3
  RenderingServerInstanceFlagMax RenderingServerInstanceFlags = 4
)

type RenderingServerShadowCastingSetting int
const (
  RenderingServerShadowCastingSettingOff RenderingServerShadowCastingSetting = 0
  RenderingServerShadowCastingSettingOn RenderingServerShadowCastingSetting = 1
  RenderingServerShadowCastingSettingDoubleSided RenderingServerShadowCastingSetting = 2
  RenderingServerShadowCastingSettingShadowsOnly RenderingServerShadowCastingSetting = 3
)

type RenderingServerVisibilityRangeFadeMode int
const (
  RenderingServerVisibilityRangeFadeDisabled RenderingServerVisibilityRangeFadeMode = 0
  RenderingServerVisibilityRangeFadeSelf RenderingServerVisibilityRangeFadeMode = 1
  RenderingServerVisibilityRangeFadeDependencies RenderingServerVisibilityRangeFadeMode = 2
)

type RenderingServerBakeChannels int
const (
  RenderingServerBakeChannelAlbedoAlpha RenderingServerBakeChannels = 0
  RenderingServerBakeChannelNormal RenderingServerBakeChannels = 1
  RenderingServerBakeChannelOrm RenderingServerBakeChannels = 2
  RenderingServerBakeChannelEmission RenderingServerBakeChannels = 3
)

type RenderingServerCanvasTextureChannel int
const (
  RenderingServerCanvasTextureChannelDiffuse RenderingServerCanvasTextureChannel = 0
  RenderingServerCanvasTextureChannelNormal RenderingServerCanvasTextureChannel = 1
  RenderingServerCanvasTextureChannelSpecular RenderingServerCanvasTextureChannel = 2
)

type RenderingServerNinePatchAxisMode int
const (
  RenderingServerNinePatchStretch RenderingServerNinePatchAxisMode = 0
  RenderingServerNinePatchTile RenderingServerNinePatchAxisMode = 1
  RenderingServerNinePatchTileFit RenderingServerNinePatchAxisMode = 2
)

type RenderingServerCanvasItemTextureFilter int
const (
  RenderingServerCanvasItemTextureFilterDefault RenderingServerCanvasItemTextureFilter = 0
  RenderingServerCanvasItemTextureFilterNearest RenderingServerCanvasItemTextureFilter = 1
  RenderingServerCanvasItemTextureFilterLinear RenderingServerCanvasItemTextureFilter = 2
  RenderingServerCanvasItemTextureFilterNearestWithMipmaps RenderingServerCanvasItemTextureFilter = 3
  RenderingServerCanvasItemTextureFilterLinearWithMipmaps RenderingServerCanvasItemTextureFilter = 4
  RenderingServerCanvasItemTextureFilterNearestWithMipmapsAnisotropic RenderingServerCanvasItemTextureFilter = 5
  RenderingServerCanvasItemTextureFilterLinearWithMipmapsAnisotropic RenderingServerCanvasItemTextureFilter = 6
  RenderingServerCanvasItemTextureFilterMax RenderingServerCanvasItemTextureFilter = 7
)

type RenderingServerCanvasItemTextureRepeat int
const (
  RenderingServerCanvasItemTextureRepeatDefault RenderingServerCanvasItemTextureRepeat = 0
  RenderingServerCanvasItemTextureRepeatDisabled RenderingServerCanvasItemTextureRepeat = 1
  RenderingServerCanvasItemTextureRepeatEnabled RenderingServerCanvasItemTextureRepeat = 2
  RenderingServerCanvasItemTextureRepeatMirror RenderingServerCanvasItemTextureRepeat = 3
  RenderingServerCanvasItemTextureRepeatMax RenderingServerCanvasItemTextureRepeat = 4
)

type RenderingServerCanvasGroupMode int
const (
  RenderingServerCanvasGroupModeDisabled RenderingServerCanvasGroupMode = 0
  RenderingServerCanvasGroupModeClipOnly RenderingServerCanvasGroupMode = 1
  RenderingServerCanvasGroupModeClipAndDraw RenderingServerCanvasGroupMode = 2
  RenderingServerCanvasGroupModeTransparent RenderingServerCanvasGroupMode = 3
)

type RenderingServerCanvasLightMode int
const (
  RenderingServerCanvasLightModePoint RenderingServerCanvasLightMode = 0
  RenderingServerCanvasLightModeDirectional RenderingServerCanvasLightMode = 1
)

type RenderingServerCanvasLightBlendMode int
const (
  RenderingServerCanvasLightBlendModeAdd RenderingServerCanvasLightBlendMode = 0
  RenderingServerCanvasLightBlendModeSub RenderingServerCanvasLightBlendMode = 1
  RenderingServerCanvasLightBlendModeMix RenderingServerCanvasLightBlendMode = 2
)

type RenderingServerCanvasLightShadowFilter int
const (
  RenderingServerCanvasLightFilterNone RenderingServerCanvasLightShadowFilter = 0
  RenderingServerCanvasLightFilterPcf5 RenderingServerCanvasLightShadowFilter = 1
  RenderingServerCanvasLightFilterPcf13 RenderingServerCanvasLightShadowFilter = 2
  RenderingServerCanvasLightFilterMax RenderingServerCanvasLightShadowFilter = 3
)

type RenderingServerCanvasOccluderPolygonCullMode int
const (
  RenderingServerCanvasOccluderPolygonCullDisabled RenderingServerCanvasOccluderPolygonCullMode = 0
  RenderingServerCanvasOccluderPolygonCullClockwise RenderingServerCanvasOccluderPolygonCullMode = 1
  RenderingServerCanvasOccluderPolygonCullCounterClockwise RenderingServerCanvasOccluderPolygonCullMode = 2
)

type RenderingServerGlobalShaderParameterType int
const (
  RenderingServerGlobalVarTypeBool RenderingServerGlobalShaderParameterType = 0
  RenderingServerGlobalVarTypeBvec2 RenderingServerGlobalShaderParameterType = 1
  RenderingServerGlobalVarTypeBvec3 RenderingServerGlobalShaderParameterType = 2
  RenderingServerGlobalVarTypeBvec4 RenderingServerGlobalShaderParameterType = 3
  RenderingServerGlobalVarTypeInt RenderingServerGlobalShaderParameterType = 4
  RenderingServerGlobalVarTypeIvec2 RenderingServerGlobalShaderParameterType = 5
  RenderingServerGlobalVarTypeIvec3 RenderingServerGlobalShaderParameterType = 6
  RenderingServerGlobalVarTypeIvec4 RenderingServerGlobalShaderParameterType = 7
  RenderingServerGlobalVarTypeRect2I RenderingServerGlobalShaderParameterType = 8
  RenderingServerGlobalVarTypeUint RenderingServerGlobalShaderParameterType = 9
  RenderingServerGlobalVarTypeUvec2 RenderingServerGlobalShaderParameterType = 10
  RenderingServerGlobalVarTypeUvec3 RenderingServerGlobalShaderParameterType = 11
  RenderingServerGlobalVarTypeUvec4 RenderingServerGlobalShaderParameterType = 12
  RenderingServerGlobalVarTypeFloat RenderingServerGlobalShaderParameterType = 13
  RenderingServerGlobalVarTypeVec2 RenderingServerGlobalShaderParameterType = 14
  RenderingServerGlobalVarTypeVec3 RenderingServerGlobalShaderParameterType = 15
  RenderingServerGlobalVarTypeVec4 RenderingServerGlobalShaderParameterType = 16
  RenderingServerGlobalVarTypeColor RenderingServerGlobalShaderParameterType = 17
  RenderingServerGlobalVarTypeRect2 RenderingServerGlobalShaderParameterType = 18
  RenderingServerGlobalVarTypeMat2 RenderingServerGlobalShaderParameterType = 19
  RenderingServerGlobalVarTypeMat3 RenderingServerGlobalShaderParameterType = 20
  RenderingServerGlobalVarTypeMat4 RenderingServerGlobalShaderParameterType = 21
  RenderingServerGlobalVarTypeTransform2D RenderingServerGlobalShaderParameterType = 22
  RenderingServerGlobalVarTypeTransform RenderingServerGlobalShaderParameterType = 23
  RenderingServerGlobalVarTypeSampler2D RenderingServerGlobalShaderParameterType = 24
  RenderingServerGlobalVarTypeSampler2Darray RenderingServerGlobalShaderParameterType = 25
  RenderingServerGlobalVarTypeSampler3D RenderingServerGlobalShaderParameterType = 26
  RenderingServerGlobalVarTypeSamplercube RenderingServerGlobalShaderParameterType = 27
  RenderingServerGlobalVarTypeMax RenderingServerGlobalShaderParameterType = 28
)

type RenderingServerRenderingInfo int
const (
  RenderingServerRenderingInfoTotalObjectsInFrame RenderingServerRenderingInfo = 0
  RenderingServerRenderingInfoTotalPrimitivesInFrame RenderingServerRenderingInfo = 1
  RenderingServerRenderingInfoTotalDrawCallsInFrame RenderingServerRenderingInfo = 2
  RenderingServerRenderingInfoTextureMemUsed RenderingServerRenderingInfo = 3
  RenderingServerRenderingInfoBufferMemUsed RenderingServerRenderingInfo = 4
  RenderingServerRenderingInfoVideoMemUsed RenderingServerRenderingInfo = 5
)

type RenderingServerFeatures int
const (
  RenderingServerFeatureShaders RenderingServerFeatures = 0
  RenderingServerFeatureMultithreaded RenderingServerFeatures = 1
)
