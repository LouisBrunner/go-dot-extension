// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  PortableCompressedTexture2DCompressionModeCompressionModeLossless PortableCompressedTexture2DCompressionMode = 0
  PortableCompressedTexture2DCompressionModeCompressionModeLossy PortableCompressedTexture2DCompressionMode = 1
  PortableCompressedTexture2DCompressionModeCompressionModeBasisUniversal PortableCompressedTexture2DCompressionMode = 2
  PortableCompressedTexture2DCompressionModeCompressionModeS3Tc PortableCompressedTexture2DCompressionMode = 3
  PortableCompressedTexture2DCompressionModeCompressionModeEtc2 PortableCompressedTexture2DCompressionMode = 4
  PortableCompressedTexture2DCompressionModeCompressionModeBptc PortableCompressedTexture2DCompressionMode = 5
)

func  (me *PortableCompressedTexture2D) CreateFromImage(image Image, compression_mode PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) GetFormat() { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) GetCompressionMode() { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) SetSizeOverride(size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) GetSizeOverride() { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) SetKeepCompressedBuffer(keep bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PortableCompressedTexture2D) IsKeepingCompressedBuffer() { // TODO: return value
  // TODO: implement
}

func  PortableCompressedTexture2DSetKeepAllCompressedBuffers(keep bool, ) { // TODO: return value
  // TODO: implement
}

func  PortableCompressedTexture2DIsKeepingAllCompressedBuffers() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
