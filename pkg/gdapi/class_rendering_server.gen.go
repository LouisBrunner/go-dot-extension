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

type ptrsForRenderingServerList struct {
  fnTexture2DCreate gdc.MethodBindPtr
  fnTexture2DLayeredCreate gdc.MethodBindPtr
  fnTexture3DCreate gdc.MethodBindPtr
  fnTextureProxyCreate gdc.MethodBindPtr
  fnTexture2DUpdate gdc.MethodBindPtr
  fnTexture3DUpdate gdc.MethodBindPtr
  fnTextureProxyUpdate gdc.MethodBindPtr
  fnTexture2DPlaceholderCreate gdc.MethodBindPtr
  fnTexture2DLayeredPlaceholderCreate gdc.MethodBindPtr
  fnTexture3DPlaceholderCreate gdc.MethodBindPtr
  fnTexture2DGet gdc.MethodBindPtr
  fnTexture2DLayerGet gdc.MethodBindPtr
  fnTexture3DGet gdc.MethodBindPtr
  fnTextureReplace gdc.MethodBindPtr
  fnTextureSetSizeOverride gdc.MethodBindPtr
  fnTextureSetPath gdc.MethodBindPtr
  fnTextureGetPath gdc.MethodBindPtr
  fnTextureGetFormat gdc.MethodBindPtr
  fnTextureSetForceRedrawIfVisible gdc.MethodBindPtr
  fnTextureRdCreate gdc.MethodBindPtr
  fnTextureGetRdTexture gdc.MethodBindPtr
  fnTextureGetNativeHandle gdc.MethodBindPtr
  fnShaderCreate gdc.MethodBindPtr
  fnShaderSetCode gdc.MethodBindPtr
  fnShaderSetPathHint gdc.MethodBindPtr
  fnShaderGetCode gdc.MethodBindPtr
  fnGetShaderParameterList gdc.MethodBindPtr
  fnShaderGetParameterDefault gdc.MethodBindPtr
  fnShaderSetDefaultTextureParameter gdc.MethodBindPtr
  fnShaderGetDefaultTextureParameter gdc.MethodBindPtr
  fnMaterialCreate gdc.MethodBindPtr
  fnMaterialSetShader gdc.MethodBindPtr
  fnMaterialSetParam gdc.MethodBindPtr
  fnMaterialGetParam gdc.MethodBindPtr
  fnMaterialSetRenderPriority gdc.MethodBindPtr
  fnMaterialSetNextPass gdc.MethodBindPtr
  fnMeshCreateFromSurfaces gdc.MethodBindPtr
  fnMeshCreate gdc.MethodBindPtr
  fnMeshSurfaceGetFormatOffset gdc.MethodBindPtr
  fnMeshSurfaceGetFormatVertexStride gdc.MethodBindPtr
  fnMeshSurfaceGetFormatNormalTangentStride gdc.MethodBindPtr
  fnMeshSurfaceGetFormatAttributeStride gdc.MethodBindPtr
  fnMeshSurfaceGetFormatSkinStride gdc.MethodBindPtr
  fnMeshAddSurface gdc.MethodBindPtr
  fnMeshAddSurfaceFromArrays gdc.MethodBindPtr
  fnMeshGetBlendShapeCount gdc.MethodBindPtr
  fnMeshSetBlendShapeMode gdc.MethodBindPtr
  fnMeshGetBlendShapeMode gdc.MethodBindPtr
  fnMeshSurfaceSetMaterial gdc.MethodBindPtr
  fnMeshSurfaceGetMaterial gdc.MethodBindPtr
  fnMeshGetSurface gdc.MethodBindPtr
  fnMeshSurfaceGetArrays gdc.MethodBindPtr
  fnMeshSurfaceGetBlendShapeArrays gdc.MethodBindPtr
  fnMeshGetSurfaceCount gdc.MethodBindPtr
  fnMeshSetCustomAabb gdc.MethodBindPtr
  fnMeshGetCustomAabb gdc.MethodBindPtr
  fnMeshClear gdc.MethodBindPtr
  fnMeshSurfaceUpdateVertexRegion gdc.MethodBindPtr
  fnMeshSurfaceUpdateAttributeRegion gdc.MethodBindPtr
  fnMeshSurfaceUpdateSkinRegion gdc.MethodBindPtr
  fnMeshSetShadowMesh gdc.MethodBindPtr
  fnMultimeshCreate gdc.MethodBindPtr
  fnMultimeshAllocateData gdc.MethodBindPtr
  fnMultimeshGetInstanceCount gdc.MethodBindPtr
  fnMultimeshSetMesh gdc.MethodBindPtr
  fnMultimeshInstanceSetTransform gdc.MethodBindPtr
  fnMultimeshInstanceSetTransform2D gdc.MethodBindPtr
  fnMultimeshInstanceSetColor gdc.MethodBindPtr
  fnMultimeshInstanceSetCustomData gdc.MethodBindPtr
  fnMultimeshGetMesh gdc.MethodBindPtr
  fnMultimeshGetAabb gdc.MethodBindPtr
  fnMultimeshInstanceGetTransform gdc.MethodBindPtr
  fnMultimeshInstanceGetTransform2D gdc.MethodBindPtr
  fnMultimeshInstanceGetColor gdc.MethodBindPtr
  fnMultimeshInstanceGetCustomData gdc.MethodBindPtr
  fnMultimeshSetVisibleInstances gdc.MethodBindPtr
  fnMultimeshGetVisibleInstances gdc.MethodBindPtr
  fnMultimeshSetBuffer gdc.MethodBindPtr
  fnMultimeshGetBuffer gdc.MethodBindPtr
  fnSkeletonCreate gdc.MethodBindPtr
  fnSkeletonAllocateData gdc.MethodBindPtr
  fnSkeletonGetBoneCount gdc.MethodBindPtr
  fnSkeletonBoneSetTransform gdc.MethodBindPtr
  fnSkeletonBoneGetTransform gdc.MethodBindPtr
  fnSkeletonBoneSetTransform2D gdc.MethodBindPtr
  fnSkeletonBoneGetTransform2D gdc.MethodBindPtr
  fnSkeletonSetBaseTransform2D gdc.MethodBindPtr
  fnDirectionalLightCreate gdc.MethodBindPtr
  fnOmniLightCreate gdc.MethodBindPtr
  fnSpotLightCreate gdc.MethodBindPtr
  fnLightSetColor gdc.MethodBindPtr
  fnLightSetParam gdc.MethodBindPtr
  fnLightSetShadow gdc.MethodBindPtr
  fnLightSetProjector gdc.MethodBindPtr
  fnLightSetNegative gdc.MethodBindPtr
  fnLightSetCullMask gdc.MethodBindPtr
  fnLightSetDistanceFade gdc.MethodBindPtr
  fnLightSetReverseCullFaceMode gdc.MethodBindPtr
  fnLightSetBakeMode gdc.MethodBindPtr
  fnLightSetMaxSdfgiCascade gdc.MethodBindPtr
  fnLightOmniSetShadowMode gdc.MethodBindPtr
  fnLightDirectionalSetShadowMode gdc.MethodBindPtr
  fnLightDirectionalSetBlendSplits gdc.MethodBindPtr
  fnLightDirectionalSetSkyMode gdc.MethodBindPtr
  fnLightProjectorsSetFilter gdc.MethodBindPtr
  fnPositionalSoftShadowFilterSetQuality gdc.MethodBindPtr
  fnDirectionalSoftShadowFilterSetQuality gdc.MethodBindPtr
  fnDirectionalShadowAtlasSetSize gdc.MethodBindPtr
  fnReflectionProbeCreate gdc.MethodBindPtr
  fnReflectionProbeSetUpdateMode gdc.MethodBindPtr
  fnReflectionProbeSetIntensity gdc.MethodBindPtr
  fnReflectionProbeSetAmbientMode gdc.MethodBindPtr
  fnReflectionProbeSetAmbientColor gdc.MethodBindPtr
  fnReflectionProbeSetAmbientEnergy gdc.MethodBindPtr
  fnReflectionProbeSetMaxDistance gdc.MethodBindPtr
  fnReflectionProbeSetSize gdc.MethodBindPtr
  fnReflectionProbeSetOriginOffset gdc.MethodBindPtr
  fnReflectionProbeSetAsInterior gdc.MethodBindPtr
  fnReflectionProbeSetEnableBoxProjection gdc.MethodBindPtr
  fnReflectionProbeSetEnableShadows gdc.MethodBindPtr
  fnReflectionProbeSetCullMask gdc.MethodBindPtr
  fnReflectionProbeSetResolution gdc.MethodBindPtr
  fnReflectionProbeSetMeshLodThreshold gdc.MethodBindPtr
  fnDecalCreate gdc.MethodBindPtr
  fnDecalSetSize gdc.MethodBindPtr
  fnDecalSetTexture gdc.MethodBindPtr
  fnDecalSetEmissionEnergy gdc.MethodBindPtr
  fnDecalSetAlbedoMix gdc.MethodBindPtr
  fnDecalSetModulate gdc.MethodBindPtr
  fnDecalSetCullMask gdc.MethodBindPtr
  fnDecalSetDistanceFade gdc.MethodBindPtr
  fnDecalSetFade gdc.MethodBindPtr
  fnDecalSetNormalFade gdc.MethodBindPtr
  fnDecalsSetFilter gdc.MethodBindPtr
  fnGiSetUseHalfResolution gdc.MethodBindPtr
  fnVoxelGiCreate gdc.MethodBindPtr
  fnVoxelGiAllocateData gdc.MethodBindPtr
  fnVoxelGiGetOctreeSize gdc.MethodBindPtr
  fnVoxelGiGetOctreeCells gdc.MethodBindPtr
  fnVoxelGiGetDataCells gdc.MethodBindPtr
  fnVoxelGiGetDistanceField gdc.MethodBindPtr
  fnVoxelGiGetLevelCounts gdc.MethodBindPtr
  fnVoxelGiGetToCellXform gdc.MethodBindPtr
  fnVoxelGiSetDynamicRange gdc.MethodBindPtr
  fnVoxelGiSetPropagation gdc.MethodBindPtr
  fnVoxelGiSetEnergy gdc.MethodBindPtr
  fnVoxelGiSetBakedExposureNormalization gdc.MethodBindPtr
  fnVoxelGiSetBias gdc.MethodBindPtr
  fnVoxelGiSetNormalBias gdc.MethodBindPtr
  fnVoxelGiSetInterior gdc.MethodBindPtr
  fnVoxelGiSetUseTwoBounces gdc.MethodBindPtr
  fnVoxelGiSetQuality gdc.MethodBindPtr
  fnLightmapCreate gdc.MethodBindPtr
  fnLightmapSetTextures gdc.MethodBindPtr
  fnLightmapSetProbeBounds gdc.MethodBindPtr
  fnLightmapSetProbeInterior gdc.MethodBindPtr
  fnLightmapSetProbeCaptureData gdc.MethodBindPtr
  fnLightmapGetProbeCapturePoints gdc.MethodBindPtr
  fnLightmapGetProbeCaptureSh gdc.MethodBindPtr
  fnLightmapGetProbeCaptureTetrahedra gdc.MethodBindPtr
  fnLightmapGetProbeCaptureBspTree gdc.MethodBindPtr
  fnLightmapSetBakedExposureNormalization gdc.MethodBindPtr
  fnLightmapSetProbeCaptureUpdateSpeed gdc.MethodBindPtr
  fnParticlesCreate gdc.MethodBindPtr
  fnParticlesSetMode gdc.MethodBindPtr
  fnParticlesSetEmitting gdc.MethodBindPtr
  fnParticlesGetEmitting gdc.MethodBindPtr
  fnParticlesSetAmount gdc.MethodBindPtr
  fnParticlesSetAmountRatio gdc.MethodBindPtr
  fnParticlesSetLifetime gdc.MethodBindPtr
  fnParticlesSetOneShot gdc.MethodBindPtr
  fnParticlesSetPreProcessTime gdc.MethodBindPtr
  fnParticlesSetExplosivenessRatio gdc.MethodBindPtr
  fnParticlesSetRandomnessRatio gdc.MethodBindPtr
  fnParticlesSetInterpToEnd gdc.MethodBindPtr
  fnParticlesSetEmitterVelocity gdc.MethodBindPtr
  fnParticlesSetCustomAabb gdc.MethodBindPtr
  fnParticlesSetSpeedScale gdc.MethodBindPtr
  fnParticlesSetUseLocalCoordinates gdc.MethodBindPtr
  fnParticlesSetProcessMaterial gdc.MethodBindPtr
  fnParticlesSetFixedFps gdc.MethodBindPtr
  fnParticlesSetInterpolate gdc.MethodBindPtr
  fnParticlesSetFractionalDelta gdc.MethodBindPtr
  fnParticlesSetCollisionBaseSize gdc.MethodBindPtr
  fnParticlesSetTransformAlign gdc.MethodBindPtr
  fnParticlesSetTrails gdc.MethodBindPtr
  fnParticlesSetTrailBindPoses gdc.MethodBindPtr
  fnParticlesIsInactive gdc.MethodBindPtr
  fnParticlesRequestProcess gdc.MethodBindPtr
  fnParticlesRestart gdc.MethodBindPtr
  fnParticlesSetSubemitter gdc.MethodBindPtr
  fnParticlesEmit gdc.MethodBindPtr
  fnParticlesSetDrawOrder gdc.MethodBindPtr
  fnParticlesSetDrawPasses gdc.MethodBindPtr
  fnParticlesSetDrawPassMesh gdc.MethodBindPtr
  fnParticlesGetCurrentAabb gdc.MethodBindPtr
  fnParticlesSetEmissionTransform gdc.MethodBindPtr
  fnParticlesCollisionCreate gdc.MethodBindPtr
  fnParticlesCollisionSetCollisionType gdc.MethodBindPtr
  fnParticlesCollisionSetCullMask gdc.MethodBindPtr
  fnParticlesCollisionSetSphereRadius gdc.MethodBindPtr
  fnParticlesCollisionSetBoxExtents gdc.MethodBindPtr
  fnParticlesCollisionSetAttractorStrength gdc.MethodBindPtr
  fnParticlesCollisionSetAttractorDirectionality gdc.MethodBindPtr
  fnParticlesCollisionSetAttractorAttenuation gdc.MethodBindPtr
  fnParticlesCollisionSetFieldTexture gdc.MethodBindPtr
  fnParticlesCollisionHeightFieldUpdate gdc.MethodBindPtr
  fnParticlesCollisionSetHeightFieldResolution gdc.MethodBindPtr
  fnFogVolumeCreate gdc.MethodBindPtr
  fnFogVolumeSetShape gdc.MethodBindPtr
  fnFogVolumeSetSize gdc.MethodBindPtr
  fnFogVolumeSetMaterial gdc.MethodBindPtr
  fnVisibilityNotifierCreate gdc.MethodBindPtr
  fnVisibilityNotifierSetAabb gdc.MethodBindPtr
  fnVisibilityNotifierSetCallbacks gdc.MethodBindPtr
  fnOccluderCreate gdc.MethodBindPtr
  fnOccluderSetMesh gdc.MethodBindPtr
  fnCameraCreate gdc.MethodBindPtr
  fnCameraSetPerspective gdc.MethodBindPtr
  fnCameraSetOrthogonal gdc.MethodBindPtr
  fnCameraSetFrustum gdc.MethodBindPtr
  fnCameraSetTransform gdc.MethodBindPtr
  fnCameraSetCullMask gdc.MethodBindPtr
  fnCameraSetEnvironment gdc.MethodBindPtr
  fnCameraSetCameraAttributes gdc.MethodBindPtr
  fnCameraSetUseVerticalAspect gdc.MethodBindPtr
  fnViewportCreate gdc.MethodBindPtr
  fnViewportSetUseXr gdc.MethodBindPtr
  fnViewportSetSize gdc.MethodBindPtr
  fnViewportSetActive gdc.MethodBindPtr
  fnViewportSetParentViewport gdc.MethodBindPtr
  fnViewportAttachToScreen gdc.MethodBindPtr
  fnViewportSetRenderDirectToScreen gdc.MethodBindPtr
  fnViewportSetCanvasCullMask gdc.MethodBindPtr
  fnViewportSetScaling3DMode gdc.MethodBindPtr
  fnViewportSetScaling3DScale gdc.MethodBindPtr
  fnViewportSetFsrSharpness gdc.MethodBindPtr
  fnViewportSetTextureMipmapBias gdc.MethodBindPtr
  fnViewportSetUpdateMode gdc.MethodBindPtr
  fnViewportSetClearMode gdc.MethodBindPtr
  fnViewportGetRenderTarget gdc.MethodBindPtr
  fnViewportGetTexture gdc.MethodBindPtr
  fnViewportSetDisable3D gdc.MethodBindPtr
  fnViewportSetDisable2D gdc.MethodBindPtr
  fnViewportSetEnvironmentMode gdc.MethodBindPtr
  fnViewportAttachCamera gdc.MethodBindPtr
  fnViewportSetScenario gdc.MethodBindPtr
  fnViewportAttachCanvas gdc.MethodBindPtr
  fnViewportRemoveCanvas gdc.MethodBindPtr
  fnViewportSetSnap2DTransformsToPixel gdc.MethodBindPtr
  fnViewportSetSnap2DVerticesToPixel gdc.MethodBindPtr
  fnViewportSetDefaultCanvasItemTextureFilter gdc.MethodBindPtr
  fnViewportSetDefaultCanvasItemTextureRepeat gdc.MethodBindPtr
  fnViewportSetCanvasTransform gdc.MethodBindPtr
  fnViewportSetCanvasStacking gdc.MethodBindPtr
  fnViewportSetTransparentBackground gdc.MethodBindPtr
  fnViewportSetGlobalCanvasTransform gdc.MethodBindPtr
  fnViewportSetSdfOversizeAndScale gdc.MethodBindPtr
  fnViewportSetPositionalShadowAtlasSize gdc.MethodBindPtr
  fnViewportSetPositionalShadowAtlasQuadrantSubdivision gdc.MethodBindPtr
  fnViewportSetMsaa3D gdc.MethodBindPtr
  fnViewportSetMsaa2D gdc.MethodBindPtr
  fnViewportSetUseHdr2D gdc.MethodBindPtr
  fnViewportSetScreenSpaceAa gdc.MethodBindPtr
  fnViewportSetUseTaa gdc.MethodBindPtr
  fnViewportSetUseDebanding gdc.MethodBindPtr
  fnViewportSetUseOcclusionCulling gdc.MethodBindPtr
  fnViewportSetOcclusionRaysPerThread gdc.MethodBindPtr
  fnViewportSetOcclusionCullingBuildQuality gdc.MethodBindPtr
  fnViewportGetRenderInfo gdc.MethodBindPtr
  fnViewportSetDebugDraw gdc.MethodBindPtr
  fnViewportSetMeasureRenderTime gdc.MethodBindPtr
  fnViewportGetMeasuredRenderTimeCpu gdc.MethodBindPtr
  fnViewportGetMeasuredRenderTimeGpu gdc.MethodBindPtr
  fnViewportSetVrsMode gdc.MethodBindPtr
  fnViewportSetVrsTexture gdc.MethodBindPtr
  fnSkyCreate gdc.MethodBindPtr
  fnSkySetRadianceSize gdc.MethodBindPtr
  fnSkySetMode gdc.MethodBindPtr
  fnSkySetMaterial gdc.MethodBindPtr
  fnSkyBakePanorama gdc.MethodBindPtr
  fnEnvironmentCreate gdc.MethodBindPtr
  fnEnvironmentSetBackground gdc.MethodBindPtr
  fnEnvironmentSetSky gdc.MethodBindPtr
  fnEnvironmentSetSkyCustomFov gdc.MethodBindPtr
  fnEnvironmentSetSkyOrientation gdc.MethodBindPtr
  fnEnvironmentSetBgColor gdc.MethodBindPtr
  fnEnvironmentSetBgEnergy gdc.MethodBindPtr
  fnEnvironmentSetCanvasMaxLayer gdc.MethodBindPtr
  fnEnvironmentSetAmbientLight gdc.MethodBindPtr
  fnEnvironmentSetGlow gdc.MethodBindPtr
  fnEnvironmentSetTonemap gdc.MethodBindPtr
  fnEnvironmentSetAdjustment gdc.MethodBindPtr
  fnEnvironmentSetSsr gdc.MethodBindPtr
  fnEnvironmentSetSsao gdc.MethodBindPtr
  fnEnvironmentSetFog gdc.MethodBindPtr
  fnEnvironmentSetSdfgi gdc.MethodBindPtr
  fnEnvironmentSetVolumetricFog gdc.MethodBindPtr
  fnEnvironmentGlowSetUseBicubicUpscale gdc.MethodBindPtr
  fnEnvironmentSetSsrRoughnessQuality gdc.MethodBindPtr
  fnEnvironmentSetSsaoQuality gdc.MethodBindPtr
  fnEnvironmentSetSsilQuality gdc.MethodBindPtr
  fnEnvironmentSetSdfgiRayCount gdc.MethodBindPtr
  fnEnvironmentSetSdfgiFramesToConverge gdc.MethodBindPtr
  fnEnvironmentSetSdfgiFramesToUpdateLight gdc.MethodBindPtr
  fnEnvironmentSetVolumetricFogVolumeSize gdc.MethodBindPtr
  fnEnvironmentSetVolumetricFogFilterActive gdc.MethodBindPtr
  fnEnvironmentBakePanorama gdc.MethodBindPtr
  fnScreenSpaceRoughnessLimiterSetActive gdc.MethodBindPtr
  fnSubSurfaceScatteringSetQuality gdc.MethodBindPtr
  fnSubSurfaceScatteringSetScale gdc.MethodBindPtr
  fnCameraAttributesCreate gdc.MethodBindPtr
  fnCameraAttributesSetDofBlurQuality gdc.MethodBindPtr
  fnCameraAttributesSetDofBlurBokehShape gdc.MethodBindPtr
  fnCameraAttributesSetDofBlur gdc.MethodBindPtr
  fnCameraAttributesSetExposure gdc.MethodBindPtr
  fnCameraAttributesSetAutoExposure gdc.MethodBindPtr
  fnScenarioCreate gdc.MethodBindPtr
  fnScenarioSetEnvironment gdc.MethodBindPtr
  fnScenarioSetFallbackEnvironment gdc.MethodBindPtr
  fnScenarioSetCameraAttributes gdc.MethodBindPtr
  fnInstanceCreate2 gdc.MethodBindPtr
  fnInstanceCreate gdc.MethodBindPtr
  fnInstanceSetBase gdc.MethodBindPtr
  fnInstanceSetScenario gdc.MethodBindPtr
  fnInstanceSetLayerMask gdc.MethodBindPtr
  fnInstanceSetPivotData gdc.MethodBindPtr
  fnInstanceSetTransform gdc.MethodBindPtr
  fnInstanceAttachObjectInstanceId gdc.MethodBindPtr
  fnInstanceSetBlendShapeWeight gdc.MethodBindPtr
  fnInstanceSetSurfaceOverrideMaterial gdc.MethodBindPtr
  fnInstanceSetVisible gdc.MethodBindPtr
  fnInstanceGeometrySetTransparency gdc.MethodBindPtr
  fnInstanceSetCustomAabb gdc.MethodBindPtr
  fnInstanceAttachSkeleton gdc.MethodBindPtr
  fnInstanceSetExtraVisibilityMargin gdc.MethodBindPtr
  fnInstanceSetVisibilityParent gdc.MethodBindPtr
  fnInstanceSetIgnoreCulling gdc.MethodBindPtr
  fnInstanceGeometrySetFlag gdc.MethodBindPtr
  fnInstanceGeometrySetCastShadowsSetting gdc.MethodBindPtr
  fnInstanceGeometrySetMaterialOverride gdc.MethodBindPtr
  fnInstanceGeometrySetMaterialOverlay gdc.MethodBindPtr
  fnInstanceGeometrySetVisibilityRange gdc.MethodBindPtr
  fnInstanceGeometrySetLightmap gdc.MethodBindPtr
  fnInstanceGeometrySetLodBias gdc.MethodBindPtr
  fnInstanceGeometrySetShaderParameter gdc.MethodBindPtr
  fnInstanceGeometryGetShaderParameter gdc.MethodBindPtr
  fnInstanceGeometryGetShaderParameterDefaultValue gdc.MethodBindPtr
  fnInstanceGeometryGetShaderParameterList gdc.MethodBindPtr
  fnInstancesCullAabb gdc.MethodBindPtr
  fnInstancesCullRay gdc.MethodBindPtr
  fnInstancesCullConvex gdc.MethodBindPtr
  fnBakeRenderUv2 gdc.MethodBindPtr
  fnCanvasCreate gdc.MethodBindPtr
  fnCanvasSetItemMirroring gdc.MethodBindPtr
  fnCanvasSetModulate gdc.MethodBindPtr
  fnCanvasSetDisableScale gdc.MethodBindPtr
  fnCanvasTextureCreate gdc.MethodBindPtr
  fnCanvasTextureSetChannel gdc.MethodBindPtr
  fnCanvasTextureSetShadingParameters gdc.MethodBindPtr
  fnCanvasTextureSetTextureFilter gdc.MethodBindPtr
  fnCanvasTextureSetTextureRepeat gdc.MethodBindPtr
  fnCanvasItemCreate gdc.MethodBindPtr
  fnCanvasItemSetParent gdc.MethodBindPtr
  fnCanvasItemSetDefaultTextureFilter gdc.MethodBindPtr
  fnCanvasItemSetDefaultTextureRepeat gdc.MethodBindPtr
  fnCanvasItemSetVisible gdc.MethodBindPtr
  fnCanvasItemSetLightMask gdc.MethodBindPtr
  fnCanvasItemSetVisibilityLayer gdc.MethodBindPtr
  fnCanvasItemSetTransform gdc.MethodBindPtr
  fnCanvasItemSetClip gdc.MethodBindPtr
  fnCanvasItemSetDistanceFieldMode gdc.MethodBindPtr
  fnCanvasItemSetCustomRect gdc.MethodBindPtr
  fnCanvasItemSetModulate gdc.MethodBindPtr
  fnCanvasItemSetSelfModulate gdc.MethodBindPtr
  fnCanvasItemSetDrawBehindParent gdc.MethodBindPtr
  fnCanvasItemAddLine gdc.MethodBindPtr
  fnCanvasItemAddPolyline gdc.MethodBindPtr
  fnCanvasItemAddMultiline gdc.MethodBindPtr
  fnCanvasItemAddRect gdc.MethodBindPtr
  fnCanvasItemAddCircle gdc.MethodBindPtr
  fnCanvasItemAddTextureRect gdc.MethodBindPtr
  fnCanvasItemAddMsdfTextureRectRegion gdc.MethodBindPtr
  fnCanvasItemAddLcdTextureRectRegion gdc.MethodBindPtr
  fnCanvasItemAddTextureRectRegion gdc.MethodBindPtr
  fnCanvasItemAddNinePatch gdc.MethodBindPtr
  fnCanvasItemAddPrimitive gdc.MethodBindPtr
  fnCanvasItemAddPolygon gdc.MethodBindPtr
  fnCanvasItemAddTriangleArray gdc.MethodBindPtr
  fnCanvasItemAddMesh gdc.MethodBindPtr
  fnCanvasItemAddMultimesh gdc.MethodBindPtr
  fnCanvasItemAddParticles gdc.MethodBindPtr
  fnCanvasItemAddSetTransform gdc.MethodBindPtr
  fnCanvasItemAddClipIgnore gdc.MethodBindPtr
  fnCanvasItemAddAnimationSlice gdc.MethodBindPtr
  fnCanvasItemSetSortChildrenByY gdc.MethodBindPtr
  fnCanvasItemSetZIndex gdc.MethodBindPtr
  fnCanvasItemSetZAsRelativeToParent gdc.MethodBindPtr
  fnCanvasItemSetCopyToBackbuffer gdc.MethodBindPtr
  fnCanvasItemClear gdc.MethodBindPtr
  fnCanvasItemSetDrawIndex gdc.MethodBindPtr
  fnCanvasItemSetMaterial gdc.MethodBindPtr
  fnCanvasItemSetUseParentMaterial gdc.MethodBindPtr
  fnCanvasItemSetVisibilityNotifier gdc.MethodBindPtr
  fnCanvasItemSetCanvasGroupMode gdc.MethodBindPtr
  fnCanvasLightCreate gdc.MethodBindPtr
  fnCanvasLightAttachToCanvas gdc.MethodBindPtr
  fnCanvasLightSetEnabled gdc.MethodBindPtr
  fnCanvasLightSetTextureScale gdc.MethodBindPtr
  fnCanvasLightSetTransform gdc.MethodBindPtr
  fnCanvasLightSetTexture gdc.MethodBindPtr
  fnCanvasLightSetTextureOffset gdc.MethodBindPtr
  fnCanvasLightSetColor gdc.MethodBindPtr
  fnCanvasLightSetHeight gdc.MethodBindPtr
  fnCanvasLightSetEnergy gdc.MethodBindPtr
  fnCanvasLightSetZRange gdc.MethodBindPtr
  fnCanvasLightSetLayerRange gdc.MethodBindPtr
  fnCanvasLightSetItemCullMask gdc.MethodBindPtr
  fnCanvasLightSetItemShadowCullMask gdc.MethodBindPtr
  fnCanvasLightSetMode gdc.MethodBindPtr
  fnCanvasLightSetShadowEnabled gdc.MethodBindPtr
  fnCanvasLightSetShadowFilter gdc.MethodBindPtr
  fnCanvasLightSetShadowColor gdc.MethodBindPtr
  fnCanvasLightSetShadowSmooth gdc.MethodBindPtr
  fnCanvasLightSetBlendMode gdc.MethodBindPtr
  fnCanvasLightOccluderCreate gdc.MethodBindPtr
  fnCanvasLightOccluderAttachToCanvas gdc.MethodBindPtr
  fnCanvasLightOccluderSetEnabled gdc.MethodBindPtr
  fnCanvasLightOccluderSetPolygon gdc.MethodBindPtr
  fnCanvasLightOccluderSetAsSdfCollision gdc.MethodBindPtr
  fnCanvasLightOccluderSetTransform gdc.MethodBindPtr
  fnCanvasLightOccluderSetLightMask gdc.MethodBindPtr
  fnCanvasOccluderPolygonCreate gdc.MethodBindPtr
  fnCanvasOccluderPolygonSetShape gdc.MethodBindPtr
  fnCanvasOccluderPolygonSetCullMode gdc.MethodBindPtr
  fnCanvasSetShadowTextureSize gdc.MethodBindPtr
  fnGlobalShaderParameterAdd gdc.MethodBindPtr
  fnGlobalShaderParameterRemove gdc.MethodBindPtr
  fnGlobalShaderParameterGetList gdc.MethodBindPtr
  fnGlobalShaderParameterSet gdc.MethodBindPtr
  fnGlobalShaderParameterSetOverride gdc.MethodBindPtr
  fnGlobalShaderParameterGet gdc.MethodBindPtr
  fnGlobalShaderParameterGetType gdc.MethodBindPtr
  fnFreeRid gdc.MethodBindPtr
  fnRequestFrameDrawnCallback gdc.MethodBindPtr
  fnHasChanged gdc.MethodBindPtr
  fnGetRenderingInfo gdc.MethodBindPtr
  fnGetVideoAdapterName gdc.MethodBindPtr
  fnGetVideoAdapterVendor gdc.MethodBindPtr
  fnGetVideoAdapterType gdc.MethodBindPtr
  fnGetVideoAdapterApiVersion gdc.MethodBindPtr
  fnMakeSphereMesh gdc.MethodBindPtr
  fnGetTestCube gdc.MethodBindPtr
  fnGetTestTexture gdc.MethodBindPtr
  fnGetWhiteTexture gdc.MethodBindPtr
  fnSetBootImage gdc.MethodBindPtr
  fnGetDefaultClearColor gdc.MethodBindPtr
  fnSetDefaultClearColor gdc.MethodBindPtr
  fnHasFeature gdc.MethodBindPtr
  fnHasOsFeature gdc.MethodBindPtr
  fnSetDebugGenerateWireframes gdc.MethodBindPtr
  fnIsRenderLoopEnabled gdc.MethodBindPtr
  fnSetRenderLoopEnabled gdc.MethodBindPtr
  fnGetFrameSetupTimeCpu gdc.MethodBindPtr
  fnForceSync gdc.MethodBindPtr
  fnForceDraw gdc.MethodBindPtr
  fnGetRenderingDevice gdc.MethodBindPtr
  fnCreateLocalRenderingDevice gdc.MethodBindPtr
  fnCallOnRenderThread gdc.MethodBindPtr
}

