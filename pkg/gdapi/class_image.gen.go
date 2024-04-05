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

type ptrsForImageList struct {
  fnGetWidth gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnHasMipmaps gdc.MethodBindPtr
  fnGetFormat gdc.MethodBindPtr
  fnGetData gdc.MethodBindPtr
  fnConvert gdc.MethodBindPtr
  fnGetMipmapCount gdc.MethodBindPtr
  fnGetMipmapOffset gdc.MethodBindPtr
  fnResizeToPo2 gdc.MethodBindPtr
  fnResize gdc.MethodBindPtr
  fnShrinkX2 gdc.MethodBindPtr
  fnCrop gdc.MethodBindPtr
  fnFlipX gdc.MethodBindPtr
  fnFlipY gdc.MethodBindPtr
  fnGenerateMipmaps gdc.MethodBindPtr
  fnClearMipmaps gdc.MethodBindPtr
  fnCreate gdc.MethodBindPtr
  fnCreateFromData gdc.MethodBindPtr
  fnSetData gdc.MethodBindPtr
  fnIsEmpty gdc.MethodBindPtr
  fnLoad gdc.MethodBindPtr
  fnLoadFromFile gdc.MethodBindPtr
  fnSavePng gdc.MethodBindPtr
  fnSavePngToBuffer gdc.MethodBindPtr
  fnSaveJpg gdc.MethodBindPtr
  fnSaveJpgToBuffer gdc.MethodBindPtr
  fnSaveExr gdc.MethodBindPtr
  fnSaveExrToBuffer gdc.MethodBindPtr
  fnSaveWebp gdc.MethodBindPtr
  fnSaveWebpToBuffer gdc.MethodBindPtr
  fnDetectAlpha gdc.MethodBindPtr
  fnIsInvisible gdc.MethodBindPtr
  fnDetectUsedChannels gdc.MethodBindPtr
  fnCompress gdc.MethodBindPtr
  fnCompressFromChannels gdc.MethodBindPtr
  fnDecompress gdc.MethodBindPtr
  fnIsCompressed gdc.MethodBindPtr
  fnRotate90 gdc.MethodBindPtr
  fnRotate180 gdc.MethodBindPtr
  fnFixAlphaEdges gdc.MethodBindPtr
  fnPremultiplyAlpha gdc.MethodBindPtr
  fnSrgbToLinear gdc.MethodBindPtr
  fnNormalMapToXy gdc.MethodBindPtr
  fnRgbeToSrgb gdc.MethodBindPtr
  fnBumpMapToNormalMap gdc.MethodBindPtr
  fnComputeImageMetrics gdc.MethodBindPtr
  fnBlitRect gdc.MethodBindPtr
  fnBlitRectMask gdc.MethodBindPtr
  fnBlendRect gdc.MethodBindPtr
  fnBlendRectMask gdc.MethodBindPtr
  fnFill gdc.MethodBindPtr
  fnFillRect gdc.MethodBindPtr
  fnGetUsedRect gdc.MethodBindPtr
  fnGetRegion gdc.MethodBindPtr
  fnCopyFrom gdc.MethodBindPtr
  fnGetPixelv gdc.MethodBindPtr
  fnGetPixel gdc.MethodBindPtr
  fnSetPixelv gdc.MethodBindPtr
  fnSetPixel gdc.MethodBindPtr
  fnAdjustBcs gdc.MethodBindPtr
  fnLoadPngFromBuffer gdc.MethodBindPtr
  fnLoadJpgFromBuffer gdc.MethodBindPtr
  fnLoadWebpFromBuffer gdc.MethodBindPtr
  fnLoadTgaFromBuffer gdc.MethodBindPtr
  fnLoadBmpFromBuffer gdc.MethodBindPtr
  fnLoadKtxFromBuffer gdc.MethodBindPtr
  fnLoadSvgFromBuffer gdc.MethodBindPtr
  fnLoadSvgFromString gdc.MethodBindPtr
}

