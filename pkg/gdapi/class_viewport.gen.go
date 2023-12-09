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



// Enums

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

func (me *Viewport) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Viewport) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Viewport) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Viewport) SetWorld2D(world_2d World2D, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetWorld2D()  {
  panic("TODO: implement")
}

func  (me *Viewport) FindWorld2D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetCanvasTransform(xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetCanvasTransform()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetGlobalCanvasTransform(xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetGlobalCanvasTransform()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetFinalTransform()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetScreenTransform()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetVisibleRect()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetTransparentBackground(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) HasTransparentBackground()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetMsaa2D(msaa ViewportMSAA, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetMsaa2D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetMsaa3D(msaa ViewportMSAA, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetMsaa3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetScreenSpaceAa(screen_space_aa ViewportScreenSpaceAA, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetScreenSpaceAa()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetUseTaa(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsUsingTaa()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetUseDebanding(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsUsingDebanding()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetUseOcclusionCulling(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsUsingOcclusionCulling()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetDebugDraw(debug_draw ViewportDebugDraw, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetDebugDraw()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetRenderInfo(type_ ViewportRenderInfoType, info ViewportRenderInfo, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetTexture()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetPhysicsObjectPicking(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetPhysicsObjectPicking()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetPhysicsObjectPickingSort(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetPhysicsObjectPickingSort()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetViewportRid()  {
  panic("TODO: implement")
}

func  (me *Viewport) PushTextInput(text String, )  {
  panic("TODO: implement")
}

func  (me *Viewport) PushInput(event InputEvent, in_local_coords bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) PushUnhandledInput(event InputEvent, in_local_coords bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetCamera2D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetAsAudioListener2D(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsAudioListener2D()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetMousePosition()  {
  panic("TODO: implement")
}

func  (me *Viewport) WarpMouse(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Viewport) UpdateMouseCursorState()  {
  panic("TODO: implement")
}

func  (me *Viewport) GuiGetDragData()  {
  panic("TODO: implement")
}

func  (me *Viewport) GuiIsDragging()  {
  panic("TODO: implement")
}

func  (me *Viewport) GuiIsDragSuccessful()  {
  panic("TODO: implement")
}

func  (me *Viewport) GuiReleaseFocus()  {
  panic("TODO: implement")
}

func  (me *Viewport) GuiGetFocusOwner()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetDisableInput(disable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsInputDisabled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetPositionalShadowAtlasSize(size int, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetPositionalShadowAtlasSize()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetPositionalShadowAtlas16Bits(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetPositionalShadowAtlas16Bits()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetSnapControlsToPixels(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsSnapControlsToPixelsEnabled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetSnap2DTransformsToPixel(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsSnap2DTransformsToPixelEnabled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetSnap2DVerticesToPixel(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsSnap2DVerticesToPixelEnabled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetPositionalShadowAtlasQuadrantSubdiv(quadrant int, subdiv ViewportPositionalShadowAtlasQuadrantSubdiv, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetPositionalShadowAtlasQuadrantSubdiv(quadrant int, )  {
  panic("TODO: implement")
}

func  (me *Viewport) SetInputAsHandled()  {
  panic("TODO: implement")
}

func  (me *Viewport) IsInputHandled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetHandleInputLocally(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsHandlingInputLocally()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetDefaultCanvasItemTextureFilter(mode ViewportDefaultCanvasItemTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetDefaultCanvasItemTextureFilter()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetEmbeddingSubwindows(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsEmbeddingSubwindows()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetCanvasCullMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetCanvasCullMask()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetCanvasCullMaskBit(layer int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetCanvasCullMaskBit(layer int, )  {
  panic("TODO: implement")
}

func  (me *Viewport) SetDefaultCanvasItemTextureRepeat(mode ViewportDefaultCanvasItemTextureRepeat, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetDefaultCanvasItemTextureRepeat()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetSdfOversize(oversize ViewportSDFOversize, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetSdfOversize()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetSdfScale(scale ViewportSDFScale, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetSdfScale()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetMeshLodThreshold(pixels float32, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetMeshLodThreshold()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetWorld3D(world_3d World3D, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetWorld3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) FindWorld3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetUseOwnWorld3D(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsUsingOwnWorld3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) GetCamera3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetAsAudioListener3D(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsAudioListener3D()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetDisable3D(disable bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) Is3DDisabled()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetUseXr(use bool, )  {
  panic("TODO: implement")
}

func  (me *Viewport) IsUsingXr()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetScaling3DMode(scaling_3d_mode ViewportScaling3DMode, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetScaling3DMode()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetScaling3DScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetScaling3DScale()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetFsrSharpness(fsr_sharpness float32, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetFsrSharpness()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetTextureMipmapBias(texture_mipmap_bias float32, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetTextureMipmapBias()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetVrsMode(mode ViewportVRSMode, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetVrsMode()  {
  panic("TODO: implement")
}

func  (me *Viewport) SetVrsTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Viewport) GetVrsTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
