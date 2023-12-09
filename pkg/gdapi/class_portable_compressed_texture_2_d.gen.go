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



// Enums

type PortableCompressedTexture2DCompressionMode int
const (
  PortableCompressedTexture2DCompressionModeCompressionModeLossless PortableCompressedTexture2DCompressionMode = 0
  PortableCompressedTexture2DCompressionModeCompressionModeLossy PortableCompressedTexture2DCompressionMode = 1
  PortableCompressedTexture2DCompressionModeCompressionModeBasisUniversal PortableCompressedTexture2DCompressionMode = 2
  PortableCompressedTexture2DCompressionModeCompressionModeS3Tc PortableCompressedTexture2DCompressionMode = 3
  PortableCompressedTexture2DCompressionModeCompressionModeEtc2 PortableCompressedTexture2DCompressionMode = 4
  PortableCompressedTexture2DCompressionModeCompressionModeBptc PortableCompressedTexture2DCompressionMode = 5
)

func (me *PortableCompressedTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PortableCompressedTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PortableCompressedTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PortableCompressedTexture2D) CreateFromImage(image Image, compression_mode PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float32, )  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) GetFormat()  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) GetCompressionMode()  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) SetSizeOverride(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) GetSizeOverride()  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) SetKeepCompressedBuffer(keep bool, )  {
  panic("TODO: implement")
}

func  (me *PortableCompressedTexture2D) IsKeepingCompressedBuffer()  {
  panic("TODO: implement")
}

func  PortableCompressedTexture2DSetKeepAllCompressedBuffers(keep bool, )  {
  panic("TODO: implement")
}

func  PortableCompressedTexture2DIsKeepingAllCompressedBuffers()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
