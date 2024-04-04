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

type Viewport struct {
  Node
}

func (me *Viewport) BaseClass() string {
  return "Viewport"
}

func NewViewport() *Viewport {
  str := StringNameFromStr("Viewport") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Viewport{}
  obj.SetBaseObject(objPtr)
  return obj
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
  ViewportScaling3DModeScaling3DModeFsr2 ViewportScaling3DMode = 2
  ViewportScaling3DModeScaling3DModeMax ViewportScaling3DMode = 3
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
  ViewportDebugDrawDebugDrawInternalBuffer ViewportDebugDraw = 26
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
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2736080068) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{world_2d.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetWorld2D() World2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339128592) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) FindWorld2D() World2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_world_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339128592) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetCanvasTransform(xform Transform2D, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetCanvasTransform() Transform2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetGlobalCanvasTransform(xform Transform2D, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetGlobalCanvasTransform() Transform2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) GetFinalTransform() Transform2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_final_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) GetScreenTransform() Transform2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) GetVisibleRect() Rect2 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetTransparentBackground(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transparent_background")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) HasTransparentBackground() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_transparent_background")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetUseHdr2D(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_hdr_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingHdr2D() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_hdr_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetMsaa2D(msaa ViewportMSAA, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msaa_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3330258708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetMsaa2D() ViewportMSAA {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msaa_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2542055527) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportMSAA

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetMsaa3D(msaa ViewportMSAA, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3330258708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetMsaa3D() ViewportMSAA {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_msaa_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2542055527) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportMSAA

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetScreenSpaceAa(screen_space_aa ViewportScreenSpaceAA, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3544169389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_space_aa) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetScreenSpaceAa() ViewportScreenSpaceAA {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_space_aa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1390814124) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportScreenSpaceAA

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetUseTaa(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_taa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingTaa() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_taa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetUseDebanding(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingDebanding() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetUseOcclusionCulling(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_occlusion_culling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingOcclusionCulling() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_occlusion_culling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetDebugDraw(debug_draw ViewportDebugDraw, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1970246205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&debug_draw) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetDebugDraw() ViewportDebugDraw {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 579191299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportDebugDraw

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) GetRenderInfo(type_ ViewportRenderInfoType, info ViewportRenderInfo, ) int64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 481977019) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&info) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)
  pinner.Pin(&info)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GetTexture() ViewportTexture {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1746695840) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewViewportTexture()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetPhysicsObjectPicking(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_object_picking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetPhysicsObjectPicking() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_object_picking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetPhysicsObjectPickingSort(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_object_picking_sort")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetPhysicsObjectPickingSort() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_object_picking_sort")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GetViewportRid() RID {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) PushTextInput(text String, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_text_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) PushInput(event InputEvent, in_local_coords bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3644664830) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&in_local_coords) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) PushUnhandledInput(event InputEvent, in_local_coords bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_unhandled_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3644664830) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&in_local_coords) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetCamera2D() Camera2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3551466917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCamera2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetAsAudioListener2D(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_audio_listener_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsAudioListener2D() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_audio_listener_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GetMousePosition() Vector2 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mouse_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) WarpMouse(position Vector2, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("warp_mouse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) UpdateMouseCursorState()  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_mouse_cursor_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GuiGetDragData() Variant {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("gui_get_drag_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) GuiIsDragging() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("gui_is_dragging")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GuiIsDragSuccessful() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("gui_is_drag_successful")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GuiReleaseFocus()  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("gui_release_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GuiGetFocusOwner() Control {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("gui_get_focus_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewControl()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetDisableInput(disable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsInputDisabled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetPositionalShadowAtlasSize(size int64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_positional_shadow_atlas_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetPositionalShadowAtlasSize() int64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_positional_shadow_atlas_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetPositionalShadowAtlas16Bits(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_positional_shadow_atlas_16_bits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetPositionalShadowAtlas16Bits() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_positional_shadow_atlas_16_bits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetSnapControlsToPixels(enabled bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap_controls_to_pixels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsSnapControlsToPixelsEnabled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_snap_controls_to_pixels_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetSnap2DTransformsToPixel(enabled bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap_2d_transforms_to_pixel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsSnap2DTransformsToPixelEnabled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_snap_2d_transforms_to_pixel_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetSnap2DVerticesToPixel(enabled bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap_2d_vertices_to_pixel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsSnap2DVerticesToPixelEnabled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_snap_2d_vertices_to_pixel_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetPositionalShadowAtlasQuadrantSubdiv(quadrant int64, subdiv ViewportPositionalShadowAtlasQuadrantSubdiv, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_positional_shadow_atlas_quadrant_subdiv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2596956071) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quadrant) , gdc.ConstTypePtr(&subdiv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetPositionalShadowAtlasQuadrantSubdiv(quadrant int64, ) ViewportPositionalShadowAtlasQuadrantSubdiv {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_positional_shadow_atlas_quadrant_subdiv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2676778355) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quadrant) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportPositionalShadowAtlasQuadrantSubdiv
  pinner.Pin(&quadrant)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetInputAsHandled()  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_as_handled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsInputHandled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_handled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetHandleInputLocally(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handle_input_locally")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsHandlingInputLocally() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_handling_input_locally")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetDefaultCanvasItemTextureFilter(mode ViewportDefaultCanvasItemTextureFilter, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_canvas_item_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2815160100) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetDefaultCanvasItemTextureFilter() ViewportDefaultCanvasItemTextureFilter {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_canvas_item_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 896601198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportDefaultCanvasItemTextureFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetEmbeddingSubwindows(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_embedding_subwindows")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsEmbeddingSubwindows() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_embedding_subwindows")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GetEmbeddedSubwindows() []Window {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_embedded_subwindows")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Window](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Viewport) SetCanvasCullMask(mask int64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetCanvasCullMask() int64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetCanvasCullMaskBit(layer int64, enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_cull_mask_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetCanvasCullMaskBit(layer int64, ) bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_cull_mask_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetDefaultCanvasItemTextureRepeat(mode ViewportDefaultCanvasItemTextureRepeat, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_canvas_item_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1658513413) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetDefaultCanvasItemTextureRepeat() ViewportDefaultCanvasItemTextureRepeat {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_canvas_item_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4049774160) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportDefaultCanvasItemTextureRepeat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetSdfOversize(oversize ViewportSDFOversize, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdf_oversize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2574159017) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversize) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetSdfOversize() ViewportSDFOversize {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdf_oversize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2631427510) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportSDFOversize

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetSdfScale(scale ViewportSDFScale, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdf_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1402773951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetSdfScale() ViewportSDFScale {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdf_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3162688184) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportSDFScale

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetMeshLodThreshold(pixels float64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh_lod_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetMeshLodThreshold() float64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh_lod_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetWorld3D(world_3d World3D, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1400875337) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{world_3d.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetWorld3D() World3D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 317588385) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) FindWorld3D() World3D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 317588385) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetUseOwnWorld3D(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_own_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingOwnWorld3D() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_own_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) GetCamera3D() Camera3D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285090890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCamera3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Viewport) SetAsAudioListener3D(enable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_audio_listener_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsAudioListener3D() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_audio_listener_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetDisable3D(disable bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) Is3DDisabled() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_3d_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetUseXr(use bool, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_xr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) IsUsingXr() bool {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_xr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetScaling3DMode(scaling_3d_mode ViewportScaling3DMode, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1531597597) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scaling_3d_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetScaling3DMode() ViewportScaling3DMode {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scaling_3d_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2597660574) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportScaling3DMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetScaling3DScale(scale float64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scaling_3d_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetScaling3DScale() float64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scaling_3d_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetFsrSharpness(fsr_sharpness float64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fsr_sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetFsrSharpness() float64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fsr_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetTextureMipmapBias(texture_mipmap_bias float64, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mipmap_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetTextureMipmapBias() float64 {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_mipmap_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Viewport) SetVrsMode(mode ViewportVRSMode, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vrs_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2749867817) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetVrsMode() ViewportVRSMode {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vrs_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 349660525) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ViewportVRSMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Viewport) SetVrsTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vrs_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Viewport) GetVrsTexture() Texture2D {
  classNameV := StringNameFromStr("Viewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vrs_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ViewportSizeChangedSignalFn func()

func (me *Viewport) ConnectSizeChanged(subs SignalSubscribers, fn ViewportSizeChangedSignalFn) {
  sig := StringNameFromStr("size_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Viewport) DisconnectSizeChanged(subs SignalSubscribers, fn ViewportSizeChangedSignalFn) {
  sig := StringNameFromStr("size_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ViewportGuiFocusChangedSignalFn func(node Control, )

func (me *Viewport) ConnectGuiFocusChanged(subs SignalSubscribers, fn ViewportGuiFocusChangedSignalFn) {
  sig := StringNameFromStr("gui_focus_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Viewport) DisconnectGuiFocusChanged(subs SignalSubscribers, fn ViewportGuiFocusChangedSignalFn) {
  sig := StringNameFromStr("gui_focus_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