var ptrsForImage ptrsForImageList

func initImagePtrs(iface gdc.Interface) {

  className := StringNameFromStr("Image")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_width")
    defer methodName.Destroy()
    ptrsForImage.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForImage.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForImage.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("has_mipmaps")
    defer methodName.Destroy()
    ptrsForImage.fnHasMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_format")
    defer methodName.Destroy()
    ptrsForImage.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3847873762))
  }
  {
    methodName := StringNameFromStr("get_data")
    defer methodName.Destroy()
    ptrsForImage.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
  }
  {
    methodName := StringNameFromStr("convert")
    defer methodName.Destroy()
    ptrsForImage.fnConvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2120693146))
  }
  {
    methodName := StringNameFromStr("get_mipmap_count")
    defer methodName.Destroy()
    ptrsForImage.fnGetMipmapCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_mipmap_offset")
    defer methodName.Destroy()
    ptrsForImage.fnGetMipmapOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("resize_to_po2")
    defer methodName.Destroy()
    ptrsForImage.fnResizeToPo2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4189212329))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForImage.fnResize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 994498151))
  }
  {
    methodName := StringNameFromStr("shrink_x2")
    defer methodName.Destroy()
    ptrsForImage.fnShrinkX2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("crop")
    defer methodName.Destroy()
    ptrsForImage.fnCrop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("flip_x")
    defer methodName.Destroy()
    ptrsForImage.fnFlipX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("flip_y")
    defer methodName.Destroy()
    ptrsForImage.fnFlipY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("generate_mipmaps")
    defer methodName.Destroy()
    ptrsForImage.fnGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1633102583))
  }
  {
    methodName := StringNameFromStr("clear_mipmaps")
    defer methodName.Destroy()
    ptrsForImage.fnClearMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("create")
    defer methodName.Destroy()
    ptrsForImage.fnCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 986942177))
  }
  {
    methodName := StringNameFromStr("create_from_data")
    defer methodName.Destroy()
    ptrsForImage.fnCreateFromData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 299398494))
  }
  {
    methodName := StringNameFromStr("set_data")
    defer methodName.Destroy()
    ptrsForImage.fnSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2740482212))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForImage.fnIsEmpty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("load")
    defer methodName.Destroy()
    ptrsForImage.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("load_from_file")
    defer methodName.Destroy()
    ptrsForImage.fnLoadFromFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 736337515))
  }
  {
    methodName := StringNameFromStr("save_png")
    defer methodName.Destroy()
    ptrsForImage.fnSavePng = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2113323047))
  }
  {
    methodName := StringNameFromStr("save_png_to_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnSavePngToBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
  }
  {
    methodName := StringNameFromStr("save_jpg")
    defer methodName.Destroy()
    ptrsForImage.fnSaveJpg = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2800019068))
  }
  {
    methodName := StringNameFromStr("save_jpg_to_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnSaveJpgToBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 592235273))
  }
  {
    methodName := StringNameFromStr("save_exr")
    defer methodName.Destroy()
    ptrsForImage.fnSaveExr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3108122999))
  }
  {
    methodName := StringNameFromStr("save_exr_to_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnSaveExrToBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3178917920))
  }
  {
    methodName := StringNameFromStr("save_webp")
    defer methodName.Destroy()
    ptrsForImage.fnSaveWebp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2781156876))
  }
  {
    methodName := StringNameFromStr("save_webp_to_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnSaveWebpToBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214628238))
  }
  {
    methodName := StringNameFromStr("detect_alpha")
    defer methodName.Destroy()
    ptrsForImage.fnDetectAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2030116505))
  }
  {
    methodName := StringNameFromStr("is_invisible")
    defer methodName.Destroy()
    ptrsForImage.fnIsInvisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("detect_used_channels")
    defer methodName.Destroy()
    ptrsForImage.fnDetectUsedChannels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2703139984))
  }
  {
    methodName := StringNameFromStr("compress")
    defer methodName.Destroy()
    ptrsForImage.fnCompress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2975424957))
  }
  {
    methodName := StringNameFromStr("compress_from_channels")
    defer methodName.Destroy()
    ptrsForImage.fnCompressFromChannels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4212890953))
  }
  {
    methodName := StringNameFromStr("decompress")
    defer methodName.Destroy()
    ptrsForImage.fnDecompress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("is_compressed")
    defer methodName.Destroy()
    ptrsForImage.fnIsCompressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("rotate_90")
    defer methodName.Destroy()
    ptrsForImage.fnRotate90 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1901204267))
  }
  {
    methodName := StringNameFromStr("rotate_180")
    defer methodName.Destroy()
    ptrsForImage.fnRotate180 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("fix_alpha_edges")
    defer methodName.Destroy()
    ptrsForImage.fnFixAlphaEdges = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("premultiply_alpha")
    defer methodName.Destroy()
    ptrsForImage.fnPremultiplyAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("srgb_to_linear")
    defer methodName.Destroy()
    ptrsForImage.fnSrgbToLinear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("normal_map_to_xy")
    defer methodName.Destroy()
    ptrsForImage.fnNormalMapToXy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("rgbe_to_srgb")
    defer methodName.Destroy()
    ptrsForImage.fnRgbeToSrgb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 564927088))
  }
  {
    methodName := StringNameFromStr("bump_map_to_normal_map")
    defer methodName.Destroy()
    ptrsForImage.fnBumpMapToNormalMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3423495036))
  }
  {
    methodName := StringNameFromStr("compute_image_metrics")
    defer methodName.Destroy()
    ptrsForImage.fnComputeImageMetrics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3080961247))
  }
  {
    methodName := StringNameFromStr("blit_rect")
    defer methodName.Destroy()
    ptrsForImage.fnBlitRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2903928755))
  }
  {
    methodName := StringNameFromStr("blit_rect_mask")
    defer methodName.Destroy()
    ptrsForImage.fnBlitRectMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383581145))
  }
  {
    methodName := StringNameFromStr("blend_rect")
    defer methodName.Destroy()
    ptrsForImage.fnBlendRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2903928755))
  }
  {
    methodName := StringNameFromStr("blend_rect_mask")
    defer methodName.Destroy()
    ptrsForImage.fnBlendRectMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383581145))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForImage.fnFill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("fill_rect")
    defer methodName.Destroy()
    ptrsForImage.fnFillRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 514693913))
  }
  {
    methodName := StringNameFromStr("get_used_rect")
    defer methodName.Destroy()
    ptrsForImage.fnGetUsedRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410525958))
  }
  {
    methodName := StringNameFromStr("get_region")
    defer methodName.Destroy()
    ptrsForImage.fnGetRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2601441065))
  }
  {
    methodName := StringNameFromStr("copy_from")
    defer methodName.Destroy()
    ptrsForImage.fnCopyFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
  }
  {
    methodName := StringNameFromStr("get_pixelv")
    defer methodName.Destroy()
    ptrsForImage.fnGetPixelv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1532707496))
  }
  {
    methodName := StringNameFromStr("get_pixel")
    defer methodName.Destroy()
    ptrsForImage.fnGetPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2165839948))
  }
  {
    methodName := StringNameFromStr("set_pixelv")
    defer methodName.Destroy()
    ptrsForImage.fnSetPixelv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 287851464))
  }
  {
    methodName := StringNameFromStr("set_pixel")
    defer methodName.Destroy()
    ptrsForImage.fnSetPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3733378741))
  }
  {
    methodName := StringNameFromStr("adjust_bcs")
    defer methodName.Destroy()
    ptrsForImage.fnAdjustBcs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2385087082))
  }
  {
    methodName := StringNameFromStr("load_png_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadPngFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_jpg_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadJpgFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_webp_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadWebpFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_tga_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadTgaFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_bmp_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadBmpFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_ktx_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadKtxFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("load_svg_from_buffer")
    defer methodName.Destroy()
    ptrsForImage.fnLoadSvgFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 311853421))
  }
  {
    methodName := StringNameFromStr("load_svg_from_string")
    defer methodName.Destroy()
    ptrsForImage.fnLoadSvgFromString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3254053600))
  }
}

