// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  ImageMAX_WIDTH = 16777216
  ImageMAX_HEIGHT = 16777216
)

type ImageFormat int
const (
  ImageFormatL8 ImageFormat = 0
  ImageFormatLa8 ImageFormat = 1
  ImageFormatR8 ImageFormat = 2
  ImageFormatRg8 ImageFormat = 3
  ImageFormatRgb8 ImageFormat = 4
  ImageFormatRgba8 ImageFormat = 5
  ImageFormatRgba4444 ImageFormat = 6
  ImageFormatRgb565 ImageFormat = 7
  ImageFormatRf ImageFormat = 8
  ImageFormatRgf ImageFormat = 9
  ImageFormatRgbf ImageFormat = 10
  ImageFormatRgbaf ImageFormat = 11
  ImageFormatRh ImageFormat = 12
  ImageFormatRgh ImageFormat = 13
  ImageFormatRgbh ImageFormat = 14
  ImageFormatRgbah ImageFormat = 15
  ImageFormatRgbe9995 ImageFormat = 16
  ImageFormatDxt1 ImageFormat = 17
  ImageFormatDxt3 ImageFormat = 18
  ImageFormatDxt5 ImageFormat = 19
  ImageFormatRgtcR ImageFormat = 20
  ImageFormatRgtcRg ImageFormat = 21
  ImageFormatBptcRgba ImageFormat = 22
  ImageFormatBptcRgbf ImageFormat = 23
  ImageFormatBptcRgbfu ImageFormat = 24
  ImageFormatEtc ImageFormat = 25
  ImageFormatEtc2R11 ImageFormat = 26
  ImageFormatEtc2R11S ImageFormat = 27
  ImageFormatEtc2Rg11 ImageFormat = 28
  ImageFormatEtc2Rg11S ImageFormat = 29
  ImageFormatEtc2Rgb8 ImageFormat = 30
  ImageFormatEtc2Rgba8 ImageFormat = 31
  ImageFormatEtc2Rgb8A1 ImageFormat = 32
  ImageFormatEtc2RaAsRg ImageFormat = 33
  ImageFormatDxt5RaAsRg ImageFormat = 34
  ImageFormatAstc4X4 ImageFormat = 35
  ImageFormatAstc4X4Hdr ImageFormat = 36
  ImageFormatAstc8X8 ImageFormat = 37
  ImageFormatAstc8X8Hdr ImageFormat = 38
  ImageFormatMax ImageFormat = 39
)

type ImageInterpolation int
const (
  ImageInterpolateNearest ImageInterpolation = 0
  ImageInterpolateBilinear ImageInterpolation = 1
  ImageInterpolateCubic ImageInterpolation = 2
  ImageInterpolateTrilinear ImageInterpolation = 3
  ImageInterpolateLanczos ImageInterpolation = 4
)

type ImageAlphaMode int
const (
  ImageAlphaNone ImageAlphaMode = 0
  ImageAlphaBit ImageAlphaMode = 1
  ImageAlphaBlend ImageAlphaMode = 2
)

type ImageCompressMode int
const (
  ImageCompressS3Tc ImageCompressMode = 0
  ImageCompressEtc ImageCompressMode = 1
  ImageCompressEtc2 ImageCompressMode = 2
  ImageCompressBptc ImageCompressMode = 3
  ImageCompressAstc ImageCompressMode = 4
  ImageCompressMax ImageCompressMode = 5
)

type ImageUsedChannels int
const (
  ImageUsedChannelsL ImageUsedChannels = 0
  ImageUsedChannelsLa ImageUsedChannels = 1
  ImageUsedChannelsR ImageUsedChannels = 2
  ImageUsedChannelsRg ImageUsedChannels = 3
  ImageUsedChannelsRgb ImageUsedChannels = 4
  ImageUsedChannelsRgba ImageUsedChannels = 5
)

type ImageCompressSource int
const (
  ImageCompressSourceGeneric ImageCompressSource = 0
  ImageCompressSourceSrgb ImageCompressSource = 1
  ImageCompressSourceNormal ImageCompressSource = 2
)

type ImageASTCFormat int
const (
  ImageAstcFormat4X4 ImageASTCFormat = 0
  ImageAstcFormat8X8 ImageASTCFormat = 1
)
