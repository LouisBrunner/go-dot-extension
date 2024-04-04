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

type PortableCompressedTexture2D struct {
  Texture2D
}

func (me *PortableCompressedTexture2D) BaseClass() string {
  return "PortableCompressedTexture2D"
}

func NewPortableCompressedTexture2D() *PortableCompressedTexture2D {
  str := StringNameFromStr("PortableCompressedTexture2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PortableCompressedTexture2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PortableCompressedTexture2D) CreateFromImage(image Image, compression_mode PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float64, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3679243433) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&compression_mode) , gdc.ConstTypePtr(&normal_map) , gdc.ConstTypePtr(&lossy_quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PortableCompressedTexture2D) GetFormat() ImageFormat {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3847873762) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PortableCompressedTexture2D) GetCompressionMode() PortableCompressedTexture2DCompressionMode {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_compression_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3265612739) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PortableCompressedTexture2DCompressionMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PortableCompressedTexture2D) SetSizeOverride(size Vector2, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PortableCompressedTexture2D) GetSizeOverride() Vector2 {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PortableCompressedTexture2D) SetKeepCompressedBuffer(keep bool, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_compressed_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PortableCompressedTexture2D) IsKeepingCompressedBuffer() bool {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keeping_compressed_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  PortableCompressedTexture2DSetKeepAllCompressedBuffers(keep bool, )  {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_all_compressed_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}

func  PortableCompressedTexture2DIsKeepingAllCompressedBuffers() bool {
  classNameV := StringNameFromStr("PortableCompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keeping_all_compressed_buffers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