type Image struct {
  Resource
}

func (me *Image) BaseClass() string {
  return "Image"
}

func NewImage() *Image {
  str := StringNameFromStr("Image") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Image{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  ImageMaxWidth = "16777216" // TODO: construct correctly
  ImageMaxHeight = "16777216" // TODO: construct correctly
)

// Enums

type ImageFormat int
const (
  ImageFormatFormatL8 ImageFormat = 0
  ImageFormatFormatLa8 ImageFormat = 1
  ImageFormatFormatR8 ImageFormat = 2
  ImageFormatFormatRg8 ImageFormat = 3
  ImageFormatFormatRgb8 ImageFormat = 4
  ImageFormatFormatRgba8 ImageFormat = 5
  ImageFormatFormatRgba4444 ImageFormat = 6
  ImageFormatFormatRgb565 ImageFormat = 7
  ImageFormatFormatRf ImageFormat = 8
  ImageFormatFormatRgf ImageFormat = 9
  ImageFormatFormatRgbf ImageFormat = 10
  ImageFormatFormatRgbaf ImageFormat = 11
  ImageFormatFormatRh ImageFormat = 12
  ImageFormatFormatRgh ImageFormat = 13
  ImageFormatFormatRgbh ImageFormat = 14
  ImageFormatFormatRgbah ImageFormat = 15
  ImageFormatFormatRgbe9995 ImageFormat = 16
  ImageFormatFormatDxt1 ImageFormat = 17
  ImageFormatFormatDxt3 ImageFormat = 18
  ImageFormatFormatDxt5 ImageFormat = 19
  ImageFormatFormatRgtcR ImageFormat = 20
  ImageFormatFormatRgtcRg ImageFormat = 21
  ImageFormatFormatBptcRgba ImageFormat = 22
  ImageFormatFormatBptcRgbf ImageFormat = 23
  ImageFormatFormatBptcRgbfu ImageFormat = 24
  ImageFormatFormatEtc ImageFormat = 25
  ImageFormatFormatEtc2R11 ImageFormat = 26
  ImageFormatFormatEtc2R11S ImageFormat = 27
  ImageFormatFormatEtc2Rg11 ImageFormat = 28
  ImageFormatFormatEtc2Rg11S ImageFormat = 29
  ImageFormatFormatEtc2Rgb8 ImageFormat = 30
  ImageFormatFormatEtc2Rgba8 ImageFormat = 31
  ImageFormatFormatEtc2Rgb8A1 ImageFormat = 32
  ImageFormatFormatEtc2RaAsRg ImageFormat = 33
  ImageFormatFormatDxt5RaAsRg ImageFormat = 34
  ImageFormatFormatAstc4X4 ImageFormat = 35
  ImageFormatFormatAstc4X4Hdr ImageFormat = 36
  ImageFormatFormatAstc8X8 ImageFormat = 37
  ImageFormatFormatAstc8X8Hdr ImageFormat = 38
  ImageFormatFormatMax ImageFormat = 39
)

type ImageInterpolation int
const (
  ImageInterpolationInterpolateNearest ImageInterpolation = 0
  ImageInterpolationInterpolateBilinear ImageInterpolation = 1
  ImageInterpolationInterpolateCubic ImageInterpolation = 2
  ImageInterpolationInterpolateTrilinear ImageInterpolation = 3
  ImageInterpolationInterpolateLanczos ImageInterpolation = 4
)

type ImageAlphaMode int
const (
  ImageAlphaModeAlphaNone ImageAlphaMode = 0
  ImageAlphaModeAlphaBit ImageAlphaMode = 1
  ImageAlphaModeAlphaBlend ImageAlphaMode = 2
)

type ImageCompressMode int
const (
  ImageCompressModeCompressS3Tc ImageCompressMode = 0
  ImageCompressModeCompressEtc ImageCompressMode = 1
  ImageCompressModeCompressEtc2 ImageCompressMode = 2
  ImageCompressModeCompressBptc ImageCompressMode = 3
  ImageCompressModeCompressAstc ImageCompressMode = 4
  ImageCompressModeCompressMax ImageCompressMode = 5
)

type ImageUsedChannels int
const (
  ImageUsedChannelsUsedChannelsL ImageUsedChannels = 0
  ImageUsedChannelsUsedChannelsLa ImageUsedChannels = 1
  ImageUsedChannelsUsedChannelsR ImageUsedChannels = 2
  ImageUsedChannelsUsedChannelsRg ImageUsedChannels = 3
  ImageUsedChannelsUsedChannelsRgb ImageUsedChannels = 4
  ImageUsedChannelsUsedChannelsRgba ImageUsedChannels = 5
)

type ImageCompressSource int
const (
  ImageCompressSourceCompressSourceGeneric ImageCompressSource = 0
  ImageCompressSourceCompressSourceSrgb ImageCompressSource = 1
  ImageCompressSourceCompressSourceNormal ImageCompressSource = 2
)

type ImageASTCFormat int
const (
  ImageASTCFormatAstcFormat4X4 ImageASTCFormat = 0
  ImageASTCFormatAstcFormat8X8 ImageASTCFormat = 1
)

func (me *Image) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Image) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Image) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Image) GetWidth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) GetHeight() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) GetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) HasMipmaps() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnHasMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) GetFormat() ImageFormat {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) GetData() PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) Convert(format ImageFormat, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnConvert), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) GetMipmapCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetMipmapCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) GetMipmapOffset(mipmap int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mipmap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&mipmap)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetMipmapOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) ResizeToPo2(square bool, interpolation ImageInterpolation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&square) , gdc.ConstTypePtr(&interpolation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnResizeToPo2), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) Resize(width int64, height int64, interpolation ImageInterpolation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&interpolation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnResize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) ShrinkX2()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnShrinkX2), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) Crop(width int64, height int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCrop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) FlipX()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnFlipX), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) FlipY()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnFlipY), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) GenerateMipmaps(renormalize bool, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&renormalize) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&renormalize)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGenerateMipmaps), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) ClearMipmaps()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnClearMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func  ImageCreate(width int64, height int64, use_mipmaps bool, format ImageFormat, ) Image {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&use_mipmaps) , gdc.ConstTypePtr(&format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()
  pinner.Pin(&width)
  pinner.Pin(&height)
  pinner.Pin(&use_mipmaps)
  pinner.Pin(&format)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCreate), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  ImageCreateFromData(width int64, height int64, use_mipmaps bool, format ImageFormat, data PackedByteArray, ) Image {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&use_mipmaps) , gdc.ConstTypePtr(&format) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()
  pinner.Pin(&width)
  pinner.Pin(&height)
  pinner.Pin(&use_mipmaps)
  pinner.Pin(&format)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCreateFromData), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SetData(width int64, height int64, use_mipmaps bool, format ImageFormat, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&use_mipmaps) , gdc.ConstTypePtr(&format) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) IsEmpty() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnIsEmpty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) Load(path String, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  ImageLoadFromFile(path String, ) Image {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadFromFile), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SavePng(path String, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSavePng), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) SavePngToBuffer() PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSavePngToBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SaveJpg(path String, quality float64, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&quality)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveJpg), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) SaveJpgToBuffer(quality float64, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&quality)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveJpgToBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SaveExr(path String, grayscale bool, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&grayscale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&grayscale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveExr), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) SaveExrToBuffer(grayscale bool, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&grayscale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&grayscale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveExrToBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SaveWebp(path String, lossy bool, quality float64, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&lossy) , gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&lossy)
  pinner.Pin(&quality)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveWebp), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) SaveWebpToBuffer(lossy bool, quality float64, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lossy) , gdc.ConstTypePtr(&quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&lossy)
  pinner.Pin(&quality)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSaveWebpToBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) DetectAlpha() ImageAlphaMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageAlphaMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnDetectAlpha), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) IsInvisible() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnIsInvisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) DetectUsedChannels(source ImageCompressSource, ) ImageUsedChannels {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageUsedChannels
  pinner.Pin(&source)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnDetectUsedChannels), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) Compress(mode ImageCompressMode, source ImageCompressSource, astc_format ImageASTCFormat, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&source) , gdc.ConstTypePtr(&astc_format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&mode)
  pinner.Pin(&source)
  pinner.Pin(&astc_format)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCompress), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) CompressFromChannels(mode ImageCompressMode, channels ImageUsedChannels, astc_format ImageASTCFormat, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&channels) , gdc.ConstTypePtr(&astc_format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&mode)
  pinner.Pin(&channels)
  pinner.Pin(&astc_format)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCompressFromChannels), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) Decompress() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnDecompress), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) IsCompressed() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnIsCompressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Image) Rotate90(direction ClockDirection, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnRotate90), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) Rotate180()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnRotate180), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) FixAlphaEdges()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnFixAlphaEdges), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) PremultiplyAlpha()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnPremultiplyAlpha), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) SrgbToLinear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSrgbToLinear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) NormalMapToXy()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnNormalMapToXy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) RgbeToSrgb() Image {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnRgbeToSrgb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) BumpMapToNormalMap(bump_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bump_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnBumpMapToNormalMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) ComputeImageMetrics(compared_image Image, use_luma bool, ) Dictionary {
  cargs := []gdc.ConstTypePtr{compared_image.AsCTypePtr(), gdc.ConstTypePtr(&use_luma) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&use_luma)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnComputeImageMetrics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) BlitRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  cargs := []gdc.ConstTypePtr{src.AsCTypePtr(), src_rect.AsCTypePtr(), dst.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnBlitRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) BlitRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  cargs := []gdc.ConstTypePtr{src.AsCTypePtr(), mask.AsCTypePtr(), src_rect.AsCTypePtr(), dst.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnBlitRectMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) BlendRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  cargs := []gdc.ConstTypePtr{src.AsCTypePtr(), src_rect.AsCTypePtr(), dst.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnBlendRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) BlendRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  cargs := []gdc.ConstTypePtr{src.AsCTypePtr(), mask.AsCTypePtr(), src_rect.AsCTypePtr(), dst.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnBlendRectMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) Fill(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnFill), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) FillRect(rect Rect2i, color Color, )  {
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnFillRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) GetUsedRect() Rect2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetUsedRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) GetRegion(region Rect2i, ) Image {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) CopyFrom(src Image, )  {
  cargs := []gdc.ConstTypePtr{src.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnCopyFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) GetPixelv(point Vector2i, ) Color {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetPixelv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) GetPixel(x int64, y int64, ) Color {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&x)
  pinner.Pin(&y)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnGetPixel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Image) SetPixelv(point Vector2i, color Color, )  {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSetPixelv), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) SetPixel(x int64, y int64, color Color, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnSetPixel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) AdjustBcs(brightness float64, contrast float64, saturation float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brightness) , gdc.ConstTypePtr(&contrast) , gdc.ConstTypePtr(&saturation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnAdjustBcs), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Image) LoadPngFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadPngFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadJpgFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadJpgFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadWebpFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadWebpFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadTgaFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadTgaFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadBmpFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadBmpFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadKtxFromBuffer(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadKtxFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadSvgFromBuffer(buffer PackedByteArray, scale float64, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&scale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadSvgFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Image) LoadSvgFromString(svg_str String, scale float64, ) Error {
  cargs := []gdc.ConstTypePtr{svg_str.AsCTypePtr(), gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&scale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImage.fnLoadSvgFromString), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