var ptrsForRenderingServer ptrsForRenderingServerList

func initRenderingServerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RenderingServer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("texture_2d_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2010018390))
  }
  {
    methodName := StringNameFromStr("texture_2d_layered_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DLayeredCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 913689023))
  }
  {
    methodName := StringNameFromStr("texture_3d_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture3DCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4036838706))
  }
  {
    methodName := StringNameFromStr("texture_proxy_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureProxyCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 41030802))
  }
  {
    methodName := StringNameFromStr("texture_2d_update")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 999539803))
  }
  {
    methodName := StringNameFromStr("texture_3d_update")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture3DUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 684822712))
  }
  {
    methodName := StringNameFromStr("texture_proxy_update")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureProxyUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("texture_2d_placeholder_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DPlaceholderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("texture_2d_layered_placeholder_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DLayeredPlaceholderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1394585590))
  }
  {
    methodName := StringNameFromStr("texture_3d_placeholder_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture3DPlaceholderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("texture_2d_get")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4206205781))
  }
  {
    methodName := StringNameFromStr("texture_2d_layer_get")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture2DLayerGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2705440895))
  }
  {
    methodName := StringNameFromStr("texture_3d_get")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTexture3DGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("texture_replace")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureReplace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("texture_set_size_override")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureSetSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
  }
  {
    methodName := StringNameFromStr("texture_set_path")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureSetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
  }
  {
    methodName := StringNameFromStr("texture_get_path")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
  }
  {
    methodName := StringNameFromStr("texture_get_format")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1932918979))
  }
  {
    methodName := StringNameFromStr("texture_set_force_redraw_if_visible")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureSetForceRedrawIfVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("texture_rd_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureRdCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1434128712))
  }
  {
    methodName := StringNameFromStr("texture_get_rd_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureGetRdTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2790148051))
  }
  {
    methodName := StringNameFromStr("texture_get_native_handle")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnTextureGetNativeHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1834114100))
  }
  {
    methodName := StringNameFromStr("shader_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("shader_set_code")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderSetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
  }
  {
    methodName := StringNameFromStr("shader_set_path_hint")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderSetPathHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
  }
  {
    methodName := StringNameFromStr("shader_get_code")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderGetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
  }
  {
    methodName := StringNameFromStr("get_shader_parameter_list")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetShaderParameterList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("shader_get_parameter_default")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderGetParameterDefault = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2621281810))
  }
  {
    methodName := StringNameFromStr("shader_set_default_texture_parameter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderSetDefaultTextureParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4094001817))
  }
  {
    methodName := StringNameFromStr("shader_get_default_texture_parameter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnShaderGetDefaultTextureParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1464608890))
  }
  {
    methodName := StringNameFromStr("material_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("material_set_shader")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialSetShader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("material_set_param")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3477296213))
  }
  {
    methodName := StringNameFromStr("material_get_param")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2621281810))
  }
  {
    methodName := StringNameFromStr("material_set_render_priority")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialSetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("material_set_next_pass")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMaterialSetNextPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("mesh_create_from_surfaces")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshCreateFromSurfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291747531))
  }
  {
    methodName := StringNameFromStr("mesh_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_format_offset")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetFormatOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981368685))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_format_vertex_stride")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetFormatVertexStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3188363337))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_format_normal_tangent_stride")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetFormatNormalTangentStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3188363337))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_format_attribute_stride")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetFormatAttributeStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3188363337))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_format_skin_stride")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetFormatSkinStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3188363337))
  }
  {
    methodName := StringNameFromStr("mesh_add_surface")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshAddSurface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1217542888))
  }
  {
    methodName := StringNameFromStr("mesh_add_surface_from_arrays")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshAddSurfaceFromArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2342446560))
  }
  {
    methodName := StringNameFromStr("mesh_get_blend_shape_count")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshGetBlendShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("mesh_set_blend_shape_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1294662092))
  }
  {
    methodName := StringNameFromStr("mesh_get_blend_shape_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshGetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4282291819))
  }
  {
    methodName := StringNameFromStr("mesh_surface_set_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
  }
  {
    methodName := StringNameFromStr("mesh_get_surface")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshGetSurface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 186674697))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_arrays")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1778388067))
  }
  {
    methodName := StringNameFromStr("mesh_surface_get_blend_shape_arrays")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceGetBlendShapeArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1778388067))
  }
  {
    methodName := StringNameFromStr("mesh_get_surface_count")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshGetSurfaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("mesh_set_custom_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3696536120))
  }
  {
    methodName := StringNameFromStr("mesh_get_custom_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshGetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974181306))
  }
  {
    methodName := StringNameFromStr("mesh_clear")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("mesh_surface_update_vertex_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceUpdateVertexRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2900195149))
  }
  {
    methodName := StringNameFromStr("mesh_surface_update_attribute_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceUpdateAttributeRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2900195149))
  }
  {
    methodName := StringNameFromStr("mesh_surface_update_skin_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSurfaceUpdateSkinRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2900195149))
  }
  {
    methodName := StringNameFromStr("mesh_set_shadow_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMeshSetShadowMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("multimesh_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("multimesh_allocate_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshAllocateData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 283685892))
  }
  {
    methodName := StringNameFromStr("multimesh_get_instance_count")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshGetInstanceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("multimesh_set_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675327471))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_set_transform_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceSetTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 736082694))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_set_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 176975443))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_set_custom_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceSetCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 176975443))
  }
  {
    methodName := StringNameFromStr("multimesh_get_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("multimesh_get_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshGetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974181306))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_get_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1050775521))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_get_transform_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceGetTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1324854622))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_get_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2946315076))
  }
  {
    methodName := StringNameFromStr("multimesh_instance_get_custom_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshInstanceGetCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2946315076))
  }
  {
    methodName := StringNameFromStr("multimesh_set_visible_instances")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshSetVisibleInstances = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("multimesh_get_visible_instances")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshGetVisibleInstances = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("multimesh_set_buffer")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshSetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2960552364))
  }
  {
    methodName := StringNameFromStr("multimesh_get_buffer")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMultimeshGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3964669176))
  }
  {
    methodName := StringNameFromStr("skeleton_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("skeleton_allocate_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonAllocateData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1904426712))
  }
  {
    methodName := StringNameFromStr("skeleton_get_bone_count")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonGetBoneCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("skeleton_bone_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonBoneSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675327471))
  }
  {
    methodName := StringNameFromStr("skeleton_bone_get_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonBoneGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1050775521))
  }
  {
    methodName := StringNameFromStr("skeleton_bone_set_transform_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonBoneSetTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 736082694))
  }
  {
    methodName := StringNameFromStr("skeleton_bone_get_transform_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonBoneGetTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1324854622))
  }
  {
    methodName := StringNameFromStr("skeleton_set_base_transform_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkeletonSetBaseTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("directional_light_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDirectionalLightCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("omni_light_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnOmniLightCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("spot_light_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSpotLightCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("light_set_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("light_set_param")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501936875))
  }
  {
    methodName := StringNameFromStr("light_set_shadow")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetShadow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("light_set_projector")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetProjector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("light_set_negative")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetNegative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("light_set_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("light_set_distance_fade")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1622292572))
  }
  {
    methodName := StringNameFromStr("light_set_reverse_cull_face_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetReverseCullFaceMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("light_set_bake_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetBakeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1048525260))
  }
  {
    methodName := StringNameFromStr("light_set_max_sdfgi_cascade")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightSetMaxSdfgiCascade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("light_omni_set_shadow_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightOmniSetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2552677200))
  }
  {
    methodName := StringNameFromStr("light_directional_set_shadow_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightDirectionalSetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 380462970))
  }
  {
    methodName := StringNameFromStr("light_directional_set_blend_splits")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightDirectionalSetBlendSplits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("light_directional_set_sky_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightDirectionalSetSkyMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2559740754))
  }
  {
    methodName := StringNameFromStr("light_projectors_set_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightProjectorsSetFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 43944325))
  }
  {
    methodName := StringNameFromStr("positional_soft_shadow_filter_set_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnPositionalSoftShadowFilterSetQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3613045266))
  }
  {
    methodName := StringNameFromStr("directional_soft_shadow_filter_set_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDirectionalSoftShadowFilterSetQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3613045266))
  }
  {
    methodName := StringNameFromStr("directional_shadow_atlas_set_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDirectionalShadowAtlasSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("reflection_probe_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_update_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3853670147))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_intensity")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_ambient_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetAmbientMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 184163074))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_ambient_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetAmbientColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_ambient_energy")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetAmbientEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_max_distance")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_origin_offset")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetOriginOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_as_interior")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetAsInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_enable_box_projection")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetEnableBoxProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_enable_shadows")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetEnableShadows = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_resolution")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("reflection_probe_set_mesh_lod_threshold")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnReflectionProbeSetMeshLodThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("decal_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("decal_set_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("decal_set_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3953344054))
  }
  {
    methodName := StringNameFromStr("decal_set_emission_energy")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetEmissionEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("decal_set_albedo_mix")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetAlbedoMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("decal_set_modulate")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("decal_set_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("decal_set_distance_fade")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2972769666))
  }
  {
    methodName := StringNameFromStr("decal_set_fade")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2513314492))
  }
  {
    methodName := StringNameFromStr("decal_set_normal_fade")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalSetNormalFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("decals_set_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnDecalsSetFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3519875702))
  }
  {
    methodName := StringNameFromStr("gi_set_use_half_resolution")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGiSetUseHalfResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("voxel_gi_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("voxel_gi_allocate_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiAllocateData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108223027))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_octree_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetOctreeSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2607699645))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_octree_cells")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetOctreeCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3348040486))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_data_cells")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetDataCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3348040486))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_distance_field")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3348040486))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_level_counts")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetLevelCounts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788230395))
  }
  {
    methodName := StringNameFromStr("voxel_gi_get_to_cell_xform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiGetToCellXform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1128465797))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_dynamic_range")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetDynamicRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_propagation")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetPropagation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_energy")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_baked_exposure_normalization")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetBakedExposureNormalization = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_bias")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_normal_bias")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetNormalBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_interior")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_use_two_bounces")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetUseTwoBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("voxel_gi_set_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVoxelGiSetQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1538689978))
  }
  {
    methodName := StringNameFromStr("lightmap_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("lightmap_set_textures")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2646464759))
  }
  {
    methodName := StringNameFromStr("lightmap_set_probe_bounds")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetProbeBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3696536120))
  }
  {
    methodName := StringNameFromStr("lightmap_set_probe_interior")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetProbeInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("lightmap_set_probe_capture_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetProbeCaptureData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3217845880))
  }
  {
    methodName := StringNameFromStr("lightmap_get_probe_capture_points")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapGetProbeCapturePoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 808965560))
  }
  {
    methodName := StringNameFromStr("lightmap_get_probe_capture_sh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapGetProbeCaptureSh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1569415609))
  }
  {
    methodName := StringNameFromStr("lightmap_get_probe_capture_tetrahedra")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapGetProbeCaptureTetrahedra = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788230395))
  }
  {
    methodName := StringNameFromStr("lightmap_get_probe_capture_bsp_tree")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapGetProbeCaptureBspTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788230395))
  }
  {
    methodName := StringNameFromStr("lightmap_set_baked_exposure_normalization")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetBakedExposureNormalization = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("lightmap_set_probe_capture_update_speed")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnLightmapSetProbeCaptureUpdateSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("particles_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("particles_set_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3492270028))
  }
  {
    methodName := StringNameFromStr("particles_set_emitting")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("particles_get_emitting")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesGetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
  }
  {
    methodName := StringNameFromStr("particles_set_amount")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("particles_set_amount_ratio")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetAmountRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_lifetime")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_one_shot")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("particles_set_pre_process_time")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_explosiveness_ratio")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_randomness_ratio")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_interp_to_end")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetInterpToEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_emitter_velocity")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetEmitterVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("particles_set_custom_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3696536120))
  }
  {
    methodName := StringNameFromStr("particles_set_speed_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_use_local_coordinates")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("particles_set_process_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetProcessMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("particles_set_fixed_fps")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("particles_set_interpolate")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("particles_set_fractional_delta")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("particles_set_collision_base_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetCollisionBaseSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_set_transform_align")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetTransformAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3264971368))
  }
  {
    methodName := StringNameFromStr("particles_set_trails")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetTrails = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2010054925))
  }
  {
    methodName := StringNameFromStr("particles_set_trail_bind_poses")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetTrailBindPoses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 684822712))
  }
  {
    methodName := StringNameFromStr("particles_is_inactive")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesIsInactive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
  }
  {
    methodName := StringNameFromStr("particles_request_process")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesRequestProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("particles_restart")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesRestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("particles_set_subemitter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetSubemitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("particles_emit")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesEmit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4043136117))
  }
  {
    methodName := StringNameFromStr("particles_set_draw_order")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 935028487))
  }
  {
    methodName := StringNameFromStr("particles_set_draw_passes")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetDrawPasses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("particles_set_draw_pass_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetDrawPassMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
  }
  {
    methodName := StringNameFromStr("particles_get_current_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesGetCurrentAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3952830260))
  }
  {
    methodName := StringNameFromStr("particles_set_emission_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesSetEmissionTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935195649))
  }
  {
    methodName := StringNameFromStr("particles_collision_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_collision_type")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetCollisionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497044930))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_sphere_radius")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetSphereRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_box_extents")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetBoxExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_attractor_strength")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetAttractorStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_attractor_directionality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetAttractorDirectionality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_attractor_attenuation")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetAttractorAttenuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_field_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetFieldTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("particles_collision_height_field_update")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionHeightFieldUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("particles_collision_set_height_field_resolution")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnParticlesCollisionSetHeightFieldResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 962977297))
  }
  {
    methodName := StringNameFromStr("fog_volume_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnFogVolumeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("fog_volume_set_shape")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnFogVolumeSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3818703106))
  }
  {
    methodName := StringNameFromStr("fog_volume_set_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnFogVolumeSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
  }
  {
    methodName := StringNameFromStr("fog_volume_set_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnFogVolumeSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("visibility_notifier_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVisibilityNotifierCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("visibility_notifier_set_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVisibilityNotifierSetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3696536120))
  }
  {
    methodName := StringNameFromStr("visibility_notifier_set_callbacks")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnVisibilityNotifierSetCallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2689735388))
  }
  {
    methodName := StringNameFromStr("occluder_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnOccluderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("occluder_set_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnOccluderSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3854404263))
  }
  {
    methodName := StringNameFromStr("camera_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("camera_set_perspective")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 157498339))
  }
  {
    methodName := StringNameFromStr("camera_set_orthogonal")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetOrthogonal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 157498339))
  }
  {
    methodName := StringNameFromStr("camera_set_frustum")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetFrustum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1889878953))
  }
  {
    methodName := StringNameFromStr("camera_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935195649))
  }
  {
    methodName := StringNameFromStr("camera_set_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("camera_set_environment")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("camera_set_camera_attributes")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("camera_set_use_vertical_aspect")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraSetUseVerticalAspect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("viewport_set_use_xr")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUseXr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
  }
  {
    methodName := StringNameFromStr("viewport_set_active")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_parent_viewport")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetParentViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("viewport_attach_to_screen")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportAttachToScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1062245816))
  }
  {
    methodName := StringNameFromStr("viewport_set_render_direct_to_screen")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetRenderDirectToScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_canvas_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetCanvasCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("viewport_set_scaling_3d_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetScaling3DMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2386524376))
  }
  {
    methodName := StringNameFromStr("viewport_set_scaling_3d_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetScaling3DScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("viewport_set_fsr_sharpness")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetFsrSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("viewport_set_texture_mipmap_bias")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetTextureMipmapBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("viewport_set_update_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3161116010))
  }
  {
    methodName := StringNameFromStr("viewport_set_clear_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetClearMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3628367896))
  }
  {
    methodName := StringNameFromStr("viewport_get_render_target")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportGetRenderTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("viewport_get_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("viewport_set_disable_3d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetDisable3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_disable_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetDisable2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_environment_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetEnvironmentMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2196892182))
  }
  {
    methodName := StringNameFromStr("viewport_attach_camera")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportAttachCamera = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("viewport_set_scenario")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetScenario = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("viewport_attach_canvas")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportAttachCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("viewport_remove_canvas")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportRemoveCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("viewport_set_snap_2d_transforms_to_pixel")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetSnap2DTransformsToPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_snap_2d_vertices_to_pixel")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetSnap2DVerticesToPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_default_canvas_item_texture_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetDefaultCanvasItemTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1155129294))
  }
  {
    methodName := StringNameFromStr("viewport_set_default_canvas_item_texture_repeat")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetDefaultCanvasItemTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1652956681))
  }
  {
    methodName := StringNameFromStr("viewport_set_canvas_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3608606053))
  }
  {
    methodName := StringNameFromStr("viewport_set_canvas_stacking")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetCanvasStacking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3713930247))
  }
  {
    methodName := StringNameFromStr("viewport_set_transparent_background")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetTransparentBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_global_canvas_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetGlobalCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("viewport_set_sdf_oversize_and_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetSdfOversizeAndScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1329198632))
  }
  {
    methodName := StringNameFromStr("viewport_set_positional_shadow_atlas_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetPositionalShadowAtlasSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1904426712))
  }
  {
    methodName := StringNameFromStr("viewport_set_positional_shadow_atlas_quadrant_subdivision")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetPositionalShadowAtlasQuadrantSubdivision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
  }
  {
    methodName := StringNameFromStr("viewport_set_msaa_3d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetMsaa3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3764433340))
  }
  {
    methodName := StringNameFromStr("viewport_set_msaa_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetMsaa2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3764433340))
  }
  {
    methodName := StringNameFromStr("viewport_set_use_hdr_2d")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUseHdr2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_screen_space_aa")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetScreenSpaceAa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1447279591))
  }
  {
    methodName := StringNameFromStr("viewport_set_use_taa")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUseTaa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_use_debanding")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_use_occlusion_culling")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetUseOcclusionCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_set_occlusion_rays_per_thread")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetOcclusionRaysPerThread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("viewport_set_occlusion_culling_build_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetOcclusionCullingBuildQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2069725696))
  }
  {
    methodName := StringNameFromStr("viewport_get_render_info")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportGetRenderInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2041262392))
  }
  {
    methodName := StringNameFromStr("viewport_set_debug_draw")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetDebugDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2089420930))
  }
  {
    methodName := StringNameFromStr("viewport_set_measure_render_time")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetMeasureRenderTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("viewport_get_measured_render_time_cpu")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportGetMeasuredRenderTimeCpu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("viewport_get_measured_render_time_gpu")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportGetMeasuredRenderTimeGpu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("viewport_set_vrs_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetVrsMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 398809874))
  }
  {
    methodName := StringNameFromStr("viewport_set_vrs_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnViewportSetVrsTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("sky_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkyCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("sky_set_radiance_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkySetRadianceSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("sky_set_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkySetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3279019937))
  }
  {
    methodName := StringNameFromStr("sky_set_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkySetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("sky_bake_panorama")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSkyBakePanorama = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3875285818))
  }
  {
    methodName := StringNameFromStr("environment_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("environment_set_background")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937328877))
  }
  {
    methodName := StringNameFromStr("environment_set_sky")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("environment_set_sky_custom_fov")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSkyCustomFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("environment_set_sky_orientation")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSkyOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1735850857))
  }
  {
    methodName := StringNameFromStr("environment_set_bg_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("environment_set_bg_energy")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetBgEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2513314492))
  }
  {
    methodName := StringNameFromStr("environment_set_canvas_max_layer")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetCanvasMaxLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("environment_set_ambient_light")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetAmbientLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214961493))
  }
  {
    methodName := StringNameFromStr("environment_set_glow")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetGlow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2421724940))
  }
  {
    methodName := StringNameFromStr("environment_set_tonemap")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetTonemap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2914312638))
  }
  {
    methodName := StringNameFromStr("environment_set_adjustment")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetAdjustment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 876799838))
  }
  {
    methodName := StringNameFromStr("environment_set_ssr")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSsr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3607294374))
  }
  {
    methodName := StringNameFromStr("environment_set_ssao")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSsao = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3994732740))
  }
  {
    methodName := StringNameFromStr("environment_set_fog")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetFog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2793577733))
  }
  {
    methodName := StringNameFromStr("environment_set_sdfgi")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSdfgi = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3519144388))
  }
  {
    methodName := StringNameFromStr("environment_set_volumetric_fog")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetVolumetricFog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1553633833))
  }
  {
    methodName := StringNameFromStr("environment_glow_set_use_bicubic_upscale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentGlowSetUseBicubicUpscale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("environment_set_ssr_roughness_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSsrRoughnessQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1190026788))
  }
  {
    methodName := StringNameFromStr("environment_set_ssao_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSsaoQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 189753569))
  }
  {
    methodName := StringNameFromStr("environment_set_ssil_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSsilQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1713836683))
  }
  {
    methodName := StringNameFromStr("environment_set_sdfgi_ray_count")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSdfgiRayCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 340137951))
  }
  {
    methodName := StringNameFromStr("environment_set_sdfgi_frames_to_converge")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSdfgiFramesToConverge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2182444374))
  }
  {
    methodName := StringNameFromStr("environment_set_sdfgi_frames_to_update_light")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetSdfgiFramesToUpdateLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1251144068))
  }
  {
    methodName := StringNameFromStr("environment_set_volumetric_fog_volume_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetVolumetricFogVolumeSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("environment_set_volumetric_fog_filter_active")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentSetVolumetricFogFilterActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("environment_bake_panorama")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnEnvironmentBakePanorama = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2452908646))
  }
  {
    methodName := StringNameFromStr("screen_space_roughness_limiter_set_active")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnScreenSpaceRoughnessLimiterSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 916716790))
  }
  {
    methodName := StringNameFromStr("sub_surface_scattering_set_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSubSurfaceScatteringSetQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 64571803))
  }
  {
    methodName := StringNameFromStr("sub_surface_scattering_set_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSubSurfaceScatteringSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1017552074))
  }
  {
    methodName := StringNameFromStr("camera_attributes_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("camera_attributes_set_dof_blur_quality")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesSetDofBlurQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2220136795))
  }
  {
    methodName := StringNameFromStr("camera_attributes_set_dof_blur_bokeh_shape")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesSetDofBlurBokehShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1205058394))
  }
  {
    methodName := StringNameFromStr("camera_attributes_set_dof_blur")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesSetDofBlur = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 316272616))
  }
  {
    methodName := StringNameFromStr("camera_attributes_set_exposure")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesSetExposure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2513314492))
  }
  {
    methodName := StringNameFromStr("camera_attributes_set_auto_exposure")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCameraAttributesSetAutoExposure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4266986332))
  }
  {
    methodName := StringNameFromStr("scenario_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnScenarioCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("scenario_set_environment")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnScenarioSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("scenario_set_fallback_environment")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnScenarioSetFallbackEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("scenario_set_camera_attributes")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnScenarioSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_create2")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceCreate2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 746547085))
  }
  {
    methodName := StringNameFromStr("instance_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("instance_set_base")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetBase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_set_scenario")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetScenario = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_set_layer_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetLayerMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("instance_set_pivot_data")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetPivotData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1280615259))
  }
  {
    methodName := StringNameFromStr("instance_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935195649))
  }
  {
    methodName := StringNameFromStr("instance_attach_object_instance_id")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceAttachObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("instance_set_blend_shape_weight")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetBlendShapeWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
  }
  {
    methodName := StringNameFromStr("instance_set_surface_override_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetSurfaceOverrideMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
  }
  {
    methodName := StringNameFromStr("instance_set_visible")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_transparency")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetTransparency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("instance_set_custom_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3696536120))
  }
  {
    methodName := StringNameFromStr("instance_attach_skeleton")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceAttachSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_set_extra_visibility_margin")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetExtraVisibilityMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("instance_set_visibility_parent")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetVisibilityParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_set_ignore_culling")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceSetIgnoreCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_flag")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1014989537))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_cast_shadows_setting")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetCastShadowsSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3768836020))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_material_override")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_material_overlay")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetMaterialOverlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_visibility_range")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetVisibilityRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4263925858))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_lightmap")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetLightmap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 536974962))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_lod_bias")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetLodBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("instance_geometry_set_shader_parameter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometrySetShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3477296213))
  }
  {
    methodName := StringNameFromStr("instance_geometry_get_shader_parameter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometryGetShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2621281810))
  }
  {
    methodName := StringNameFromStr("instance_geometry_get_shader_parameter_default_value")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometryGetShaderParameterDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2621281810))
  }
  {
    methodName := StringNameFromStr("instance_geometry_get_shader_parameter_list")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstanceGeometryGetShaderParameterList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("instances_cull_aabb")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstancesCullAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2570105777))
  }
  {
    methodName := StringNameFromStr("instances_cull_ray")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstancesCullRay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2208759584))
  }
  {
    methodName := StringNameFromStr("instances_cull_convex")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnInstancesCullConvex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2488539944))
  }
  {
    methodName := StringNameFromStr("bake_render_uv2")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnBakeRenderUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1904608558))
  }
  {
    methodName := StringNameFromStr("canvas_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_set_item_mirroring")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasSetItemMirroring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2343975398))
  }
  {
    methodName := StringNameFromStr("canvas_set_modulate")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("canvas_set_disable_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasSetDisableScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("canvas_texture_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasTextureCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_texture_set_channel")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasTextureSetChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3822119138))
  }
  {
    methodName := StringNameFromStr("canvas_texture_set_shading_parameters")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasTextureSetShadingParameters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2124967469))
  }
  {
    methodName := StringNameFromStr("canvas_texture_set_texture_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasTextureSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1155129294))
  }
  {
    methodName := StringNameFromStr("canvas_texture_set_texture_repeat")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasTextureSetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1652956681))
  }
  {
    methodName := StringNameFromStr("canvas_item_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_parent")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_default_texture_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetDefaultTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1155129294))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_default_texture_repeat")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetDefaultTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1652956681))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_visible")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_light_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_visibility_layer")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetVisibilityLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_clip")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_distance_field_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetDistanceFieldMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_custom_rect")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetCustomRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1333997032))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_modulate")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_self_modulate")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetSelfModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_draw_behind_parent")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetDrawBehindParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_line")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1819681853))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_polyline")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddPolyline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3098767073))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_multiline")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddMultiline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2088642721))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_rect")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 934531857))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_circle")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddCircle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2439351960))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_texture_rect")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddTextureRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 324864032))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_msdf_texture_rect_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddMsdfTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 97408773))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_lcd_texture_rect_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddLcdTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 359793297))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_texture_rect_region")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 485157892))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_nine_patch")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddNinePatch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 389957886))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_primitive")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddPrimitive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3731601077))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_polygon")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3580000528))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_triangle_array")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddTriangleArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 660261329))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 316450961))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_multimesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2131855138))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_particles")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddParticles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2575754278))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_clip_ignore")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddClipIgnore = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_add_animation_slice")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemAddAnimationSlice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2646834499))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_sort_children_by_y")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetSortChildrenByY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_z_index")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_z_as_relative_to_parent")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetZAsRelativeToParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_copy_to_backbuffer")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetCopyToBackbuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2429202503))
  }
  {
    methodName := StringNameFromStr("canvas_item_clear")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_draw_index")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetDrawIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_use_parent_material")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetUseParentMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_visibility_notifier")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetVisibilityNotifier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3568945579))
  }
  {
    methodName := StringNameFromStr("canvas_item_set_canvas_group_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasItemSetCanvasGroupMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3973586316))
  }
  {
    methodName := StringNameFromStr("canvas_light_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_light_attach_to_canvas")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightAttachToCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_enabled")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_texture_scale")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetTextureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_texture_offset")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetTextureOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_height")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_energy")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_z_range")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetZRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_layer_range")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetLayerRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288446313))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_item_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetItemCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_item_shadow_cull_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetItemShadowCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2957564891))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_shadow_enabled")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetShadowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_shadow_filter")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetShadowFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 393119659))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_shadow_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_shadow_smooth")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetShadowSmooth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("canvas_light_set_blend_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 804895945))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_attach_to_canvas")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderAttachToCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_set_enabled")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_set_polygon")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_set_as_sdf_collision")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderSetAsSdfCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_set_transform")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("canvas_light_occluder_set_light_mask")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasLightOccluderSetLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("canvas_occluder_polygon_create")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasOccluderPolygonCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("canvas_occluder_polygon_set_shape")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasOccluderPolygonSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2103882027))
  }
  {
    methodName := StringNameFromStr("canvas_occluder_polygon_set_cull_mode")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasOccluderPolygonSetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1839404663))
  }
  {
    methodName := StringNameFromStr("canvas_set_shadow_texture_size")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCanvasSetShadowTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_add")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterAdd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 463390080))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_remove")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterRemove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_get_list")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterGetList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_set")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_set_override")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterSetOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_get")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
  }
  {
    methodName := StringNameFromStr("global_shader_parameter_get_type")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGlobalShaderParameterGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1601414142))
  }
  {
    methodName := StringNameFromStr("free_rid")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("request_frame_drawn_callback")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnRequestFrameDrawnCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
  {
    methodName := StringNameFromStr("has_changed")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnHasChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_rendering_info")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetRenderingInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3763192241))
  }
  {
    methodName := StringNameFromStr("get_video_adapter_name")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetVideoAdapterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_video_adapter_vendor")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetVideoAdapterVendor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_video_adapter_type")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetVideoAdapterType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3099547011))
  }
  {
    methodName := StringNameFromStr("get_video_adapter_api_version")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetVideoAdapterApiVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("make_sphere_mesh")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnMakeSphereMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2251015897))
  }
  {
    methodName := StringNameFromStr("get_test_cube")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetTestCube = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("get_test_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetTestTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("get_white_texture")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetWhiteTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("set_boot_image")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSetBootImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3759744527))
  }
  {
    methodName := StringNameFromStr("get_default_clear_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetDefaultClearColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200896285))
  }
  {
    methodName := StringNameFromStr("set_default_clear_color")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSetDefaultClearColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("has_feature")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnHasFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 598462696))
  }
  {
    methodName := StringNameFromStr("has_os_feature")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnHasOsFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("set_debug_generate_wireframes")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSetDebugGenerateWireframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_render_loop_enabled")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnIsRenderLoopEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_render_loop_enabled")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnSetRenderLoopEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_frame_setup_time_cpu")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetFrameSetupTimeCpu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("force_sync")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnForceSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("force_draw")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnForceDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1076185472))
  }
  {
    methodName := StringNameFromStr("get_rendering_device")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnGetRenderingDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1405107940))
  }
  {
    methodName := StringNameFromStr("create_local_rendering_device")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCreateLocalRenderingDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1405107940))
  }
  {
    methodName := StringNameFromStr("call_on_render_thread")
    defer methodName.Destroy()
    ptrsForRenderingServer.fnCallOnRenderThread = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
}

