// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdivDisabled ViewportPositionalShadowAtlasQuadrantSubdiv = 0
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv1 ViewportPositionalShadowAtlasQuadrantSubdiv = 1
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv4 ViewportPositionalShadowAtlasQuadrantSubdiv = 2
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv16 ViewportPositionalShadowAtlasQuadrantSubdiv = 3
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv64 ViewportPositionalShadowAtlasQuadrantSubdiv = 4
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv256 ViewportPositionalShadowAtlasQuadrantSubdiv = 5
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv1024 ViewportPositionalShadowAtlasQuadrantSubdiv = 6
  ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdivMax ViewportPositionalShadowAtlasQuadrantSubdiv = 7
)

type ViewportScaling3DMode int
const (
  ViewportScaling3DModeScaling3DModeBilinear ViewportScaling3DMode = 0
  ViewportScaling3DModeScaling3DModeFsr ViewportScaling3DMode = 1
  ViewportScaling3DModeScaling3DModeMax ViewportScaling3DMode = 2
)

type ViewportMSAA int
const (
  ViewportMSAAMsaaDisabled ViewportMSAA = 0
  ViewportMSAAMsaa2X ViewportMSAA = 1
  ViewportMSAAMsaa4X ViewportMSAA = 2
  ViewportMSAAMsaa8X ViewportMSAA = 3
  ViewportMSAAMsaaMax ViewportMSAA = 4
)

type ViewportScreenSpaceAA int
const (
  ViewportScreenSpaceAAScreenSpaceAaDisabled ViewportScreenSpaceAA = 0
  ViewportScreenSpaceAAScreenSpaceAaFxaa ViewportScreenSpaceAA = 1
  ViewportScreenSpaceAAScreenSpaceAaMax ViewportScreenSpaceAA = 2
)

type ViewportRenderInfo int
const (
  ViewportRenderInfoRenderInfoObjectsInFrame ViewportRenderInfo = 0
  ViewportRenderInfoRenderInfoPrimitivesInFrame ViewportRenderInfo = 1
  ViewportRenderInfoRenderInfoDrawCallsInFrame ViewportRenderInfo = 2
  ViewportRenderInfoRenderInfoMax ViewportRenderInfo = 3
)

type ViewportRenderInfoType int
const (
  ViewportRenderInfoTypeRenderInfoTypeVisible ViewportRenderInfoType = 0
  ViewportRenderInfoTypeRenderInfoTypeShadow ViewportRenderInfoType = 1
  ViewportRenderInfoTypeRenderInfoTypeMax ViewportRenderInfoType = 2
)

type ViewportDebugDraw int
const (
  ViewportDebugDrawDebugDrawDisabled ViewportDebugDraw = 0
  ViewportDebugDrawDebugDrawUnshaded ViewportDebugDraw = 1
  ViewportDebugDrawDebugDrawLighting ViewportDebugDraw = 2
  ViewportDebugDrawDebugDrawOverdraw ViewportDebugDraw = 3
  ViewportDebugDrawDebugDrawWireframe ViewportDebugDraw = 4
  ViewportDebugDrawDebugDrawNormalBuffer ViewportDebugDraw = 5
  ViewportDebugDrawDebugDrawVoxelGiAlbedo ViewportDebugDraw = 6
  ViewportDebugDrawDebugDrawVoxelGiLighting ViewportDebugDraw = 7
  ViewportDebugDrawDebugDrawVoxelGiEmission ViewportDebugDraw = 8
  ViewportDebugDrawDebugDrawShadowAtlas ViewportDebugDraw = 9
  ViewportDebugDrawDebugDrawDirectionalShadowAtlas ViewportDebugDraw = 10
  ViewportDebugDrawDebugDrawSceneLuminance ViewportDebugDraw = 11
  ViewportDebugDrawDebugDrawSsao ViewportDebugDraw = 12
  ViewportDebugDrawDebugDrawSsil ViewportDebugDraw = 13
  ViewportDebugDrawDebugDrawPssmSplits ViewportDebugDraw = 14
  ViewportDebugDrawDebugDrawDecalAtlas ViewportDebugDraw = 15
  ViewportDebugDrawDebugDrawSdfgi ViewportDebugDraw = 16
  ViewportDebugDrawDebugDrawSdfgiProbes ViewportDebugDraw = 17
  ViewportDebugDrawDebugDrawGiBuffer ViewportDebugDraw = 18
  ViewportDebugDrawDebugDrawDisableLod ViewportDebugDraw = 19
  ViewportDebugDrawDebugDrawClusterOmniLights ViewportDebugDraw = 20
  ViewportDebugDrawDebugDrawClusterSpotLights ViewportDebugDraw = 21
  ViewportDebugDrawDebugDrawClusterDecals ViewportDebugDraw = 22
  ViewportDebugDrawDebugDrawClusterReflectionProbes ViewportDebugDraw = 23
  ViewportDebugDrawDebugDrawOccluders ViewportDebugDraw = 24
  ViewportDebugDrawDebugDrawMotionVectors ViewportDebugDraw = 25
)

type ViewportDefaultCanvasItemTextureFilter int
const (
  ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterNearest ViewportDefaultCanvasItemTextureFilter = 0
  ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterLinear ViewportDefaultCanvasItemTextureFilter = 1
  ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterLinearWithMipmaps ViewportDefaultCanvasItemTextureFilter = 2
  ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterNearestWithMipmaps ViewportDefaultCanvasItemTextureFilter = 3
  ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterMax ViewportDefaultCanvasItemTextureFilter = 4
)

