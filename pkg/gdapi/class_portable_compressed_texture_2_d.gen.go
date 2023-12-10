// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 97251393) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(image.AsCTypePtr()), gdc.ConstTypePtr(&compression_mode), gdc.ConstTypePtr(&normal_map), gdc.ConstTypePtr(&lossy_quality), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PortableCompressedTexture2D) GetFormat() ImageFormat {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3847873762) // FIXME: should cache?
  var ret ImageFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PortableCompressedTexture2D) GetCompressionMode() PortableCompressedTexture2DCompressionMode {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_compression_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3265612739) // FIXME: should cache?
  var ret PortableCompressedTexture2DCompressionMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PortableCompressedTexture2D) SetSizeOverride(size Vector2, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PortableCompressedTexture2D) GetSizeOverride() Vector2 {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PortableCompressedTexture2D) SetKeepCompressedBuffer(keep bool, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_compressed_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PortableCompressedTexture2D) IsKeepingCompressedBuffer() bool {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keeping_compressed_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  PortableCompressedTexture2DSetKeepAllCompressedBuffers(keep bool, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_all_compressed_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

func  PortableCompressedTexture2DIsKeepingAllCompressedBuffers() bool {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keeping_all_compressed_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PortableCompressedTexture2D) GetPropSizeOverride() Vector2 {
  panic("TODO: implement")
}

func (me *PortableCompressedTexture2D) SetPropSizeOverride(value Vector2) {
  panic("TODO: implement")
}

func (me *PortableCompressedTexture2D) GetPropKeepCompressedBuffer() bool {
  panic("TODO: implement")
}

func (me *PortableCompressedTexture2D) SetPropKeepCompressedBuffer(value bool) {
  panic("TODO: implement")
}