type RenderingServer struct {
  Object
}

func (me *RenderingServer) BaseClass() string {
  return "RenderingServer"
}

func NewRenderingServer() *RenderingServer {
  str := StringNameFromStr("RenderingServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RenderingServer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  RenderingServerNoIndexArray = "-1" // TODO: construct correctly
  RenderingServerArrayWeightsSize = "4" // TODO: construct correctly
  RenderingServerCanvasItemZMin = "-4096" // TODO: construct correctly
  RenderingServerCanvasItemZMax = "4096" // TODO: construct correctly
  RenderingServerMaxGlowLevels = "7" // TODO: construct correctly
  RenderingServerMaxCursors = "8" // TODO: construct correctly
  RenderingServerMax2DDirectionalLights = "8" // TODO: construct correctly
  RenderingServerMaterialRenderPriorityMin = "-128" // TODO: construct correctly
  RenderingServerMaterialRenderPriorityMax = "127" // TODO: construct correctly
  RenderingServerArrayCustomCount = "4" // TODO: construct correctly
  RenderingServerParticlesEmitFlagPosition = "1" // TODO: construct correctly
  RenderingServerParticlesEmitFlagRotationScale = "2" // TODO: construct correctly
  RenderingServerParticlesEmitFlagVelocity = "4" // TODO: construct correctly
  RenderingServerParticlesEmitFlagColor = "8" // TODO: construct correctly
  RenderingServerParticlesEmitFlagCustom = "16" // TODO: construct correctly
)

// Enums

type RenderingServerTextureLayeredType int
const (
  RenderingServerTextureLayeredTypeTextureLayered2DArray RenderingServerTextureLayeredType = 0
  RenderingServerTextureLayeredTypeTextureLayeredCubemap RenderingServerTextureLayeredType = 1
  RenderingServerTextureLayeredTypeTextureLayeredCubemapArray RenderingServerTextureLayeredType = 2
)

type RenderingServerCubeMapLayer int
const (
  RenderingServerCubeMapLayerCubemapLayerLeft RenderingServerCubeMapLayer = 0
  RenderingServerCubeMapLayerCubemapLayerRight RenderingServerCubeMapLayer = 1
  RenderingServerCubeMapLayerCubemapLayerBottom RenderingServerCubeMapLayer = 2
  RenderingServerCubeMapLayerCubemapLayerTop RenderingServerCubeMapLayer = 3
  RenderingServerCubeMapLayerCubemapLayerFront RenderingServerCubeMapLayer = 4
  RenderingServerCubeMapLayerCubemapLayerBack RenderingServerCubeMapLayer = 5
)

type RenderingServerShaderMode int
const (
  RenderingServerShaderModeShaderSpatial RenderingServerShaderMode = 0
  RenderingServerShaderModeShaderCanvasItem RenderingServerShaderMode = 1
  RenderingServerShaderModeShaderParticles RenderingServerShaderMode = 2
  RenderingServerShaderModeShaderSky RenderingServerShaderMode = 3
  RenderingServerShaderModeShaderFog RenderingServerShaderMode = 4
  RenderingServerShaderModeShaderMax RenderingServerShaderMode = 5
)

type RenderingServerArrayType int
const (
  RenderingServerArrayTypeArrayVertex RenderingServerArrayType = 0
  RenderingServerArrayTypeArrayNormal RenderingServerArrayType = 1
  RenderingServerArrayTypeArrayTangent RenderingServerArrayType = 2
  RenderingServerArrayTypeArrayColor RenderingServerArrayType = 3
  RenderingServerArrayTypeArrayTexUv RenderingServerArrayType = 4
  RenderingServerArrayTypeArrayTexUv2 RenderingServerArrayType = 5
  RenderingServerArrayTypeArrayCustom0 RenderingServerArrayType = 6
  RenderingServerArrayTypeArrayCustom1 RenderingServerArrayType = 7
  RenderingServerArrayTypeArrayCustom2 RenderingServerArrayType = 8
  RenderingServerArrayTypeArrayCustom3 RenderingServerArrayType = 9
  RenderingServerArrayTypeArrayBones RenderingServerArrayType = 10
  RenderingServerArrayTypeArrayWeights RenderingServerArrayType = 11
  RenderingServerArrayTypeArrayIndex RenderingServerArrayType = 12
  RenderingServerArrayTypeArrayMax RenderingServerArrayType = 13
)

type RenderingServerArrayCustomFormat int
const (
  RenderingServerArrayCustomFormatArrayCustomRgba8Unorm RenderingServerArrayCustomFormat = 0
  RenderingServerArrayCustomFormatArrayCustomRgba8Snorm RenderingServerArrayCustomFormat = 1
  RenderingServerArrayCustomFormatArrayCustomRgHalf RenderingServerArrayCustomFormat = 2
  RenderingServerArrayCustomFormatArrayCustomRgbaHalf RenderingServerArrayCustomFormat = 3
  RenderingServerArrayCustomFormatArrayCustomRFloat RenderingServerArrayCustomFormat = 4
  RenderingServerArrayCustomFormatArrayCustomRgFloat RenderingServerArrayCustomFormat = 5
  RenderingServerArrayCustomFormatArrayCustomRgbFloat RenderingServerArrayCustomFormat = 6
  RenderingServerArrayCustomFormatArrayCustomRgbaFloat RenderingServerArrayCustomFormat = 7
  RenderingServerArrayCustomFormatArrayCustomMax RenderingServerArrayCustomFormat = 8
)

type RenderingServerArrayFormat int
const (
  RenderingServerArrayFormatArrayFormatVertex RenderingServerArrayFormat = 1
  RenderingServerArrayFormatArrayFormatNormal RenderingServerArrayFormat = 2
  RenderingServerArrayFormatArrayFormatTangent RenderingServerArrayFormat = 4
  RenderingServerArrayFormatArrayFormatColor RenderingServerArrayFormat = 8
  RenderingServerArrayFormatArrayFormatTexUv RenderingServerArrayFormat = 16
  RenderingServerArrayFormatArrayFormatTexUv2 RenderingServerArrayFormat = 32
  RenderingServerArrayFormatArrayFormatCustom0 RenderingServerArrayFormat = 64
  RenderingServerArrayFormatArrayFormatCustom1 RenderingServerArrayFormat = 128
  RenderingServerArrayFormatArrayFormatCustom2 RenderingServerArrayFormat = 256
  RenderingServerArrayFormatArrayFormatCustom3 RenderingServerArrayFormat = 512
  RenderingServerArrayFormatArrayFormatBones RenderingServerArrayFormat = 1024
  RenderingServerArrayFormatArrayFormatWeights RenderingServerArrayFormat = 2048
  RenderingServerArrayFormatArrayFormatIndex RenderingServerArrayFormat = 4096
  RenderingServerArrayFormatArrayFormatBlendShapeMask RenderingServerArrayFormat = 7
  RenderingServerArrayFormatArrayFormatCustomBase RenderingServerArrayFormat = 13
  RenderingServerArrayFormatArrayFormatCustomBits RenderingServerArrayFormat = 3
  RenderingServerArrayFormatArrayFormatCustom0Shift RenderingServerArrayFormat = 13
  RenderingServerArrayFormatArrayFormatCustom1Shift RenderingServerArrayFormat = 16
  RenderingServerArrayFormatArrayFormatCustom2Shift RenderingServerArrayFormat = 19
  RenderingServerArrayFormatArrayFormatCustom3Shift RenderingServerArrayFormat = 22
  RenderingServerArrayFormatArrayFormatCustomMask RenderingServerArrayFormat = 7
  RenderingServerArrayFormatArrayCompressFlagsBase RenderingServerArrayFormat = 25
  RenderingServerArrayFormatArrayFlagUse2DVertices RenderingServerArrayFormat = 33554432
  RenderingServerArrayFormatArrayFlagUseDynamicUpdate RenderingServerArrayFormat = 67108864
  RenderingServerArrayFormatArrayFlagUse8BoneWeights RenderingServerArrayFormat = 134217728
  RenderingServerArrayFormatArrayFlagUsesEmptyVertexArray RenderingServerArrayFormat = 268435456
  RenderingServerArrayFormatArrayFlagCompressAttributes RenderingServerArrayFormat = 536870912
  RenderingServerArrayFormatArrayFlagFormatVersionBase RenderingServerArrayFormat = 35
  RenderingServerArrayFormatArrayFlagFormatVersionShift RenderingServerArrayFormat = 35
  RenderingServerArrayFormatArrayFlagFormatVersion1 RenderingServerArrayFormat = 0
  RenderingServerArrayFormatArrayFlagFormatVersion2 RenderingServerArrayFormat = 34359738368
  RenderingServerArrayFormatArrayFlagFormatCurrentVersion RenderingServerArrayFormat = 34359738368
  RenderingServerArrayFormatArrayFlagFormatVersionMask RenderingServerArrayFormat = 255
)

type RenderingServerPrimitiveType int
const (
  RenderingServerPrimitiveTypePrimitivePoints RenderingServerPrimitiveType = 0
  RenderingServerPrimitiveTypePrimitiveLines RenderingServerPrimitiveType = 1
  RenderingServerPrimitiveTypePrimitiveLineStrip RenderingServerPrimitiveType = 2
  RenderingServerPrimitiveTypePrimitiveTriangles RenderingServerPrimitiveType = 3
  RenderingServerPrimitiveTypePrimitiveTriangleStrip RenderingServerPrimitiveType = 4
  RenderingServerPrimitiveTypePrimitiveMax RenderingServerPrimitiveType = 5
)

type RenderingServerBlendShapeMode int
const (
  RenderingServerBlendShapeModeBlendShapeModeNormalized RenderingServerBlendShapeMode = 0
  RenderingServerBlendShapeModeBlendShapeModeRelative RenderingServerBlendShapeMode = 1
)

type RenderingServerMultimeshTransformFormat int
const (
  RenderingServerMultimeshTransformFormatMultimeshTransform2D RenderingServerMultimeshTransformFormat = 0
  RenderingServerMultimeshTransformFormatMultimeshTransform3D RenderingServerMultimeshTransformFormat = 1
)

type RenderingServerLightProjectorFilter int
const (
  RenderingServerLightProjectorFilterLightProjectorFilterNearest RenderingServerLightProjectorFilter = 0
  RenderingServerLightProjectorFilterLightProjectorFilterLinear RenderingServerLightProjectorFilter = 1
  RenderingServerLightProjectorFilterLightProjectorFilterNearestMipmaps RenderingServerLightProjectorFilter = 2
  RenderingServerLightProjectorFilterLightProjectorFilterLinearMipmaps RenderingServerLightProjectorFilter = 3
  RenderingServerLightProjectorFilterLightProjectorFilterNearestMipmapsAnisotropic RenderingServerLightProjectorFilter = 4
  RenderingServerLightProjectorFilterLightProjectorFilterLinearMipmapsAnisotropic RenderingServerLightProjectorFilter = 5
)

type RenderingServerLightType int
const (
  RenderingServerLightTypeLightDirectional RenderingServerLightType = 0
  RenderingServerLightTypeLightOmni RenderingServerLightType = 1
  RenderingServerLightTypeLightSpot RenderingServerLightType = 2
)

type RenderingServerLightParam int
const (
  RenderingServerLightParamLightParamEnergy RenderingServerLightParam = 0
  RenderingServerLightParamLightParamIndirectEnergy RenderingServerLightParam = 1
  RenderingServerLightParamLightParamVolumetricFogEnergy RenderingServerLightParam = 2
  RenderingServerLightParamLightParamSpecular RenderingServerLightParam = 3
  RenderingServerLightParamLightParamRange RenderingServerLightParam = 4
  RenderingServerLightParamLightParamSize RenderingServerLightParam = 5
  RenderingServerLightParamLightParamAttenuation RenderingServerLightParam = 6
  RenderingServerLightParamLightParamSpotAngle RenderingServerLightParam = 7
  RenderingServerLightParamLightParamSpotAttenuation RenderingServerLightParam = 8
  RenderingServerLightParamLightParamShadowMaxDistance RenderingServerLightParam = 9
  RenderingServerLightParamLightParamShadowSplit1Offset RenderingServerLightParam = 10
  RenderingServerLightParamLightParamShadowSplit2Offset RenderingServerLightParam = 11
  RenderingServerLightParamLightParamShadowSplit3Offset RenderingServerLightParam = 12
  RenderingServerLightParamLightParamShadowFadeStart RenderingServerLightParam = 13
  RenderingServerLightParamLightParamShadowNormalBias RenderingServerLightParam = 14
  RenderingServerLightParamLightParamShadowBias RenderingServerLightParam = 15
  RenderingServerLightParamLightParamShadowPancakeSize RenderingServerLightParam = 16
  RenderingServerLightParamLightParamShadowOpacity RenderingServerLightParam = 17
  RenderingServerLightParamLightParamShadowBlur RenderingServerLightParam = 18
  RenderingServerLightParamLightParamTransmittanceBias RenderingServerLightParam = 19
  RenderingServerLightParamLightParamIntensity RenderingServerLightParam = 20
  RenderingServerLightParamLightParamMax RenderingServerLightParam = 21
)

type RenderingServerLightBakeMode int
const (
  RenderingServerLightBakeModeLightBakeDisabled RenderingServerLightBakeMode = 0
  RenderingServerLightBakeModeLightBakeStatic RenderingServerLightBakeMode = 1
  RenderingServerLightBakeModeLightBakeDynamic RenderingServerLightBakeMode = 2
)

type RenderingServerLightOmniShadowMode int
const (
  RenderingServerLightOmniShadowModeLightOmniShadowDualParaboloid RenderingServerLightOmniShadowMode = 0
  RenderingServerLightOmniShadowModeLightOmniShadowCube RenderingServerLightOmniShadowMode = 1
)

type RenderingServerLightDirectionalShadowMode int
const (
  RenderingServerLightDirectionalShadowModeLightDirectionalShadowOrthogonal RenderingServerLightDirectionalShadowMode = 0
  RenderingServerLightDirectionalShadowModeLightDirectionalShadowParallel2Splits RenderingServerLightDirectionalShadowMode = 1
  RenderingServerLightDirectionalShadowModeLightDirectionalShadowParallel4Splits RenderingServerLightDirectionalShadowMode = 2
)

type RenderingServerLightDirectionalSkyMode int
const (
  RenderingServerLightDirectionalSkyModeLightDirectionalSkyModeLightAndSky RenderingServerLightDirectionalSkyMode = 0
  RenderingServerLightDirectionalSkyModeLightDirectionalSkyModeLightOnly RenderingServerLightDirectionalSkyMode = 1
  RenderingServerLightDirectionalSkyModeLightDirectionalSkyModeSkyOnly RenderingServerLightDirectionalSkyMode = 2
)

type RenderingServerShadowQuality int
const (
  RenderingServerShadowQualityShadowQualityHard RenderingServerShadowQuality = 0
  RenderingServerShadowQualityShadowQualitySoftVeryLow RenderingServerShadowQuality = 1
  RenderingServerShadowQualityShadowQualitySoftLow RenderingServerShadowQuality = 2
  RenderingServerShadowQualityShadowQualitySoftMedium RenderingServerShadowQuality = 3
  RenderingServerShadowQualityShadowQualitySoftHigh RenderingServerShadowQuality = 4
  RenderingServerShadowQualityShadowQualitySoftUltra RenderingServerShadowQuality = 5
  RenderingServerShadowQualityShadowQualityMax RenderingServerShadowQuality = 6
)

type RenderingServerReflectionProbeUpdateMode int
const (
  RenderingServerReflectionProbeUpdateModeReflectionProbeUpdateOnce RenderingServerReflectionProbeUpdateMode = 0
  RenderingServerReflectionProbeUpdateModeReflectionProbeUpdateAlways RenderingServerReflectionProbeUpdateMode = 1
)

type RenderingServerReflectionProbeAmbientMode int
const (
  RenderingServerReflectionProbeAmbientModeReflectionProbeAmbientDisabled RenderingServerReflectionProbeAmbientMode = 0
  RenderingServerReflectionProbeAmbientModeReflectionProbeAmbientEnvironment RenderingServerReflectionProbeAmbientMode = 1
  RenderingServerReflectionProbeAmbientModeReflectionProbeAmbientColor RenderingServerReflectionProbeAmbientMode = 2
)

type RenderingServerDecalTexture int
const (
  RenderingServerDecalTextureDecalTextureAlbedo RenderingServerDecalTexture = 0
  RenderingServerDecalTextureDecalTextureNormal RenderingServerDecalTexture = 1
  RenderingServerDecalTextureDecalTextureOrm RenderingServerDecalTexture = 2
  RenderingServerDecalTextureDecalTextureEmission RenderingServerDecalTexture = 3
  RenderingServerDecalTextureDecalTextureMax RenderingServerDecalTexture = 4
)

type RenderingServerDecalFilter int
const (
  RenderingServerDecalFilterDecalFilterNearest RenderingServerDecalFilter = 0
  RenderingServerDecalFilterDecalFilterLinear RenderingServerDecalFilter = 1
  RenderingServerDecalFilterDecalFilterNearestMipmaps RenderingServerDecalFilter = 2
  RenderingServerDecalFilterDecalFilterLinearMipmaps RenderingServerDecalFilter = 3
  RenderingServerDecalFilterDecalFilterNearestMipmapsAnisotropic RenderingServerDecalFilter = 4
  RenderingServerDecalFilterDecalFilterLinearMipmapsAnisotropic RenderingServerDecalFilter = 5
)

type RenderingServerVoxelGIQuality int
const (
  RenderingServerVoxelGIQualityVoxelGiQualityLow RenderingServerVoxelGIQuality = 0
  RenderingServerVoxelGIQualityVoxelGiQualityHigh RenderingServerVoxelGIQuality = 1
)

type RenderingServerParticlesMode int
const (
  RenderingServerParticlesModeParticlesMode2D RenderingServerParticlesMode = 0
  RenderingServerParticlesModeParticlesMode3D RenderingServerParticlesMode = 1
)

type RenderingServerParticlesTransformAlign int
const (
  RenderingServerParticlesTransformAlignParticlesTransformAlignDisabled RenderingServerParticlesTransformAlign = 0
  RenderingServerParticlesTransformAlignParticlesTransformAlignZBillboard RenderingServerParticlesTransformAlign = 1
  RenderingServerParticlesTransformAlignParticlesTransformAlignYToVelocity RenderingServerParticlesTransformAlign = 2
  RenderingServerParticlesTransformAlignParticlesTransformAlignZBillboardYToVelocity RenderingServerParticlesTransformAlign = 3
)

type RenderingServerParticlesDrawOrder int
const (
  RenderingServerParticlesDrawOrderParticlesDrawOrderIndex RenderingServerParticlesDrawOrder = 0
  RenderingServerParticlesDrawOrderParticlesDrawOrderLifetime RenderingServerParticlesDrawOrder = 1
  RenderingServerParticlesDrawOrderParticlesDrawOrderReverseLifetime RenderingServerParticlesDrawOrder = 2
  RenderingServerParticlesDrawOrderParticlesDrawOrderViewDepth RenderingServerParticlesDrawOrder = 3
)

type RenderingServerParticlesCollisionType int
const (
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeSphereAttract RenderingServerParticlesCollisionType = 0
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeBoxAttract RenderingServerParticlesCollisionType = 1
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeVectorFieldAttract RenderingServerParticlesCollisionType = 2
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeSphereCollide RenderingServerParticlesCollisionType = 3
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeBoxCollide RenderingServerParticlesCollisionType = 4
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeSdfCollide RenderingServerParticlesCollisionType = 5
  RenderingServerParticlesCollisionTypeParticlesCollisionTypeHeightfieldCollide RenderingServerParticlesCollisionType = 6
)

type RenderingServerParticlesCollisionHeightfieldResolution int
const (
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution256 RenderingServerParticlesCollisionHeightfieldResolution = 0
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution512 RenderingServerParticlesCollisionHeightfieldResolution = 1
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution1024 RenderingServerParticlesCollisionHeightfieldResolution = 2
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution2048 RenderingServerParticlesCollisionHeightfieldResolution = 3
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution4096 RenderingServerParticlesCollisionHeightfieldResolution = 4
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolution8192 RenderingServerParticlesCollisionHeightfieldResolution = 5
  RenderingServerParticlesCollisionHeightfieldResolutionParticlesCollisionHeightfieldResolutionMax RenderingServerParticlesCollisionHeightfieldResolution = 6
)

type RenderingServerFogVolumeShape int
const (
  RenderingServerFogVolumeShapeFogVolumeShapeEllipsoid RenderingServerFogVolumeShape = 0
  RenderingServerFogVolumeShapeFogVolumeShapeCone RenderingServerFogVolumeShape = 1
  RenderingServerFogVolumeShapeFogVolumeShapeCylinder RenderingServerFogVolumeShape = 2
  RenderingServerFogVolumeShapeFogVolumeShapeBox RenderingServerFogVolumeShape = 3
  RenderingServerFogVolumeShapeFogVolumeShapeWorld RenderingServerFogVolumeShape = 4
  RenderingServerFogVolumeShapeFogVolumeShapeMax RenderingServerFogVolumeShape = 5
)

type RenderingServerViewportScaling3DMode int
const (
  RenderingServerViewportScaling3DModeViewportScaling3DModeBilinear RenderingServerViewportScaling3DMode = 0
  RenderingServerViewportScaling3DModeViewportScaling3DModeFsr RenderingServerViewportScaling3DMode = 1
  RenderingServerViewportScaling3DModeViewportScaling3DModeFsr2 RenderingServerViewportScaling3DMode = 2
  RenderingServerViewportScaling3DModeViewportScaling3DModeMax RenderingServerViewportScaling3DMode = 3
)

type RenderingServerViewportUpdateMode int
const (
  RenderingServerViewportUpdateModeViewportUpdateDisabled RenderingServerViewportUpdateMode = 0
  RenderingServerViewportUpdateModeViewportUpdateOnce RenderingServerViewportUpdateMode = 1
  RenderingServerViewportUpdateModeViewportUpdateWhenVisible RenderingServerViewportUpdateMode = 2
  RenderingServerViewportUpdateModeViewportUpdateWhenParentVisible RenderingServerViewportUpdateMode = 3
  RenderingServerViewportUpdateModeViewportUpdateAlways RenderingServerViewportUpdateMode = 4
)

type RenderingServerViewportClearMode int
const (
  RenderingServerViewportClearModeViewportClearAlways RenderingServerViewportClearMode = 0
  RenderingServerViewportClearModeViewportClearNever RenderingServerViewportClearMode = 1
  RenderingServerViewportClearModeViewportClearOnlyNextFrame RenderingServerViewportClearMode = 2
)

type RenderingServerViewportEnvironmentMode int
const (
  RenderingServerViewportEnvironmentModeViewportEnvironmentDisabled RenderingServerViewportEnvironmentMode = 0
  RenderingServerViewportEnvironmentModeViewportEnvironmentEnabled RenderingServerViewportEnvironmentMode = 1
  RenderingServerViewportEnvironmentModeViewportEnvironmentInherit RenderingServerViewportEnvironmentMode = 2
  RenderingServerViewportEnvironmentModeViewportEnvironmentMax RenderingServerViewportEnvironmentMode = 3
)

type RenderingServerViewportSDFOversize int
const (
  RenderingServerViewportSDFOversizeViewportSdfOversize100Percent RenderingServerViewportSDFOversize = 0
  RenderingServerViewportSDFOversizeViewportSdfOversize120Percent RenderingServerViewportSDFOversize = 1
  RenderingServerViewportSDFOversizeViewportSdfOversize150Percent RenderingServerViewportSDFOversize = 2
  RenderingServerViewportSDFOversizeViewportSdfOversize200Percent RenderingServerViewportSDFOversize = 3
  RenderingServerViewportSDFOversizeViewportSdfOversizeMax RenderingServerViewportSDFOversize = 4
)

type RenderingServerViewportSDFScale int
const (
  RenderingServerViewportSDFScaleViewportSdfScale100Percent RenderingServerViewportSDFScale = 0
  RenderingServerViewportSDFScaleViewportSdfScale50Percent RenderingServerViewportSDFScale = 1
  RenderingServerViewportSDFScaleViewportSdfScale25Percent RenderingServerViewportSDFScale = 2
  RenderingServerViewportSDFScaleViewportSdfScaleMax RenderingServerViewportSDFScale = 3
)

type RenderingServerViewportMSAA int
const (
  RenderingServerViewportMSAAViewportMsaaDisabled RenderingServerViewportMSAA = 0
  RenderingServerViewportMSAAViewportMsaa2X RenderingServerViewportMSAA = 1
  RenderingServerViewportMSAAViewportMsaa4X RenderingServerViewportMSAA = 2
  RenderingServerViewportMSAAViewportMsaa8X RenderingServerViewportMSAA = 3
  RenderingServerViewportMSAAViewportMsaaMax RenderingServerViewportMSAA = 4
)

type RenderingServerViewportScreenSpaceAA int
const (
  RenderingServerViewportScreenSpaceAAViewportScreenSpaceAaDisabled RenderingServerViewportScreenSpaceAA = 0
  RenderingServerViewportScreenSpaceAAViewportScreenSpaceAaFxaa RenderingServerViewportScreenSpaceAA = 1
  RenderingServerViewportScreenSpaceAAViewportScreenSpaceAaMax RenderingServerViewportScreenSpaceAA = 2
)

type RenderingServerViewportOcclusionCullingBuildQuality int
const (
  RenderingServerViewportOcclusionCullingBuildQualityViewportOcclusionBuildQualityLow RenderingServerViewportOcclusionCullingBuildQuality = 0
  RenderingServerViewportOcclusionCullingBuildQualityViewportOcclusionBuildQualityMedium RenderingServerViewportOcclusionCullingBuildQuality = 1
  RenderingServerViewportOcclusionCullingBuildQualityViewportOcclusionBuildQualityHigh RenderingServerViewportOcclusionCullingBuildQuality = 2
)

type RenderingServerViewportRenderInfo int
const (
  RenderingServerViewportRenderInfoViewportRenderInfoObjectsInFrame RenderingServerViewportRenderInfo = 0
  RenderingServerViewportRenderInfoViewportRenderInfoPrimitivesInFrame RenderingServerViewportRenderInfo = 1
  RenderingServerViewportRenderInfoViewportRenderInfoDrawCallsInFrame RenderingServerViewportRenderInfo = 2
  RenderingServerViewportRenderInfoViewportRenderInfoMax RenderingServerViewportRenderInfo = 3
)

type RenderingServerViewportRenderInfoType int
const (
  RenderingServerViewportRenderInfoTypeViewportRenderInfoTypeVisible RenderingServerViewportRenderInfoType = 0
  RenderingServerViewportRenderInfoTypeViewportRenderInfoTypeShadow RenderingServerViewportRenderInfoType = 1
  RenderingServerViewportRenderInfoTypeViewportRenderInfoTypeMax RenderingServerViewportRenderInfoType = 2
)

type RenderingServerViewportDebugDraw int
const (
  RenderingServerViewportDebugDrawViewportDebugDrawDisabled RenderingServerViewportDebugDraw = 0
  RenderingServerViewportDebugDrawViewportDebugDrawUnshaded RenderingServerViewportDebugDraw = 1
  RenderingServerViewportDebugDrawViewportDebugDrawLighting RenderingServerViewportDebugDraw = 2
  RenderingServerViewportDebugDrawViewportDebugDrawOverdraw RenderingServerViewportDebugDraw = 3
  RenderingServerViewportDebugDrawViewportDebugDrawWireframe RenderingServerViewportDebugDraw = 4
  RenderingServerViewportDebugDrawViewportDebugDrawNormalBuffer RenderingServerViewportDebugDraw = 5
  RenderingServerViewportDebugDrawViewportDebugDrawVoxelGiAlbedo RenderingServerViewportDebugDraw = 6
  RenderingServerViewportDebugDrawViewportDebugDrawVoxelGiLighting RenderingServerViewportDebugDraw = 7
  RenderingServerViewportDebugDrawViewportDebugDrawVoxelGiEmission RenderingServerViewportDebugDraw = 8
  RenderingServerViewportDebugDrawViewportDebugDrawShadowAtlas RenderingServerViewportDebugDraw = 9
  RenderingServerViewportDebugDrawViewportDebugDrawDirectionalShadowAtlas RenderingServerViewportDebugDraw = 10
  RenderingServerViewportDebugDrawViewportDebugDrawSceneLuminance RenderingServerViewportDebugDraw = 11
  RenderingServerViewportDebugDrawViewportDebugDrawSsao RenderingServerViewportDebugDraw = 12
  RenderingServerViewportDebugDrawViewportDebugDrawSsil RenderingServerViewportDebugDraw = 13
  RenderingServerViewportDebugDrawViewportDebugDrawPssmSplits RenderingServerViewportDebugDraw = 14
  RenderingServerViewportDebugDrawViewportDebugDrawDecalAtlas RenderingServerViewportDebugDraw = 15
  RenderingServerViewportDebugDrawViewportDebugDrawSdfgi RenderingServerViewportDebugDraw = 16
  RenderingServerViewportDebugDrawViewportDebugDrawSdfgiProbes RenderingServerViewportDebugDraw = 17
  RenderingServerViewportDebugDrawViewportDebugDrawGiBuffer RenderingServerViewportDebugDraw = 18
  RenderingServerViewportDebugDrawViewportDebugDrawDisableLod RenderingServerViewportDebugDraw = 19
  RenderingServerViewportDebugDrawViewportDebugDrawClusterOmniLights RenderingServerViewportDebugDraw = 20
  RenderingServerViewportDebugDrawViewportDebugDrawClusterSpotLights RenderingServerViewportDebugDraw = 21
  RenderingServerViewportDebugDrawViewportDebugDrawClusterDecals RenderingServerViewportDebugDraw = 22
  RenderingServerViewportDebugDrawViewportDebugDrawClusterReflectionProbes RenderingServerViewportDebugDraw = 23
  RenderingServerViewportDebugDrawViewportDebugDrawOccluders RenderingServerViewportDebugDraw = 24
  RenderingServerViewportDebugDrawViewportDebugDrawMotionVectors RenderingServerViewportDebugDraw = 25
  RenderingServerViewportDebugDrawViewportDebugDrawInternalBuffer RenderingServerViewportDebugDraw = 26
)

type RenderingServerViewportVRSMode int
const (
  RenderingServerViewportVRSModeViewportVrsDisabled RenderingServerViewportVRSMode = 0
  RenderingServerViewportVRSModeViewportVrsTexture RenderingServerViewportVRSMode = 1
  RenderingServerViewportVRSModeViewportVrsXr RenderingServerViewportVRSMode = 2
  RenderingServerViewportVRSModeViewportVrsMax RenderingServerViewportVRSMode = 3
)

type RenderingServerSkyMode int
const (
  RenderingServerSkyModeSkyModeAutomatic RenderingServerSkyMode = 0
  RenderingServerSkyModeSkyModeQuality RenderingServerSkyMode = 1
  RenderingServerSkyModeSkyModeIncremental RenderingServerSkyMode = 2
  RenderingServerSkyModeSkyModeRealtime RenderingServerSkyMode = 3
)

type RenderingServerEnvironmentBG int
const (
  RenderingServerEnvironmentBGEnvBgClearColor RenderingServerEnvironmentBG = 0
  RenderingServerEnvironmentBGEnvBgColor RenderingServerEnvironmentBG = 1
  RenderingServerEnvironmentBGEnvBgSky RenderingServerEnvironmentBG = 2
  RenderingServerEnvironmentBGEnvBgCanvas RenderingServerEnvironmentBG = 3
  RenderingServerEnvironmentBGEnvBgKeep RenderingServerEnvironmentBG = 4
  RenderingServerEnvironmentBGEnvBgCameraFeed RenderingServerEnvironmentBG = 5
  RenderingServerEnvironmentBGEnvBgMax RenderingServerEnvironmentBG = 6
)

type RenderingServerEnvironmentAmbientSource int
const (
  RenderingServerEnvironmentAmbientSourceEnvAmbientSourceBg RenderingServerEnvironmentAmbientSource = 0
  RenderingServerEnvironmentAmbientSourceEnvAmbientSourceDisabled RenderingServerEnvironmentAmbientSource = 1
  RenderingServerEnvironmentAmbientSourceEnvAmbientSourceColor RenderingServerEnvironmentAmbientSource = 2
  RenderingServerEnvironmentAmbientSourceEnvAmbientSourceSky RenderingServerEnvironmentAmbientSource = 3
)

type RenderingServerEnvironmentReflectionSource int
const (
  RenderingServerEnvironmentReflectionSourceEnvReflectionSourceBg RenderingServerEnvironmentReflectionSource = 0
  RenderingServerEnvironmentReflectionSourceEnvReflectionSourceDisabled RenderingServerEnvironmentReflectionSource = 1
  RenderingServerEnvironmentReflectionSourceEnvReflectionSourceSky RenderingServerEnvironmentReflectionSource = 2
)

type RenderingServerEnvironmentGlowBlendMode int
const (
  RenderingServerEnvironmentGlowBlendModeEnvGlowBlendModeAdditive RenderingServerEnvironmentGlowBlendMode = 0
  RenderingServerEnvironmentGlowBlendModeEnvGlowBlendModeScreen RenderingServerEnvironmentGlowBlendMode = 1
  RenderingServerEnvironmentGlowBlendModeEnvGlowBlendModeSoftlight RenderingServerEnvironmentGlowBlendMode = 2
  RenderingServerEnvironmentGlowBlendModeEnvGlowBlendModeReplace RenderingServerEnvironmentGlowBlendMode = 3
  RenderingServerEnvironmentGlowBlendModeEnvGlowBlendModeMix RenderingServerEnvironmentGlowBlendMode = 4
)

type RenderingServerEnvironmentToneMapper int
const (
  RenderingServerEnvironmentToneMapperEnvToneMapperLinear RenderingServerEnvironmentToneMapper = 0
  RenderingServerEnvironmentToneMapperEnvToneMapperReinhard RenderingServerEnvironmentToneMapper = 1
  RenderingServerEnvironmentToneMapperEnvToneMapperFilmic RenderingServerEnvironmentToneMapper = 2
  RenderingServerEnvironmentToneMapperEnvToneMapperAces RenderingServerEnvironmentToneMapper = 3
)

type RenderingServerEnvironmentSSRRoughnessQuality int
const (
  RenderingServerEnvironmentSSRRoughnessQualityEnvSsrRoughnessQualityDisabled RenderingServerEnvironmentSSRRoughnessQuality = 0
  RenderingServerEnvironmentSSRRoughnessQualityEnvSsrRoughnessQualityLow RenderingServerEnvironmentSSRRoughnessQuality = 1
  RenderingServerEnvironmentSSRRoughnessQualityEnvSsrRoughnessQualityMedium RenderingServerEnvironmentSSRRoughnessQuality = 2
  RenderingServerEnvironmentSSRRoughnessQualityEnvSsrRoughnessQualityHigh RenderingServerEnvironmentSSRRoughnessQuality = 3
)

type RenderingServerEnvironmentSSAOQuality int
const (
  RenderingServerEnvironmentSSAOQualityEnvSsaoQualityVeryLow RenderingServerEnvironmentSSAOQuality = 0
  RenderingServerEnvironmentSSAOQualityEnvSsaoQualityLow RenderingServerEnvironmentSSAOQuality = 1
  RenderingServerEnvironmentSSAOQualityEnvSsaoQualityMedium RenderingServerEnvironmentSSAOQuality = 2
  RenderingServerEnvironmentSSAOQualityEnvSsaoQualityHigh RenderingServerEnvironmentSSAOQuality = 3
  RenderingServerEnvironmentSSAOQualityEnvSsaoQualityUltra RenderingServerEnvironmentSSAOQuality = 4
)

type RenderingServerEnvironmentSSILQuality int
const (
  RenderingServerEnvironmentSSILQualityEnvSsilQualityVeryLow RenderingServerEnvironmentSSILQuality = 0
  RenderingServerEnvironmentSSILQualityEnvSsilQualityLow RenderingServerEnvironmentSSILQuality = 1
  RenderingServerEnvironmentSSILQualityEnvSsilQualityMedium RenderingServerEnvironmentSSILQuality = 2
  RenderingServerEnvironmentSSILQualityEnvSsilQualityHigh RenderingServerEnvironmentSSILQuality = 3
  RenderingServerEnvironmentSSILQualityEnvSsilQualityUltra RenderingServerEnvironmentSSILQuality = 4
)

type RenderingServerEnvironmentSDFGIYScale int
const (
  RenderingServerEnvironmentSDFGIYScaleEnvSdfgiYScale50Percent RenderingServerEnvironmentSDFGIYScale = 0
  RenderingServerEnvironmentSDFGIYScaleEnvSdfgiYScale75Percent RenderingServerEnvironmentSDFGIYScale = 1
  RenderingServerEnvironmentSDFGIYScaleEnvSdfgiYScale100Percent RenderingServerEnvironmentSDFGIYScale = 2
)

type RenderingServerEnvironmentSDFGIRayCount int
const (
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount4 RenderingServerEnvironmentSDFGIRayCount = 0
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount8 RenderingServerEnvironmentSDFGIRayCount = 1
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount16 RenderingServerEnvironmentSDFGIRayCount = 2
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount32 RenderingServerEnvironmentSDFGIRayCount = 3
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount64 RenderingServerEnvironmentSDFGIRayCount = 4
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount96 RenderingServerEnvironmentSDFGIRayCount = 5
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCount128 RenderingServerEnvironmentSDFGIRayCount = 6
  RenderingServerEnvironmentSDFGIRayCountEnvSdfgiRayCountMax RenderingServerEnvironmentSDFGIRayCount = 7
)

type RenderingServerEnvironmentSDFGIFramesToConverge int
const (
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn5Frames RenderingServerEnvironmentSDFGIFramesToConverge = 0
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn10Frames RenderingServerEnvironmentSDFGIFramesToConverge = 1
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn15Frames RenderingServerEnvironmentSDFGIFramesToConverge = 2
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn20Frames RenderingServerEnvironmentSDFGIFramesToConverge = 3
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn25Frames RenderingServerEnvironmentSDFGIFramesToConverge = 4
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeIn30Frames RenderingServerEnvironmentSDFGIFramesToConverge = 5
  RenderingServerEnvironmentSDFGIFramesToConvergeEnvSdfgiConvergeMax RenderingServerEnvironmentSDFGIFramesToConverge = 6
)

type RenderingServerEnvironmentSDFGIFramesToUpdateLight int
const (
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightIn1Frame RenderingServerEnvironmentSDFGIFramesToUpdateLight = 0
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightIn2Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 1
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightIn4Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 2
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightIn8Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 3
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightIn16Frames RenderingServerEnvironmentSDFGIFramesToUpdateLight = 4
  RenderingServerEnvironmentSDFGIFramesToUpdateLightEnvSdfgiUpdateLightMax RenderingServerEnvironmentSDFGIFramesToUpdateLight = 5
)

type RenderingServerSubSurfaceScatteringQuality int
const (
  RenderingServerSubSurfaceScatteringQualitySubSurfaceScatteringQualityDisabled RenderingServerSubSurfaceScatteringQuality = 0
  RenderingServerSubSurfaceScatteringQualitySubSurfaceScatteringQualityLow RenderingServerSubSurfaceScatteringQuality = 1
  RenderingServerSubSurfaceScatteringQualitySubSurfaceScatteringQualityMedium RenderingServerSubSurfaceScatteringQuality = 2
  RenderingServerSubSurfaceScatteringQualitySubSurfaceScatteringQualityHigh RenderingServerSubSurfaceScatteringQuality = 3
)

type RenderingServerDOFBokehShape int
const (
  RenderingServerDOFBokehShapeDofBokehBox RenderingServerDOFBokehShape = 0
  RenderingServerDOFBokehShapeDofBokehHexagon RenderingServerDOFBokehShape = 1
  RenderingServerDOFBokehShapeDofBokehCircle RenderingServerDOFBokehShape = 2
)

type RenderingServerDOFBlurQuality int
const (
  RenderingServerDOFBlurQualityDofBlurQualityVeryLow RenderingServerDOFBlurQuality = 0
  RenderingServerDOFBlurQualityDofBlurQualityLow RenderingServerDOFBlurQuality = 1
  RenderingServerDOFBlurQualityDofBlurQualityMedium RenderingServerDOFBlurQuality = 2
  RenderingServerDOFBlurQualityDofBlurQualityHigh RenderingServerDOFBlurQuality = 3
)

type RenderingServerInstanceType int
const (
  RenderingServerInstanceTypeInstanceNone RenderingServerInstanceType = 0
  RenderingServerInstanceTypeInstanceMesh RenderingServerInstanceType = 1
  RenderingServerInstanceTypeInstanceMultimesh RenderingServerInstanceType = 2
  RenderingServerInstanceTypeInstanceParticles RenderingServerInstanceType = 3
  RenderingServerInstanceTypeInstanceParticlesCollision RenderingServerInstanceType = 4
  RenderingServerInstanceTypeInstanceLight RenderingServerInstanceType = 5
  RenderingServerInstanceTypeInstanceReflectionProbe RenderingServerInstanceType = 6
  RenderingServerInstanceTypeInstanceDecal RenderingServerInstanceType = 7
  RenderingServerInstanceTypeInstanceVoxelGi RenderingServerInstanceType = 8
  RenderingServerInstanceTypeInstanceLightmap RenderingServerInstanceType = 9
  RenderingServerInstanceTypeInstanceOccluder RenderingServerInstanceType = 10
  RenderingServerInstanceTypeInstanceVisiblityNotifier RenderingServerInstanceType = 11
  RenderingServerInstanceTypeInstanceFogVolume RenderingServerInstanceType = 12
  RenderingServerInstanceTypeInstanceMax RenderingServerInstanceType = 13
  RenderingServerInstanceTypeInstanceGeometryMask RenderingServerInstanceType = 14
)

type RenderingServerInstanceFlags int
const (
  RenderingServerInstanceFlagsInstanceFlagUseBakedLight RenderingServerInstanceFlags = 0
  RenderingServerInstanceFlagsInstanceFlagUseDynamicGi RenderingServerInstanceFlags = 1
  RenderingServerInstanceFlagsInstanceFlagDrawNextFrameIfVisible RenderingServerInstanceFlags = 2
  RenderingServerInstanceFlagsInstanceFlagIgnoreOcclusionCulling RenderingServerInstanceFlags = 3
  RenderingServerInstanceFlagsInstanceFlagMax RenderingServerInstanceFlags = 4
)

type RenderingServerShadowCastingSetting int
const (
  RenderingServerShadowCastingSettingShadowCastingSettingOff RenderingServerShadowCastingSetting = 0
  RenderingServerShadowCastingSettingShadowCastingSettingOn RenderingServerShadowCastingSetting = 1
  RenderingServerShadowCastingSettingShadowCastingSettingDoubleSided RenderingServerShadowCastingSetting = 2
  RenderingServerShadowCastingSettingShadowCastingSettingShadowsOnly RenderingServerShadowCastingSetting = 3
)

type RenderingServerVisibilityRangeFadeMode int
const (
  RenderingServerVisibilityRangeFadeModeVisibilityRangeFadeDisabled RenderingServerVisibilityRangeFadeMode = 0
  RenderingServerVisibilityRangeFadeModeVisibilityRangeFadeSelf RenderingServerVisibilityRangeFadeMode = 1
  RenderingServerVisibilityRangeFadeModeVisibilityRangeFadeDependencies RenderingServerVisibilityRangeFadeMode = 2
)

type RenderingServerBakeChannels int
const (
  RenderingServerBakeChannelsBakeChannelAlbedoAlpha RenderingServerBakeChannels = 0
  RenderingServerBakeChannelsBakeChannelNormal RenderingServerBakeChannels = 1
  RenderingServerBakeChannelsBakeChannelOrm RenderingServerBakeChannels = 2
  RenderingServerBakeChannelsBakeChannelEmission RenderingServerBakeChannels = 3
)

type RenderingServerCanvasTextureChannel int
const (
  RenderingServerCanvasTextureChannelCanvasTextureChannelDiffuse RenderingServerCanvasTextureChannel = 0
  RenderingServerCanvasTextureChannelCanvasTextureChannelNormal RenderingServerCanvasTextureChannel = 1
  RenderingServerCanvasTextureChannelCanvasTextureChannelSpecular RenderingServerCanvasTextureChannel = 2
)

type RenderingServerNinePatchAxisMode int
const (
  RenderingServerNinePatchAxisModeNinePatchStretch RenderingServerNinePatchAxisMode = 0
  RenderingServerNinePatchAxisModeNinePatchTile RenderingServerNinePatchAxisMode = 1
  RenderingServerNinePatchAxisModeNinePatchTileFit RenderingServerNinePatchAxisMode = 2
)

type RenderingServerCanvasItemTextureFilter int
const (
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterDefault RenderingServerCanvasItemTextureFilter = 0
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterNearest RenderingServerCanvasItemTextureFilter = 1
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterLinear RenderingServerCanvasItemTextureFilter = 2
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterNearestWithMipmaps RenderingServerCanvasItemTextureFilter = 3
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterLinearWithMipmaps RenderingServerCanvasItemTextureFilter = 4
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterNearestWithMipmapsAnisotropic RenderingServerCanvasItemTextureFilter = 5
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterLinearWithMipmapsAnisotropic RenderingServerCanvasItemTextureFilter = 6
  RenderingServerCanvasItemTextureFilterCanvasItemTextureFilterMax RenderingServerCanvasItemTextureFilter = 7
)

type RenderingServerCanvasItemTextureRepeat int
const (
  RenderingServerCanvasItemTextureRepeatCanvasItemTextureRepeatDefault RenderingServerCanvasItemTextureRepeat = 0
  RenderingServerCanvasItemTextureRepeatCanvasItemTextureRepeatDisabled RenderingServerCanvasItemTextureRepeat = 1
  RenderingServerCanvasItemTextureRepeatCanvasItemTextureRepeatEnabled RenderingServerCanvasItemTextureRepeat = 2
  RenderingServerCanvasItemTextureRepeatCanvasItemTextureRepeatMirror RenderingServerCanvasItemTextureRepeat = 3
  RenderingServerCanvasItemTextureRepeatCanvasItemTextureRepeatMax RenderingServerCanvasItemTextureRepeat = 4
)

type RenderingServerCanvasGroupMode int
const (
  RenderingServerCanvasGroupModeCanvasGroupModeDisabled RenderingServerCanvasGroupMode = 0
  RenderingServerCanvasGroupModeCanvasGroupModeClipOnly RenderingServerCanvasGroupMode = 1
  RenderingServerCanvasGroupModeCanvasGroupModeClipAndDraw RenderingServerCanvasGroupMode = 2
  RenderingServerCanvasGroupModeCanvasGroupModeTransparent RenderingServerCanvasGroupMode = 3
)

type RenderingServerCanvasLightMode int
const (
  RenderingServerCanvasLightModeCanvasLightModePoint RenderingServerCanvasLightMode = 0
  RenderingServerCanvasLightModeCanvasLightModeDirectional RenderingServerCanvasLightMode = 1
)

type RenderingServerCanvasLightBlendMode int
const (
  RenderingServerCanvasLightBlendModeCanvasLightBlendModeAdd RenderingServerCanvasLightBlendMode = 0
  RenderingServerCanvasLightBlendModeCanvasLightBlendModeSub RenderingServerCanvasLightBlendMode = 1
  RenderingServerCanvasLightBlendModeCanvasLightBlendModeMix RenderingServerCanvasLightBlendMode = 2
)

type RenderingServerCanvasLightShadowFilter int
const (
  RenderingServerCanvasLightShadowFilterCanvasLightFilterNone RenderingServerCanvasLightShadowFilter = 0
  RenderingServerCanvasLightShadowFilterCanvasLightFilterPcf5 RenderingServerCanvasLightShadowFilter = 1
  RenderingServerCanvasLightShadowFilterCanvasLightFilterPcf13 RenderingServerCanvasLightShadowFilter = 2
  RenderingServerCanvasLightShadowFilterCanvasLightFilterMax RenderingServerCanvasLightShadowFilter = 3
)

type RenderingServerCanvasOccluderPolygonCullMode int
const (
  RenderingServerCanvasOccluderPolygonCullModeCanvasOccluderPolygonCullDisabled RenderingServerCanvasOccluderPolygonCullMode = 0
  RenderingServerCanvasOccluderPolygonCullModeCanvasOccluderPolygonCullClockwise RenderingServerCanvasOccluderPolygonCullMode = 1
  RenderingServerCanvasOccluderPolygonCullModeCanvasOccluderPolygonCullCounterClockwise RenderingServerCanvasOccluderPolygonCullMode = 2
)

type RenderingServerGlobalShaderParameterType int
const (
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeBool RenderingServerGlobalShaderParameterType = 0
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeBvec2 RenderingServerGlobalShaderParameterType = 1
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeBvec3 RenderingServerGlobalShaderParameterType = 2
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeBvec4 RenderingServerGlobalShaderParameterType = 3
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeInt RenderingServerGlobalShaderParameterType = 4
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeIvec2 RenderingServerGlobalShaderParameterType = 5
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeIvec3 RenderingServerGlobalShaderParameterType = 6
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeIvec4 RenderingServerGlobalShaderParameterType = 7
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeRect2I RenderingServerGlobalShaderParameterType = 8
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeUint RenderingServerGlobalShaderParameterType = 9
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeUvec2 RenderingServerGlobalShaderParameterType = 10
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeUvec3 RenderingServerGlobalShaderParameterType = 11
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeUvec4 RenderingServerGlobalShaderParameterType = 12
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeFloat RenderingServerGlobalShaderParameterType = 13
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeVec2 RenderingServerGlobalShaderParameterType = 14
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeVec3 RenderingServerGlobalShaderParameterType = 15
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeVec4 RenderingServerGlobalShaderParameterType = 16
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeColor RenderingServerGlobalShaderParameterType = 17
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeRect2 RenderingServerGlobalShaderParameterType = 18
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeMat2 RenderingServerGlobalShaderParameterType = 19
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeMat3 RenderingServerGlobalShaderParameterType = 20
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeMat4 RenderingServerGlobalShaderParameterType = 21
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeTransform2D RenderingServerGlobalShaderParameterType = 22
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeTransform RenderingServerGlobalShaderParameterType = 23
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeSampler2D RenderingServerGlobalShaderParameterType = 24
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeSampler2Darray RenderingServerGlobalShaderParameterType = 25
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeSampler3D RenderingServerGlobalShaderParameterType = 26
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeSamplercube RenderingServerGlobalShaderParameterType = 27
  RenderingServerGlobalShaderParameterTypeGlobalVarTypeMax RenderingServerGlobalShaderParameterType = 28
)

type RenderingServerRenderingInfo int
const (
  RenderingServerRenderingInfoRenderingInfoTotalObjectsInFrame RenderingServerRenderingInfo = 0
  RenderingServerRenderingInfoRenderingInfoTotalPrimitivesInFrame RenderingServerRenderingInfo = 1
  RenderingServerRenderingInfoRenderingInfoTotalDrawCallsInFrame RenderingServerRenderingInfo = 2
  RenderingServerRenderingInfoRenderingInfoTextureMemUsed RenderingServerRenderingInfo = 3
  RenderingServerRenderingInfoRenderingInfoBufferMemUsed RenderingServerRenderingInfo = 4
  RenderingServerRenderingInfoRenderingInfoVideoMemUsed RenderingServerRenderingInfo = 5
)

type RenderingServerFeatures int
const (
  RenderingServerFeaturesFeatureShaders RenderingServerFeatures = 0
  RenderingServerFeaturesFeatureMultithreaded RenderingServerFeatures = 1
)

func (me *RenderingServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RenderingServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RenderingServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RenderingServer) Texture2DCreate(image Image, ) RID {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture2DLayeredCreate(layers []Image, layered_type RenderingServerTextureLayeredType, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , gdc.ConstTypePtr(&layered_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layers)
  pinner.Pin(&layered_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DLayeredCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture3DCreate(format ImageFormat, width int64, height int64, depth int64, mipmaps bool, data []Image, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&depth) , gdc.ConstTypePtr(&mipmaps) , gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&format)
  pinner.Pin(&width)
  pinner.Pin(&height)
  pinner.Pin(&depth)
  pinner.Pin(&mipmaps)
  pinner.Pin(&data)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture3DCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) TextureProxyCreate(base RID, ) RID {
  cargs := []gdc.ConstTypePtr{base.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureProxyCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture2DUpdate(texture RID, image Image, layer int64, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), image.AsCTypePtr(), gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) Texture3DUpdate(texture RID, data []Image, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture3DUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) TextureProxyUpdate(texture RID, proxy_to RID, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), proxy_to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureProxyUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) Texture2DPlaceholderCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DPlaceholderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture2DLayeredPlaceholderCreate(layered_type RenderingServerTextureLayeredType, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layered_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layered_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DLayeredPlaceholderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture3DPlaceholderCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture3DPlaceholderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture2DGet(texture RID, ) Image {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture2DLayerGet(texture RID, layer int64, ) Image {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture2DLayerGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) Texture3DGet(texture RID, ) []Image {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTexture3DGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Image](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) TextureReplace(texture RID, by_texture RID, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), by_texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureReplace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) TextureSetSizeOverride(texture RID, width int64, height int64, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureSetSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) TextureSetPath(texture RID, path String, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureSetPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) TextureGetPath(texture RID, ) String {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) TextureGetFormat(texture RID, ) ImageFormat {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderingServer) TextureSetForceRedrawIfVisible(texture RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureSetForceRedrawIfVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) TextureRdCreate(rd_texture RID, layer_type RenderingServerTextureLayeredType, ) RID {
  cargs := []gdc.ConstTypePtr{rd_texture.AsCTypePtr(), gdc.ConstTypePtr(&layer_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&layer_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureRdCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) TextureGetRdTexture(texture RID, srgb bool, ) RID {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&srgb) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&srgb)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureGetRdTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) TextureGetNativeHandle(texture RID, srgb bool, ) int64 {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&srgb) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&srgb)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnTextureGetNativeHandle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ShaderCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ShaderSetCode(shader RID, code String, )  {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), code.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderSetCode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ShaderSetPathHint(shader RID, path String, )  {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderSetPathHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ShaderGetCode(shader RID, ) String {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderGetCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetShaderParameterList(shader RID, ) []Dictionary {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetShaderParameterList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) ShaderGetParameterDefault(shader RID, name StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderGetParameterDefault), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ShaderSetDefaultTextureParameter(shader RID, name StringName, texture RID, index int64, )  {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), name.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderSetDefaultTextureParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ShaderGetDefaultTextureParameter(shader RID, name StringName, index int64, ) RID {
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnShaderGetDefaultTextureParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MaterialCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MaterialSetShader(shader_material RID, shader RID, )  {
  cargs := []gdc.ConstTypePtr{shader_material.AsCTypePtr(), shader.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialSetShader), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MaterialSetParam(material RID, parameter StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), parameter.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MaterialGetParam(material RID, parameter StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), parameter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MaterialSetRenderPriority(material RID, priority int64, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialSetRenderPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MaterialSetNextPass(material RID, next_material RID, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), next_material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMaterialSetNextPass), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshCreateFromSurfaces(surfaces []Dictionary, blend_shape_count int64, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surfaces) , gdc.ConstTypePtr(&blend_shape_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&surfaces)
  pinner.Pin(&blend_shape_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshCreateFromSurfaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshSurfaceGetFormatOffset(format RenderingServerArrayFormat, vertex_count int64, array_index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&vertex_count) , gdc.ConstTypePtr(&array_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&format)
  pinner.Pin(&vertex_count)
  pinner.Pin(&array_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetFormatOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSurfaceGetFormatVertexStride(format RenderingServerArrayFormat, vertex_count int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&vertex_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&format)
  pinner.Pin(&vertex_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetFormatVertexStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSurfaceGetFormatNormalTangentStride(format RenderingServerArrayFormat, vertex_count int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&vertex_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&format)
  pinner.Pin(&vertex_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetFormatNormalTangentStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSurfaceGetFormatAttributeStride(format RenderingServerArrayFormat, vertex_count int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&vertex_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&format)
  pinner.Pin(&vertex_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetFormatAttributeStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSurfaceGetFormatSkinStride(format RenderingServerArrayFormat, vertex_count int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , gdc.ConstTypePtr(&vertex_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&format)
  pinner.Pin(&vertex_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetFormatSkinStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshAddSurface(mesh RID, surface Dictionary, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), surface.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshAddSurface), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshAddSurfaceFromArrays(mesh RID, primitive RenderingServerPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, compress_format RenderingServerArrayFormat, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&primitive) , arrays.AsCTypePtr(), blend_shapes.AsCTypePtr(), lods.AsCTypePtr(), gdc.ConstTypePtr(&compress_format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshAddSurfaceFromArrays), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshGetBlendShapeCount(mesh RID, ) int64 {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshGetBlendShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSetBlendShapeMode(mesh RID, mode RenderingServerBlendShapeMode, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSetBlendShapeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshGetBlendShapeMode(mesh RID, ) RenderingServerBlendShapeMode {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerBlendShapeMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshGetBlendShapeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderingServer) MeshSurfaceSetMaterial(mesh RID, surface int64, material RID, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshSurfaceGetMaterial(mesh RID, surface int64, ) RID {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshGetSurface(mesh RID, surface int64, ) Dictionary {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshGetSurface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshSurfaceGetArrays(mesh RID, surface int64, ) Array {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshSurfaceGetBlendShapeArrays(mesh RID, surface int64, ) []Array {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceGetBlendShapeArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Array](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) MeshGetSurfaceCount(mesh RID, ) int64 {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshGetSurfaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MeshSetCustomAabb(mesh RID, aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshGetCustomAabb(mesh RID, ) AABB {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshGetCustomAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MeshClear(mesh RID, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshSurfaceUpdateVertexRegion(mesh RID, surface int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceUpdateVertexRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshSurfaceUpdateAttributeRegion(mesh RID, surface int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceUpdateAttributeRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshSurfaceUpdateSkinRegion(mesh RID, surface int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSurfaceUpdateSkinRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MeshSetShadowMesh(mesh RID, shadow_mesh RID, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), shadow_mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMeshSetShadowMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshAllocateData(multimesh RID, instances int64, transform_format RenderingServerMultimeshTransformFormat, color_format bool, custom_data_format bool, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&instances) , gdc.ConstTypePtr(&transform_format) , gdc.ConstTypePtr(&color_format) , gdc.ConstTypePtr(&custom_data_format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshAllocateData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshGetInstanceCount(multimesh RID, ) int64 {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshGetInstanceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MultimeshSetMesh(multimesh RID, mesh RID, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshInstanceSetTransform(multimesh RID, index int64, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshInstanceSetTransform2D(multimesh RID, index int64, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceSetTransform2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshInstanceSetColor(multimesh RID, index int64, color Color, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshInstanceSetCustomData(multimesh RID, index int64, custom_data Color, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , custom_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceSetCustomData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshGetMesh(multimesh RID, ) RID {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshGetAabb(multimesh RID, ) AABB {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshGetAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshInstanceGetTransform(multimesh RID, index int64, ) Transform3D {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshInstanceGetTransform2D(multimesh RID, index int64, ) Transform2D {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceGetTransform2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshInstanceGetColor(multimesh RID, index int64, ) Color {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshInstanceGetCustomData(multimesh RID, index int64, ) Color {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshInstanceGetCustomData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MultimeshSetVisibleInstances(multimesh RID, visible int64, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshSetVisibleInstances), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshGetVisibleInstances(multimesh RID, ) int64 {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshGetVisibleInstances), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) MultimeshSetBuffer(multimesh RID, buffer PackedFloat32Array, )  {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshSetBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) MultimeshGetBuffer(multimesh RID, ) PackedFloat32Array {
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMultimeshGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SkeletonCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SkeletonAllocateData(skeleton RID, bones int64, is_2d_skeleton bool, )  {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), gdc.ConstTypePtr(&bones) , gdc.ConstTypePtr(&is_2d_skeleton) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonAllocateData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkeletonGetBoneCount(skeleton RID, ) int64 {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonGetBoneCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) SkeletonBoneSetTransform(skeleton RID, bone int64, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), gdc.ConstTypePtr(&bone) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonBoneSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkeletonBoneGetTransform(skeleton RID, bone int64, ) Transform3D {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), gdc.ConstTypePtr(&bone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonBoneGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SkeletonBoneSetTransform2D(skeleton RID, bone int64, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), gdc.ConstTypePtr(&bone) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonBoneSetTransform2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkeletonBoneGetTransform2D(skeleton RID, bone int64, ) Transform2D {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), gdc.ConstTypePtr(&bone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()
  pinner.Pin(&bone)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonBoneGetTransform2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SkeletonSetBaseTransform2D(skeleton RID, base_transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{skeleton.AsCTypePtr(), base_transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkeletonSetBaseTransform2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DirectionalLightCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDirectionalLightCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) OmniLightCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnOmniLightCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SpotLightCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSpotLightCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightSetColor(light RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetParam(light RID, param RenderingServerLightParam, value float64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetShadow(light RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetShadow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetProjector(light RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetProjector), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetNegative(light RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetNegative), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetCullMask(light RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetDistanceFade(decal RID, enabled bool, begin float64, shadow float64, length float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , gdc.ConstTypePtr(&begin) , gdc.ConstTypePtr(&shadow) , gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetDistanceFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetReverseCullFaceMode(light RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetReverseCullFaceMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetBakeMode(light RID, bake_mode RenderingServerLightBakeMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&bake_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetBakeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightSetMaxSdfgiCascade(light RID, cascade int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&cascade) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightSetMaxSdfgiCascade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightOmniSetShadowMode(light RID, mode RenderingServerLightOmniShadowMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightOmniSetShadowMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightDirectionalSetShadowMode(light RID, mode RenderingServerLightDirectionalShadowMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightDirectionalSetShadowMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightDirectionalSetBlendSplits(light RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightDirectionalSetBlendSplits), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightDirectionalSetSkyMode(light RID, mode RenderingServerLightDirectionalSkyMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightDirectionalSetSkyMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightProjectorsSetFilter(filter RenderingServerLightProjectorFilter, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightProjectorsSetFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) PositionalSoftShadowFilterSetQuality(quality RenderingServerShadowQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnPositionalSoftShadowFilterSetQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DirectionalSoftShadowFilterSetQuality(quality RenderingServerShadowQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDirectionalSoftShadowFilterSetQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DirectionalShadowAtlasSetSize(size int64, is_16bits bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , gdc.ConstTypePtr(&is_16bits) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDirectionalShadowAtlasSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ReflectionProbeSetUpdateMode(probe RID, mode RenderingServerReflectionProbeUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetIntensity(probe RID, intensity float64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetAmbientMode(probe RID, mode RenderingServerReflectionProbeAmbientMode, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetAmbientMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetAmbientColor(probe RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetAmbientColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetAmbientEnergy(probe RID, energy float64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetAmbientEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetMaxDistance(probe RID, distance float64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetSize(probe RID, size Vector3, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetOriginOffset(probe RID, offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetOriginOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetAsInterior(probe RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetAsInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetEnableBoxProjection(probe RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetEnableBoxProjection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetEnableShadows(probe RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetEnableShadows), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetCullMask(probe RID, layers int64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetResolution(probe RID, resolution int64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ReflectionProbeSetMeshLodThreshold(probe RID, pixels float64, )  {
  cargs := []gdc.ConstTypePtr{probe.AsCTypePtr(), gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnReflectionProbeSetMeshLodThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) DecalSetSize(decal RID, size Vector3, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetTexture(decal RID, type_ RenderingServerDecalTexture, texture RID, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&type_) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetEmissionEnergy(decal RID, energy float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetEmissionEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetAlbedoMix(decal RID, albedo_mix float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&albedo_mix) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetAlbedoMix), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetModulate(decal RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetCullMask(decal RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetDistanceFade(decal RID, enabled bool, begin float64, length float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , gdc.ConstTypePtr(&begin) , gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetDistanceFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetFade(decal RID, above float64, below float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&above) , gdc.ConstTypePtr(&below) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalSetNormalFade(decal RID, fade float64, )  {
  cargs := []gdc.ConstTypePtr{decal.AsCTypePtr(), gdc.ConstTypePtr(&fade) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalSetNormalFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) DecalsSetFilter(filter RenderingServerDecalFilter, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnDecalsSetFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GiSetUseHalfResolution(half_resolution bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&half_resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGiSetUseHalfResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiAllocateData(voxel_gi RID, to_cell_xform Transform3D, aabb AABB, octree_size Vector3i, octree_cells PackedByteArray, data_cells PackedByteArray, distance_field PackedByteArray, level_counts PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), to_cell_xform.AsCTypePtr(), aabb.AsCTypePtr(), octree_size.AsCTypePtr(), octree_cells.AsCTypePtr(), data_cells.AsCTypePtr(), distance_field.AsCTypePtr(), level_counts.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiAllocateData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiGetOctreeSize(voxel_gi RID, ) Vector3i {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetOctreeSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiGetOctreeCells(voxel_gi RID, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetOctreeCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiGetDataCells(voxel_gi RID, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetDataCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiGetDistanceField(voxel_gi RID, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetDistanceField), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiGetLevelCounts(voxel_gi RID, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetLevelCounts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiGetToCellXform(voxel_gi RID, ) Transform3D {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiGetToCellXform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VoxelGiSetDynamicRange(voxel_gi RID, range_ float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&range_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetDynamicRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetPropagation(voxel_gi RID, amount float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetPropagation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetEnergy(voxel_gi RID, energy float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetBakedExposureNormalization(voxel_gi RID, baked_exposure float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&baked_exposure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetBakedExposureNormalization), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetBias(voxel_gi RID, bias float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetNormalBias(voxel_gi RID, bias float64, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetNormalBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetInterior(voxel_gi RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetUseTwoBounces(voxel_gi RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{voxel_gi.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetUseTwoBounces), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VoxelGiSetQuality(quality RenderingServerVoxelGIQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVoxelGiSetQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightmapSetTextures(lightmap RID, light RID, uses_sh bool, )  {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), light.AsCTypePtr(), gdc.ConstTypePtr(&uses_sh) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetTextures), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapSetProbeBounds(lightmap RID, bounds AABB, )  {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), bounds.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetProbeBounds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapSetProbeInterior(lightmap RID, interior bool, )  {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), gdc.ConstTypePtr(&interior) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetProbeInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapSetProbeCaptureData(lightmap RID, points PackedVector3Array, point_sh PackedColorArray, tetrahedra PackedInt32Array, bsp_tree PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), points.AsCTypePtr(), point_sh.AsCTypePtr(), tetrahedra.AsCTypePtr(), bsp_tree.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetProbeCaptureData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapGetProbeCapturePoints(lightmap RID, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapGetProbeCapturePoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightmapGetProbeCaptureSh(lightmap RID, ) PackedColorArray {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapGetProbeCaptureSh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightmapGetProbeCaptureTetrahedra(lightmap RID, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapGetProbeCaptureTetrahedra), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightmapGetProbeCaptureBspTree(lightmap RID, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapGetProbeCaptureBspTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) LightmapSetBakedExposureNormalization(lightmap RID, baked_exposure float64, )  {
  cargs := []gdc.ConstTypePtr{lightmap.AsCTypePtr(), gdc.ConstTypePtr(&baked_exposure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetBakedExposureNormalization), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) LightmapSetProbeCaptureUpdateSpeed(speed float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnLightmapSetProbeCaptureUpdateSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ParticlesSetMode(particles RID, mode RenderingServerParticlesMode, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetEmitting(particles RID, emitting bool, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&emitting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetEmitting), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesGetEmitting(particles RID, ) bool {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesGetEmitting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ParticlesSetAmount(particles RID, amount int64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetAmountRatio(particles RID, ratio float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetAmountRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetLifetime(particles RID, lifetime float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&lifetime) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetOneShot(particles RID, one_shot bool, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&one_shot) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetPreProcessTime(particles RID, time float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetPreProcessTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetExplosivenessRatio(particles RID, ratio float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetRandomnessRatio(particles RID, ratio float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetRandomnessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetInterpToEnd(particles RID, factor float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetInterpToEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetEmitterVelocity(particles RID, velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetEmitterVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetCustomAabb(particles RID, aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetSpeedScale(particles RID, scale float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetUseLocalCoordinates(particles RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetProcessMaterial(particles RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetProcessMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetFixedFps(particles RID, fps int64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&fps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetFixedFps), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetInterpolate(particles RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetInterpolate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetFractionalDelta(particles RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetFractionalDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetCollisionBaseSize(particles RID, size float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetCollisionBaseSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetTransformAlign(particles RID, align RenderingServerParticlesTransformAlign, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&align) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetTransformAlign), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetTrails(particles RID, enable bool, length_sec float64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&length_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetTrails), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetTrailBindPoses(particles RID, bind_poses []Transform3D, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&bind_poses) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetTrailBindPoses), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesIsInactive(particles RID, ) bool {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesIsInactive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ParticlesRequestProcess(particles RID, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesRequestProcess), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesRestart(particles RID, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesRestart), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetSubemitter(particles RID, subemitter_particles RID, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), subemitter_particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetSubemitter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesEmit(particles RID, transform Transform3D, velocity Vector3, color Color, custom Color, emit_flags int64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), transform.AsCTypePtr(), velocity.AsCTypePtr(), color.AsCTypePtr(), custom.AsCTypePtr(), gdc.ConstTypePtr(&emit_flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesEmit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetDrawOrder(particles RID, order RenderingServerParticlesDrawOrder, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetDrawOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetDrawPasses(particles RID, count int64, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetDrawPasses), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesSetDrawPassMesh(particles RID, pass int64, mesh RID, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), gdc.ConstTypePtr(&pass) , mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetDrawPassMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesGetCurrentAabb(particles RID, ) AABB {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesGetCurrentAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ParticlesSetEmissionTransform(particles RID, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesSetEmissionTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ParticlesCollisionSetCollisionType(particles_collision RID, type_ RenderingServerParticlesCollisionType, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetCollisionType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetCullMask(particles_collision RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetSphereRadius(particles_collision RID, radius float64, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetSphereRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetBoxExtents(particles_collision RID, extents Vector3, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), extents.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetBoxExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetAttractorStrength(particles_collision RID, strength float64, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetAttractorStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetAttractorDirectionality(particles_collision RID, amount float64, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetAttractorDirectionality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetAttractorAttenuation(particles_collision RID, curve float64, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetAttractorAttenuation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetFieldTexture(particles_collision RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetFieldTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionHeightFieldUpdate(particles_collision RID, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionHeightFieldUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ParticlesCollisionSetHeightFieldResolution(particles_collision RID, resolution RenderingServerParticlesCollisionHeightfieldResolution, )  {
  cargs := []gdc.ConstTypePtr{particles_collision.AsCTypePtr(), gdc.ConstTypePtr(&resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnParticlesCollisionSetHeightFieldResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) FogVolumeCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnFogVolumeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) FogVolumeSetShape(fog_volume RID, shape RenderingServerFogVolumeShape, )  {
  cargs := []gdc.ConstTypePtr{fog_volume.AsCTypePtr(), gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnFogVolumeSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) FogVolumeSetSize(fog_volume RID, size Vector3, )  {
  cargs := []gdc.ConstTypePtr{fog_volume.AsCTypePtr(), size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnFogVolumeSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) FogVolumeSetMaterial(fog_volume RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{fog_volume.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnFogVolumeSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VisibilityNotifierCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVisibilityNotifierCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) VisibilityNotifierSetAabb(notifier RID, aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{notifier.AsCTypePtr(), aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVisibilityNotifierSetAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) VisibilityNotifierSetCallbacks(notifier RID, enter_callable Callable, exit_callable Callable, )  {
  cargs := []gdc.ConstTypePtr{notifier.AsCTypePtr(), enter_callable.AsCTypePtr(), exit_callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnVisibilityNotifierSetCallbacks), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) OccluderCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnOccluderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) OccluderSetMesh(occluder RID, vertices PackedVector3Array, indices PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), vertices.AsCTypePtr(), indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnOccluderSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CameraSetPerspective(camera RID, fovy_degrees float64, z_near float64, z_far float64, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), gdc.ConstTypePtr(&fovy_degrees) , gdc.ConstTypePtr(&z_near) , gdc.ConstTypePtr(&z_far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetPerspective), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetOrthogonal(camera RID, size float64, z_near float64, z_far float64, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), gdc.ConstTypePtr(&size) , gdc.ConstTypePtr(&z_near) , gdc.ConstTypePtr(&z_far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetOrthogonal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetFrustum(camera RID, size float64, offset Vector2, z_near float64, z_far float64, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), gdc.ConstTypePtr(&size) , offset.AsCTypePtr(), gdc.ConstTypePtr(&z_near) , gdc.ConstTypePtr(&z_far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetFrustum), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetTransform(camera RID, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetCullMask(camera RID, layers int64, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetEnvironment(camera RID, env RID, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), env.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetCameraAttributes(camera RID, effects RID, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), effects.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraSetUseVerticalAspect(camera RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{camera.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraSetUseVerticalAspect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ViewportSetUseXr(viewport RID, use_xr bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&use_xr) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUseXr), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetSize(viewport RID, width int64, height int64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetActive(viewport RID, active bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetParentViewport(viewport RID, parent_viewport RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), parent_viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetParentViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportAttachToScreen(viewport RID, rect Rect2, screen int64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), rect.AsCTypePtr(), gdc.ConstTypePtr(&screen) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportAttachToScreen), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetRenderDirectToScreen(viewport RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetRenderDirectToScreen), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetCanvasCullMask(viewport RID, canvas_cull_mask int64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&canvas_cull_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetCanvasCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetScaling3DMode(viewport RID, scaling_3d_mode RenderingServerViewportScaling3DMode, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&scaling_3d_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetScaling3DMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetScaling3DScale(viewport RID, scale float64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetScaling3DScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetFsrSharpness(viewport RID, sharpness float64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetFsrSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetTextureMipmapBias(viewport RID, mipmap_bias float64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&mipmap_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetTextureMipmapBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetUpdateMode(viewport RID, update_mode RenderingServerViewportUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&update_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetClearMode(viewport RID, clear_mode RenderingServerViewportClearMode, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&clear_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetClearMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportGetRenderTarget(viewport RID, ) RID {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportGetRenderTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ViewportGetTexture(viewport RID, ) RID {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ViewportSetDisable3D(viewport RID, disable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetDisable3D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetDisable2D(viewport RID, disable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetDisable2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetEnvironmentMode(viewport RID, mode RenderingServerViewportEnvironmentMode, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetEnvironmentMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportAttachCamera(viewport RID, camera RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), camera.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportAttachCamera), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetScenario(viewport RID, scenario RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetScenario), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportAttachCanvas(viewport RID, canvas RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), canvas.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportAttachCanvas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportRemoveCanvas(viewport RID, canvas RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), canvas.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportRemoveCanvas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetSnap2DTransformsToPixel(viewport RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetSnap2DTransformsToPixel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetSnap2DVerticesToPixel(viewport RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetSnap2DVerticesToPixel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetDefaultCanvasItemTextureFilter(viewport RID, filter RenderingServerCanvasItemTextureFilter, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetDefaultCanvasItemTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetDefaultCanvasItemTextureRepeat(viewport RID, repeat RenderingServerCanvasItemTextureRepeat, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&repeat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetDefaultCanvasItemTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetCanvasTransform(viewport RID, canvas RID, offset Transform2D, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), canvas.AsCTypePtr(), offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetCanvasTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetCanvasStacking(viewport RID, canvas RID, layer int64, sublayer int64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), canvas.AsCTypePtr(), gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&sublayer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetCanvasStacking), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetTransparentBackground(viewport RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetTransparentBackground), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetGlobalCanvasTransform(viewport RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetGlobalCanvasTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetSdfOversizeAndScale(viewport RID, oversize RenderingServerViewportSDFOversize, scale RenderingServerViewportSDFScale, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&oversize) , gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetSdfOversizeAndScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetPositionalShadowAtlasSize(viewport RID, size int64, use_16_bits bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&size) , gdc.ConstTypePtr(&use_16_bits) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetPositionalShadowAtlasSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetPositionalShadowAtlasQuadrantSubdivision(viewport RID, quadrant int64, subdivision int64, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&quadrant) , gdc.ConstTypePtr(&subdivision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetPositionalShadowAtlasQuadrantSubdivision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetMsaa3D(viewport RID, msaa RenderingServerViewportMSAA, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&msaa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetMsaa3D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetMsaa2D(viewport RID, msaa RenderingServerViewportMSAA, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&msaa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetMsaa2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetUseHdr2D(viewport RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUseHdr2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetScreenSpaceAa(viewport RID, mode RenderingServerViewportScreenSpaceAA, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetScreenSpaceAa), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetUseTaa(viewport RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUseTaa), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetUseDebanding(viewport RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUseDebanding), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetUseOcclusionCulling(viewport RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetUseOcclusionCulling), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetOcclusionRaysPerThread(rays_per_thread int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rays_per_thread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetOcclusionRaysPerThread), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetOcclusionCullingBuildQuality(quality RenderingServerViewportOcclusionCullingBuildQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetOcclusionCullingBuildQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportGetRenderInfo(viewport RID, type_ RenderingServerViewportRenderInfoType, info RenderingServerViewportRenderInfo, ) int64 {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&info) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)
  pinner.Pin(&info)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportGetRenderInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ViewportSetDebugDraw(viewport RID, draw RenderingServerViewportDebugDraw, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&draw) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetDebugDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetMeasureRenderTime(viewport RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetMeasureRenderTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportGetMeasuredRenderTimeCpu(viewport RID, ) float64 {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportGetMeasuredRenderTimeCpu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ViewportGetMeasuredRenderTimeGpu(viewport RID, ) float64 {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportGetMeasuredRenderTimeGpu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ViewportSetVrsMode(viewport RID, mode RenderingServerViewportVRSMode, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetVrsMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ViewportSetVrsTexture(viewport RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnViewportSetVrsTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkyCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkyCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SkySetRadianceSize(sky RID, radiance_size int64, )  {
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), gdc.ConstTypePtr(&radiance_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkySetRadianceSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkySetMode(sky RID, mode RenderingServerSkyMode, )  {
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkySetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkySetMaterial(sky RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkySetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SkyBakePanorama(sky RID, energy float64, bake_irradiance bool, size Vector2i, ) Image {
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), gdc.ConstTypePtr(&energy) , gdc.ConstTypePtr(&bake_irradiance) , size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()
  pinner.Pin(&energy)
  pinner.Pin(&bake_irradiance)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSkyBakePanorama), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) EnvironmentCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) EnvironmentSetBackground(env RID, bg RenderingServerEnvironmentBG, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&bg) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetBackground), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSky(env RID, sky RID, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), sky.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSky), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSkyCustomFov(env RID, scale float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSkyCustomFov), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSkyOrientation(env RID, orientation Basis, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), orientation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSkyOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetBgColor(env RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetBgEnergy(env RID, multiplier float64, exposure_value float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&multiplier) , gdc.ConstTypePtr(&exposure_value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetBgEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetCanvasMaxLayer(env RID, max_layer int64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&max_layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetCanvasMaxLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetAmbientLight(env RID, color Color, ambient RenderingServerEnvironmentAmbientSource, energy float64, sky_contibution float64, reflection_source RenderingServerEnvironmentReflectionSource, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&ambient) , gdc.ConstTypePtr(&energy) , gdc.ConstTypePtr(&sky_contibution) , gdc.ConstTypePtr(&reflection_source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetAmbientLight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetGlow(env RID, enable bool, levels PackedFloat32Array, intensity float64, strength float64, mix float64, bloom_threshold float64, blend_mode RenderingServerEnvironmentGlowBlendMode, hdr_bleed_threshold float64, hdr_bleed_scale float64, hdr_luminance_cap float64, glow_map_strength float64, glow_map RID, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , levels.AsCTypePtr(), gdc.ConstTypePtr(&intensity) , gdc.ConstTypePtr(&strength) , gdc.ConstTypePtr(&mix) , gdc.ConstTypePtr(&bloom_threshold) , gdc.ConstTypePtr(&blend_mode) , gdc.ConstTypePtr(&hdr_bleed_threshold) , gdc.ConstTypePtr(&hdr_bleed_scale) , gdc.ConstTypePtr(&hdr_luminance_cap) , gdc.ConstTypePtr(&glow_map_strength) , glow_map.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetGlow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetTonemap(env RID, tone_mapper RenderingServerEnvironmentToneMapper, exposure float64, white float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&tone_mapper) , gdc.ConstTypePtr(&exposure) , gdc.ConstTypePtr(&white) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetTonemap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetAdjustment(env RID, enable bool, brightness float64, contrast float64, saturation float64, use_1d_color_correction bool, color_correction RID, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&brightness) , gdc.ConstTypePtr(&contrast) , gdc.ConstTypePtr(&saturation) , gdc.ConstTypePtr(&use_1d_color_correction) , color_correction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetAdjustment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSsr(env RID, enable bool, max_steps int64, fade_in float64, fade_out float64, depth_tolerance float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&max_steps) , gdc.ConstTypePtr(&fade_in) , gdc.ConstTypePtr(&fade_out) , gdc.ConstTypePtr(&depth_tolerance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSsr), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSsao(env RID, enable bool, radius float64, intensity float64, power float64, detail float64, horizon float64, sharpness float64, light_affect float64, ao_channel_affect float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&radius) , gdc.ConstTypePtr(&intensity) , gdc.ConstTypePtr(&power) , gdc.ConstTypePtr(&detail) , gdc.ConstTypePtr(&horizon) , gdc.ConstTypePtr(&sharpness) , gdc.ConstTypePtr(&light_affect) , gdc.ConstTypePtr(&ao_channel_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSsao), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetFog(env RID, enable bool, light_color Color, light_energy float64, sun_scatter float64, density float64, height float64, height_density float64, aerial_perspective float64, sky_affect float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , light_color.AsCTypePtr(), gdc.ConstTypePtr(&light_energy) , gdc.ConstTypePtr(&sun_scatter) , gdc.ConstTypePtr(&density) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&height_density) , gdc.ConstTypePtr(&aerial_perspective) , gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetFog), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSdfgi(env RID, enable bool, cascades int64, min_cell_size float64, y_scale RenderingServerEnvironmentSDFGIYScale, use_occlusion bool, bounce_feedback float64, read_sky bool, energy float64, normal_bias float64, probe_bias float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&cascades) , gdc.ConstTypePtr(&min_cell_size) , gdc.ConstTypePtr(&y_scale) , gdc.ConstTypePtr(&use_occlusion) , gdc.ConstTypePtr(&bounce_feedback) , gdc.ConstTypePtr(&read_sky) , gdc.ConstTypePtr(&energy) , gdc.ConstTypePtr(&normal_bias) , gdc.ConstTypePtr(&probe_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSdfgi), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetVolumetricFog(env RID, enable bool, density float64, albedo Color, emission Color, emission_energy float64, anisotropy float64, length float64, p_detail_spread float64, gi_inject float64, temporal_reprojection bool, temporal_reprojection_amount float64, ambient_inject float64, sky_affect float64, )  {
  cargs := []gdc.ConstTypePtr{env.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&density) , albedo.AsCTypePtr(), emission.AsCTypePtr(), gdc.ConstTypePtr(&emission_energy) , gdc.ConstTypePtr(&anisotropy) , gdc.ConstTypePtr(&length) , gdc.ConstTypePtr(&p_detail_spread) , gdc.ConstTypePtr(&gi_inject) , gdc.ConstTypePtr(&temporal_reprojection) , gdc.ConstTypePtr(&temporal_reprojection_amount) , gdc.ConstTypePtr(&ambient_inject) , gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetVolumetricFog), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentGlowSetUseBicubicUpscale(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentGlowSetUseBicubicUpscale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSsrRoughnessQuality(quality RenderingServerEnvironmentSSRRoughnessQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSsrRoughnessQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSsaoQuality(quality RenderingServerEnvironmentSSAOQuality, half_size bool, adaptive_target float64, blur_passes int64, fadeout_from float64, fadeout_to float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , gdc.ConstTypePtr(&half_size) , gdc.ConstTypePtr(&adaptive_target) , gdc.ConstTypePtr(&blur_passes) , gdc.ConstTypePtr(&fadeout_from) , gdc.ConstTypePtr(&fadeout_to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSsaoQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSsilQuality(quality RenderingServerEnvironmentSSILQuality, half_size bool, adaptive_target float64, blur_passes int64, fadeout_from float64, fadeout_to float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , gdc.ConstTypePtr(&half_size) , gdc.ConstTypePtr(&adaptive_target) , gdc.ConstTypePtr(&blur_passes) , gdc.ConstTypePtr(&fadeout_from) , gdc.ConstTypePtr(&fadeout_to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSsilQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSdfgiRayCount(ray_count RenderingServerEnvironmentSDFGIRayCount, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSdfgiRayCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSdfgiFramesToConverge(frames RenderingServerEnvironmentSDFGIFramesToConverge, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSdfgiFramesToConverge), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetSdfgiFramesToUpdateLight(frames RenderingServerEnvironmentSDFGIFramesToUpdateLight, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetSdfgiFramesToUpdateLight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetVolumetricFogVolumeSize(size int64, depth int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , gdc.ConstTypePtr(&depth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetVolumetricFogVolumeSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentSetVolumetricFogFilterActive(active bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentSetVolumetricFogFilterActive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) EnvironmentBakePanorama(environment RID, bake_irradiance bool, size Vector2i, ) Image {
  cargs := []gdc.ConstTypePtr{environment.AsCTypePtr(), gdc.ConstTypePtr(&bake_irradiance) , size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()
  pinner.Pin(&bake_irradiance)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnEnvironmentBakePanorama), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ScreenSpaceRoughnessLimiterSetActive(enable bool, amount float64, limit float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&amount) , gdc.ConstTypePtr(&limit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnScreenSpaceRoughnessLimiterSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SubSurfaceScatteringSetQuality(quality RenderingServerSubSurfaceScatteringQuality, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSubSurfaceScatteringSetQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) SubSurfaceScatteringSetScale(scale float64, depth_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , gdc.ConstTypePtr(&depth_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSubSurfaceScatteringSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraAttributesCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CameraAttributesSetDofBlurQuality(quality RenderingServerDOFBlurQuality, use_jitter bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , gdc.ConstTypePtr(&use_jitter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesSetDofBlurQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraAttributesSetDofBlurBokehShape(shape RenderingServerDOFBokehShape, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesSetDofBlurBokehShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraAttributesSetDofBlur(camera_attributes RID, far_enable bool, far_distance float64, far_transition float64, near_enable bool, near_distance float64, near_transition float64, amount float64, )  {
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), gdc.ConstTypePtr(&far_enable) , gdc.ConstTypePtr(&far_distance) , gdc.ConstTypePtr(&far_transition) , gdc.ConstTypePtr(&near_enable) , gdc.ConstTypePtr(&near_distance) , gdc.ConstTypePtr(&near_transition) , gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesSetDofBlur), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraAttributesSetExposure(camera_attributes RID, multiplier float64, normalization float64, )  {
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), gdc.ConstTypePtr(&multiplier) , gdc.ConstTypePtr(&normalization) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesSetExposure), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CameraAttributesSetAutoExposure(camera_attributes RID, enable bool, min_sensitivity float64, max_sensitivity float64, speed float64, scale float64, )  {
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), gdc.ConstTypePtr(&enable) , gdc.ConstTypePtr(&min_sensitivity) , gdc.ConstTypePtr(&max_sensitivity) , gdc.ConstTypePtr(&speed) , gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCameraAttributesSetAutoExposure), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ScenarioCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnScenarioCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) ScenarioSetEnvironment(scenario RID, environment RID, )  {
  cargs := []gdc.ConstTypePtr{scenario.AsCTypePtr(), environment.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnScenarioSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ScenarioSetFallbackEnvironment(scenario RID, environment RID, )  {
  cargs := []gdc.ConstTypePtr{scenario.AsCTypePtr(), environment.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnScenarioSetFallbackEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ScenarioSetCameraAttributes(scenario RID, effects RID, )  {
  cargs := []gdc.ConstTypePtr{scenario.AsCTypePtr(), effects.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnScenarioSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceCreate2(base RID, scenario RID, ) RID {
  cargs := []gdc.ConstTypePtr{base.AsCTypePtr(), scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceCreate2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstanceCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstanceSetBase(instance RID, base RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), base.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetBase), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetScenario(instance RID, scenario RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetScenario), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetLayerMask(instance RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetLayerMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetPivotData(instance RID, sorting_offset float64, use_aabb_center bool, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&sorting_offset) , gdc.ConstTypePtr(&use_aabb_center) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetPivotData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetTransform(instance RID, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceAttachObjectInstanceId(instance RID, id int64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceAttachObjectInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetBlendShapeWeight(instance RID, shape int64, weight float64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&shape) , gdc.ConstTypePtr(&weight) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetBlendShapeWeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetSurfaceOverrideMaterial(instance RID, surface int64, material RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&surface) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetSurfaceOverrideMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetVisible(instance RID, visible bool, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetTransparency(instance RID, transparency float64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&transparency) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetTransparency), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetCustomAabb(instance RID, aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceAttachSkeleton(instance RID, skeleton RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceAttachSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetExtraVisibilityMargin(instance RID, margin float64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetExtraVisibilityMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetVisibilityParent(instance RID, parent RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), parent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetVisibilityParent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceSetIgnoreCulling(instance RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceSetIgnoreCulling), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetFlag(instance RID, flag RenderingServerInstanceFlags, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&flag) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetCastShadowsSetting(instance RID, shadow_casting_setting RenderingServerShadowCastingSetting, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&shadow_casting_setting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetCastShadowsSetting), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetMaterialOverride(instance RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetMaterialOverlay(instance RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetMaterialOverlay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetVisibilityRange(instance RID, min float64, max float64, min_margin float64, max_margin float64, fade_mode RenderingServerVisibilityRangeFadeMode, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&min) , gdc.ConstTypePtr(&max) , gdc.ConstTypePtr(&min_margin) , gdc.ConstTypePtr(&max_margin) , gdc.ConstTypePtr(&fade_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetVisibilityRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetLightmap(instance RID, lightmap RID, lightmap_uv_scale Rect2, lightmap_slice int64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), lightmap.AsCTypePtr(), lightmap_uv_scale.AsCTypePtr(), gdc.ConstTypePtr(&lightmap_slice) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetLightmap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetLodBias(instance RID, lod_bias float64, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), gdc.ConstTypePtr(&lod_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetLodBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometrySetShaderParameter(instance RID, parameter StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), parameter.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometrySetShaderParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) InstanceGeometryGetShaderParameter(instance RID, parameter StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), parameter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometryGetShaderParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstanceGeometryGetShaderParameterDefaultValue(instance RID, parameter StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), parameter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometryGetShaderParameterDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstanceGeometryGetShaderParameterList(instance RID, ) []Dictionary {
  cargs := []gdc.ConstTypePtr{instance.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstanceGeometryGetShaderParameterList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) InstancesCullAabb(aabb AABB, scenario RID, ) PackedInt64Array {
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstancesCullAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstancesCullRay(from Vector3, to Vector3, scenario RID, ) PackedInt64Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstancesCullRay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) InstancesCullConvex(convex []Plane, scenario RID, ) PackedInt64Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex) , scenario.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()
  pinner.Pin(&convex)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnInstancesCullConvex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) BakeRenderUv2(base RID, material_overrides []RID, image_size Vector2i, ) []Image {
  cargs := []gdc.ConstTypePtr{base.AsCTypePtr(), gdc.ConstTypePtr(&material_overrides) , image_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&material_overrides)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnBakeRenderUv2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Image](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) CanvasCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasSetItemMirroring(canvas RID, item RID, mirroring Vector2, )  {
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), item.AsCTypePtr(), mirroring.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasSetItemMirroring), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasSetModulate(canvas RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasSetDisableScale(disable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasSetDisableScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasTextureCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasTextureCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasTextureSetChannel(canvas_texture RID, channel RenderingServerCanvasTextureChannel, texture RID, )  {
  cargs := []gdc.ConstTypePtr{canvas_texture.AsCTypePtr(), gdc.ConstTypePtr(&channel) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasTextureSetChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasTextureSetShadingParameters(canvas_texture RID, base_color Color, shininess float64, )  {
  cargs := []gdc.ConstTypePtr{canvas_texture.AsCTypePtr(), base_color.AsCTypePtr(), gdc.ConstTypePtr(&shininess) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasTextureSetShadingParameters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasTextureSetTextureFilter(canvas_texture RID, filter RenderingServerCanvasItemTextureFilter, )  {
  cargs := []gdc.ConstTypePtr{canvas_texture.AsCTypePtr(), gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasTextureSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasTextureSetTextureRepeat(canvas_texture RID, repeat RenderingServerCanvasItemTextureRepeat, )  {
  cargs := []gdc.ConstTypePtr{canvas_texture.AsCTypePtr(), gdc.ConstTypePtr(&repeat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasTextureSetTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasItemSetParent(item RID, parent RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), parent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetParent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetDefaultTextureFilter(item RID, filter RenderingServerCanvasItemTextureFilter, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetDefaultTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetDefaultTextureRepeat(item RID, repeat RenderingServerCanvasItemTextureRepeat, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&repeat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetDefaultTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetVisible(item RID, visible bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetLightMask(item RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetLightMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetVisibilityLayer(item RID, visibility_layer int64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&visibility_layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetVisibilityLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetTransform(item RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetClip(item RID, clip bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&clip) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetClip), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetDistanceFieldMode(item RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetDistanceFieldMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetCustomRect(item RID, use_custom_rect bool, rect Rect2, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&use_custom_rect) , rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetCustomRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetModulate(item RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetSelfModulate(item RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetSelfModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetDrawBehindParent(item RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetDrawBehindParent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddLine(item RID, from Vector2, to Vector2, color Color, width float64, antialiased bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), from.AsCTypePtr(), to.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddPolyline(item RID, points PackedVector2Array, colors PackedColorArray, width float64, antialiased bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddPolyline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddMultiline(item RID, points PackedVector2Array, colors PackedColorArray, width float64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddMultiline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddRect(item RID, rect Rect2, color Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddCircle(item RID, pos Vector2, radius float64, color Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&radius) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddCircle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddTextureRect(item RID, rect Rect2, texture RID, tile bool, modulate Color, transpose bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&tile) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddTextureRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddMsdfTextureRectRegion(item RID, rect Rect2, texture RID, src_rect Rect2, modulate Color, outline_size int64, px_range float64, scale float64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), texture.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&outline_size) , gdc.ConstTypePtr(&px_range) , gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddMsdfTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddLcdTextureRectRegion(item RID, rect Rect2, texture RID, src_rect Rect2, modulate Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), texture.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddLcdTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddTextureRectRegion(item RID, rect Rect2, texture RID, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), texture.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , gdc.ConstTypePtr(&clip_uv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddNinePatch(item RID, rect Rect2, source Rect2, texture RID, topleft Vector2, bottomright Vector2, x_axis_mode RenderingServerNinePatchAxisMode, y_axis_mode RenderingServerNinePatchAxisMode, draw_center bool, modulate Color, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), rect.AsCTypePtr(), source.AsCTypePtr(), texture.AsCTypePtr(), topleft.AsCTypePtr(), bottomright.AsCTypePtr(), gdc.ConstTypePtr(&x_axis_mode) , gdc.ConstTypePtr(&y_axis_mode) , gdc.ConstTypePtr(&draw_center) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddNinePatch), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddPrimitive(item RID, points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddPrimitive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddPolygon(item RID, points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddTriangleArray(item RID, indices PackedInt32Array, points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, bones PackedInt32Array, weights PackedFloat32Array, texture RID, count int64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), indices.AsCTypePtr(), points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), bones.AsCTypePtr(), weights.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddTriangleArray), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddMesh(item RID, mesh RID, transform Transform2D, modulate Color, texture RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), mesh.AsCTypePtr(), transform.AsCTypePtr(), modulate.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddMultimesh(item RID, mesh RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), mesh.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddMultimesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddParticles(item RID, particles RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), particles.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddParticles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddSetTransform(item RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddClipIgnore(item RID, ignore bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&ignore) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddClipIgnore), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemAddAnimationSlice(item RID, animation_length float64, slice_begin float64, slice_end float64, offset float64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&animation_length) , gdc.ConstTypePtr(&slice_begin) , gdc.ConstTypePtr(&slice_end) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemAddAnimationSlice), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetSortChildrenByY(item RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetSortChildrenByY), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetZIndex(item RID, z_index int64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&z_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetZIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetZAsRelativeToParent(item RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetZAsRelativeToParent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetCopyToBackbuffer(item RID, enabled bool, rect Rect2, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetCopyToBackbuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemClear(item RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetDrawIndex(item RID, index int64, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetDrawIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetMaterial(item RID, material RID, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetUseParentMaterial(item RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetUseParentMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetVisibilityNotifier(item RID, enable bool, area Rect2, enter_callable Callable, exit_callable Callable, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&enable) , area.AsCTypePtr(), enter_callable.AsCTypePtr(), exit_callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetVisibilityNotifier), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasItemSetCanvasGroupMode(item RID, mode RenderingServerCanvasGroupMode, clear_margin float64, fit_empty bool, fit_margin float64, blur_mipmaps bool, )  {
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&clear_margin) , gdc.ConstTypePtr(&fit_empty) , gdc.ConstTypePtr(&fit_margin) , gdc.ConstTypePtr(&blur_mipmaps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasItemSetCanvasGroupMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasLightAttachToCanvas(light RID, canvas RID, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), canvas.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightAttachToCanvas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetEnabled(light RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetTextureScale(light RID, scale float64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetTextureScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetTransform(light RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetTexture(light RID, texture RID, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetTextureOffset(light RID, offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetTextureOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetColor(light RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetHeight(light RID, height float64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetEnergy(light RID, energy float64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetZRange(light RID, min_z int64, max_z int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&min_z) , gdc.ConstTypePtr(&max_z) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetZRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetLayerRange(light RID, min_layer int64, max_layer int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&min_layer) , gdc.ConstTypePtr(&max_layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetLayerRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetItemCullMask(light RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetItemCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetItemShadowCullMask(light RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetItemShadowCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetMode(light RID, mode RenderingServerCanvasLightMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetShadowEnabled(light RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetShadowEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetShadowFilter(light RID, filter RenderingServerCanvasLightShadowFilter, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetShadowFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetShadowColor(light RID, color Color, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetShadowColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetShadowSmooth(light RID, smooth float64, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&smooth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetShadowSmooth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightSetBlendMode(light RID, mode RenderingServerCanvasLightBlendMode, )  {
  cargs := []gdc.ConstTypePtr{light.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasLightOccluderAttachToCanvas(occluder RID, canvas RID, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), canvas.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderAttachToCanvas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderSetEnabled(occluder RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderSetPolygon(occluder RID, polygon RID, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderSetAsSdfCollision(occluder RID, enable bool, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderSetAsSdfCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderSetTransform(occluder RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasLightOccluderSetLightMask(occluder RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasLightOccluderSetLightMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasOccluderPolygonCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasOccluderPolygonCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CanvasOccluderPolygonSetShape(occluder_polygon RID, shape PackedVector2Array, closed bool, )  {
  cargs := []gdc.ConstTypePtr{occluder_polygon.AsCTypePtr(), shape.AsCTypePtr(), gdc.ConstTypePtr(&closed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasOccluderPolygonSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasOccluderPolygonSetCullMode(occluder_polygon RID, mode RenderingServerCanvasOccluderPolygonCullMode, )  {
  cargs := []gdc.ConstTypePtr{occluder_polygon.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasOccluderPolygonSetCullMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) CanvasSetShadowTextureSize(size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCanvasSetShadowTextureSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GlobalShaderParameterAdd(name StringName, type_ RenderingServerGlobalShaderParameterType, default_value Variant, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&type_) , default_value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterAdd), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GlobalShaderParameterRemove(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterRemove), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GlobalShaderParameterGetList() []StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterGetList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[StringName](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RenderingServer) GlobalShaderParameterSet(name StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterSet), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GlobalShaderParameterSetOverride(name StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterSetOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GlobalShaderParameterGet(name StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GlobalShaderParameterGetType(name StringName, ) RenderingServerGlobalShaderParameterType {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerGlobalShaderParameterType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGlobalShaderParameterGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderingServer) FreeRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) RequestFrameDrawnCallback(callable Callable, )  {
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnRequestFrameDrawnCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) HasChanged() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnHasChanged), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) GetRenderingInfo(info RenderingServerRenderingInfo, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&info) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&info)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetRenderingInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) GetVideoAdapterName() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetVideoAdapterName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetVideoAdapterVendor() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetVideoAdapterVendor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetVideoAdapterType() RenderingDeviceDeviceType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceDeviceType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetVideoAdapterType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RenderingServer) GetVideoAdapterApiVersion() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetVideoAdapterApiVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) MakeSphereMesh(latitudes int64, longitudes int64, radius float64, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&latitudes) , gdc.ConstTypePtr(&longitudes) , gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&latitudes)
  pinner.Pin(&longitudes)
  pinner.Pin(&radius)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnMakeSphereMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetTestCube() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetTestCube), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetTestTexture() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetTestTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) GetWhiteTexture() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetWhiteTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SetBootImage(image Image, color Color, scale bool, use_filter bool, )  {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&scale) , gdc.ConstTypePtr(&use_filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSetBootImage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GetDefaultClearColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetDefaultClearColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) SetDefaultClearColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSetDefaultClearColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) HasFeature(feature RenderingServerFeatures, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&feature)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnHasFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) HasOsFeature(feature String, ) bool {
  cargs := []gdc.ConstTypePtr{feature.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnHasOsFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) SetDebugGenerateWireframes(generate bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&generate) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSetDebugGenerateWireframes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) IsRenderLoopEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnIsRenderLoopEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) SetRenderLoopEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnSetRenderLoopEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GetFrameSetupTimeCpu() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetFrameSetupTimeCpu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RenderingServer) ForceSync()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnForceSync), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) ForceDraw(swap_buffers bool, frame_step float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&swap_buffers) , gdc.ConstTypePtr(&frame_step) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnForceDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RenderingServer) GetRenderingDevice() RenderingDevice {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRenderingDevice()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnGetRenderingDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CreateLocalRenderingDevice() RenderingDevice {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRenderingDevice()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCreateLocalRenderingDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RenderingServer) CallOnRenderThread(callable Callable, )  {
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingServer.fnCallOnRenderThread), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RenderingServerFramePreDrawSignalFn func()

func (me *RenderingServer) ConnectFramePreDraw(subs SignalSubscribers, fn RenderingServerFramePreDrawSignalFn) {
  sig := StringNameFromStr("frame_pre_draw")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RenderingServer) DisconnectFramePreDraw(subs SignalSubscribers, fn RenderingServerFramePreDrawSignalFn) {
  sig := StringNameFromStr("frame_pre_draw")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RenderingServerFramePostDrawSignalFn func()

func (me *RenderingServer) ConnectFramePostDraw(subs SignalSubscribers, fn RenderingServerFramePostDrawSignalFn) {
  sig := StringNameFromStr("frame_post_draw")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RenderingServer) DisconnectFramePostDraw(subs SignalSubscribers, fn RenderingServerFramePostDrawSignalFn) {
  sig := StringNameFromStr("frame_post_draw")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
