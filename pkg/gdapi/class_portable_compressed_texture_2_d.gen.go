// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PortableCompressedTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *PortableCompressedTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PortableCompressedTexture2D) BaseClass() string {
  return "PortableCompressedTexture2D"
}

type PortableCompressedTexture2DCompressionMode int
const (
  PortableCompressedTexture2DCompressionModeLossless PortableCompressedTexture2DCompressionMode = 0
  PortableCompressedTexture2DCompressionModeLossy PortableCompressedTexture2DCompressionMode = 1
  PortableCompressedTexture2DCompressionModeBasisUniversal PortableCompressedTexture2DCompressionMode = 2
  PortableCompressedTexture2DCompressionModeS3Tc PortableCompressedTexture2DCompressionMode = 3
  PortableCompressedTexture2DCompressionModeEtc2 PortableCompressedTexture2DCompressionMode = 4
  PortableCompressedTexture2DCompressionModeBptc PortableCompressedTexture2DCompressionMode = 5
)
