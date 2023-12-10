// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *Image) GetWidth() int {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetHeight() int {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetSize() Vector2i {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) HasMipmaps() bool {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetFormat() ImageFormat {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3847873762) // FIXME: should cache?
  var ret ImageFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetData() PackedByteArray {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) Convert(format ImageFormat, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2120693146) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) GetMipmapOffset(mipmap int, ) int {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mipmap_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mipmap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) ResizeToPo2(square bool, interpolation ImageInterpolation, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize_to_po2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4189212329) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&square), gdc.ConstTypePtr(&interpolation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) Resize(width int, height int, interpolation ImageInterpolation, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2461393748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&interpolation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) ShrinkX2()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shrink_x2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) Crop(width int, height int, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("crop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) FlipX()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flip_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) FlipY()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flip_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) GenerateMipmaps(renormalize bool, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1633102583) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&renormalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) ClearMipmaps()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  ImageCreate(width int, height int, use_mipmaps bool, format ImageFormat, ) Image {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 986942177) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&use_mipmaps), gdc.ConstTypePtr(&format), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  ImageCreateFromData(width int, height int, use_mipmaps bool, format ImageFormat, data PackedByteArray, ) Image {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 299398494) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&use_mipmaps), gdc.ConstTypePtr(&format), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SetData(width int, height int, use_mipmaps bool, format ImageFormat, data PackedByteArray, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2740482212) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&use_mipmaps), gdc.ConstTypePtr(&format), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) IsEmpty() bool {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_empty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) Load(path String, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  ImageLoadFromFile(path String, ) Image {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 736337515) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SavePng(path String, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_png")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2113323047) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SavePngToBuffer() PackedByteArray {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_png_to_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveJpg(path String, quality float32, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_jpg")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 578836491) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&quality), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveJpgToBuffer(quality float32, ) PackedByteArray {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_jpg_to_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 592235273) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quality), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveExr(path String, grayscale bool, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_exr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3108122999) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&grayscale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveExrToBuffer(grayscale bool, ) PackedByteArray {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_exr_to_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3178917920) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&grayscale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveWebp(path String, lossy bool, quality float32, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_webp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3594949219) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&lossy), gdc.ConstTypePtr(&quality), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SaveWebpToBuffer(lossy bool, quality float32, ) PackedByteArray {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_webp_to_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214628238) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lossy), gdc.ConstTypePtr(&quality), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) DetectAlpha() ImageAlphaMode {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("detect_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2030116505) // FIXME: should cache?
  var ret ImageAlphaMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) IsInvisible() bool {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_invisible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) DetectUsedChannels(source ImageCompressSource, ) ImageUsedChannels {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("detect_used_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2703139984) // FIXME: should cache?
  var ret ImageUsedChannels
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) Compress(mode ImageCompressMode, source ImageCompressSource, astc_format ImageASTCFormat, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4094210332) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&source), gdc.ConstTypePtr(&astc_format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) CompressFromChannels(mode ImageCompressMode, channels ImageUsedChannels, astc_format ImageASTCFormat, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compress_from_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 279105990) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&channels), gdc.ConstTypePtr(&astc_format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) Decompress() Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("decompress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) IsCompressed() bool {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_compressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) Rotate90(direction ClockDirection, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_90")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1901204267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) Rotate180()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_180")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) FixAlphaEdges()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fix_alpha_edges")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) PremultiplyAlpha()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("premultiply_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) SrgbToLinear()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("srgb_to_linear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) NormalMapToXy()  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("normal_map_to_xy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) RgbeToSrgb() Image {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rgbe_to_srgb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 564927088) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) BumpMapToNormalMap(bump_scale float32, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bump_map_to_normal_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3423495036) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bump_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) ComputeImageMetrics(compared_image Image, use_luma bool, ) Dictionary {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compute_image_metrics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3080961247) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(compared_image.AsCTypePtr()), gdc.ConstTypePtr(&use_luma), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) BlitRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blit_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2903928755) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(dst.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) BlitRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blit_rect_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3383581145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), gdc.ConstTypePtr(mask.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(dst.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) BlendRect(src Image, src_rect Rect2i, dst Vector2i, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2903928755) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(dst.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) BlendRectMask(src Image, mask Image, src_rect Rect2i, dst Vector2i, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_rect_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3383581145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), gdc.ConstTypePtr(mask.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(dst.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) Fill(color Color, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) FillRect(rect Rect2i, color Color, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fill_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 514693913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) GetUsedRect() Rect2i {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410525958) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetRegion(region Rect2i, ) Image {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2601441065) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) CopyFrom(src Image, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("copy_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) GetPixelv(point Vector2i, ) Color {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pixelv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1532707496) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) GetPixel(x int, y int, ) Color {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pixel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2165839948) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) SetPixelv(point Vector2i, color Color, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pixelv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 287851464) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) SetPixel(x int, y int, color Color, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pixel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3733378741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) AdjustBcs(brightness float32, contrast float32, saturation float32, )  {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("adjust_bcs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2385087082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brightness), gdc.ConstTypePtr(&contrast), gdc.ConstTypePtr(&saturation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Image) LoadPngFromBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_png_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) LoadJpgFromBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_jpg_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) LoadWebpFromBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_webp_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) LoadTgaFromBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_tga_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Image) LoadBmpFromBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("Image")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_bmp_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Image) GetPropData() Dictionary {
  panic("TODO: implement")
}

func (me *Image) SetPropData(value Dictionary) {
  panic("TODO: implement")
}