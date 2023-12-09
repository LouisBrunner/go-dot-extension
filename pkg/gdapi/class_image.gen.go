// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Image struct {
  obj gdc.ObjectPtr
}

func (me *Image) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Image) BaseClass() string {
  return "Image"
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

func  (me *Image) GetWidth()  {
  panic("TODO: implement")
}

func  (me *Image) GetHeight()  {
  panic("TODO: implement")
}

func  (me *Image) GetSize()  {
  panic("TODO: implement")
}

func  (me *Image) HasMipmaps()  {
  panic("TODO: implement")
}

func  (me *Image) GetFormat()  {
  panic("TODO: implement")
}

func  (me *Image) GetData()  {
  panic("TODO: implement")
}

func  (me *Image) Convert(format ImageFormat, )  {
  panic("TODO: implement")
}

func  (me *Image) GetMipmapOffset(mipmap int, )  {
  panic("TODO: implement")
}

func  (me *Image) ResizeToPo2(square bool, interpolation ImageInterpolation, )  {
  panic("TODO: implement")
}

func  (me *Image) Resize(width int, height int, interpolation ImageInterpolation, )  {
  panic("TODO: implement")
}

func  (me *Image) ShrinkX2()  {
  panic("TODO: implement")
}

func  (me *Image) Crop(width int, height int, )  {
  panic("TODO: implement")
}

func  (me *Image) FlipX()  {
  panic("TODO: implement")
}

func  (me *Image) FlipY()  {
  panic("TODO: implement")
}

func  (me *Image) GenerateMipmaps(renormalize bool, )  {
  panic("TODO: implement")
}

func  (me *Image) ClearMipmaps()  {
  panic("TODO: implement")
}

func  ImageCreate(width int, height int, use_mipmaps bool, format ImageFormat, )  {
  panic("TODO: implement")
}

func  ImageCreateFromData(width int, height int, use_mipmaps bool, format ImageFormat, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) SetData(width int, height int, use_mipmaps bool, format ImageFormat, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) IsEmpty()  {
  panic("TODO: implement")
}

func  (me *Image) Load(path String, )  {
  panic("TODO: implement")
}

func  ImageLoadFromFile(path String, )  {
  panic("TODO: implement")
}

func  (me *Image) SavePng(path String, )  {
  panic("TODO: implement")
}

func  (me *Image) SavePngToBuffer()  {
  panic("TODO: implement")
}

func  (me *Image) SaveJpg(path String, quality float32, )  {
  panic("TODO: implement")
}

func  (me *Image) SaveJpgToBuffer(quality float32, )  {
  panic("TODO: implement")
}

func  (me *Image) SaveExr(path String, grayscale bool, )  {
  panic("TODO: implement")
}

func  (me *Image) SaveExrToBuffer(grayscale bool, )  {
  panic("TODO: implement")
}

func  (me *Image) SaveWebp(path String, lossy bool, quality float32, )  {
  panic("TODO: implement")
}

func  (me *Image) SaveWebpToBuffer(lossy bool, quality float32, )  {
  panic("TODO: implement")
}

func  (me *Image) DetectAlpha()  {
  panic("TODO: implement")
}

func  (me *Image) IsInvisible()  {
  panic("TODO: implement")
}

func  (me *Image) DetectUsedChannels(source ImageCompressSource, )  {
  panic("TODO: implement")
}

func  (me *Image) Compress(mode ImageCompressMode, source ImageCompressSource, astc_format ImageASTCFormat, )  {
  panic("TODO: implement")
}

func  (me *Image) CompressFromChannels(mode ImageCompressMode, channels ImageUsedChannels, astc_format ImageASTCFormat, )  {
  panic("TODO: implement")
}

func  (me *Image) Decompress()  {
  panic("TODO: implement")
}

func  (me *Image) IsCompressed()  {
  panic("TODO: implement")
}

func  (me *Image) Rotate90(direction ClockDirection, )  {
  panic("TODO: implement")
}

func  (me *Image) Rotate180()  {
  panic("TODO: implement")
}

func  (me *Image) FixAlphaEdges()  {
  panic("TODO: implement")
}

func  (me *Image) PremultiplyAlpha()  {
  panic("TODO: implement")
}

func  (me *Image) SrgbToLinear()  {
  panic("TODO: implement")
}

func  (me *Image) NormalMapToXy()  {
  panic("TODO: implement")
}

func  (me *Image) RgbeToSrgb()  {
  panic("TODO: implement")
}

func  (me *Image) BumpMapToNormalMap(bump_scale float32, )  {
  panic("TODO: implement")
}

func  (me *Image) ComputeImageMetrics(compared_image Image, use_luma bool, )  {
  panic("TODO: implement")
}

func  (me *Image) BlitRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Image) BlitRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Image) BlendRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Image) BlendRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Image) Fill(color Color, )  {
  panic("TODO: implement")
}

func  (me *Image) FillRect(rect Rect2i, color Color, )  {
  panic("TODO: implement")
}

func  (me *Image) GetUsedRect()  {
  panic("TODO: implement")
}

func  (me *Image) GetRegion(region Rect2i, )  {
  panic("TODO: implement")
}

func  (me *Image) CopyFrom(src Image, )  {
  panic("TODO: implement")
}

func  (me *Image) GetPixelv(point Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Image) GetPixel(x int, y int, )  {
  panic("TODO: implement")
}

func  (me *Image) SetPixelv(point Vector2i, color Color, )  {
  panic("TODO: implement")
}

func  (me *Image) SetPixel(x int, y int, color Color, )  {
  panic("TODO: implement")
}

func  (me *Image) AdjustBcs(brightness float32, contrast float32, saturation float32, )  {
  panic("TODO: implement")
}

func  (me *Image) LoadPngFromBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) LoadJpgFromBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) LoadWebpFromBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) LoadTgaFromBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Image) LoadBmpFromBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