type ViewportDefaultCanvasItemTextureRepeat int
const (
  ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatDisabled ViewportDefaultCanvasItemTextureRepeat = 0
  ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatEnabled ViewportDefaultCanvasItemTextureRepeat = 1
  ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatMirror ViewportDefaultCanvasItemTextureRepeat = 2
  ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatMax ViewportDefaultCanvasItemTextureRepeat = 3
)

type ViewportSDFOversize int
const (
  ViewportSDFOversizeSdfOversize100Percent ViewportSDFOversize = 0
  ViewportSDFOversizeSdfOversize120Percent ViewportSDFOversize = 1
  ViewportSDFOversizeSdfOversize150Percent ViewportSDFOversize = 2
  ViewportSDFOversizeSdfOversize200Percent ViewportSDFOversize = 3
  ViewportSDFOversizeSdfOversizeMax ViewportSDFOversize = 4
)

type ViewportSDFScale int
const (
  ViewportSDFScaleSdfScale100Percent ViewportSDFScale = 0
  ViewportSDFScaleSdfScale50Percent ViewportSDFScale = 1
  ViewportSDFScaleSdfScale25Percent ViewportSDFScale = 2
  ViewportSDFScaleSdfScaleMax ViewportSDFScale = 3
)

type ViewportVRSMode int
const (
  ViewportVRSModeVrsDisabled ViewportVRSMode = 0
  ViewportVRSModeVrsTexture ViewportVRSMode = 1
  ViewportVRSModeVrsXr ViewportVRSMode = 2
  ViewportVRSModeVrsMax ViewportVRSMode = 3
)

func  (me *Viewport) SetWorld2D(world_2d World2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetWorld2D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) FindWorld2D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetCanvasTransform(xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetCanvasTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetGlobalCanvasTransform(xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetGlobalCanvasTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetFinalTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetScreenTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetVisibleRect() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetTransparentBackground(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) HasTransparentBackground() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetMsaa2D(msaa ViewportMSAA, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetMsaa2D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetMsaa3D(msaa ViewportMSAA, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetMsaa3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetScreenSpaceAa(screen_space_aa ViewportScreenSpaceAA, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetScreenSpaceAa() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetUseTaa(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsUsingTaa() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetUseDebanding(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsUsingDebanding() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetUseOcclusionCulling(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsUsingOcclusionCulling() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetDebugDraw(debug_draw ViewportDebugDraw, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetDebugDraw() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetRenderInfo(type_ ViewportRenderInfoType, info ViewportRenderInfo, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetTexture() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetPhysicsObjectPicking(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetPhysicsObjectPicking() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetPhysicsObjectPickingSort(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetPhysicsObjectPickingSort() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetViewportRid() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) PushTextInput(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) PushInput(event InputEvent, in_local_coords bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) PushUnhandledInput(event InputEvent, in_local_coords bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetCamera2D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetAsAudioListener2D(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsAudioListener2D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetMousePosition() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) WarpMouse(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) UpdateMouseCursorState() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GuiGetDragData() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GuiIsDragging() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GuiIsDragSuccessful() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GuiReleaseFocus() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GuiGetFocusOwner() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetDisableInput(disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsInputDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetPositionalShadowAtlasSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetPositionalShadowAtlasSize() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetPositionalShadowAtlas16Bits(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetPositionalShadowAtlas16Bits() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetSnapControlsToPixels(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsSnapControlsToPixelsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetSnap2DTransformsToPixel(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsSnap2DTransformsToPixelEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetSnap2DVerticesToPixel(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsSnap2DVerticesToPixelEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetPositionalShadowAtlasQuadrantSubdiv(quadrant int, subdiv ViewportPositionalShadowAtlasQuadrantSubdiv, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetPositionalShadowAtlasQuadrantSubdiv(quadrant int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetInputAsHandled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsInputHandled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetHandleInputLocally(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsHandlingInputLocally() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetDefaultCanvasItemTextureFilter(mode ViewportDefaultCanvasItemTextureFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetDefaultCanvasItemTextureFilter() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetEmbeddingSubwindows(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsEmbeddingSubwindows() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetCanvasCullMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetCanvasCullMask() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetCanvasCullMaskBit(layer int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetCanvasCullMaskBit(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetDefaultCanvasItemTextureRepeat(mode ViewportDefaultCanvasItemTextureRepeat, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetDefaultCanvasItemTextureRepeat() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetSdfOversize(oversize ViewportSDFOversize, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetSdfOversize() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetSdfScale(scale ViewportSDFScale, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetSdfScale() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetMeshLodThreshold(pixels float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetMeshLodThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetWorld3D(world_3d World3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetWorld3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) FindWorld3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetUseOwnWorld3D(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsUsingOwnWorld3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetCamera3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetAsAudioListener3D(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsAudioListener3D() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetDisable3D(disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) Is3DDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetUseXr(use bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) IsUsingXr() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetScaling3DMode(scaling_3d_mode ViewportScaling3DMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetScaling3DMode() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetScaling3DScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetScaling3DScale() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetFsrSharpness(fsr_sharpness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetFsrSharpness() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetTextureMipmapBias(texture_mipmap_bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetTextureMipmapBias() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetVrsMode(mode ViewportVRSMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetVrsMode() { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) SetVrsTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Viewport) GetVrsTexture() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
