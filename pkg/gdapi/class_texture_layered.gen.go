// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureLayered struct {
  Texture
}

func (me *TextureLayered) BaseClass() string {
  return "TextureLayered"
}

func NewTextureLayered() *TextureLayered {
  str := StringNameFromStr("TextureLayered") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextureLayered{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TextureLayeredLayeredType int
const (
  TextureLayeredLayeredTypeLayeredType2DArray TextureLayeredLayeredType = 0
  TextureLayeredLayeredTypeLayeredTypeCubemap TextureLayeredLayeredType = 1
  TextureLayeredLayeredTypeLayeredTypeCubemapArray TextureLayeredLayeredType = 2
)

func (me *TextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextureLayered) GetFormat() ImageFormat {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3847873762) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextureLayered) GetLayeredType() TextureLayeredLayeredType {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layered_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518123893) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextureLayeredLayeredType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextureLayered) GetWidth() int64 {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureLayered) GetHeight() int64 {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureLayered) GetLayers() int64 {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureLayered) HasMipmaps() bool {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureLayered) GetLayerData(layer int64, ) Image {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3655284255) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
