// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Viewport struct {
  obj gdc.ObjectPtr
}

func (me *Viewport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Viewport) BaseClass() string {
  return "Viewport"
}

type ViewportPositionalShadowAtlasQuadrantSubdiv int
const (
  ViewportShadowAtlasQuadrantSubdivDisabled ViewportPositionalShadowAtlasQuadrantSubdiv = 0
  ViewportShadowAtlasQuadrantSubdiv1 ViewportPositionalShadowAtlasQuadrantSubdiv = 1
  ViewportShadowAtlasQuadrantSubdiv4 ViewportPositionalShadowAtlasQuadrantSubdiv = 2
  ViewportShadowAtlasQuadrantSubdiv16 ViewportPositionalShadowAtlasQuadrantSubdiv = 3
  ViewportShadowAtlasQuadrantSubdiv64 ViewportPositionalShadowAtlasQuadrantSubdiv = 4
  ViewportShadowAtlasQuadrantSubdiv256 ViewportPositionalShadowAtlasQuadrantSubdiv = 5
  ViewportShadowAtlasQuadrantSubdiv1024 ViewportPositionalShadowAtlasQuadrantSubdiv = 6
  ViewportShadowAtlasQuadrantSubdivMax ViewportPositionalShadowAtlasQuadrantSubdiv = 7
)

type ViewportScaling3DMode int
const (
  ViewportScaling3DModeBilinear ViewportScaling3DMode = 0
  ViewportScaling3DModeFsr ViewportScaling3DMode = 1
  ViewportScaling3DModeMax ViewportScaling3DMode = 2
)

type ViewportMSAA int
const (
  ViewportMsaaDisabled ViewportMSAA = 0
  ViewportMsaa2X ViewportMSAA = 1
  ViewportMsaa4X ViewportMSAA = 2
  ViewportMsaa8X ViewportMSAA = 3
  ViewportMsaaMax ViewportMSAA = 4
)

type ViewportScreenSpaceAA int
const (
  ViewportScreenSpaceAaDisabled ViewportScreenSpaceAA = 0
  ViewportScreenSpaceAaFxaa ViewportScreenSpaceAA = 1
  ViewportScreenSpaceAaMax ViewportScreenSpaceAA = 2
)

type ViewportRenderInfo int
const (
  ViewportRenderInfoObjectsInFrame ViewportRenderInfo = 0
  ViewportRenderInfoPrimitivesInFrame ViewportRenderInfo = 1
  ViewportRenderInfoDrawCallsInFrame ViewportRenderInfo = 2
  ViewportRenderInfoMax ViewportRenderInfo = 3
)

type ViewportRenderInfoType int
const (
  ViewportRenderInfoTypeVisible ViewportRenderInfoType = 0
  ViewportRenderInfoTypeShadow ViewportRenderInfoType = 1
  ViewportRenderInfoTypeMax ViewportRenderInfoType = 2
)

type ViewportDebugDraw int
const (
  ViewportDebugDrawDisabled ViewportDebugDraw = 0
  ViewportDebugDrawUnshaded ViewportDebugDraw = 1
  ViewportDebugDrawLighting ViewportDebugDraw = 2
  ViewportDebugDrawOverdraw ViewportDebugDraw = 3
  ViewportDebugDrawWireframe ViewportDebugDraw = 4
  ViewportDebugDrawNormalBuffer ViewportDebugDraw = 5
  ViewportDebugDrawVoxelGiAlbedo ViewportDebugDraw = 6
  ViewportDebugDrawVoxelGiLighting ViewportDebugDraw = 7
  ViewportDebugDrawVoxelGiEmission ViewportDebugDraw = 8
  ViewportDebugDrawShadowAtlas ViewportDebugDraw = 9
  ViewportDebugDrawDirectionalShadowAtlas ViewportDebugDraw = 10
  ViewportDebugDrawSceneLuminance ViewportDebugDraw = 11
  ViewportDebugDrawSsao ViewportDebugDraw = 12
  ViewportDebugDrawSsil ViewportDebugDraw = 13
  ViewportDebugDrawPssmSplits ViewportDebugDraw = 14
  ViewportDebugDrawDecalAtlas ViewportDebugDraw = 15
  ViewportDebugDrawSdfgi ViewportDebugDraw = 16
  ViewportDebugDrawSdfgiProbes ViewportDebugDraw = 17
  ViewportDebugDrawGiBuffer ViewportDebugDraw = 18
  ViewportDebugDrawDisableLod ViewportDebugDraw = 19
  ViewportDebugDrawClusterOmniLights ViewportDebugDraw = 20
  ViewportDebugDrawClusterSpotLights ViewportDebugDraw = 21
  ViewportDebugDrawClusterDecals ViewportDebugDraw = 22
  ViewportDebugDrawClusterReflectionProbes ViewportDebugDraw = 23
  ViewportDebugDrawOccluders ViewportDebugDraw = 24
  ViewportDebugDrawMotionVectors ViewportDebugDraw = 25
)

type ViewportDefaultCanvasItemTextureFilter int
const (
  ViewportDefaultCanvasItemTextureFilterNearest ViewportDefaultCanvasItemTextureFilter = 0
  ViewportDefaultCanvasItemTextureFilterLinear ViewportDefaultCanvasItemTextureFilter = 1
  ViewportDefaultCanvasItemTextureFilterLinearWithMipmaps ViewportDefaultCanvasItemTextureFilter = 2
  ViewportDefaultCanvasItemTextureFilterNearestWithMipmaps ViewportDefaultCanvasItemTextureFilter = 3
  ViewportDefaultCanvasItemTextureFilterMax ViewportDefaultCanvasItemTextureFilter = 4
)

type ViewportDefaultCanvasItemTextureRepeat int
const (
  ViewportDefaultCanvasItemTextureRepeatDisabled ViewportDefaultCanvasItemTextureRepeat = 0
  ViewportDefaultCanvasItemTextureRepeatEnabled ViewportDefaultCanvasItemTextureRepeat = 1
  ViewportDefaultCanvasItemTextureRepeatMirror ViewportDefaultCanvasItemTextureRepeat = 2
  ViewportDefaultCanvasItemTextureRepeatMax ViewportDefaultCanvasItemTextureRepeat = 3
)

type ViewportSDFOversize int
const (
  ViewportSdfOversize100Percent ViewportSDFOversize = 0
  ViewportSdfOversize120Percent ViewportSDFOversize = 1
  ViewportSdfOversize150Percent ViewportSDFOversize = 2
  ViewportSdfOversize200Percent ViewportSDFOversize = 3
  ViewportSdfOversizeMax ViewportSDFOversize = 4
)

type ViewportSDFScale int
const (
  ViewportSdfScale100Percent ViewportSDFScale = 0
  ViewportSdfScale50Percent ViewportSDFScale = 1
  ViewportSdfScale25Percent ViewportSDFScale = 2
  ViewportSdfScaleMax ViewportSDFScale = 3
)

type ViewportVRSMode int
const (
  ViewportVrsDisabled ViewportVRSMode = 0
  ViewportVrsTexture ViewportVRSMode = 1
  ViewportVrsXr ViewportVRSMode = 2
  ViewportVrsMax ViewportVRSMode = 3
)
